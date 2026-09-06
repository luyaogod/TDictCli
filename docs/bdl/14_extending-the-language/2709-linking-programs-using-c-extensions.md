---
title: "Linking programs using C-Extensions"
source: "fgl-topics/c_fgl_CExtensions_007.html"
breadcrumb: "Extending the language > C-Extensions > Linking programs using C-Extensions"
type: "concept"
---

# Linking programs using C-Extensions

> When creating a 42r program or 42x library, the linker needs to resolve all function names, including C-Extension functions.

If extension modules are not specified explicitly in the source files with the [`IMPORT`](2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack.") directive, you must give the
extension modules with the `-e` option in the command line:

```
fgllink -e myext1,myext2,myext3 -o myprog.42r moduleA.42m moduleB.42m ...
```

The `-e` option does not include C-Extension references into the
.42r file: At runtime, also use the `-e` argument with the
fglrun command, in order to load the extension libraries when executing
programs.

The `-e` option is not needed when using the default
`userextension` module, or if C-Extensions are specified with the
`IMPORT` directive.

## Related links

**Related concepts**  

[fgllink](../13_programming-tools/2517-fgllink.md "The fgllink tool assembles p-code modules produced with fglcomp into a .42r program or a .42x library.")

[Loading C-Extensions at runtime](2710-loading-c-extensions-at-runtime.md "The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files.")

[Compiling source files](../13_programming-tools/2530-compiling-source-files.md "Describes how to build the runtime files from source files.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")
