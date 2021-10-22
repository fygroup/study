### 移进/归约
```
bison 采用自底向上 (bottom-up) 的分析方法。它用到一个分析栈 (parser stack)，关键有两个动作：
移进 (shift)
    读取的 token 移进到分析栈中。
归约 (reduce)
    当分析栈顶的 n 个符号匹配某规则的右端时，用该规则的左端取代之

1）移进：从句子左端将一个终结符移到栈顶
2）归约：根据规则，将栈顶的若干个字符替换为一个符号
3）接受：句子中所有词语都已移进栈中，且栈中只剩下一个符号  ，分析成功，结束
4）拒绝：句子中所有词语都已移进栈中，栈中并非只有一个符号  ，也无法进行任何归约操作，分析失败，结束
```

### 一些语法
```
%{
#include<stdio.h>
#include<stdlib.h>
%}

%union{
    struct ast* a;
    double b;
}

/* 声明 */
%token<b> NUMBER

%type<a> exp factor term

/* 规则 */
%%

exp: factor
    | exp '+' factor { $$ = newast('+', $1, $3); }
    | exp '-' factor { $$ = newast('-', $1, $3); }
    ;

factor: term
    | factor * term { $$ = newast('*', $1, $3); }
    | factor / term { $$ = newast('/', $1, $3); }
    ;

term: NUMBER { $$ = newnum($1); }
    | '|' term { $$ = newast('|', $2, NULL); }
    | '(' exp ')' { $$ = $2; }
    | '-' term { $$ = newast('M', $2, NULL); }
    ;
%%
```

### 优先级
```
%left %right %noassoc出现的顺序决定由低到高优先级
左结合 右结合 无结合性

%left '+' '-'
%left '*' '/'
%noassoc '|' UMINUS

/*
* %nonassoc的含义是没有结合性。它一般与%prec结合使用表示该操作有同样的优先级
* expr: '-' expr %prec UMINUS { $$ = node(UMINUS, 1, $2); }
* 表示该操作的优先级与UMINUS相同，在上面的定义中，UMINUS的优先级高于其他操作符，所以该操作的优先级也高于其他操作符计算
*/

%type<a> exp

%%

/* 规则可以糅杂在非终结符中，bison通过上面定义的优先级进行分析 */
exp: exp '+' exp { $$ = newast('+', $1, $3); }
    | exp '-' exp { $$ = newast('-', $1, $3); }
    | exp '*' exp { $$ = newast('*', $1, $3); }
    | exp '/' exp { $$ = newast('/', $1, $3); }
    | '|' exp { $$ = newast('|', $1, NULL); }
    | '(' exp ')' { $$ = $2; }
    | '-' exp %pec UMINUS { $$ = newast('M', $2, NULL); }
    | NUMBER { $$ = newnum($1); }
    ;
%%

注意：谨慎使用bison的优先级规则，尽量用语法来规定优先级
```
