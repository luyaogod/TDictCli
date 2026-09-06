---
title: "Passing static arrays as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_static_array.html"
breadcrumb: "Advanced features > Runtime stack > Passing static arrays as parameter"
type: "concept"
description: "It is possible to pass a complete static array as a function parameter, but this is not recommended. When passing a static array to a function, the complete array is copied on the stack and every ..."
---

# Passing static arrays as parameter

It is possible to pass a complete static array as a function parameter, but this is not
recommended. When passing a static array to a function, the complete array is copied on the stack
and every element is passed by value. The receiving local variables in the function must be defined
with the same static array definition as the caller:

```
MAIN
  DEFINE arr ARRAY[5] OF INTEGER
  LET arr[3] = 111
  CALL func(arr)
  DISPLAY arr[3]  -- shows 111 (local array unchanged)
END MAIN

-- function defining same static array as the caller 
FUNCTION func(p_arr ARRAY[5] OF INTEGER) RETURNS ()
  DISPLAY p_arr[3]  -- shows 111
  LET p_arr[3] = 999
END FUNCTION
```

Consider using dynamic arrays: these are passed by reference. It is more efficient, and arrays
can be changed in called functions.

## Related links

**Related concepts**  

[Static arrays](../08_language-basics/0732-static-arrays.md "Static arrays have a predefined and limited size.")
