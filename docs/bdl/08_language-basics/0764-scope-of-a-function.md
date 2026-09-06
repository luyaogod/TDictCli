---
title: "Scope of a function"
source: "fgl-topics/c_fgl_Functions_scope.html"
breadcrumb: "Language basics > Functions > Scope of a function"
type: "concept"
---

# Scope of a function

> A functions can be isolated to control its visibility to other modules.

A `FUNCTION` block cannot appear within the `MAIN` block, in a
`REPORT` block, or within another `FUNCTION` block. A function must be
declared at the root level in the source
code:

```
MAIN
  ...
END MAIN

FUNCTION myfunc( ... )
  ...
END FUNCTION

REPORT myrep( ... )
  ...
END REPORT
```

By default, functions are `PUBLIC`; They can be called by any other module of the
program.

If a function is only used by the current module, you can hide that function to other modules, to
make sure that it will not be called by mistake.

To keep a function local to the module, add the `PRIVATE` keyword before the
function header.

Private functions are only hidden to external modules, all function of the current module can
still call local private
functions.

```
PRIVATE FUNCTION check_number(n)
  ...
END FUNCTION
```

For better code readability, you can use the `PUBLIC` keyword for functions that
are global and visible to all
modules:

```
PUBLIC FUNCTION initialize()
  ...
END FUNCTION
```

When not using the module prefix, function symbols are global to all modules and must be unique.
However, with the `IMPORT FGL` method, you can use the module prefix when invoking a
function and simplify function naming conventions, since different modules can the define functions
with the same
name:

```
-- Module svgutils.4gl:
PUBLIC FUNCTION initialize()
  ...
END FUNCTION

-- Module sqlutils.4gl:
PUBLIC FUNCTION initialize()
  ...
END FUNCTION

-- Main program:
IMPORT FGL svgutils
IMPORT FGL sqlutils
MAIN
    CALL svgutils.initialize()
    CALL sqlutils.initialize()
    ...
END MAIN
```

## Related links

**Related concepts**  

[Importing modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
