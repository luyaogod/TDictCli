---
title: "Copying and assigning arrays"
source: "fgl-topics/c_fgl_Arrays_011.html"
breadcrumb: "Language basics > Arrays > Copying and assigning arrays"
type: "concept"
---

# Copying and assigning arrays

> Arrays can be fully copied or assigned by reference.

## Prerequisites to copy and assign arrays

To make a copy of a static or dynamic array, or when assigning a reference of a dynamic array to
another variable, the source and destination arrays must have the same structure.

## The `copyTo()` method

An array can be copied to another array with the [`copyTo()`](../15_library-reference/2947-dynamic-array-copyto.md "Copies a complete array to the destination array passed as parameter.") method. This method makes a full copy of all elements of the
source array:

```
MAIN
  DEFINE left, right DYNAMIC ARRAY OF RECORD
             key INTEGER
         END RECORD
  LET left[1].key = 123
  CALL left.copyTo(right) -- copies the array
  DISPLAY right[1].key    -- shows 123
  LET right[1].key = 456
  DISPLAY left[1].key     -- shows 123
END MAIN
```

The `copyTo()` method can be used with dynamic and static arrays.

Depending on the type of the array element fields, the copy is done by value or by reference. For
more details, see the reference page of the [`copyTo()`](../15_library-reference/2947-dynamic-array-copyto.md "Copies a complete array to the destination array passed as parameter.") method.

## Assigning dynamic array references

The reference of a dynamic array can be copied to another variable defined with the same type, by
specifying the name of the array, as if it was a simple data type.

After assigning a dynamic array to another variable, both variables will reference the same
dynamic array object and elements:

```
MAIN
  DEFINE left, right DYNAMIC ARRAY OF RECORD
             key INTEGER
         END RECORD
  LET left[1].key = 123
  LET right = left        -- copies the reference
  DISPLAY right[1].key    -- shows 123
  LET right[1].key = 456
  DISPLAY left[1].key     -- shows 456
END MAIN
```

## The `.*` notation with static arrays

For backward compatibility, the compiler allows the `.*` notation to copy static
arrays using the same structure.
> **Important:**
>
> Copying static arrays with the
> `.*` notation is supported for backward compatibility. Consider using the array names
> without the `.*` suffix: `LET arr2 = arr1`, this syntax has the same
> effect with static arrays. Consider also using dynamic arrays instead of static arrays.

When using the `.*` notation, static array elements are copied by value
(except objects and `BYTE`/`TEXT` members):

```
MAIN
  DEFINE left, right ARRAY[10] OF RECORD
             key INTEGER
         END RECORD
  LET left[1].key = 123
  LET right.* = left.*    -- copies the static array
  DISPLAY right[1].key    -- shows 123
  LET right[1].key = 456
  DISPLAY left[1].key     -- shows 123
END MAIN
```

When using the `.*` notation with dynamic arrays, the reference to the dynamic
array is copied.
