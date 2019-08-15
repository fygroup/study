### 4个基本数据结构
```
链表
队列
映射
红黑树
```

### 链表
```
(1) 内核API
    #include <include/linux/list.h>
    1) 结构体
        struct MY{
            .....
            struct list_head list;
        };
        struct MY* mys;
        mys = kmalloc(sizeof(struct MY), GFP_KERNEL);
        //注意：
            > 在内核空间中，没有进程堆的概念，用kmalloc()分配内存，实际上是Linux内核统一管理的，一般用slab分配器，也就是一个内存缓存池，管理所有可以kmalloc()分配的内存。在Linux内核态，kmalloc分配的所有的内存，都是可以被所有运行在Linux内核态的task访问到的。
            > 内核态的栈在task_struct->stack里面描述，其底部是thread_info对象，thread_info可以用来快速获取task_struct对象。整个stack区域一般只有一个内存页(可配置)
            > kmalloc的flag
                GFP_KERNEL
                当所请求分配的内存不够一页的时候,GFP_KERNEL让当前进程睡眠,来等待够一页内存大小的时候,才能获得正确分配到的内存
                GFP_ATOMIC
                用来从中断处理和进程上下文之外的其他代码中分配内存,从不睡眠.
                GFP_USER
                为用户空间页来分配内存,可能睡眠.
    2) 新建链表head
        LIST_HEAD(stu_head);
    3) 初始化每个结构体中的list
        INIT_LIST_HEAD(&mys->list);
    4) 将节点加入链表head
        list_add(&mys->list, &stu_head);
    5) 正向遍历
        struct MY* my_tmp;
        list_for_each_entry(my_tmp, &stu_head, list)       //这是个宏！！！
        {
            printk (KERN_ALERT "id  =%d\n", my_tmp->id);
        }
        //注意：这一步通过链表节点获得用户数据，用到的宏如下
        #define list_entry(ptr, type, member) container_of(ptr, type, member)
        #define container_of(ptr, type, member) ({               \
            const typeof(((type *)0)->member)*__mptr = (ptr);    \
            (type *)((char *)__mptr - offsetof(type, member)); })
    6) 反向遍历
        list_for_each_entry_reverse(my_tmp, &stu_head, list)
        {
        }
    7) 替换
        struct MY* my1 = kmalloc(sizeof(struct MY), );
        struct MY* my2;

        list_replace(&stu3->list, &stu2->list);
    8) 删除

(2) 实例代码
    #include<linux/init.h>
    #include<linux/slab.h>
    #include<linux/module.h>
    #include<linux/kernel.h>
    #include<linux/list.h>

    MODULE_LICENSE("Dual BSD/GPL");
    struct student
    {
        int id;
        char* name;
        struct list_head list;
    };

    void print_student(struct student*);

    static int testlist_init(void)
    {
        struct student *stu1, *stu2, *stu3, *stu4;
        struct student *stu;
        
        // init a list head
        LIST_HEAD(stu_head);

        // init four list nodes
        stu1 = kmalloc(sizeof(*stu1), GFP_KERNEL);
        stu1->id = 1;
        stu1->name = "wyb";
        INIT_LIST_HEAD(&stu1->list);

        stu2 = kmalloc(sizeof(*stu2), GFP_KERNEL);
        stu2->id = 2;
        stu2->name = "wyb2";
        INIT_LIST_HEAD(&stu2->list);

        stu3 = kmalloc(sizeof(*stu3), GFP_KERNEL);
        stu3->id = 3;
        stu3->name = "wyb3";
        INIT_LIST_HEAD(&stu3->list);

        stu4 = kmalloc(sizeof(*stu4), GFP_KERNEL);
        stu4->id = 4;
        stu4->name = "wyb4";
        INIT_LIST_HEAD(&stu4->list);

        // add the four nodes to head
        list_add (&stu1->list, &stu_head);
        list_add (&stu2->list, &stu_head);
        list_add (&stu3->list, &stu_head);
        list_add (&stu4->list, &stu_head);

        // print each student from 4 to 1
        list_for_each_entry(stu, &stu_head, list)
        {
            print_student(stu);
        }
        // print each student from 1 to 4
        list_for_each_entry_reverse(stu, &stu_head, list)
        {
            print_student(stu);
        }

        // delete a entry stu2
        list_del(&stu2->list);
        list_for_each_entry(stu, &stu_head, list)
        {
            print_student(stu);
        }

        // replace stu3 with stu2
        list_replace(&stu3->list, &stu2->list);
        list_for_each_entry(stu, &stu_head, list)
        {
            print_student(stu);
        }

        return 0;
    }

    static void testlist_exit(void)
    {
        printk(KERN_ALERT "*************************\n");
        printk(KERN_ALERT "testlist is exited!\n");
        printk(KERN_ALERT "*************************\n");
    }

    void print_student(struct student *stu)
    {
        printk (KERN_ALERT "======================\n");
        printk (KERN_ALERT "id  =%d\n", stu->id);
        printk (KERN_ALERT "name=%s\n", stu->name);
        printk (KERN_ALERT "======================\n");
    }

    module_init(testlist_init);
    module_exit(testlist_exit);

(2) makefile
    obj-m += testlist.o

    #generate the path
    CURRENT_PATH:=$(shell pwd)
    #the current kernel version number
    LINUX_KERNEL:=$(shell uname -r)
    #the absolute path
    LINUX_KERNEL_PATH:=/usr/src/kernels/$(LINUX_KERNEL)
    #complie object

    all:
        make -C $(LINUX_KERNEL_PATH) M=$(CURRENT_PATH) modules
        rm -rf modules.order Module.symvers .*.cmd *.o *.mod.c .tmp_versions *.unsigned

    clean:
        rm -rf modules.order Module.symvers .*.cmd *.o *.mod.c *.ko .tmp_versions *.unsigned

(3) 安装与卸载
    insmod testlist.ko
    rmmod testlist
    dmesg | tail -100       //查看内核输出
```

### 队列
```
// 注意
    队列的size在初始化时，始终设定为2的n次方
    使用队列之前将队列结构体中的锁(spinlock)释放

```