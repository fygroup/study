### gcc 内存检测编译
```
// 编译选项
-g -O1
-fsanitize=address
    开启内存越界检测
-fsanitize=leak
    开启泄露保护？
-fsanitize-recover=address
    保证稳定性，不能遇到错误就简单退出，而是继续运行。需要叠加设置ASAN_OPTIONS=halt_on_error=0
-fno-stack-protector
    去使能栈溢出保护
-fno-omit-frame-pointer
    去使能栈溢出保护
-fno-var-tracking
    默认选项为-fvar-tracking，会导致运行非常慢
-g1
    表示最小调试信息，通常debug版本用-g即-g2

// 推荐设置
-g -O1 -fsanitize=address -fsanitize=leak -fno-omit-frame-pointer 


// 运行环境配置
ASAN_OPTIONS是Address-Sanitizier的运行选项环境变量
halt_on_error=0         检测内存错误后继续运行
use_sigaltstack=0
detect_leaks=1          使能内存泄露检测
malloc_context_size=15  内存错误发生时，显示的调用栈层数为15
log_path=asan.log       内存检查问题日志存放文件路径
suppressions=$SUPP_FILE 屏蔽打印某些内存错误
new_delete_type_mismatch=0
alloc_dealloc_mismatch=0

#export ASAN_OPTIONS=$ASAN_OPTIONS:halt_on_error=0:use_sigaltstack=0:detect_leaks=1:malloc_context_size=15log_path=asan.log:suppressions=$SUPP_FILE:new_delete_type_mismatch=0:alloc_dealloc_mismatch=0
```