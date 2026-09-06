---
title: "Microsoft Windows platform notes"
source: "fgl-topics/c_fgl_installation_023.html"
breadcrumb: "Installation > Platform specific notes > Microsoft™ Windows® platform notes"
type: "concept"
description: "Microsoft™ Visual C++ version When using C-Extensions, you need Microsoft Visual C++ compiler to compile and link your C sources. Make sure you have installed the software package corresponding to the ..."
---

# Microsoft Windows platform notes

## Microsoft™ Visual C++ version

When using C-Extensions, you need Microsoft Visual C++ compiler
to compile and link your C sources. Make sure you have installed
the software package corresponding to the MSVC version installed
on your system. The MSVC version is identified in the software package
name.

## Checking binary dependencies

Microsoft Visual C++ provides the
dumpbin utility to extract information from a binary file.

Use the `/dependents` option to check for DLL
dependencies:

```
C:\> dumpbin /dependents mylib.dll 
Microsoft (R) COFF/PE Dumper Version 7.10.3077 
Copyright (C) Microsoft Corporation.  All rights reserved.

Dump of file mylib.dll

File Type: EXECUTABLE IMAGE

Image has the following dependencies:

isqlt09a.dll
MSVCR71.dll
KERNEL32.dll

Summary

1000 .data
1000 .rdata
1000 .text
```

## Changing the stack size of fglrun

On Windows® platforms, the
fglrun.exe binary has a predefined C stack size. In some rare cases (for
example, when programs do deep recursion), the stack size of fglrun.exe binary
needs to be changed to avoid a stack overflow. The stack size of fglrun can be changed permanently
by patching the EXE file with the Microsoft Visual C++
editbin utility. Check the stack size by running the dumpbin
utility on fglrun.exe as follows:

```
C:\> dumpbin /headers %FGLDIR%\bin\fglrun.exe
```

Search for the line containing "stack reserve" words in the OPTIONAL
HEADER VALUES section:

```
OPTIONAL HEADER VALUES
    ...
    100000 size of stack reserve
```

The stack size is displayed in hexadecimal value. So for example, a value 100,000 means
1,048,567 bytes = 1MB.

In order to modify the stack size of fglrun.exe, run the
editbin utility on fglrun.exe with the /stack
option:

```
C:\> editbin /stack:1000000 %FGLDIR%\bin\fglrun.exe
```

See Microsoft Visual C++ documentation for more details.

## Java Interface

Check [Microsoft
Windows platform notes](../14_extending-the-language/2660-platform-specific-notes-for-the-jvm.md) when using the [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") on this operating system.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
