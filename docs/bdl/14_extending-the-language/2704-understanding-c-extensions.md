---
title: "Understanding C-Extensions"
source: "fgl-topics/c_fgl_CExtensions_002.html"
breadcrumb: "Extending the language > C-Extensions > Understanding C-Extensions"
type: "concept"
---

# Understanding C-Extensions

> C-Extensions allow you to write custom shared libraries in the C language.

Using C-Extensions, C functions implemented in shared libraries can be called from
the Genero application code. This feature allows you to extend the language with custom libraries,
or existing standard libraries, by writing some 'wrapper functions' to interface with the Genero
language.

On regular platforms, C-Extensions are implemented with shared libraries, that are loaded by the
[fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") program on demand.

Platforms such as iOS mobile devices do not allow you to load shared libraries. In this case, you
must re-link the virtual machine. For more details, see [Implementing C-Extensions for GMI](2717-implementing-c-extensions-for-gmi.md "This section describes how to program C-Extensions for the GMI VM.").

Function parameters and returned values are passed/returned on
the runtime stack, using [pop/push
functions](2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack."). Be sure to pop and push the exact number of parameters/returns
expected by the caller; otherwise, a fatal stack error will be raised
at runtime.

In order to use a C-Extension in your program, you typically specify the library name with the
[`IMPORT`](2710-loading-c-extensions-at-runtime.md "The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files.") instruction at the beginning
of the module calling the C-Extension functions. The compiler can then check for the existence of
the functions and the library will be automatically loaded at runtime.

The C code written in C-Extensions is usually platform specific, which does not ease the migration
of your application to a different operating system, especially when doing a lot of system calls.
Additionally, C data types are defined differently depending on the processor architecture (32 / 64
bits issues). This can also be an issue.

Make sure that the functions defined in your C-Extensions do not conflict with program functions.
In case of conflict, you will get a compiler or a runtime error, depending on the [loading technique used](2710-loading-c-extensions-at-runtime.md "The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files.").

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")
