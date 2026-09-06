---
title: "Variable assignment in function calls"
source: "fgl-topics/c_fgl_MigI4GL_043.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Variable assignment in function calls"
type: "concept"
---

# Variable assignment in function calls

> Genero BDL behaves differently to I4GL when variables are changed in expressions evaluated in function parameters.

When calling a function with expressions as parameters, these expressions must be evaluated
before calling the function. If an expression changes a program variable that is also used as
function parameter, IBM® Informix® 4GL pushes the evaluated values sequentially on the stack, while Genero BDL first
evaluates all parameters, then pushes the list of values on the stack.

To illustrate this difference, see the following example:

```
DEFINE gv INT

MAIN
    LET gv = 1
    CALL func(gv, add(), gv)
END MAIN

FUNCTION add()
    LET gv = gv + 1
    RETURN gv
END FUNCTION

FUNCTION func(p1,p2,p3)
    DEFINE p1,p2,p3 INT
    DISPLAY p1,p2,p3
END FUNCTION
```

Result with I4GL:

```
          1          2          2
```

Result with BDL:

```
          2          2          2
```

The behavior of Genero BDL is better than I4GL, using a consistent value of the modified
variables, for all parameters passed to the function.

Replace the above code by:

```
DEFINE gv INT
# ...
MAIN
    LET gv = 1
    LET gv = add()
    CALL func(gv, gv, gv)
END MAIN
```

## Related links

**Related concepts**  

[Using program variables in static SQL](../10_sql-support/1117-using-program-variables-in-static-sql.md "Static SQL syntax supports the usage of program variables as SQL parameters.")
