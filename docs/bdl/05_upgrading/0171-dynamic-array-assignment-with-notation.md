---
title: "Dynamic array assignment with .* notation"
source: "fgl-topics/c_fgl_Migrate_to_320_dyn_arr_assign.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Dynamic array assignment with .* notation"
type: "concept"
---

# Dynamic array assignment with .* notation

> The .* notation to assign dynamic arrays is discouraged.

Before Genero BDL version 3.20, the syntax to assign a dynamic array reference to another dynamic
array variable was done with the `.*`
notation:

```
LET dynarr2.* = dynarr1.*
```

Starting with version 3.20, the proper syntax to assign a dynamic array reference to another
dynamic array (or to copy a static array to another static array), is to use the array variable
names without the `.*` suffix:

```
LET dynarr2 = dynarr1
```

> **Important:**
>
> The `.*` notation to assign arrays is still supported for
> backward compatibility, to be able to compile legacy code. There is no need to change existing code.
> Use the new syntax without `.*` in new development.

## Related links

**Related concepts**  

[Copying and assigning arrays](../08_language-basics/0738-copying-and-assigning-arrays.md "Arrays can be fully copied or assigned by reference.")
