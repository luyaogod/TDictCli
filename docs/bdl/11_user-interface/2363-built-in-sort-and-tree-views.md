---
title: "Built-in sort and tree-views"
source: "fgl-topics/c_fgl_treeviews_011.html"
breadcrumb: "User interface > User interface programming > Tree views > Built-in sort and tree-views"
type: "concept"
---

# Built-in sort and tree-views

> Build-in sort on tree-views can be disabled if needed.

By default, the built-in sort is enabled in a `TREE` container;
when the end user clicks on column headers, the runtime system
sorts the visual representation of the program array. Tree nodes
are ordered by levels; the children nodes are ordered inside a given
parent node.

This is a powerful built-in feature. However, in some cases, the tree structure must be static
(the order of the nodes must not change) and you don't want the end user to sort the rows. To
prevent the built-in sort, use the `UNSORTABLECOLUMNS` attribute for the
`TREE` container
definition:

```
LAYOUT
...
END
ATTRIBUTES
TREE tv: mytree, UNSORTABLECOLUMNS, ...
...
```

## Related links

**Related concepts**  

[UNSORTABLECOLUMNS attribute](1837-unsortablecolumns-attribute.md "The UNSORTABLECOLUMNS attribute indicates that the columns of the table cannot be selected by the user for sorting.")

[Sorting rows in a list](2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required.")
