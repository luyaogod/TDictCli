---
title: "C-Extensions"
source: "fgl-topics/c_fgl_CExtensions_001.html"
breadcrumb: "Extending the language > C-Extensions"
type: "concept"
---

# C-Extensions

> With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.


## Child topics

- [Understanding C-Extensions](2704-understanding-c-extensions.md): C-Extensions allow you to write custom shared libraries in the C language.
- [Header files for ESQL/C typedefs](2705-header-files-for-esql-c-typedefs.md): C header files (.h) are required to define C structures for complex data types used in a C-Extension.
- [Creating C-Extensions](2706-creating-c-extensions.md): Custom C-Extensions must be provided to the runtime system as Shared Objects (.so) on UNIX™, and as Dynamically Loadable Libraries (.DLL) on Windows®.
- [Creating Informix ESQL/C Extensions](2707-creating-informix-esql-c-extensions.md): C-Extension libraries can be created from ESQL/C sources, as long as you have an Informix® ESQL/C compiler which is compatible with your Genero runtime system.
- [The C interface file](2708-the-c-interface-file.md): To make your C functions visible to the runtime system, you must define all the functions in the C interface file.
- [Linking programs using C-Extensions](2709-linking-programs-using-c-extensions.md): When creating a 42r program or 42x library, the linker needs to resolve all function names, including C-Extension functions.
- [Loading C-Extensions at runtime](2710-loading-c-extensions-at-runtime.md): The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files.
- [Runtime stack functions](2711-runtime-stack-functions.md): To pass values between a C function and a program, the C function and the runtime system use the runtime stack.
- [Data types and structures](2712-data-types-and-structures.md): C types are used to write C-Extensions.
- [NULL handing](2713-null-handing.md): Handling NULL in C-Extensions.
- [Calling C functions from programs](2714-calling-c-functions-from-programs.md): C-Extensions functions can be called from the program in the same way that you call a BDL function.
- [Calling program functions from C](2715-calling-program-functions-from-c.md): It is possible to call a BDL function from a C-Extension function.
- [Simple C-Extension example](2716-simple-c-extension-example.md): This example shows how to create a C-Extension library on Linux® using gcc.
- [Implementing C-Extensions for GMI](2717-implementing-c-extensions-for-gmi.md): This section describes how to program C-Extensions for the GMI VM.
