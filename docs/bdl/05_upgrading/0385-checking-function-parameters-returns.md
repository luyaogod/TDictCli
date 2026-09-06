---
title: "Checking function parameters/returns"
source: "fgl-topics/c_fgl_MigI4GL_035.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Checking function parameters/returns"
type: "concept"
---

# Checking function parameters/returns

> With Genero BDL, a parameter count mismatch is detected at link time.

IBM® Informix® 4GL
makes function parameter checking at runtime. It is possible to compile a source that calls a
function of another module with more (or less) parameters (or return values) as specified in the
function definition.

Unlike I4GL, the Genero BDL linker can detect such programming errors when building the program.
This is much better as detecting the programming error a runtime.

See the following example:

```
-- main.4gl
MAIN
    CALL func(1,2)
END MAIN
```

```
-- mod1.4gl
FUNCTION func(a,b,c)
    DEFINE a,b,c INT
END FUNCTION
```

With I4GL, the programming error can only be detected at
runtime:

```
$ c4gl -c mod1.4gl 
$ c4gl -c main.4gl 
$ c4gl -o main.bin main.o mod1.o
$ ./main.bin  
Program stopped at "mod1.4gl", line number 2.
FORMS statement error number -1318.
A parameter count mismatch has occurred between the calling 
function and the called function.
```

With Genero BDL, the programming error is detected at link time:

```
$ fglcomp mod1.4gl
$ fglcomp main.4gl
$ fgllink -o main.42r main.42m mod1.42m
ERROR(-6200):Module 'main': The function mod1.func(3,0) will be called as func(2,0).
```

> **Tip:**
>
> Consider using IMPORT FGL to get rid of program linking with Genero!

## Related links

**Related concepts**  

[IMPORT FGL](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.")

[Linking programs](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.")
