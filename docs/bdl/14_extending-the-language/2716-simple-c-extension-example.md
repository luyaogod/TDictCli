---
title: "Simple C-Extension example"
source: "fgl-topics/c_fgl_CExtensions_019.html"
breadcrumb: "Extending the language > C-Extensions > Simple C-Extension example"
type: "concept"
description: "This example shows how to create a C-Extension library on Linux using gcc."
---

# Simple C-Extension example

> This example shows how to create a C-Extension library on Linux® using gcc.

The command line options used in this sample to compile and link shared libraries are for a Linux
operating system.

## The "splitext.c" C interface file

```
#include "f2c/fglExt.h"

int next_token(int);

UsrFunction usrFunctions[]={
  { "next_token", next_token, 1, 2, 0 },
  { 0,0,0,0,0 }
};
```

## The "split.c" file

```
#include <string.h>
#include <assert.h>
#include "f2c/fglExt.h"

int next_token( int in_num );
int next_token( int in_num )
{
       char src[2049];
       char *p;
       assert(in_num == 1);
       popvchar(src, sizeof(src));
       if (*src == '\0') {
           pushvchar("", 0);
           pushvchar("", 0);
       } else {
           p = strchr(src, ' ');
           if (p == NULL) {
               pushvchar(src, strlen(src));
               pushvchar("", 0);
           } else {
               *p = '\0';
               pushvchar(src, strlen(src));
               p++;
               pushvchar(p, strlen(p));
           }
       }
       return 2;
}
```

## The program "split.4gl"

```
IMPORT libsplit
MAIN
    DEFINE t, r STRING -- Or VARCHAR(100)
    LET r = "This is my first C Extension"
    WHILE TRUE
        CALL next_token(r) RETURNING t, r
        IF t IS NULL THEN
            EXIT WHILE
        END IF
        DISPLAY "token: ", t
    END WHILE
END MAIN
```

## Compiling and running the sample

Create the C extension shared library with the [fglmkext](../13_programming-tools/2519-fglmkext.md "The fglmkext tool compiles and links a user C Extension.")
utility:

```
fglmkext -o libsplit.so split.c splitext.c
```

Compile the .4gl module:

```
$ fglcomp split.4gl
```

Run the sample:

```
$ fglrun split.42m
token: This
token: is
token: my
token: first
token: C
token: Extension
```
