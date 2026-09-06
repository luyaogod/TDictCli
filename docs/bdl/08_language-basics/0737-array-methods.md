---
title: "Array methods"
source: "fgl-topics/c_fgl_Arrays_003.html"
breadcrumb: "Language basics > Arrays > Array methods"
type: "concept"
---

# Array methods

> Native BDL arrays and Java arrays can be used to invoke built-in methods.

For example, to clear a dynamic array, use the `clear()`
method:

```
DEFINE arr DYNAMIC ARRAY OF ...
CALL arr.clear()
```

> **Note:**
>
> While designed for dynamic arrays, these methods can also be used on static arrays. However, the
> semantics differ when applied to static arrays; these variations are outside the scope of this
> documentation.

For the list of native array methods, see [DYNAMIC ARRAY methods](../15_library-reference/2944-dynamic-array-methods.md).

For the list of Java array methods, see [Java Array type methods](../15_library-reference/2964-java-array-type-methods.md).
