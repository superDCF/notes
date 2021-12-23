# Replace ziplist with listpack in quicklist

导语——最近看到redis的官方消息，用listpack替换ziplist，迭代了多个版本的listpack终于登场了（当前最新6.2.6版本，仍未替换，只在unstable分支出现，预计下个版本发布）。恐怕有同学还不了解listpack有什么提升的地方，本篇将结合图片+源码的方式，图解listpack。

## 为什么要替换ziplist？
ziplist最大特点就是，是一种内存紧凑型的数据结构，占用一块连续的内存，针对不同长度类型的数据，进行不同的编码。既可以利用CPU缓存，也非常节省内存。
当然这种连续性的内存结构，在元素大小变化、删除、增加，也有着通用的毛病。

ziplist结构布局如下：
``` 
ziplist.c文件
 <zlbytes> <zltail> <zllen> <entry> <entry> ... <entry> <zlend>

uint32_t zlbytes; // 整个压缩列表占用字节数，包括zlbytes字段本身4个字节
uint32_t zltail; // 最后一个entry元素距离压缩列表起始位置的偏移量，用于快速定位到最后一个节点而不需要遍历整个链表
uint16_t zllen; // entry元素个数
void* entries; // 元素内容列表，挨个挨个紧凑存储
uint8_t zlend; // 代表ziplist的结束，等于255

//entry 结构布局
<prevlen> <encoding> <entry-data>
prevlen：前一个entry的长度，如果这个长度小于254字节，只占用1字节。当长度大于或等于254时，它将占用5个字节。第一个字节被设置为254（FE），以表示后面有一个更大的值。其余的4个字节以前一个entry的长度为值。通过这个值，可以进行指针计算，从而跳转到上一个节点。
encoding：该字段表示entry的数据类型以及数据长度，当是整形时，只占用1字节，字节的前2 bit都是1，例如【11xxxxxx】，后几个bit位表示不同的entry数据类型。当是字符串时，根据字符串的长度不同，分别占用1、2、5字节，同样多余的bit位，表示不同的数据长度。（多说一句，redis有很多这种可变数据编码的思想，在一定程度上节省内存）
```

**由于prevlen字段表示前一个元素的长度，当新增元素、元素变大超过现有值类型，例如增加一个大于254字节的元素，而之前存储的是小于254字节的entry，这时就会改变当前这个prevlen的值，可能出现更严重的是，会顺序性的对后面的entry的prevlen都需要修改，这会直接影响到压缩列表的操作性能，不过这种出现的概率很低。
更主要的原因是因为增加、删除、变大都需要移动后续的元素，所以压缩列表只适用于保存节点不多的情况，在3.2之前使用ziplist的数据类型有，list、hash、zset，他们分别是通过list_max_ziplist_value、hash_max_ziplist_entries、zset_max_ziplist_value设置当数据量小于多少时使用ziplist数据结构，默认值都是64。**

了解了ziplist的设计和使用的取舍，在redis 3.2版本使用quicklist替换redis List类型的底层数据结构ziplist[1]。

ziplist数据结构在3.2版本之前，作为list类型底层数据结构之一，3.2版本之后[1]，被替换成quicklist。

## quicklist是什么

```
typedef struct quicklist {
    quicklistNode *head; // 头节点指针
    quicklistNode *tail; // 尾节点指针
    unsigned long count;   // 所有entry的数量     /* total count of all entries in all listpacks */
    unsigned long len;         /* number of quicklistNodes */
    signed int fill : QL_FILL_BITS;  // 每个节点的填充系数，可以主动设置。正数代表个数，负数代表级别（后面讲）。也就是设置ziplist最大存储，存放list-max-ziplist-size参数的值     /* fill factor for individual nodes */
    unsigned int compress : QL_COMP_BITS; // 节点压缩深度 /* depth of end nodes not to compress;0=off */
    unsigned int bookmark_count: QL_BM_BITS;
    quicklistBookmark bookmarks[]; // 可选字段，在quicklist重新分配内存空间时使用，不使用时不占用空间
} quicklist;

typedef struct quicklistNode {
    struct quicklistNode *prev;
    struct quicklistNode *next;
    unsigned char *entry; // quicklistNode指向的listpack
    size_t sz;       // listpack的字节大小      /* entry size in bytes */
    unsigned int count : 16;     /* count of items in listpack */
    unsigned int encoding : 2;   /* RAW==1 or LZF==2 */ // 表示listpack是否压缩了，1-RAW表示原生没被压缩，2-LZF压缩算法
    unsigned int container : 2;  /* PLAIN==1 or PACKED==2 */，当前源码默认是2，表示使用PACKED作为数据容器（其实就是listpack）
    unsigned int recompress : 1; /* was this node previous compressed? */ // 表示节点是否被压缩
    unsigned int attempted_compress : 1; /* node can't compress; too small */ // 测试使用，不用管
    unsigned int extra : 10; /* more bits to steal for future usage */ // 保留字段
} quicklistNode;

typedef struct quicklistEntry {
    const quicklist *quicklist; // 属于哪一个quicklist
    quicklistNode *node; // 属于哪一个listpack
    unsigned char *zi; // 当要插入、删除、替换entry，以该字段定位插入位置
    unsigned char *value; // 
    long long longval; // 整数填充值，不同的值，可表示值是字符串或者整型
    size_t sz; // 值大小
    int offset; //表示当前entry所在的listpack偏移
} quicklistEntry;
```
上面三个数据结构就是基本的quicklist数据结构。可以看到quicklist底层也是借助于listpack。
在想quicklist添加一个元素的时候，不会像普通链表那样，直接在某个偏移位置新增节点。而是会检查插入位置的listpack能否容纳新的元素，如果没超过fill设定限制，就会直接插入，否者会创建一个新的quicklistNode。

## listpack是什么

## list insert源码解读
下面以一个list insert操作为列，窜起quicklist、packlist
```

```




参考资料：


1. 3.2版本Release notes[https://raw.githubusercontent.com/antirez/redis/3.2/00-RELEASENOTES](https://raw.githubusercontent.com/antirez/redis/3.2/00-RELEASENOTES)
