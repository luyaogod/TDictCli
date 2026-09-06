---
title: "ui.DragDrop.addPossibleOperation"
source: "fgl-topics/c_fgl_ClassDragDrop_addPossibleOperation.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.addPossibleOperation"
type: "concept"
---

# ui.DragDrop.addPossibleOperation

> Add a possible operation.

## Syntax

```
addPossibleOperation(
   operation STRING )
```

1. operation is the name of a drag & drop operation.

## Usage

Drag & drop actions can be of different kinds; you can do a copy of the dragged object, or
move the dragged object from the source to the destination.

The default drag & drop operation is defined by a call to `setOperation()`
method in `ON DRAG_START`. Use the `addPossibleOperation()` method
to define additional operations that are allowed.

See [setOperation()](3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.") for
possible values.

## Related links

**Related concepts**  

[Understanding drag & drop](../11_user-interface/2371-understanding-drag-drop.md "This is an introduction to drag & drop programming.")
