---
title: "Control block execution order"
source: "fgl-topics/c_fgl_DragDrop_control_block_execution_order.html"
breadcrumb: "User interface > User interface programming > Drag & drop > Control block execution order"
type: "concept"
---

# Control block execution order

> What is the execution order of drag & drop related dialog control blocks?

The table below shows the order in which the runtime system executes the control blocks
related to drag & drop events:

| Context / User action | Control block execution order |
| --- | --- |
| The user starts to drag an object from the source dialog. | [`ON DRAG_START`](1990-on-drag-start-block.md) (in source dialog) |
| The mouse cursor enters the drop target dialog. | [`ON DRAG_ENTER`](1992-on-drag-enter-block.md) (in target dialog) |
| After entering the target dialog, the mouse cursor moves from row to row, or user chooses to change the drop operation (move or copy). | [`ON DRAG_OVER`](1993-on-drag-over-block.md) (in target dialog) |
| The user releases the mouse button over the target dialog. | [`ON DROP`](1994-on-drop-block.md) (in target dialog)`ON DRAG_FINISHED` (in source dialog) |

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")
