---
title: "Implicit data type conversion on the stack"
source: "fgl-topics/c_fgl_runtime_stack_implicit_conversions.html"
breadcrumb: "Advanced features > Runtime stack > Implicit data type conversion on the stack"
type: "concept"
description: "When a value or a reference is popped from the stack, implicit data conversion takes place. This means, for example, that you can pass a string value to a function that defines the receiving variable ..."
---

# Implicit data type conversion on the stack

When a value or a reference is popped from the stack, implicit
data conversion takes place. This means, for example, that you can
pass a string value to a function that defines the receiving variable
as a numeric data type; no compilation error will occur, but you can
get a runtime error if the string cannot be converted to a numeric.
The same principle applies to values returned from functions, since
the stack is also used in this case.

```
MAIN
  DEFINE s STRING
  LET s = "123"
  CALL display_integer(s) -- Will be accepted
  LET s = "abc"
  CALL display_integer(s) -- Will fail with conversion error
END MAIN

FUNCTION display_integer(x INTEGER) RETURNS ()
  DISPLAY x
END FUNCTION
```

When using the `RETURNS` clause in the function definition, the returned value is
typed as specified in the `RETURNS` clause.

```
MAIN
  DISPLAY int_div(5,3)    -- shows 1, not 1.6666...
END MAIN

FUNCTION int_div(n1 INTEGER, n2 INTEGER) RETURNS INTEGER
  RETURN (n1/n2)
END FUNCTION
```

## Related links

**Related concepts**  

[Type conversions](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
