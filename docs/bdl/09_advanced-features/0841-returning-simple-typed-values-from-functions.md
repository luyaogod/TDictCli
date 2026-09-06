---
title: "Returning simple typed values from functions"
source: "fgl-topics/c_fgl_runtime_stack_return_simple.html"
breadcrumb: "Advanced features > Runtime stack > Returning simple typed values from functions"
type: "concept"
description: "Simple data types such as INTEGER , DECIMAL , VARCHAR are returned by value with the RETURN instruction. When returning a simple typed value, the runtime system pushes a copy of the data on the stack. ..."
---

# Returning simple typed values from functions

Simple data types such as `INTEGER`, `DECIMAL`,
`VARCHAR` are returned by value with the [`RETURN`](../08_language-basics/0768-returning-values.md "A function can return values with the RETURN instruction.") instruction.

When returning a simple typed value, the runtime system pushes a copy of the data on the stack.
The `STRING` data type is an exception to this rule: elements of this type are return
by mutable reference: the whole string value is not copied on the stack, only the reference to the
string value is copied.

When using the `RETURNS` clause in the function definition, the returned value is
typed as specified in the `RETURNS` clause.

```
MAIN
  DISPLAY int_add(10,20)
  DISPLAY int_div(5,3)
END MAIN

FUNCTION int_add(n1 INTEGER, n2 INTEGER) RETURNS INTEGER
  RETURN (n1+n2)
END FUNCTION

FUNCTION int_div(n1 INTEGER, n2 INTEGER) RETURNS INTEGER
  RETURN (n1/n2)
END FUNCTION
```

Output:

```
         30
          1
```

## Related links

**Related concepts**  

[Type conversions](../08_language-basics/0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
