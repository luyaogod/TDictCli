---
title: "Runtime stack"
source: "fgl-topics/c_fgl_runtime_stack.html"
breadcrumb: "Advanced features > Runtime stack"
type: "concept"
---

# Runtime stack

> The runtime stack is used to pass/return values to/from functions.

When passing arguments to a function or when returning values from a function, you are using the
runtime stack. When you call a function, parameters are pushed on the stack;
before the function code executes, parameters are popped from the stack in the local
variables defined in the function. On the other hand, each parameter returned by a function
is pushed on the stack and popped into variables specified in the `RETURNING`
clause of the caller.

Elements are pushed on the stack in a given order, then popped from the stack in the reverse
order. This is transparent to the programmer. However, if you want to implement a C extension,
you must keep this in mind.

According to the data type, parameters are passed and returned by value or by
reference. When an element is passed/returned by value, a complete copy of the value
is passed. When an element is passed by reference, only the handle of the object is
passed/returned. If the type allows it, elements passed by reference can be manipulated in
the called function to modify the value.

| Mode | Data type or data structure |
| --- | --- |
| By value | `BOOLEAN`, `BIGINT`, `INTEGER` , `SMALLINT`, `TINYINT`, `FLOAT` , `SMALLFLOAT`, `DECIMAL`, `MONEY`, `CHAR`, `VARCHAR`, `DATE`, `DATETIME`, `INTERVAL`, records (by default, expanded) and static arrays (cannot be returned). |
| By reference | Dynamic arrays, dictionaries, objects (from Java, built-in or extension classes), `BYTE`/`TEXT`, `STRING` (but cannot be modified), records (in methods, and/or with `INOUT`) |

## Child topics

- [Passing simple typed values as parameter](0834-passing-simple-typed-values-as-parameter.md)
- [Passing records as parameter](0835-passing-records-as-parameter.md)
- [Passing static arrays as parameter](0836-passing-static-arrays-as-parameter.md)
- [Passing dynamic arrays as parameter](0837-passing-dynamic-arrays-as-parameter.md)
- [Passing dictionaries as parameter](0838-passing-dictionaries-as-parameter.md)
- [Passing objects as parameter](0839-passing-objects-as-parameter.md)
- [Passing TEXT/BYTE values as parameter](0840-passing-text-byte-values-as-parameter.md)
- [Returning simple typed values from functions](0841-returning-simple-typed-values-from-functions.md)
- [Returning records from functions](0842-returning-records-from-functions.md)
- [Returning dynamic arrays from functions](0843-returning-dynamic-arrays-from-functions.md)
- [Returning dictionaries from functions](0844-returning-dictionaries-from-functions.md)
- [Returning TEXT/BYTE values from functions](0845-returning-text-byte-values-from-functions.md)
- [Implicit data type conversion on the stack](0846-implicit-data-type-conversion-on-the-stack.md)
- [Various type specification](0847-various-type-specification.md): Some Genero APIs use variant types for parameters or returns.
