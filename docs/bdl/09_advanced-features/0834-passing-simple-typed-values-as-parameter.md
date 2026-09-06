---
title: "Passing simple typed values as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_simple.html"
breadcrumb: "Advanced features > Runtime stack > Passing simple typed values as parameter"
type: "concept"
description: "Simple data types such as INTEGER , DECIMAL , VARCHAR are passed by value in function parameters. When passing a function parameter by value, the runtime system pushes a copy of the data on the stack. ..."
---

# Passing simple typed values as parameter

Simple data types such as `INTEGER`, `DECIMAL`,
`VARCHAR` are passed by value in function parameters. When passing a
function parameter by value, the runtime system pushes a copy of the data on the stack.

The `STRING` data type is an exception to this rule for simple types:
elements of this type are passed by reference. In fact the runtime system passes a reference
to the string value, so the actual string data is not copied on the stack as for other
simple types. However, the value of the caller cannot be modified. If a
`STRING` parameter gets a new value in a function, a new string reference
is created. Passed `STRING` parameters improve performances compared to
`CHAR`/`VARCHAR`, with the same semantics as
`VARCHAR()`.

When passing a simple typed value to a function, the local variable receiving the value can
be changed without affecting the variable used by the caller:

```
MAIN
  DEFINE c CHAR(10), s STRING
  LET c = "abc"
  LET s = "def"
  CALL func(c,s)
  DISPLAY c -- Shows "abc"
  DISPLAY s -- Shows "def"
END MAIN

FUNCTION func(pc CHAR(10), ps CHAR(10)) RETURNS ()
  DISPLAY pc -- Shows "abc" (this is a copy of the string)
  DISPLAY ps -- Shows "def" (this is the original string)
  LET pc = "zz" -- Assigns new value to local variable
  LET ps = "zz" -- Assigns new value to local variable
END FUNCTION
```

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Type conversions](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
