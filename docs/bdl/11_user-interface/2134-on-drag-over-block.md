---
title: "ON DRAG_OVER block"
source: "fgl-topics/c_fgl_dialog_ON_DRAG_OVER_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG interaction blocks > ON DRAG_OVER block"
type: "concept"
description: "Syntax ON DRAG_OVER ( dnd-object ) instruction [...] Usage When the ON DROP control block is defined, the ON DRAG_OVER block will be executed after ON DRAG_ENTER , when the mouse cursor is moving over ..."
---

# ON DRAG_OVER block

## Syntax

```
ON DRAG_OVER (dnd-object)
   instruction [...]
```

## Usage

When the `ON DROP` control block is
defined, the `ON DRAG_OVER` block will be executed after `ON
DRAG_ENTER`, when the mouse cursor is moving over the drop target, or when the drag &
drop operation has changed (toggling copy/move).

`ON DRAG_OVER` will be called only once per row, even if the mouse cursor moves
over the row.

In the `ON DRAG_OVER` block, the method [`ui.DragDrop.getLocationRow()`](../15_library-reference/3272-ui-dragdrop-getlocationrow.md "Get the index of the target row where the object was dropped.")
returns the index of the row in the target array, and can be used to allow or disallow the drop.
When using a tree-view, you must also check the index returned by the [`ui.DragDrop.getLocationParent()`](../15_library-reference/3271-ui-dragdrop-getlocationparent.md "Get the index of the parent node where the object was dropped.") method to detect if the object was dropped
as a sibling or as a child node, and allow/disallow the drop operation
accordingly.

The program can change the drop operation at any execution of the `ON DRAG_OVER`
block. You can disallow or allow a specific drop operation with a call to [`ui.DragDrop.setOperation()`](../15_library-reference/3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.");
passing a `NULL` to the method will disallow the drop.

The current operation (returned by [`ui.DragDrop.getOperation()`](../15_library-reference/3274-ui-dragdrop-getoperation.md "Identify the type of operation on drop.")) is the value set in previous `ON
DRAG_ENTER` or `ON DRAG_OVER` events, or the operation selected by the end
user, if it can toggle between copy and move. Thus, `ON DRAG_OVER` can occur even if
the mouse position has not changed.

If dropping has been prevented with `ui.DragDrop.setOperation(NULL)` in the
previous `ON DRAG_OVER` event, the program can reset the operation to allow a drop
with a call to `ui.DragDrop.setOperation()` with the operation parameter "move" or
"copy".

`ON DRAG_OVER` will not be called if drop has been disabled in `ON
DRAG_ENTER` with `ui.DragDrop.setOperation(NULL)`

`ON DRAG_OVER` is optional, and must only be defined if the operation or the
acceptance of the drag object depends on the target row of the drop target.

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DRAG_ENTER (dnd)
    ...
  ON DRAG_OVER (dnd)
    IF arr[dnd.getLocationRow()].acceptsCopy THEN
      CALL dnd.setOperation("copy")
    ELSE
      CALL dnd.setOperation(NULL)
    END IF
  ON DROP (dnd)
    ...
END DISPLAY
```

During a drag & drop process, the end user (or the target application) can decide to modify
the type of the operation, to indicate whether the dragged object has to be copied or moved from the
source to the target. For example, in a typical file explorer, by default files are moved when doing
a drag & drop on the same disk. To make a copy of a file, you must press the Ctrl key while
doing the drag & drop with the mouse.

In the drop target dialog, you can detect such operation changes in the `ON
DRAG_OVER` trigger and query the `ui.DragDrop` object for the current
operation with `ui.DragDrop.getOperation()`. In the drag source dialog, you typically
check `ui.DragDrop.getOperation()` in the `ON DRAG_FINISHED` trigger
to know what type of operation occurred, to keep ("copy" operation) or delete ("move" operation) the
original dragged object.

This example tests the current operation in the drop target list and displays a message
accordingly:

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DRAG_ENTER (dnd)
    ...
  ON DRAG_OVER (dnd)
    CASE dnd.getOperation()
    WHEN "move"
      MESSAGE "The object will be moved to row ", dnd.getLocationRow()
    WHEN "copy"
      MESSAGE "The object will be copied to row ", dnd.getLocationRow()
    END CASE
    ...
  ON DROP (dnd)
    ...
END DISPLAY
```

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")

[Tree views](2352-tree-views.md "Describes how to implement tree views.")

[Drag & drop](2370-drag-drop.md "Explains programming techniques for the drag & drop feature.")
