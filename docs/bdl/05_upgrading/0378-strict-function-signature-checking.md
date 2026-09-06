---
title: "Strict function signature checking"
source: "fgl-topics/c_fgl_MigI4GL_026.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Strict function signature checking"
type: "concept"
---

# Strict function signature checking

> With Genero BDL, a function's signature is detected at link time.

IBM® Informix® 4GL (I4GL) is not very strict regarding
function signature. With I4GL, you can, for example, define a function
in module A that returns three values, and call that function in
module B with a returning clause specifying two variables:

Module A:

```
FUNCTION func()
   RETURN "abc", "def", "ghi"
END FUNCTION
```

Module B (main):

```
MAIN
   DEFINE v1, v2 VARCHAR(100)
   CALL func() RETURNING v1, v2
END MAIN
```

The c4gl compiler (7.32) compiles and links these modules without
error, but at execution time you get the following runtime error:

```
Program stopped at  "main.4gl", line number 3. 
FORMS statement error number -1320. 
A function has not returned the correct number of values 
expected by the calling function.
```

With Genero Business Development Language (BDL), the mistake will be detected at link
time:

```
$ fgllink -o prog.42x main.42m module_a.42m
ERROR(-6200): Module 'main': The function module_a.func(0,3) will be 
called as func(0,2).
```

Similarly, I4GL does not detect an invalid number of parameters
passed to a function defined in a different module:

Module A:

```
FUNCTION func( p )
   DEFINE p INTEGER
   DISPLAY p 
END FUNCTION
```

Module B (main):

```
MAIN
   CALL func(1,2)
END MAIN
```

The c4gl compiler (7.32) compiles and links these modules without
error, but at execution time, you get the following runtime error:

```
Program stopped at  "main.4gl", line number 2. 
FORMS statement error number -1318. 
A parameter count mismatch has occurred between the calling 
function and the called function.
```

When using Genero BDL, the error will be detected at link
time:

```
$ fgllink -o prog.42x main.42m module_a.42m
ERROR(-6200): Module 'main': The function module_a.func(1,0) will be 
called as func(2,0).
```

However, Genero BDL does not check function signatures when several `RETURN`
instructions are found by the compiler. This is necessary in order to be compatible with I4GL. The
next code example compiles and runs with both I4GL and
BDL:

```
MAIN
   DEFINE v1, v2 VARCHAR(100)
   CALL func(1) RETURNING v1
   DISPLAY v1
   CALL func(2) RETURNING v1, v2
   DISPLAY v1, v2
END MAIN

FUNCTION func( n )
   DEFINE n INTEGER
   IF n == 1 THEN
      RETURN "abc"
   ELSE
      RETURN "abc", "def"
   END IF
END FUNCTION
```

However, this type of programming is not recommended.

## Related links

**Related concepts**  

[Returning values](../08_language-basics/0768-returning-values.md "A function can return values with the RETURN instruction.")
