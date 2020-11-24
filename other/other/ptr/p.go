package main

func main() {
	var  a int = 10;
    print("%d\n", a);
    print("%p\n", &a);

    print("~~~~~~~~~~~~~~\n");
    var  b *int = &a;
    print("%p\n", b);
    print("%p\n", &b);

    print("~~~~~~~~~~~~~~\n");
    var   c = &a;
    print("%d\n", c);
    print("%p\n", &c);
}