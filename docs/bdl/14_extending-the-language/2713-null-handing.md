---
title: "NULL handing"
source: "fgl-topics/c_fgl_CExtensions_null.html"
breadcrumb: "Extending the language > C-Extensions > NULL handing"
type: "concept"
---

# NULL handing

> Handling NULL in C-Extensions.

Variables passed from the BDL code can be checked for `NULL` depending on
the data type.

Define macros as follows to check for `NULL`
values:

```
#define MY_IS_NULL_SMALLFLOAT(v) (isnan(v))
#define MY_IS_NULL_FLOAT(v) (isnan(v))
#define MY_IS_NULL_INTEGER(v) (v == (int) 0x80000000)
#define MY_IS_NULL_SMALLINT(v) (v == (short int) 0x8000)
#define MY_IS_NULL_CHAR(v) (v[0] == '\0')
#define MY_IS_NULL_DECIMAL(v) (v.dec_pos < 0)
#define MY_IS_NULL_DATETIME(v) (v.dt_dec.dec_pos < 0)
#define MY_IS_NULL_INTERVAL(v) (v.in_dec.dec_pos < 0)
#define MY_IS_NULL_LOCATOR(v) (v.loc_indicator < 0)
```

To return a `NULL` value from your C extension function, push an empty string on
the stack as follows:

```
pushquote("", 0);
```

## Example

C Extension source (ext1.c):

```
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "f2c/fglExt.h"

int check_null(int);

UsrFunction usrFunctions[]={
  { "check_null",   check_null,     1, 1 },
  { NULL,           NULL,           0, 0 }
};

#define MY_IS_NULL_DECIMAL(v) (v.dec_pos < 0)

int check_null( int pc )
{
    dec_t price;
    const char *answer;
    int len;
    if (pc != 1) exit(1);
    popdec(&price);
    if (MY_IS_NULL_DECIMAL(price)) {
        answer = "This decimal is NULL";
    } else {
        answer = "This decimal is not NULL";
    }
    len = (int) strlen(answer);
    pushquote(answer, len);
    return 1;
}
```

Genero program
(main.4gl):

```
IMPORT libext1
MAIN
    DEFINE d DECIMAL(10,2)
    LET d = NULL
    DISPLAY check_null(d)
    LET d = 12.34
    DISPLAY check_null(d)
END MAIN
```

Commands to compile and execute (on
Linux):

```
$ gcc -fPIC -c ext1.c -I $FGLDIR/include
$ gcc --shared -o libext1.so ext1.o -L$FGLDIR/lib -lfgl
$ fglcomp main.4gl
$ fglrun main.42m
This decimal is NULL
This decimal is not NULL
```
