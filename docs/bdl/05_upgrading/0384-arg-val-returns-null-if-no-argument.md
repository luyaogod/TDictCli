---
title: "arg_val() returns NULL if no argument"
source: "fgl-topics/c_fgl_MigI4GL_032.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > arg_val() returns NULL if no argument"
type: "concept"
---

# arg_val() returns NULL if no argument

> If the index passed to the function references an argument that does not exist, the arg_val() function must be handled differently between I4GL and Genero BDL.

With IBM® Informix® 4GL (I4GL), if the index passed to the function references an argument that does not
exist, the `arg_val()` function returns a blank character. With Genero BDL, if the
index passed to the function references an argument that does not exist, the
`arg_val()` function returns `NULL`:

```
$ cat arg_val.4gl
MAIN
  DEFINE v VARCHAR(10)
  LET v = arg_val(1)
  IF v IS NULL THEN
     DISPLAY "IS NULL"
  ELSE
     DISPLAY "IS NOT NULL: [", v, "]"
  END IF
END MAIN

$ fglcomp arg_val.4gl && fglrun arg_val.42m
IS NULL

$ c4gl -o arg_val.bin arg_val.4gl && ./arg_val.bin
IS NOT NULL: [ ]
```

You must review `arg_val()` conditions using the `!=` operator, and
add `NVL()` to handle `NULL`
arguments:

```
-- Default is "USER" mode when no argument is passed or when argument
-- is different from "ADMIN"
MAIN
  --IF arg_val(1)!="ADMIN" THEN
  IF NVL(arg_val(1),"?")!="ADMIN" THEN
     DISPLAY "User mode..."
  ELSE
     DISPLAY "Admin mode..."
  END IF
END MAIN
```

Normally the code should use the `NUM_ARGS()`, to check the number of parameters
passed to the program.

## Related links

**Related concepts**  

[arg\_val()](../15_library-reference/2726-arg-val.md "Returns a command line argument by position.")

[num\_args()](../15_library-reference/2731-num-args.md "Returns the number of program arguments.")
