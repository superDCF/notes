# redis替换ziplist

Replace ziplist with listpack in quicklist

导语：最近看到redis的官方消息，用listpack替换ziplist，迭代了多个版本的listpack终于登场了（当前最新6.2.6版本，仍未替换，只在unstable分支出现，预计很快发布）。恐怕有同学还不了解listpack有什么提升的地方，本篇将结合图片+源码的方式，走进quicklist、listpack。

## 特别声明：
以下阅读，都是基于unstable分支，commit:https://github.com/redis/redis/commit/4512905961b3a2f4c00e5fe7ffff8d96db82861e
该分支还处于迭代开发中，但是**Replace ziplist with listpack in quicklist**中的quicklist、listpack等替换ziplist的替换已经完成。

## 为什么要替换ziplist
ziplist在3.2版本之前，是list、hash、zset三种数据类型的底层使用的数据结构之一。ziplist最大特点就是，是一种内存紧凑型的数据结构，占用一块连续的内存，针对不同长度类型的数据，进行不同的编码。既可以利用CPU缓存，也非常节省内存。
当然这种连续性的内存结构，在元素大小变化、删除、增加，也有着通用的毛病。

**ziplist结构布局如下**

![ziplist结构布局](https://s4.ax1x.com/2021/12/25/TagGTJ.md.png)
* uint32_t zlbytes; // 整个压缩列表占用字节数，包括zlbytes字段本身4个字节
* uint32_t zltail; // 最后一个entry元素距离压缩列表起始位置的偏移量，用于快速定位到最后一个节点而不需要遍历整个链表
* uint16_t zllen; // entry元素个数
* void* entries; // 元素内容列表，挨个紧凑存储
* uint8_t zlend; // 代表ziplist的结束，固定值等于255

**entry 结构布局**

![TaWRpR.md.png](https://s4.ax1x.com/2021/12/25/TaWRpR.md.png)
* prevlen：前一个entry的长度，如果这个长度小于254字节，则只占用1字节。当长度大于或等于254时，它将占用5个字节。第一个字节被设置为254（FE），以表示后面有一个更大的值。其余的4个字节以前一个entry的长度为值，通过这个值，可以进行指针计算，从而跳转到上一个节点。
* encoding：该字段表示entry的数据类型以及数据长度，当是整形时，只占用1字节，字节的前2 bit都是1，例如【11xxxxxx】，后几个bit位表示不同的entry数据类型。当是字符串时，根据字符串的长度不同，可能占用1、2、5字节，同样多余的bit位，表示不同的数据长度。（多说一句，redis有很多这种可变数据编码的思想，在一定程度上节省内存）

**由于prevlen字段表示前一个元素的长度，当新增元素、元素变大超过现有值类型，例如增加一个大于254字节的元素，而之前存储的是小于254字节的entry，这时就会改变当前这个prevlen的值，可能出现更严重的是，会顺序性的对后面的entry的prevlen都需要修改，这会直接影响到压缩列表的操作性能，不过这种出现的概率很低。**

由于压缩列表的搜索是O(n)，增加、删除、修改都会移动内存，所以压缩列表只适用于保存节点不多的情况，在3.2之前使用ziplist的数据类型有，list、hash、zset，他们分别是通过list_max_ziplist_value、hash_max_ziplist_entries、zset_max_ziplist_value设置当数据量小于多少时使用ziplist数据结构，默认值都是64。

了解了ziplist的设计和使用的取舍，在redis 3.2版本使用quicklist替换redis List类型的底层数据结构ziplist[1]。
当前未发布版本，把hash、zset底层的ziplist实现，也替换成了listpack。

## quicklist是什么

当前未发布的最新版本，阅读源码可知，quicklist其实就是由双向链表+listpack构成。新数据的插入或者变大，把之前的单个ziplist替换成一个双向链表上的多个listtpack，来减少影响。
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
    unsigned int encoding : 2;   /* RAW==1 or LZF==2 */ // 表示listpack是否压缩了，1-RAW表示原生没被压缩，2-LZF压缩算法压缩
    unsigned int container : 2;  /* PLAIN==1 or PACKED==2 */，当前源码默认是2，表示使用PACKED作为数据容器（其实就是listpack）
    unsigned int recompress : 1; /* was this node previous compressed? */ // 表示节点之前是否被压缩
    unsigned int attempted_compress : 1; /* node can't compress; too small */ // 测试使用，不用管
    unsigned int extra : 10; /* more bits to steal for future usage */ // 保留字段
} quicklistNode;

typedef struct quicklistEntry {
    const quicklist *quicklist; // 属于哪一个quicklist
    quicklistNode *node; // 属于哪一个listpack
    unsigned char *zi; // 当要插入、删除、替换entry，以该字段为基准
    unsigned char *value; // 如果是字符串，就是值本身，否则，值存在longval字段
    long long longval; // 整数填充值
    size_t sz; // 值大小
    int offset; //表示当前entry所在的listpack偏移
} quicklistEntry;
```
上面三个数据结构就是基本的quicklist数据结构组成，可以看到quicklist底层也是借助于listpack。

![quicklist结构图](https://s4.ax1x.com/2021/12/24/TYlp5V.png)

在向quicklist添加一个元素的时候，不会像普通链表那样，直接在某个偏移位置新增节点。而是会检查插入位置的listpack能否容纳新的元素，如果没超过fill设定限制，就会直接插入，否者会创建一个新的quicklistNode。
 
## listpack是什么
listpack继承了ziplist内存紧凑的设计，重新设计了标头和单个元素数据结构。结构示意如下：

![Ta5x2j.md.png](https://s4.ax1x.com/2021/12/25/Ta5x2j.md.png)

* tot-bytes: uint32_t，表示包含listpack总字节数，total bytes.
* num-elements: uint16_t，表示listpack包含的元素数量
* listpack-end-byte：固定值，和ziplist一样都是255。终止符的主要优点是能够在不保存（并在每次迭代时比较）listpack末尾地址的情况下扫描listpack，并且可以轻松识别listpack是否格式正确或被截断。

元素示意结构：

![TaIBo8.md.png](https://s4.ax1x.com/2021/12/25/TaIBo8.md.png)

* encoding-type:编码类型，前面说过，redis有很多使用不同的编码记录不同类型数据。具体的编码规则，见参考资料2。
* element-data：元素数据本身，如整数，或者字符串的字节数组。
* element-tot-len：元素总长度，便于从后往前遍历。

**对比ziplist发现，在元素结构，listpack少了`prevlen`，之前说过正是由于该字段的存在，可能会有连锁迁移内存的风险，而在listpack，使用`element-tot-len`表达自身长度，所以新增或者是改变元素，对下一位元素的大小是没有影响的。**

## list insert源码解读
下面以一个list insert操作为列，串起quicklist、packlist。由于源码很长，本篇只摘重要代码解读。
```
// t_list.c
// List的push操作，最终会到listTypeInsert，中间有很多不必要的代码，省略。
// entry是listpack中的，要插入value的基准位置，where表示是在entry之前还是之后插入
void listTypeInsert(listTypeEntry *entry, robj *value, int where) {
    if (entry->li->encoding == OBJ_ENCODING_QUICKLIST) { // 这里显示，底层数据结构唯一使用quicklist，注意，listpack是quicklist的底层实现的一部分。
        value = getDecodedObject(value); // 获取解码的对象，以符合redisObject。
        sds str = value->ptr;
        size_t len = sdslen(str);
        if (where == LIST_TAIL) {
            quicklistInsertAfter(entry->li->iter, &entry->entry, str, len);
        } else if (where == LIST_HEAD) {
            quicklistInsertBefore(entry->li->iter, &entry->entry, str, len);
        }
        decrRefCount(value); // 解引用，释放内存
    } else {
        serverPanic("Unknown list encoding");
    }
}

//quicklist.c
// iter表示是哪个quicklist，
/*
quicklistIter，如命名一样，该对象是在遍历时使用
typedef struct quicklistIter {
    quicklist *quicklist; //标示所属quicklist
    quicklistNode *current; // 当前在哪个节点
    unsigned char *zi; //操作的基准entry
    long offset; /* offset in current listpack */
    int direction; // 什么操作
} quicklistIter;
*/
// iter参数如上结构，entry是基准方位，用来定位值插入after的位置
// value，即要插入的值
// sz表示插入值的大小
// after表示在前还是在后插入
REDIS_STATIC void _quicklistInsert(quicklistIter *iter, quicklistEntry *entry,
                                   void *value, const size_t sz, int after)
{
    quicklist *quicklist = iter->quicklist;
    int full = 0, at_tail = 0, at_head = 0, avail_next = 0, avail_prev = 0;
    int fill = quicklist->fill;
    quicklistNode *node = entry->node;
    quicklistNode *new_node = NULL;

    if (!node) {
        /* we have no reference node, so let's create only node in the list */
        D("No node given!");
        if (unlikely(isLargeElement(sz))) {
            __quicklistInsertPlainNode(quicklist, quicklist->tail, value, sz, after);
            return;
        }
        new_node = quicklistCreateNode(); // 如果entry所在quicklistNode不存在，就新建节点，把自己当作新quicklistNode第一个节点
        new_node->entry = lpPrepend(lpNew(0), value, sz); // 把value放入listpack
        __quicklistInsertNode(quicklist, NULL, new_node, after); // 把new_node加入quicklist
        new_node->count++;
        quicklist->count++;
        return;
    }
    // 如果存在quicklistNode
    /* Populate accounting flags for easier boolean checks later */
    if (!_quicklistNodeAllowInsert(node, fill, sz)) { // 当前quicklistNode已经满了，不允许插入，fill控制再能加入listpack的大小
        D("Current node is full with count %d with requested fill %d",
          node->count, fill);
        full = 1; // 标记当前quicklistNode节点已经满了
    }

    // 判断后一个quicklistNode节点是否能够插入
    if (after && (entry->offset == node->count - 1 || entry->offset == -1)) {
        D("At Tail of current listpack");
        at_tail = 1;
        if (_quicklistNodeAllowInsert(node->next, fill, sz)) {
            D("Next node is available.");
            avail_next = 1;
        }
    }

    /* Now determine where and how to insert the new element */
    if (!full && after) { // 没有满，在入参之后插入
        D("Not full, inserting after current position.");
        quicklistDecompressNodeForUse(node); // 解压quicklistNode，以便插入
        node->entry = lpInsertString(node->entry, value, sz, entry->zi, LP_AFTER, NULL); // 插入新的value
        node->count++;
        quicklistNodeUpdateSz(node);
        quicklistRecompressOnly(node);
    } else if (full && at_tail && avail_next && after) { // 满了，在当前listpack的最后，下一个quicklistNode可以插入，在after之后插入
        /* If we are: at tail, next has free space, and inserting after:
         *   - insert entry at head of next node. */
        D("Full and tail, but next isn't full; inserting next node head");
        new_node = node->next;
        quicklistDecompressNodeForUse(new_node);
        new_node->entry = lpPrepend(new_node->entry, value, sz);
        new_node->count++;
        quicklistNodeUpdateSz(new_node);
        quicklistRecompressOnly(new_node);
    }

    quicklist->count++;

    /* In any case, we reset iterator to forbid use of iterator after insert.
     * Notice: iter->current has been compressed in _quicklistInsert(). */
    resetIterator(iter); 
}
```
借由List insert操作，大致了解如何插入quicklistNode中listpack，这就与本篇标题*Replace ziplist with listpack in quicklist*完全对上了

其实redis的源码写的非常规范，注释也足够清楚，大家如果有兴趣，最好自己阅读一遍自己感兴趣的源码。




参考资料：

1. [3.2版本Release notes](https://raw.githubusercontent.com/antirez/redis/3.2/00-RELEASENOTES)
2. [listpack设计](https://github.com/antirez/listpack/blob/master/listpack.md)
