---
title: "Calling C functions from programs"
source: "fgl-topics/c_fgl_CExtensions_016.html"
breadcrumb: "Extending the language > C-Extensions > Calling C functions from programs"
type: "concept"
---

# Calling C functions from programs

> C-Extensions functions can be called from the program in the same way that you call a BDL function.

The C functions that can be called from the BDL code must use the following signature:

```
int function-name( int )
```

Here function-name must be written in lowercase letters. The
fglcomp compiler converts all BDL function names (following a
`CALL` keyword) to lowercase.

The C function must be referenced in the `usrFunctions` array of the C interface
file. For more details, see [The C interface file](2708-the-c-interface-file.md "To make your C functions visible to the runtime system, you must define all the functions in the C interface file.").

The `(int)` parameter of the C function defines the number of parameters on the
FGL runtime stack, that can be popped.

The `(int)` return value of the C function is ignored by the runtime system: the
number of returned values is defined by the number of calls to push\* functions. However, a good
practice is to return the number of values pushed on the FGL runtime stack.

Parameters and return values must be popped from / pushed on the FGL runtime stack, by using the
[FGL stack functions](2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack.").

The order of calls to pop parameter / push values is important:

- Parameters passed to the C function must be popped in the reverse order of the BDL call list:
  `CALL c_fct( A, B, C ) => pop C, B, A`.
- Values returned from the C function must be pushed in the same order as in the BDL
  `RETURNING` clause: `push A, B, C => CALL c_fct() RETURNING A, B,
  C`.

In this code example, the C-Extension module mycext.c defines the
`c_fct()` function:

```
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "f2c/fglExt.h"

int c_fct( int n );

UsrFunction usrFunctions[]={
  {"c_fct",c_fct,2,2},
  {0,0,0,0}
};

int c_fct( int n )
{
   int rc;
   float price;
   char name[31];
   if (n != 2) exit(1);
   popflo(&price);
   popvchar(name, sizeof(name));
   printf(">> [%s] price:%f\n", name, price);
   pushint( strlen(name) );
   price = price * 2;
   pushflo( &price );
   return 2;
}
```

The C-Extension library is imported by the BDL module with [`IMPORT`](2710-loading-c-extensions-at-runtime.md "The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files."):

```
IMPORT mycext

MAIN
    DEFINE len INT, price2 FLOAT
    CALL c_fct("Hand gloves", 120.50)
         RETURNING len, price2
    DISPLAY "len = ", len
    DISPLAY "price2 = ", price2
END MAIN
```

Compilation and execution example on a Linux®
system:

```
$ fglmkext -o mycext.so mycext.c

$ fglcomp myprog.4gl

$ fglrun myprog.42m
>> [Hand gloves] price:120.500000
len =          11
price2 =                   241.0
```

Variable argument list and variable return list is allowed when calling a C function, if you
specify -1 as number of arguments/returns in the function entry of the [C interface file](2708-the-c-interface-file.md "To make your C functions visible to the runtime system, you must define all the functions in the C interface file.").

C interface file "funclist.c":

```
#include "f2c/fglExt.h"
int myfunc(int);
UsrFunction usrFunctions[]={
  { "myfunc", myfunc, -1, -1, 0 },
  { 0,0,0,0,0 }
};
```

C extension source file "myext.c":

```
#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "f2c/fglExt.h"

int myfunc( int in_num );
int myfunc( int in_num )
{
       char buf[10][256];
       int x;
       for (x = in_num - 1; x >= 0; x--) {
           popvchar(buf[x], 256);
           fprintf(stderr,">> pop param  %d: %s\n", x, buf[x]);
       }
       for (x = 0; x < in_num; x++) {
           pushvchar(buf[x], strlen(buf[x]));
           fprintf(stderr,">> push param %d: %s\n", x, buf[x]);
       }
       return in_num;
}
```

The main program file "main.4gl":

```
IMPORT util
IMPORT myext
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DISPLAY "Calling myfunc with 1 arg, returning 1 value:"
    CALL myfunc("aaa") RETURNING arr[1]
    DISPLAY util.JSON.stringify(arr)
    DISPLAY "Calling myfunc with 2 args, returning 2 values:"
    CALL myfunc("aaa","bbb") RETURNING arr[1], arr[2]
    DISPLAY util.JSON.stringify(arr)
END MAIN
```

Compilation and execution (Linux/gcc):

```
$ fglmkext -o myext.so myext.c funclist.c
gcc -shared ...
$ fglcomp main.4gl && fglrun main.42m
Calling myfunc with 1 arg, returning 1 value:
>> pop param  0: aaa
>> push param 0: aaa
["aaa"]
Calling myfunc with 2 args, returning 2 values:
>> pop param  1: bbb
>> pop param  0: aaa
>> push param 0: aaa
>> push param 1: bbb
["aaa","bbb"]
```

If you do not define your C function as accepting variable arguments/returns, the compiler will
produce the error [-4333](../15_library-reference/4483-genero-bdl-errors.md), when
the same function is called with different arguments/returns.

## Related links

**Related concepts**  

[Functions](../08_language-basics/0761-functions.md "Describes user defined functions.")
