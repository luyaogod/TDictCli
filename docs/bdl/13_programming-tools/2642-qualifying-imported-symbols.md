---
title: "Qualifying imported symbols"
source: "fgl-topics/c_fgl_refactor_qual_imp.html"
breadcrumb: "Programming tools > Source refactoring > Qualifying imported symbols"
type: "concept"
---

# Qualifying imported symbols

> The fglcomp compiler can automatically qualify imported symbols.

The fglcomp compiler provides the `--qualify-imports` option,
to automatically add module prefixes to all imported symbols that are not yet qualified.

> **Important:**
>
> The `--qualify-imports` option does not compile imported
> sources. The imported modules need to be available as .42m pcode
> modules.

By default, the modified source is written to the standard output stream
(stdout). No output is produced, if no modification is required. To replace the
orignal source file directly, add the `--inplace` option. With the
`--inplace` option, fglcomp makes a copy of the original source,
using the .4gl~ file extension.

Example with two modules defining public types, constants, variables and
functions: module1.4gl:

```
PUBLIC CONSTANT const1 = 500
PUBLIC TYPE type1 INTEGER
PUBLIC FUNCTION func1( p1 type1 )
    DISPLAY "func1:", p1
END FUNCTION
```

module2.4gl:

```
PUBLIC DEFINE var2 STRING
PUBLIC CONSTANT const2 = "abc"
```

Assuming the following importing
module: main.4gl

```
IMPORT FGL module1
IMPORT FGL module2
MAIN
    DEFINE v type1
    LET v = const1
    CALL func1( v )
    LET var2 = "abc"
    DISPLAY const2
END MAIN
```

Before refactoring the sources with the `--qualify-imports` option, you can find
unqualified imported symbols with the `-W unqualified-imports`
warning:

```
$ fglcomp -W unqualified-imports main.4gl
main.4gl:4:14:4:18:warning:(-8452) Unqualified imported symbol. Expecting 'module1.type1'.
main.4gl:5:13:5:18:warning:(-8452) Unqualified imported symbol. Expecting 'module1.const1'.
main.4gl:6:10:6:14:warning:(-8452) Unqualified imported symbol. Expecting 'module1.func1'.
main.4gl:7:9:7:12:warning:(-8452) Unqualified imported symbol. Expecting 'module2.var2'.
main.4gl:8:13:8:18:warning:(-8452) Unqualified imported symbol. Expecting 'module2.const2'.
```

The fglcomp --qualify-imports will produce this output from the above source
file: main.4gl:

```
IMPORT FGL module1
IMPORT FGL module2
MAIN
    DEFINE v module1.type1
    LET v = module1.const1
    CALL module1.func1( v )
    LET module2.var2 = "abc"
    DISPLAY module2.const2
END MAIN
```

If two modules define the same symbol name, fglcomp --qualify-imports will
produce the error [-8401](../15_library-reference/4483-genero-bdl-errors.md).
