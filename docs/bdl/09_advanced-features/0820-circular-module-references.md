---
title: "Circular module references"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_circular.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Circular module references"
type: "concept"
---

# Circular module references

> Circular references between imported modules are allowed.

A circular reference occurs when several modules define an interdependence with `IMPORT
FGL`.

Circular module references can be direct or indirect, for example:

1. Direct circular reference: Module A imports module B, which in turn imports module A.
2. Indirect circular reference: Module A imports module B, which imports module C, which imports
   module A.

When .42m files are not yet available, compiling modules with circular
dependency is only possible when these modules exist in the same directory or in a package
sub-directory, and [auto-compilation](0819-auto-compilation-of-imported-modules.md "Imported local and package modules are compiled automatically if needed.") is
used (option `--implicit=none` is not used). Otherwise, all interdependent modules
must be specified in the fglcomp command.

Code example: Module "module\_a.4gl":

```
IMPORT FGL module_b

MAIN
    CALL function_a1()
END MAIN

FUNCTION function_a1()
    DISPLAY "In module_a: function_a1()"
    CALL module_b.function_b1()
END FUNCTION

FUNCTION function_a2()
    DISPLAY "In module_a: function_a2()"
END FUNCTION
```

Module "module\_b.4gl":

```
IMPORT FGL module_a

FUNCTION function_b1()
    DISPLAY "In module_b: function_b1()"
    CALL module_a.function_a2()
END FUNCTION
```

Compile the sample (to get detailed information about the compilation process, we use the [`--verbose` option](../13_programming-tools/2534-compiling-program-code-files-4gl.md) of
fglcomp):

```
$ fglcomp --verbose module_a.4gl module_b.4gl
[parsing module_a.4gl]
[parsing module_b.4gl]
[building module_b]
[writing module_b.42m]
[building module_a]
[writing module_a.42m]
...
```

Execute the sample:

```
$ fglrun module_a
In module_a: function_a1()
In module_b: function_b1()
In module_a: function_a2()
```
