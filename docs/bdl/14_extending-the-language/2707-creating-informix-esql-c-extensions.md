---
title: "Creating Informix ESQL/C Extensions"
source: "fgl-topics/c_fgl_CExtensions_005.html"
breadcrumb: "Extending the language > C-Extensions > Creating Informix® ESQL/C Extensions"
type: "concept"
description: "C-Extension libraries can be created from ESQL/C sources, as long as you have an Informix ESQL/C compiler which is compatible with your Genero runtime system."
---

# Creating Informix ESQL/C Extensions

> C-Extension libraries can be created from ESQL/C sources, as long as you have an Informix® ESQL/C compiler which is compatible with your Genero runtime system.

In order to create a C-Extension from ESQL/C sources, you must:

1. Define the list of user functions in the C interface file, by including the `fglExt.h`
   header file.
2. Compile the C interface file with your C compiler.
3. Modify your ESQL/C source modules by including the `fglExt.h` header file.
4. Compile the ESQL/C modules with the esql compiler, with the position-independent
   code option.
5. Create the shared library with the compiled C interface file and ESQL/C modules by linking with
   the `libfgl` runtime system library, and with the ESQL/C libraries (esql
   -libs), to resolve the ESQL/C symbols.

Include the `fglExt.h` header file in the following way:

```
#include "f2c/fglExt.h"
```

You can compile .ec extensions with the native Informix
esql compiler. This section describes how to use the Informix
esql compiler.

The following example shows how to compile and link an extension library with Informix
esql compiler:

Linux® command-line example:

```
esql -c -I$FGLDIR/include myext.ec 
gcc -c -fPIC -I$FGLDIR/include -I$INFORMIXDIR/incl/esql cinterf.c 
gcc -shared -o myext.so myext.o cinterf.o -L$FGLDIR/lib -lfgl \
    -L$INFORMIXDIR/lib -L$INFORMIXDIR/lib/esql `esql -libs`
```

Windows® command-line example (using Microsoft™ Visual C++):

```
esql -c myext.ec -I%FGLDIR%/include 
cl /DBUILDDLL /I%FGLDIR%/include /I%INFORMIXDIR%/incl/esql /c cintref.c 
esql -target:dll -o myext.dll myext.obj cinterf.obj %FGLDIR%\lib\libfgl.lib
```

When using Informix esql, you link the extension
library with Informix client libraries. These libraries
will be shared by the extension module and the Informix
database driver loaded by the Genero runtime system. Since both the extension functions and the
runtime database driver use the same functions to execute SQL queries, you can share the current SQL
connection opened in the Genero program to execute SQL queries in the extension functions. However,
mixing connection management instructions (`DATABASE`, `CONNECT TO`)
as well as database creation can produce unexpected results. For example, you cannot do a
`CREATE DATABASE` in your ESQL/C extension, and use this database to execute SQL
statements from the BDL code.
