---
title: "ui.DragDrop.getLocationRow"
source: "fgl-topics/c_fgl_ClassDragDrop_getLocationRow.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.getLocationRow"
type: "concept"
---

# ui.DragDrop.getLocationRow

> Get the index of the target row where the object was dropped.

## Syntax

```
getLocationRow()
  RETURNS INTEGER
```

## Usage

The `getLocationRow()` method returns the index
of the row in the drop target list pointed to by the mouse cursor.

This method is typically used in the `ON DROP` block to get the index of
the target row to be modified or replaced by the dragged object.

In order to deny the drop, the `getLocationRow()` can be used in conjunction with
[`setOperation(NULL)`](3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.") in the
`ON DRAG_ENTER` or `ON DRAG_OVER`, if the current target row returned
by `getLocationRow()` is not valid for a drop operation.

## Related links

**Related concepts**  

[ui.DragDrop.getLocationParent](3271-ui-dragdrop-getlocationparent.md "Get the index of the parent node where the object was dropped.")
