---
title: "ON DRAG_START block"
source: "fgl-topics/c_fgl_dialog_ON_DRAG_START_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG interaction blocks > ON DRAG_START block"
type: "concept"
description: "Syntax ON DRAG_START ( dnd-object ) instruction [...] Usage The ON DRAG_START block is executed when the end user begins the drag operation. If this dialog trigger has not been defined, default ..."
---

# ON DRAG_START block

## Syntax

```
ON DRAG_START (dnd-object)
   instruction [...]
```

## Usage

The `ON DRAG_START` block is executed when the end user begins the drag operation.
If this dialog trigger has not been defined, default dragging is enabled for this dialog.

In the `ON DRAG_START` block, the program typically specifies the type of drag
& drop operation by calling [`ui.DragDrop.setOperation()`](../15_library-reference/3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.") with "move" or "copy". This call will define the
default and unique drag operation. If needed, the program can allow another type of drag operation
with [`ui.DragDrop.addPossibleOperation()`](../15_library-reference/3268-ui-dragdrop-addpossibleoperation.md "Add a possible operation."). The end user can then choose to move or
copy the dragged object, if the drag & drop target allows it.

If the dragged object can be dropped outside the program, the MIME type and drag/drop data must
be defined with [`ui.DragDrop.setMimeType()`](../15_library-reference/3278-ui-dragdrop-setmimetype.md "Define the MIME type of the dragged object.") and [`ui.DragDrop.setBuffer()`](../15_library-reference/3276-ui-dragdrop-setbuffer.md "Set the text data of the dragged object.")
methods.

Example:

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DRAG_START (dnd)
    CALL dnd.setOperation("move") -- Move is the default operation 
    CALL dnd.addPossibleOperation("copy") -- User can toggle to copy if needed 
    CALL dnd.setMimeType("text/plain")
    CALL dnd.setBuffer(arr[arr_curr()].cust_name)
    ...
END DISPLAY
```

## Related links

**Related concepts**  

[Handle drag & drop data with MIME types](2375-handle-drag-drop-data-with-mime-types.md "How to handle MIME types with drag & drop?")

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")

[Drag & drop](2370-drag-drop.md "Explains programming techniques for the drag & drop feature.")
