### flex输入
```
flex 提供灵活的三层输入

YY_BUFFER_STATE 结构用来处理输入，包含字符串的缓冲区以及一些变量与标记

// 默认flex行为
YY_BUFFER_STATE bp;
extern FILE* yyin;
if (!yyin) yyin = stdin;    // flex的默认输入
bp = yy_create_buffer(yyin, YY_BUF_SIZE); // 创建一个读取yyin的缓冲区, YY_BUF_SIZE 默认16k
yy_switch_to_buffer(bp);    // 告诉它使用已创建的缓冲区
yylex();    // 词法分析器调用

yyrestart(f);  // 是词法分析器读取输入文件

```