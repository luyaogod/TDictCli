---
title: "Controlling out of bound in static arrays"
source: "fgl-topics/c_fgl_Arrays_005.html"
breadcrumb: "Language basics > Arrays > Static arrays > Controlling out of bound in static arrays"
type: "concept"
description: "By default, when an index is lower than 1 or greater than the static array length, fglrun raises error -1326 , and this error is not trappable. When using a dynamic array, new elements are ..."
---

# Controlling out of bound in static arrays

By default, when an index is lower than 1 or greater than the static array length,
fglrun raises error [-1326](../15_library-reference/4483-genero-bdl-errors.md), and this error is not
trappable.

When using a dynamic array, new elements are automatically allocated, if the index is greater
than the actual array size. The error -1326 will only occur when the index is lower than 1.

Raising an index out of bounds error is normal. However, in some situations, code must execute
without error and evaluate expressions using indexes that are greater than the size of the array,
especially with boolean expressions in `IF` statements:

```
IF index <= max_index AND arr[index].name IS NOT NULL THEN
   ...
END IF
```

In this example, by default (when not using the `OPTIONS SHORT CIRCUIT`
instruction), all parts of a boolean expression need to be evaluated, and the runtime system must
get the value of the `arr[index]` element.

If such kind of code cannot be changed and must execute, use the
`fglrun.arrayIgnoreRangeError` [FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.") entry, to control the behavior of the runtime system when an array index is out
of bounds:

```
fglrun.arrayIgnoreRangeError = true
```

When this FGLPROFILE entry is set to true, the runtime system will return the first element of
the array if the index is out of bounds, and continues with the normal program flow.

Unless existing code is relying on the relaxed behavior, it is better to keep the default
settings and get array out of bounds errors when the array index is invalid.

Change the code as follows, to detatch the index checking from its usage:

```
IF index <= max_index THEN
    IF arr[index].name IS NOT NULL THEN
       ...
    END IF
END IF
```

Or, use the [`OPTIONS SHORT
CIRCUIT`](../09_advanced-features/0923-controlling-semantics-of-and-or-operators.md "The OPTIONS SHORT CIRCUIT defines the semantics of AND/OR operators.") compiler directive, to control boolean expression evaluation. You can then
put both conditions in the same `IF` statement. When the first condition is false,
the second condition is not
evaluated:

```
OPTIONS SHORT CIRCUIT
...
IF index <= max_index AND arr[index].name IS NOT NULL THEN
   ...
END IF
```

## Related links

**Related concepts**  

[Dynamic arrays](0734-dynamic-arrays.md "Dynamic arrays")
