---
title: "C compiler for C extensions"
source: "fgl-topics/c_fgl_installation_006.html"
breadcrumb: "Installation > Software requirements > C compiler for C extensions"
type: "concept"
---

# C compiler for C extensions

> Ensure you have a C compiler and linker to compile your C-Extensions.

Applications using C extensions, need a C compiler and linker to build the C extension
library that will be loaded by the runtime system.

## C compiler on UNIX™ platforms

On UNIX platforms, you need a cc
compiler on the system where you create the C extension libraries. Note that some UNIX systems do not have a C compiler
installed by default.

## C compiler on Microsoft™ Windows® platforms

On Windows platforms, it
is mandatory to install Microsoft Visual C++
version corresponding to the installed Genero BDL package. The OS identifier of the Genero
BDL package filename identifies the Visual C++ version to be used.

## C compiler on macOS® platforms

On the macOS where you create the C
extension libraries, the minimum supported Xcode® version is version 12.2.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
