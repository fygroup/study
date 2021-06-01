### 资料
```
[flex bison使用] http://www.calvinneo.com/2016/07/29/flex%E5%92%8Cbison%E4%BD%BF%E7%94%A8/
flex & bison.pdf
```

### demo
```

```

### 二义性
```
词法分析器匹配输入时匹配尽可能多的字符串
如果两个模式都匹配的话，匹配在程序中更早出现的模式
```

### IO操作
```
默认读取标准输入

指定yyin文件句柄读取输入

词法分析器到yyin的结束位置时，调用yywrap()，如今默认返回1，或
%option noyywrap

yyrestart(f) 设置句柄输入



```