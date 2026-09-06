---
title: "ui.DragDrop.getLocationParent"
source: "fgl-topics/c_fgl_ClassDragDrop_getLocationParent.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.getLocationParent"
type: "concept"
---

# ui.DragDrop.getLocationParent

> Get the index of the parent node where the object was dropped.

## Syntax

```
getLocationParent()
  RETURNS INTEGER
```

## Usage

When using a [tree view](../11_user-interface/2352-tree-views.md "Describes how to implement tree views."), a node can be dropped as
a sibling or as a child node to another node. In order to distinguish between the
cases, you must use the `getLocationParent()` method, which returns the index of
the parent node of the drop target node returned by
[`getLocationRow()`](3272-ui-dragdrop-getlocationrow.md "Get the index of the target row where the object was dropped.").

If both methods return the same row index, you must append the dropped row as a child of the
target node. Otherwise, `getLocationParent()` identifies the parent node where the
dropped row has to be added as a child, and `getLocationRow()` is the index of a
sibling node. In the latter case the dropped node must be inserted before the node identified by
`getLocationRow()`.

These methods are typically used in the `ON DROP` block, but can also be used in
`ON DRAG_OVER` to deny the drop depending on the indexes returned; for example, the
program might only allow the drop of objects as new children for a given parent node.

## Related links

**Related concepts**  

[ui.DragDrop.getLocationRow](3272-ui-dragdrop-getlocationrow.md "Get the index of the target row where the object was dropped.")
