---
title: "Make symbols case match definition"
source: "fgl-topics/c_fgl_refactor_char_case.html"
breadcrumb: "Programming tools > Source refactoring > Make symbols case match definition"
type: "concept"
---

# Make symbols case match definition

> The fglcomp compiler can make all symbol names match their definition exactly.

The fglcomp compiler provides the `--fix-case` option, to make
all symbols match exactly the name used in the definition of the symbol.

> **Tip:**
>
> Before refactoring all sources with `--fix-case`, use the [`-Wcase`
> option](2516-fglcomp.md) of fglcomp, to find how many symbols do not match their
> definition exactly. To force a compilation error, add the `-Werror` option.

By default, the modified source is written to the standard output stream
(stdout). No output is produced, if no modification is required. To replace the
orignal source file directly, add the `--inplace` option. With the
`--inplace` option, fglcomp makes a copy of the original source,
using the .4gl~ file extension.

In the next example, the variable `custid` in the `LET` instruction
should use the same "UpperCamelCase" naming convention, as in its
definition:

```
MAIN
    DEFINE CustId INTEGER
    LET custid = 998
END MAIN
```

To fix this code, use `fglcomp
--fix-case`:

```
$ fglcomp --fix-case main.4gl
MAIN
    DEFINE CustId INTEGER
    LET CustId = 998
END MAIN
```

To replace the original file directly, use the `--inplace`
option:

```
$ fglcomp --fix-case --inplace main.4gl
```
