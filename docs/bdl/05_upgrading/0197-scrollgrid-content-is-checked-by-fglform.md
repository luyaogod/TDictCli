---
title: "SCROLLGRID content is checked by fglform"
source: "fgl-topics/c_fgl_Migrate_to_310_scrollgrid_check.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > SCROLLGRID content is checked by fglform"
type: "concept"
---

# SCROLLGRID content is checked by fglform

> When using a SCROLLGRID, fglform compiler checks that it does not hold other list containers.

Before version 3.10, [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") did not give an error when using
`TABLE`, `TREE` or `SCROLLGRID` as sub-elements in a
`SCROLLGRID` container, for example:

```
SCROLLGRID gr2
{
 <TABLE tb4   >
 [ab8        ]
 }
END
```

This layout construction is illegal: It can not be rendered by the GUI and causes undefined
front-end behavior or crash.

Starting with 3.10, fglform gives a compilation error [-6846](../15_library-reference/4483-genero-bdl-errors.md) , if
`SCROLLGRID` contains elements that are not valid in this type of container.

## Related links

**Related concepts**  

[SCROLLGRID container](../11_user-interface/1723-scrollgrid-container.md "Defines a scrollable grid view widget.")
