---
title: "The C interface file"
source: "fgl-topics/c_fgl_CExtensions_006.html"
breadcrumb: "Extending the language > C-Extensions > The C interface file"
type: "concept"
---

# The C interface file

> To make your C functions visible to the runtime system, you must define all the functions in the C interface file.

The C interface file is a C source file that defines the
`usrFunctions` array. This array defines C functions that
can be called from programs.

The last record of the `usrFunctions` array must be a
line with all the elements set to NULL/0, to define the end of the list.

Each element of the `usrFunctions` array must include
the following members:

1. The first member is the name of the function, provided as a `(const char *)`
   character string.
2. The second member is the C function symbol, provided as an `(int (*function)
   (int))` C function pointer.
3. The third member is the number of parameters passed to the function through the runtime stack,
   provided as an `(int)`; use a negative value like -1 to specify a variable number of
   arguments.
4. The fourth member is the number of values returned by the function, provided as an
   `(int)`; use a negative value like -1 to specify a variable number of return
   values.

The third and fourth member of a `UsrFunction` element can be defined as a
negative value (-1), to indicate a variable number of arguments and/or return values.

You typically do a forward declaration of your C functions, before the
`usrFunctions` array initializer:

```
#include "f2c/fglExt.h"

int c_init(int);
int c_set_trace(int);
int c_get_message(int);
int c_compare(int);
int c_generate(int);

UsrFunction usrFunctions[]={
  { "init",         c_init,         0, 0 },
  { "set_trace",    c_set_trace,    1, 0 },
  { "get_message",  c_get_message,  1, 1 },
  { "compare",      c_compare,     -1, 1 }, /* var. numb. params */
  { "generate",     c_generate,     1,-1 }, /* var. numb. returns */
  { NULL,           NULL,           0, 0 }
};
```

Note that the `UsrFunction` structure contains an additional
member, dedicated for internal use. If you experience compiler warnings because
of uninitialized structure members, simply complete the C function definitions
with a fifth zero value:

```
/* Avoids C compiler warnings because of un-initialized structure members */

UsrFunction usrFunctions[]={
  { "init",         c_init,         0, 0, 0 },
            /* member for internal use ---^ */
  ...
```
