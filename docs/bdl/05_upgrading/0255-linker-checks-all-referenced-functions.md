---
title: "Linker checks all referenced functions"
source: "fgl-topics/c_fgl_Migrate_to_240_linker_error_1338.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Linker checks all referenced functions"
type: "concept"
---

# Linker checks all referenced functions

> The linker checks definition of all functions referenced in all modules provided in the link command.

Starting with version 2.40, any reference to a function has to be resolved by the linker: When
linking a 42r program, if an unused module references an undefined function, the linker
(fglrun -l or fgllink) will stop with the error
[-1338](../15_library-reference/4483-genero-bdl-errors.md). Before
version 2.40, the undefined function was ignored.

> **Note:**
>
> Complete function reference is only checked by the linker when creating a
> 42r program file. When creating a 42x library, there can be references to
> undefined functions.

In the example, the main.4gl module does not call any function, but
the module used in the link line (module.4gl) defines an unused function
(`f1`) calling an undefined function (`f2`):

main.4gl:

```
MAIN
    DISPLAY "In main..."
END MAIN
```

module.4gl:

```
FUNCTION f1() -- Unused in program
    DISPLAY "In f1..."
    CALL f2() -- Undefined
END FUNCTION
```

Compiling and linking:

```
$ fglcomp main.4gl
$ fglcomp module.4gl
$ fgllink -o prog.42r main.42m module.42m
ERROR(-1338):The function 'f2' has not been defined in any module in the program.
```

## Related links

**Related concepts**  

[Linking programs](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.")

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
