---
title: "Default drag & drop operation"
source: "fgl-topics/c_fgl_DragDrop_default_op.html"
breadcrumb: "User interface > User interface programming > Drag & drop > Default drag & drop operation"
type: "concept"
---

# Default drag & drop operation

> DISPLAY ARRAY dialogs implement a default drag operation.

The `DISPLAY ARRAY` dialog provides a default drag operation, that copies all
selected rows to the drag & drop buffer, as a tab-separated list of values.

The user code equivalent to the default drag & drop operation would look like this:

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
...
  ON DRAG_START(dnd)
    CALL dnd.setOperation("copy")
    CALL dnd.setMimeType("text/plain")
    CALL dnd.setBuffer(DIALOG.selectionToString("sr"))
...
END DISPLAY
```

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")
