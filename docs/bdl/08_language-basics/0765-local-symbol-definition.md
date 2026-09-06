---
title: "Local symbol definition"
source: "fgl-topics/c_fgl_Functions_locals.html"
breadcrumb: "Language basics > Functions > Local symbol definition"
type: "concept"
---

# Local symbol definition

> Symbols defined inside a function body are only visible to the function.

## Purpose of local symbols

Local function symbols are part of and only to be used in the function body. They are created and
initialized when the function is invoked. They are dropped when the function terminates.

## Defining local symbols

Inside the body of a function, you can define language elements that will only be visible for the
function code:

- local constants with the [`CONSTANT`](0710-constants.md "The definition of constants allows to centralize common static values.")
  instruction,
- local user-defined types with the [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") instruction,
- local variables with the [`DEFINE`](0686-variables.md "Explains how to define program variables.")
  instruction.

```
FUNCTION check_customer( cust_id INTEGER )
  CONSTANT c_max = 1000 -- local constant
  TYPE t_cust RECORD LIKE customer.* -- local type
  DEFINE found BOOLEAN -- local variable
  ...
END FUNCTION
```

Function arguments and local symbols must use different names, it is not possible to define a
local variable with the same name as a function
parameter:

```
FUNCTION func_a(x INTEGER)
   DEFINE x INTEGER
| The symbol 'x' has been defined more than once.
| See error number -4319.
   LET x = 1
END FUNCTION
```

## Scope of local symbols

Local function symbols are not visible in other program blocks. Global or module variable can use
the same name as a local variable: The global or module variable is not visible within the function
scope of the local variable using the same
name.

```
DEFINE x INTEGER   -- Declares a module variable

FUNCTION func_a()
  DEFINE x INTEGER -- Declares a local variable
  LET x = 123      -- Assigns local variable
END FUNCTION

FUNCTION func_b()
  LET x = 123     -- Changes the module variable
END FUNCTION
```

To increase code readability, consider using different names for global, module and local
function symbols. A common practice is to use prefixes for global (g\_) and module (m\_)
variables.

## Related links

**Related concepts**  

[Variables](0686-variables.md "Explains how to define program variables.")
