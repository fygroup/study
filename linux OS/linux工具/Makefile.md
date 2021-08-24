### 通配符
```
.PHONY
这个目标的所有依赖被作为伪目标。伪目标是这样一个目标：当使用 make 命令行指定此目标时，这个目标所在的规则定义的命令、无论目标文件是否存在都会被无条件执行。

OBJ:=$(wildcard *c)
test: $(OBJ)
    gcc -o $@ $^
%.o:%.c
    gcc -o $@ $^
$@ : 代表目标 test 
$^ : 代表所有 *
%  : "%.o" 把我们需要的所有的 ".o" 文件组合成为一个列表，从列表中挨个取出的每一个文件，"%" 表示取出来文件的文件名（不包含后缀），然后找到文件中和 "%"名称相同的 ".c" 文件，然后执行下面的命令，直到列表中的文件全部被取出来为止
```

### 字符串处理函数
```
// $(patsubst <pattern>,<replacement>,<text>)
查找 text 中的单词是否符合模式 pattern，如果匹配的话，则用 replacement 替换。返回值为替换后的新字符串
OBJ=$(patsubst %.c,%.o,1.c 2.c 3.c)
all:
    @echo $(OBJ)
1.o 2.o 3.o

// $(subst <from>,<to>,<text>)
函数的功能是把字符串中的 form 替换成 to，返回值为替换后的新字符串
OBJ=$(subst ee,EE,feet on the street)
fEEt on the strEEt

// $(strip <string>)
去除空格

// $(findstring <find>,<in>)
查找 in 中的 find ,如果我们查找的目标字符串存在。返回值为目标字符串，如果不存在就返回空
OBJ=$(findstring a,a b c)
a

// $(filter <pattern>,<text>)
过滤出 text 中符合模式 pattern 的字符串，可以有多个 pattern。返回值为过滤后的字符串
OBJ=$(filter %.c %.o,1.c 2.o 3.s)
1.c 2.o

// $(filter-out <pattern>,<text>)
和filter函数正好相反，但是用法相同。去除符合模式 pattern的字符串，保留符合的字符串。返回值是保留的字符串
OBJ=$(filter-out 1.c 2.o ,1.o 2.c 3.s)
3.s

// $(sort <list>)
将 <list>中的单词排序（升序） 返回值为排列后的字符串(sort会去除重复的字符串)
OBJ=$(sort foo bar foo lost)
bar foo lost

// $(word <n>,<text>)
取出函数<text>中的第n个单词。返回值为我们取出的第 n 个单词
OBJ=$(word 2,1.c 2.c 3.c)
2.c
```

### 文件处理函数
```
// $(dir <names>)
从文件名序列 names 中取出目录部分，如果没有 names 中没有 "/" ，取出的值为 "./"
OBJ=$(dir src/foo.c hacks)
all:
    @echo $(OBJ)
src/ ./

// $(notdir <names>)
从文件名序列 names 中取出非目录的部分。非目录的部分是最后一个反斜杠之后的部分。返回值为文件非目录的部分
OBJ=$(notdir src/foo.c hacks)
foo.c hacks

// $(suffix <names>)
从文件名序列中 names 中取出各个文件的后缀名。返回值为文件名序列 names 中的后缀序列，如果文件没有后缀名，则返回空字符串
OBJ=$(suffix src/foo.c hacks)
.c

// $(basename <names>)
从文件名序列 names 中取出各个文件名的前缀部分。返回值为被取出来的文件的前缀名，如果文件没有前缀名则返回空的字符串
OBJ=$(notdir src/foo.c hacks)
src/foo hacks

// $(addsuffix <suffix>,<names>)
把后缀 suffix 加到 names 中的每个单词后面
OBJ=$(addsuffix .c,src/foo.c hacks)
sec/foo.c.c hack.c

// $(addperfix <prefix>,<names>)
把前缀 prefix 加到 names 中的每个单词的前面
OBJ=$(addprefix src/, foo.c hacks)
src/foo.c src/hacks

// $(join <list1>,<list2>)
把 list2 中的单词对应的拼接到 list1 的后面
OBJ=$(join src car,abc zxc qwe)
srcabc carzxc qwe
<list1>中的文件名比<list2>的少，所以多出来的保持不变

// $(wildcard PATTERN)
列出当前目录下所有符合模式的 PATTERN 格式的文件名
OBJ=$(wildcard *.c  *.h)
得到当前函数下所有的".c "和".h"结尾的文件
```

### 其他函数
```
// $(foreach <var>,<list>,<text>)
把参数<list>中的单词逐一取出放到参数<var>所指定的变量中，然后再执行<text>所包含的表达式
name:=a b c d
files:=$(foreach n,$(names),$(n).o)
a.o b.o c.o d.o

// $(if <condition>,<then-part>)或(if<condition>,<then-part>,<else-part>)
OBJ:=foo.c
OBJ:=$(if $(OBJ),$(OBJ),main.c)
执行 make 命令我们可以得到函数的值是 foo.c，如果变量 OBJ 的值为空的话，我们得到的 OBJ 的值就是main.c

// $(call <expression>,<parm1>,<parm2>,<parm3>,...)
call 函数是唯一一个可以用来创建新的参数化的函数
reverse = $(1) $(2)
foo = $(call reverse,a,b)
a b

reverse = $(2) $(1)
foo = $(call reverse,a,b)
b a

// $(origin <variable>)
它并不操作变量的值，它只是告诉你这个变量是哪里来的
“undefined”：如果<variable>从来没有定义过，函数将返回这个值。
“default”：如果<variable>是一个默认的定义，比如说“CC”这个变量。
“environment”：如果<variable>是一个环境变量并且当Makefile被执行的时候，“-e”参数没有被打开。
“file”：如果<variable>这个变量被定义在Makefile中，将会返回这个值。
“command line”：如果<variable>这个变量是被命令执行的，将会被返回。
“override”：如果<variable>是被override指示符重新定义的。
“automatic”：如果<variable>是一个命令运行中的自动化变量。

```

### demo
```
SRC_PATH = /home/malx/project/dev_aarch64_wukong_link_mqtt/wukong/
DIRS = $(shell find $(SRC_PATH) -maxdepth 5 -type d)

SRCS_CPP += $(foreach dir, $(DIRS), $(wildcard $(dir)/*.cpp))
SRCS_CC += $(foreach dir, $(DIRS), $(wildcard $(dir)/*.cc))
SRCS_C += $(foreach dir, $(DIRS), $(wildcard $(dir)/*.c))

all:
	@echo " ${DIRS} "
	@echo " ${SRCS_CPP} "
```


### 递归获得目录
```
DIR := ./
SRCDIR := $(shell find $(DIR) -maxdepth 5 -type d)
```

### 固定动态库路径
```
-L LINK_THIRD_LIBDIR -Wl,-rpath,$(LINK_THIRD_LIBDIR) -lwukong_link -lsqlite3 -lcurl -lpthread -lssl -lcrypto 
```

### 指定目录
```
make -C xxxx
```