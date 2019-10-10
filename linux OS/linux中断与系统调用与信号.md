
### 与内核通信
```
c API ——> 内核 ——> 硬件

陷入内核
linux不会直接调用内核函数，而采用软中断，告诉内核需要执行系统调用，来切换到内核态

系统调用号是通过%eax传入内核的
比如调用fork
#define __NR_fork      2  
mov   eax, 2        
int   0x80     软中断

Linux最多允许向系统调用传递6个参数，分别依次由%ebx，%ecx，%edx，%esi，%edi这5个寄存器完成

系统调用比较费时，可以看看系统调用函数表。尽量避免多次调用里面的函数
```

### 系统调用的实现
(1) 参数验证
指针指向的区域必须是用户空间，不能哄骗内核去读其内核空间
不能哄骗内核去读其他进程的空间
内存也分为可读、可写、可执行
```
//内核提供了两个方法
copy_to_user() //向用户空间写数据
//第一个参数时进程内存参数，第二个是内核空间源地址，最后一个是需要拷贝的长度

copy_from_user()  //用户向内核拷贝数据

//失败返回没能拷贝完的字节数
//注意上面两个函数也会发生休眠，当用户数据的页被换出到硬盘。。。
```

(2) 绑定一个系统调用
```
//1、在sys_call_table中注册你的函数
//在entry.s文件中
ENTRY(sys_call_table)
    .long sys_restart_call()  /*0*/
    .long sys_exit()
    ...
    .long sys_read()  
    ...
    .long sys_foo()   //在末尾加入你的函数

//2、把系统调用号加入<asm/unistd.h>中
#define __NR_restart_call  0
#define __NR_exit          1
...
#define __NR_read          3
...
#define __NR_foo           338

//3、实现你的函数
#include <asm/page.h>
asmlinkage long sys_foo(void){
    .....
}

```

(3) 访问我们的系统函数
```
//比如open的系统调用定义
long open(const char* file, int flag, int mode)
//不依靠库的支持
#define __NR_open 5
__syscall3(long, open, const char*, filename, int, flag, int, mode)

```
