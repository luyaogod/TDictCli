---
title: "Dynamic arrays"
source: "fgl-topics/c_fgl_MigI4GL_023.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Dynamic arrays"
type: "concept"
---

# Dynamic arrays

> Support for dynamic arrays differs between I4GL and Genero BDL.

Both IBM® Informix®
4GL (I4GL) and Genero Business Development Language (BDL) implement static arrays with a fixed
size. Static arrays cannot be extended:

```
DEFINE arr ARRAY[100] OF RECORD LIKE customer.*
```

I4GL introduced dynamic arrays in version 7.32. Unlike Genero BDL, I4GL requires explicitly to
associate memory storage with a dynamic array by using the `ALLOCATE ARRAY`
statement, and memory must be freed with `DEALLOCATE ARRAY`. I4GL dynamic arrays can
be resized with the `RESIZE ARRAY` statement. I4GL dynamic arrays cannot be used in a
interactive instructions such as `DISPLAY
ARRAY`.

```
DEFINE arr DYNAMIC ARRAY OF RECORD LIKE customer.*
ALLOCATE ARRAY arr[10]
RESIZE ARRAY arr[100]
LET arr[50].cust_name = "Smith"
DEALLOCATE ARRAY arr
```

Genero BDL supports dynamic arrays in a slightly different way
than I4GL. There are no allocation, resizing, or deallocation instructions,
because the memory for element storage is automatically allocated
when needed. Furthermore, you can use dynamic arrays with interactive
instructions, making a `DISPLAY ARRAY` or `INPUT
ARRAY` unlimited.

```
DEFINE arr DYNAMIC ARRAY OF RECORD LIKE customer.*
LET arr[50].cust_name = "Smith"
DISPLAY ARRAY arr TO sr.*
```

In Genero BDL, the main difference between static arrays and dynamic
arrays is the memory usage; when you use dynamic arrays, elements
are allocated on demand. With static arrays, memory is allocated
for the complete array when the variable is created.

> **Important:**
>
> The semantics of dynamic arrays is very similar to static arrays, but there
> are some small differences. Keep in mind that the runtime system automatically allocates a new
> element for a dynamic array when needed. For example, when a `DISPLAY arr[100].*` is
> executed with a dynamic array, the element at index 100 is automatically created if does not
> exist.

## Related links

**Related concepts**  

[Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
