---
title: "Returning dynamic arrays from functions"
source: "fgl-topics/c_fgl_runtime_stack_return_dynamic_array.html"
breadcrumb: "Advanced features > Runtime stack > Returning dynamic arrays from functions"
type: "concept"
description: "When returned by a function, dynamic arrays are pushed on the stack by reference. Therefore, you can create a dynamic array in a function and return it to the caller for usage: MAIN DEFINE arr DYNAMIC ..."
---

# Returning dynamic arrays from functions

When returned by a function, [dynamic arrays](../08_language-basics/0734-dynamic-arrays.md) are pushed
on the stack by reference.

Therefore, you can create a dynamic array in a function and return it to the caller for
usage:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  LET arr = create_array(10)
  DISPLAY arr.getLength()
END MAIN

FUNCTION create_array(n INTEGER) RETURNS DYNAMIC ARRAY OF INTEGER
  DEFINE i INTEGER
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  FOR i=1 TO n
     LET arr[i] = i
  END FOR
  RETURN arr
END FUNCTION
```

## Related links

**Related concepts**  

[RETURN](../08_language-basics/0676-return.md "The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.")
