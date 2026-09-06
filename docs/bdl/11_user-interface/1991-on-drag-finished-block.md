---
title: "ON DRAG_FINISHED block"
source: "fgl-topics/c_fgl_dialog_ON_DRAG_FINISHED.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY interaction blocks > ON DRAG_FINISHED block"
type: "concept"
description: "Syntax ON DRAG_FINISHED ( dnd-object ) instruction [...] Usage Execution of the ON DRAG_FINISHED block notifies the dialog where the drag started and that the drop operation has been completed or ..."
---

# ON DRAG_FINISHED block

## Syntax

```
ON DRAG_FINISHED (dnd-object)
   instruction [...]
```

## Usage

Execution of the `ON DRAG_FINISHED` block notifies the dialog where the drag
started and that the drop operation has been completed or terminated.

Call [`ui.DragDrop.getOperation()`](../15_library-reference/3274-ui-dragdrop-getoperation.md "Identify the type of operation on drop.") to get the final type of operation of the drop.
On successful completion, the method returns "move" or "copy"; otherwise the function returns
`NULL`. If `NULL` is returned, the `ON DRAG_FINISHED`
trigger can be ignored.

In cases of successful moves to a target out of the current `DISPLAY ARRAY`, the
application must remove the transferred data from the source model. For example, if a row was moved
from dialog A to B, dialog A will get an `ON DRAG_FINISHED` execution after the row
was dropped into B, which removes the row from the list A.

The `ON DRAG_FINISHED` interaction block is optional.

```
DEFINE dnd ui.DragDrop
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DRAG_START (dnd)
    LET last_dragged_row = arr_curr()
    ...
  ON DRAG_FINISHED (dnd)
    IF dnd.getOperation() == "move" THEN
      CALL DIALOG.deleteRow("sr",last_dragged_row)
    END IF
    ...
END DISPLAY
```

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")

[Drag & drop](2370-drag-drop.md "Explains programming techniques for the drag & drop feature.")
