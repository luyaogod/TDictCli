---
title: "Identifying modules to be imported"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_miss_imp.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Identifying modules to be imported"
type: "concept"
---

# Identifying modules to be imported

> Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions.

When migrating existing projects using traditional linking, after compiling all the
.4gl sources, consider using the `--print-missing-imports`
option of fgllink or fglrun, to print the `IMPORT
FGL` suggestions for all the modules specified in the fglrun command
line.

On the other hand, the `--print-imports` option reports all imported modules that
are really used.

Just add the `--print-missing-imports` to the fgllink or
fglrun -l commands, to identify what modules should be
imported:

```
$ head *.4gl
==> main.4gl <==
MAIN
    CALL func1()
END MAIN

==> mod1.4gl <==
FUNCTION func1()
    CALL func2()
END FUNCTION

==> mod2.4gl <==
FUNCTION func2()
END FUNCTION

$ fglcomp *.4gl

$ fgllink --print-missing-imports -o prog.42r *.42m
-- in main.4gl
IMPORT FGL mod1

-- in mod1.4gl
IMPORT FGL mod2
```

The `--print-missing-imports` option will try to resolve all (function) symbols as
done during linking, but instead of producing a .42r program, it will list the
`IMPORT FGL` instructions to be added in each module, and thus avoid linking.

The `--print-imports` option prints all `IMPORT FGL` instructions,
that are really used by a module, where at least one symbol (function, variable, constant, type,
etc) is used by the importing module.

Consider using the `--print-missing-imports` options instead of
`--print-imports`: If you are missing an `IMPORT FGL` for a
non-function symbol like a variable, constant, or type, the compiler will produce a compilation
error. Regarding function symbols, the compilation is possible even when the module was not
imported. Therefore, `--print-missing-imports` is useful to identify modules to be
imported to resolve function symbols and avoid linking.

In the next example, `mod1` can be imported in `main`, but
`mod2` already imports `mod1`:

```
$ head *.4gl
==> main.4gl <==
MAIN
    CALL func1()
END MAIN

==> mod1.4gl <==
FUNCTION func1()
   CALL func2()
END FUNCTION

==> mod2.4gl <==
IMPORT FGL mod1
FUNCTION func2()
   CALL func1()
END FUNCTION

$ fglcomp main.4gl mod1.4gl mod2.4gl

$ fglrun --print-imports main.42m mod1.42m mod2.42m
-- in main.4gl
IMPORT FGL mod1

-- in mod1.4gl
IMPORT FGL mod2

-- in mod2.4gl
IMPORT FGL mod1

$ fglrun --print-missing-imports main.42m mod1.42m mod2.42m
-- in main.4gl
IMPORT FGL mod1

-- in mod1.4gl
IMPORT FGL mod2
```
