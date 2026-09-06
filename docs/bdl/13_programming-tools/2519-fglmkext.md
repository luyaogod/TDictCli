---
title: "fglmkext"
source: "fgl-topics/c_fgl_tools_fglmkext.html"
breadcrumb: "Programming tools > Command reference > fglmkext"
type: "concept"
---

# fglmkext

> The fglmkext tool compiles and links a user C Extension.

## Syntax

```
fglmkext [options] source.c [...]
```

1. options are described in Table 1.
2. source.c is a C source file implementing C extension
   functions.

## Options

| Option | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |
| `-o libname` | Output file specification, defines the C Extension library name. |

## Usage

The fglmkext command line tool compiles and links a C Extension library.

The command can be used with a single source file, the name of the library will default to the
name of the specified source:

```
fglmkext myext.c
```

If a single C source file is provided, it must define the `usrFunctions` C
extension interface structure as well as the functions to be used from a BDL program.

In order to specify a library name, use the `-o` option, several C source files
can also be specified. For example, on a UNIX
platform:

```
fglmkext -o mycext.so module_a.c module_b.c
```

If case of error, the fglmkext command execution status is different from
zero. Consequently, it is possible to detect compilation errors in scripts and makefiles.

## Related links

**Related concepts**  

[Creating C-Extensions](../14_extending-the-language/2706-creating-c-extensions.md "Custom C-Extensions must be provided to the runtime system as Shared Objects (.so) on UNIX, and as Dynamically Loadable Libraries (.DLL) on Windows.")
