---
title: "Creating C-Extensions"
source: "fgl-topics/c_fgl_CExtensions_004.html"
breadcrumb: "Extending the language > C-Extensions > Creating C-Extensions"
type: "concept"
description: "Custom C-Extensions must be provided to the runtime system as Shared Objects (.so) on UNIX, and as Dynamically Loadable Libraries (.DLL) on Windows."
---

# Creating C-Extensions

> Custom C-Extensions must be provided to the runtime system as Shared Objects (.so) on UNIX™, and as Dynamically Loadable Libraries (.DLL) on Windows®.

In order to create a C-Extension, you must:

1. Setup you C compiler environment.
2. Modify your C source modules by including the fglExt.h header
   file:

   ```
   #include "f2c/fglExt.h"
   ```
3. Define the list of user functions in the [C interface
   file](2708-the-c-interface-file.md "To make your C functions visible to the runtime system, you must define all the functions in the C interface file."), by including the fglExt.h header file.
4. Compile the C sources and C interface file, the link these object files to create a shared
   object (.so or .DLL)

To setup the C compiler environment: On a UNIX/Linux®
operating system, check the cc or gcc command. On a Windows, you need a Visual C++ version compatible with the FGL
package. It is mandatory to setup the Visual C++ environment by executing the
vcvars64.bat command file provided in the Visual C++ installation
directory.

> **Note:**
>
> When migrating from IBM® Informix® 4GL, it is possible that existing C-Extension sources include Informix specific headers like sqlhdr.h
> or decimal.h. Replace the Informix 4GL header files by
> fglExt.h header file.

In order to compile and create the shared object of your C-Extension library, use the
fglmkext command line
tool:

```
fglmkext -o myext.so module_a.c module_b.c
```

The fglmkext command line tool contains platform-specific C compiler and
linker options required to build a C Extension library.

## Related links

**Related concepts**  

[fglmkext](../13_programming-tools/2519-fglmkext.md "The fglmkext tool compiles and links a user C Extension.")

[Header files for ESQL/C typedefs](2705-header-files-for-esql-c-typedefs.md "C header files (.h) are required to define C structures for complex data types used in a C-Extension.")
