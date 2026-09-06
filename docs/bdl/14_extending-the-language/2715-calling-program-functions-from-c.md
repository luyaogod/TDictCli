---
title: "Calling program functions from C"
source: "fgl-topics/c_fgl_CExtensions_017.html"
breadcrumb: "Extending the language > C-Extensions > Calling program functions from C"
type: "concept"
---

# Calling program functions from C

> It is possible to call a BDL function from a C-Extension function.

To call a [BDL function](../08_language-basics/0761-functions.md "Describes user defined functions.") from a C-Extension function,
use the `fgl_call`
macro:

```
fgl_call ( function-name, nb-params );
```

In this call, function-name is the name of the program function to call, and
nb-params is the number of parameters pushed on the stack for the program
function. The function-name must be written in lowercase letters; the [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler converts all program function names to
lowercase.

The `fgl_call()` macro is converted to a function that returns the number of
values returned on the stack.

Parameters and return values must be pushed/popped on the runtime stack, by using the [stack functions](2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack."). Parameters passed to the BDL function must
be pushed in the same order as the BDL parameter list: `push A, B, C => FUNCTION fct( A, B, C
)`. However, values returned from the BDL function must be popped in the reverse order of
the BDL return clause: `RETURN A, B, C => pop C, B, A`.

The myprog.4gl BDL module defining the `MAIN` block and the
`display_item()` function to be called from the C extension:

```
IMPORT mycext

MAIN
    CALL c_fct()
END MAIN

FUNCTION display_item(name, size)
   DEFINE name VARCHAR(30), size INTEGER
   DISPLAY name, size
   RETURN length(name), (size / 100)
END FUNCTION
```

The mycext.c C extension module calling the BDL function:

```
#include <stdlib.h>
#include <stdio.h>

#include "f2c/fglExt.h"

int c_fct( int n );

UsrFunction usrFunctions[]={
  {"c_fct",c_fct,0,0},
  {0,0,0,0}
};

int c_fct( int n )
{
   int rc, len;
   float size2;
   if (n != 0) exit(1);
   pushquote("Hand gloves", 11);
   pushint(54);
   rc = fgl_call( display_item, 2 );
   if (rc != 2) exit(1);
   popflo(&size2);
   popint(&len);
   printf(">> %d %f\n", len, size2);
   return 0;
}
```

Compilation and execution example on a Linux®
system:

```
$ fglmkext -o mycext.so mycext.c

$ fglcomp myprog.4gl

$ fglrun myprog.42m
Hand gloves         54
>> 11 0.540000
```
