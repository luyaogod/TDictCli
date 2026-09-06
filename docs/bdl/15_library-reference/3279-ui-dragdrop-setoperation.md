---
title: "ui.DragDrop.setOperation"
source: "fgl-topics/c_fgl_ClassDragDrop_setOperation.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.setOperation"
type: "concept"
---

# ui.DragDrop.setOperation

> Define the type of Drag & Drop operation.

## Syntax

```
setOperation(
   operation STRING )
```

1. operation is the name of a drag & drop operation.

## Usage

Drag & drop actions can be of different kinds; you can do a copy of the dragged object, or
move the dragged object from the source to the destination.

Use the `setOperation()` method to define/force the type of drag & drop
operation or to deny/cancel the drag & drop process.

The [`addPossibleOperation()`](3268-ui-dragdrop-addpossibleoperation.md "Add a possible operation.") method can be used to specify additional drag &
drop operations that are possible in this dialog.

| Parameter Value | Description |
| --- | --- |
| `NULL` | To deny/cancel the drag & drop process. |
| `copy` | To allow drag & drop as a copy of the source object. |
| `move` | To allow drag & drop as a move of the source object. |

The `setOperation()` method can be called in different drag & drop
triggers.

A common usage is to deny drag & drop by passing [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.") in the `ON DRAG_ENTER` and/or `ON
DRAG_OVER` blocks because the dragged object does not correspond to the type of objects
the target can receive.

This method is also used in `ON DRAG_START` to force a specific type of drag &
drop operation (copy or move), or to deny drag start if the context does not allow a drag &
drop action.

When called in the `ON DRAG_ENTER` block, the method forces a specific drag &
drop operation.

## Related links

**Related concepts**  

[Understanding drag & drop](../11_user-interface/2371-understanding-drag-drop.md "This is an introduction to drag & drop programming.")
