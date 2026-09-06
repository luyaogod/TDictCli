---
title: "Passing dynamic arrays as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_dynamic_array.html"
breadcrumb: "Advanced features > Runtime stack > Passing dynamic arrays as parameter"
type: "concept"
description: "Passing a dynamic array as a function parameter is legal and efficient. When passed as parameter, the runtime system pushes a reference of the dynamic array on the stack, and the receiving local ..."
---

# Passing dynamic arrays as parameter

Passing a dynamic array as a function parameter is legal and efficient. When passed as
parameter, the runtime system pushes a reference of the dynamic array on the stack, and the
receiving local variables in the function can then manipulate the original data.

Returning a dynamic array from a function is also possible: The runtime system pushes the
reference of the dynamic array on the stack.

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF INT
  DISPLAY arr.getLength()
  LET arr = init(10)
  DISPLAY arr.getLength()
  CALL modify(arr)
  DISPLAY arr[50]
  DISPLAY arr[51]
  DISPLAY arr.getLength()
END MAIN

FUNCTION init(c INTEGER) RETURNS DYNAMIC ARRAY OF INTEGER
  DEFINE x DYNAMIC ARRAY OF INT
  DEFINE i INTEGER
  FOR i=1 TO c
     LET x[i] = i
  END FOR
  RETURN x
END FUNCTION

FUNCTION modify(x DYNAMIC ARRAY OF INTEGER) RETURNS ()
  LET x[50] = 222
  LET x[51] = 333
END FUNCTION
```

Output of the program:

```
          0
         10
        222
        333
         51
```

## Related links

**Related concepts**  

[Dynamic arrays](../08_language-basics/0734-dynamic-arrays.md "Dynamic arrays")
