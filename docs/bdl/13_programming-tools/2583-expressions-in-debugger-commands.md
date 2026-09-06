---
title: "Expressions in debugger commands"
source: "fgl-topics/c_fgl_Debugger_010.html"
breadcrumb: "Programming tools > Integrated debugger > Expressions in debugger commands"
type: "concept"
---

# Expressions in debugger commands

> A limited expression syntax can be used in debugger commands.

Some debugger commands such as [`display`](2595-display.md "The display command displays the specified expression's value each time program execution stops."), [`print`](2607-print.md "The print command displays the current value of the specified expression.") and [`call`](2589-call.md "The call command calls a function in the program.") take expressions as arguments. The Genero debugger supports a reduced
syntax for command expressions described in this section. For a detailed description of comparison
operators, constant values and operands, see [Expressions](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language.").

## Syntax

```
 [[package-path.]module.]variable
| char-const
| int-const
| dec-const
| NULL
| TRUE
| FALSE
| expression IS [NOT] NULL
| expression = expression
| expression == expression
| expression <= expression
| expression => expression
| expression < expression
| expression > expression
| expression + expression
| expression - expression
| expression * expression
| expression / expression
| expression OR expression
| expression AND expression
| NOT expression
| - expression
| ( expression )
```

1. variable is a program variable name, of any type (primitive type, record,
   array, ...)
2. module is a module name, if the variable is defined in this [imported module](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.").
3. package-path is the package path, if the variable is defined in an imported
   module that belongs to this [package](../09_advanced-features/0816-package.md "Defines the package the module belongs to.").
4. char-const is a character string literal delimited by single or double
   quotes.
5. int-const is an integer literal.
6. dec-const is a decimal number literal.
7. expression is a combination of:
   - one of the above listed syntax elements,
   - an imported symbol (variable, function, ...),
   - a build-in function such as `length(expression)`,
   - a class method such as
     `base.Application.getArgument(nnn)`,
   - a type function,
   - a user-defined C extension function,
   - chaining method calls is supported: `s.subString(2,3).getLength()`

## Function calls in expressions

When calling a function in an expression, the debugger requires the exact number of parameters.
If a parameter is a `RECORD`, the function must be called with a
`RECORD` variable without the `.*` notation: It is not possible to
pass simple values as parameter to a function expecting a `RECORD`, even if the
number of simple values is same as the number of record fields.

## Example 1

Source code main1.4gl:

```
IMPORT util
TYPE t_rec RECORD pkey INT, name VARCHAR(50) END RECORD
MAIN
    VAR rec t_rec
    LET rec.pkey = 124
    LET rec.name = "aaaaaa"
    DISPLAY util.JSON.stringify(rec)
    CALL myfunc(rec)
END MAIN
FUNCTION myfunc(prec t_rec) RETURNS ()
    DISPLAY "record: ", prec.*
END FUNCTION
```

Using the debugger:

```
$ fglcomp main1.4gl
$ fglrun -d main1.42m

(fgldb) break 7
Breakpoint 1 at 0x00000000: file main1.4gl, line 7.

(fgldb) run
Breakpoint 1, main() at main1.4gl:7
   4         VAR rec t_rec
   5         LET rec.pkey = 124
   6         LET rec.name = "aaaaaa"
-> 7         DISPLAY util.JSON.stringify(rec)
   8         CALL myfunc(rec)
   9     END MAIN
   10    FUNCTION myfunc(prec t_rec) RETURNS ()

(fgldb) print base.Application.getArgument(0)
$1 = "main1.42m"

(fgldb) print rec.pkey * 100
$2 = 12400

(fgldb) print rec
$2 = {pkey = 124, name = "aaaaaa"}

(fgldb) call util.JSON.parse('"eee"',rec.name)
(fgldb) print rec
$3 = {pkey = 124, name = "eee"}

(fgldb) call myfunc(rec)
record:         124eee

(fgldb) quit
```

Example main2.4gl:

```
MAIN
    VAR arr DYNAMIC ARRAY OF STRING
    LET arr[3] = "zzzzzz"
    DISPLAY arr.getLength()
END MAIN
```

Debugging session:

```
$ fglcomp main2.4gl
$ fglrun -d main2.42m

(fgldb) break 3
Breakpoint 1 at 0x00000000: file main2.4gl, line 3.

(fgldb) run
Breakpoint 1, main() at main2.4gl:3
   1     MAIN
   2         VAR arr DYNAMIC ARRAY OF STRING
-> 3         LET arr[3] = "zzzzzz"
   4         DISPLAY arr.getLength()
   5     END MAIN
   6

(fgldb) print arr
$1 = {}

(fgldb) next
   1     MAIN
   2         VAR arr DYNAMIC ARRAY OF STRING
   3         LET arr[3] = "zzzzzz"
-> 4         DISPLAY arr.getLength()
   5     END MAIN
   6

(fgldb) print arr
$2 = {(null), (null), "zzzzzz"}

(fgldb) print arr.getLength()
$3 = 3

(fgldb) quit
```
