---
title: "ui.DragDrop.setFeedback"
source: "fgl-topics/c_fgl_ClassDragDrop_setFeedback.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.setFeedback"
type: "concept"
---

# ui.DragDrop.setFeedback

> Define the appearance of the target during Drag & Drop.

## Syntax

```
setFeedback(
   feedback STRING )
```

1. feedback is the type of feedback to display during the drag &
   drop operation.

## Usage

The `setFeedback()` method defines the appearance the target object must have
during the drag & drop process.

For example, in a table or tree view, when the mouse is flying over rows in the drop target, a
different visual indicator will appear depending on the value that was passed to
`setFeedback()`.

Possible values for the `setFeedback()` method are:

| Parameter Values | Description |
| --- | --- |
| `all` | Dragged object will be dropped somewhere on the target widget, the exact location does not matter. |
| `insert` | In lists, dragged object will be inserted in between existing rows. |
| `select` | In lists, dragged object will replace the current row under the mouse. |

## Related links

**Related concepts**  

[Understanding drag & drop](../11_user-interface/2371-understanding-drag-drop.md "This is an introduction to drag & drop programming.")
