---
title: "ui.DragDrop.getOperation"
source: "fgl-topics/c_fgl_ClassDragDrop_getOperation.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.getOperation"
type: "concept"
---

# ui.DragDrop.getOperation

> Identify the type of operation on drop.

## Syntax

```
getOperation()
  RETURNS STRING
```

## Usage

The `getOperation()` method returns the type of the current drag &
drop operation ("`copy`", "`move`", or "`none`").

Depending on the value returned by this method, the program can make the appropriate changes in
the data model. For example, after a row has been dropped into another list, the source list can
remove the original row if the operation was a "`move`", but keeps the original row
if the operation was a "`copy`".

The `getOperation()` method is typically called in the
`ON DRAG_FINISHED` block.

## Related links

**Related concepts**  

[Understanding drag & drop](../11_user-interface/2371-understanding-drag-drop.md "This is an introduction to drag & drop programming.")

[ui.DragDrop.setOperation](3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.")

[ui.DragDrop.addPossibleOperation](3268-ui-dragdrop-addpossibleoperation.md "Add a possible operation.")
