---
title: "ON DROP block"
source: "fgl-topics/c_fgl_dialog_ON_DROP_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON DROP block"
type: "concept"
description: "Syntax ON DROP ( dnd-object ) instruction [...] Usage To enable drop actions on a list, you must define the ON DROP block; otherwise the list will not accept drop actions. The ON DROP block is ..."
---

# ON DROP block

## Syntax

```
ON DROP (dnd-object)
   instruction [...]
```

## Usage

To enable drop actions on a list, you must define the `ON DROP` block; otherwise
the list will not accept drop actions.

The `ON DROP` block is executed after the end user has released the mouse button
to drop the dragged object. `ON DROP` will not occur if drop has been disallowed in
the previous `ON DRAG_OVER` event or in `ON DRAG_ENTER` with a call to
[`ui.DragDrop.setOperation(NULL)`](../15_library-reference/3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.").

When `ON DROP` executes, the MIME type of the dragged object can be checked with
[`ui.DragDrop.getSelectedMimeType()`](../15_library-reference/3273-ui-dragdrop-getselectedmimetype.md "Get the previously selected MIME type."). Then call the [`ui.DragDrop.getBuffer()`](../15_library-reference/3270-ui-dragdrop-getbuffer.md "Get drag & drop data from the buffer.") method to
retrieve drag & drop data from external applications.

Ideally, the drop operation is accepted (there is no need for additional calls to [`ui.DragDrop.setOperation()`](../15_library-reference/3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.")).

In this block, the [`ui.DragDrop.getLocationRow()`](../15_library-reference/3272-ui-dragdrop-getlocationrow.md "Get the index of the target row where the object was dropped.") method returns the index of the row in the
target array, and can be used to execute the code to get the drop data / object into the row that
has been chosen by the user.

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DROP (dnd)
     LET arr[dnd.getLocationRow()].capacity == dnd.getBuffer() 
    ...
END DISPLAY
```

If the drag & drop operations are local to the same list or tree-view controller, you can use
the [`ui.DragDrop.dropInternal()`](../15_library-reference/3269-ui-dragdrop-dropinternal.md "Perform built-in row drop in trees.") method to simplify the code. This method
implements the typical move of the dragged rows or tree-view node. This is especially useful in case
of a tree-view, but is also the preferred way to move rows around in simple tables.

This `ON DROP` code example uses the `dropInternal()`
method:

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr_tree TO sr_tree.* ...
  ...
  ON DROP (dnd)
     CALL dnd.dropInternal()
    ...
END DISPLAY
```

If you want to implement by hand the code to drop a node in a tree-view, you must check the index
returned by the [`ui.DragDrop.getLocationParent()`](../15_library-reference/3271-ui-dragdrop-getlocationparent.md "Get the index of the parent node where the object was dropped.") method to detect if the object was dropped
as a sibling or as a child node, and execute the code corresponding to the drop
operation. If the drop target row index returned by `getLocationRow()` is a child of
the parent row index returned by `getLocationParent()`, the new row must be inserted
before `getLocationRow()`; otherwise the new row must be added as a child of the
parent node identified by `getLocationParent()`.

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")

[Drag & drop](2370-drag-drop.md "Explains programming techniques for the drag & drop feature.")
