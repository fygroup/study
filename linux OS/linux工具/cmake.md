### 推荐
```
https://github.com/ttroy50/cmake-examples

```

### ADD_xxx
```
ADD_EXECUTABLE
    生成可执行程序

ADD_SUBDIRECTORY
    添加子目录
    
ADD_LIBRARY
    创建库，动态和静态库

ADD_DEFINITIONS
    ADD_DEFINITIONS(-DENABLE_DEBUG -DABC -DHASH)
    向 C/C++编译器添加-D 定义

ADD_DEPENDENCIES
    ADD_DEPENDENCIES(target-name depend-target1 depend-target2)
    定义 target 依赖的其他 target，确保在编译本 target 之前，其他的 target 已经被构建

target_include_directories(wkLinkTest PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
target_link_libraries()

include_directories(${CMAKE_FIND_ROOT_PATH}/include)
link_directories(${CMAKE_FIND_ROOT_PATH}/lib)
link_libraries
```

### set target
```
SET
SET_TARGET_PROPERT
```

### get target
```
GET_TARGET_PROPERTY
```

### 目录操作
```
aux_source_directory(<dir> <variable>)
    收集指定目录中所有源文件的名称，并将列表存储在提供的<variable>变量中
```

### function
```
function(retrieve_files out_files)
    set(source_list)
    foreach(dirname ${ARGN})
        file(GLOB_RECURSE files RELATIVE ${CMAKE_CURRENT_SOURCE_DIR}
                "${dirname}/*.cpp"
                "${dirname}/*.c"
                "${dirname}/*.h"
                "${dirname}/*.hpp"
        )
        foreach(filename ${files})
            list(APPEND source_list "${CMAKE_CURRENT_SOURCE_DIR}/${filename}")
        endforeach()
    endforeach()
    set(${out_files} ${source_list} PARENT_SCOPE)
endfunction()

retrieve_files(WUKONG_BASE_SRCS wukong/base)
```

### file
```
(1) Reading
    file(READ <filename> <out-var> [...])
        读取文件名为 <filename> 的文件并将其内容存储到 <variable> 变量中
        可选的参数： <offset> 指定起始读取位置，<max-in> 最多读取字节数，HEX 将数据转为十六进制（处理二进制数据十分有用）
    
    file(STRINGS <filename> <out-var> [...])
    file(<HASH> <filename> <out-var>)
    file(TIMESTAMP <filename> <out-var> [...])
    file(GET_RUNTIME_DEPENDENCIES [...])

(2) Writing
    file({WRITE | APPEND} <filename> <content>...)
        写入 <content> 到 <filename> 文件中
        如果文件不存在则创建，任何在 <filename> 文件路径中的不存在文件夹都将被创建
        如果文件已存在，WRITE 模式将覆盖内容
        如果为 APPEND 模式将追加内容

    file({TOUCH | TOUCH_NOCREATE} [<file>...])
    file(GENERATE OUTPUT <output-file> [...])
    file(CONFIGURE OUTPUT <output-file> CONTENT <content> [...])

(3) Filesystem
    file({GLOB | GLOB_RECURSE} <out-var> [...] [<globbing-expr>...])
        产生一个匹配 <globbing-expr> 的文件列表并将它存储到变量 <variable> 中
        文件名替代表达式和正则表达式相似，但更简单
        如果 RELATIVE 标志位被设定，将返回指定路径的相对路径
        结果将按字典顺序排序

  file(RENAME <oldname> <newname>)
  file({REMOVE | REMOVE_RECURSE } [<files>...])
  file(MAKE_DIRECTORY [<dir>...])
  file({COPY | INSTALL} <file>... DESTINATION <dir> [...])
  file(SIZE <filename> <out-var>)
  file(READ_SYMLINK <linkname> <out-var>)
  file(CREATE_LINK <original> <linkname> [...])
  file(CHMOD <files>... <directories>... PERMISSIONS <permissions>... [...])
  file(CHMOD_RECURSE <files>... <directories>... PERMISSIONS <permissions>... [...])

Path Conversion
  file(REAL_PATH <path> <out-var> [BASE_DIRECTORY <dir>])
  file(RELATIVE_PATH <out-var> <directory> <file>)
  file({TO_CMAKE_PATH | TO_NATIVE_PATH} <path> <out-var>)

Transfer
  file(DOWNLOAD <url> [<file>] [...])
  file(UPLOAD <file> <url> [...])

Locking
  file(LOCK <path> [...])

Archiving
  file(ARCHIVE_CREATE OUTPUT <archive> PATHS <paths>... [...])
  file(ARCHIVE_EXTRACT INPUT <archive> [...])
```

### install
```
为生成的target配置安装目录
cmake .
make install


(1) 目标文件的安装
    install(TARGETS targets... 
            [EXPORT <export-name>]
            [ARCHIVE|LIBRARY|RUNTIME] [DESTINATION <dir>]
    )
    // ADD_EXECUTABLE或者ADD_LIBRARY定义的target
    // 
    // 表示不同类型的target(静态库、动态库、可执行文件)，后面是相对安装路径，默认CMAKE_INSTALL_PREFIX目录下。如果以'/'开头指绝对路径

    > 示例
    INSTALL(TARGETS myrun RUNTIME DESTINATION bin)
    INSTALL(TARGETS myrun mylib mystaticlib
            RUNTIME DESTINATION bin             // 可执行文件的安装路径
            LIBRARY DESTINATION lib             // 动态库的安装路径
            ARCHIVE DESTINATION libstatic       // 静态库的安装路径
    )

(2) 普通文件的安装
    INSTALL(FILES files... DESTINATION <dir>)
    // 用于安装一般文件，并可以指定访问权限，文件名是此指令所在路径下的相对路径

(3) 非目标文件的可执行程序安装(比如脚本之类)
    INSTALL(PROGRAMS files... DESTINATION <dir>)
    // 跟上面的FILES指令使用方法一样，唯一的不同是安装后权限为755

(4) 目录的安装
    INSTALL(DIRECTORY dirs... DESTINATION <dir> FILES_MATCHING PATTERN '*h*')
    // abc和abc/的区别是abc目录下的所有内容会被安装在目标目录下

    > 示例
    INSTALL(DIRECTORY icons scripts/ DESTINATION share/myproj
            FILES_MATCHING PATTERN "CVS" EXCLUDE
            FILES_MATCHING PATTERN "scripts/*" PERMISSIONS OWNER_EXECUTE OWNER_WRITE OWNER_READ GROUP_EXECUTE GROUP_READ
    )
    // 第二行：不包含目录名为CVS的目录
    // 第三行：对于scripts/*文件指定权限为...

```


### 添加子目录
```
// 子目录必须有CMakeLists.txt
ADD_SUBDIRECTORY(src)
```

### cmake 的build路径
```
mkdir build
cd build
cmake ..
```

### gdb 内存检测编译
```
// 编译选项
-g -fsanitize=address -fno-omit-frame-pointer -fsanitize=leak 

// 运行环境配置
export ASAN_OPTIONS=$ASAN_OPTIONS:log_path=./asan.log:suppressions=$SUPP_FILE:new_delete_type_mismatch=0:alloc_dealloc_mismatch=0


```

### file GLOB(收集文件)
```
file(GLOB <variable>
     [LIST_DIRECTORIES true|false] [RELATIVE <path>]
     [<globbing-expressions>...])
file(GLOB_RECURSE <variable> [FOLLOW_SYMLINKS]
     [LIST_DIRECTORIES true|false] [RELATIVE <path>]
     [<globbing-expressions>...])

file(GLOB files  *)
    挑选出当前文件下的所有文件

file(GLOB files  LIST_DIRECTORIES false *)
    LIST_DIRECTORIES设置为false，默认true
    只列出文件

set(CUR ${CMAKE_CURRENT_SOURCE_DIR})
file(GLOB files  LIST_DIRECTORIES false RELATIVE ${CUR}/.. *)
    如我们不需要绝对路径，只需要相对某个文件夹的路径，则可以通过设置RELATIVE的值来设置
```

### find
```
CMAKE_PREFIX_PATH
```
