---
title: "Syntax of drag & drop interaction blocks"
source: "fgl-topics/c_fgl_DragDrop_syntax.html"
breadcrumb: "User interface > User interface programming > Drag & drop > Syntax of drag & drop interaction blocks"
type: "concept"
---

# Syntax of drag & drop interaction blocks

> The ON DRAG* / ON DROP interaction blocks implement drag & drop operations.

```
{ ON DRAG_START ( dnd-object )
| ON DRAG_FINISHED ( dnd-object )
| ON DRAG_ENTER( dnd-object )
| ON DRAG_OVER ( dnd-object )
| ON DROP ( dnd-object ) }
    dialog-statement
    [...]
```

1. dnd-object is a variable referencing an object of the class
   `ui.DragDrop`.

## Related links

**Related concepts**  

[ON DRAG\_START block](1990-on-drag-start-block.md "ON DRAG_START block")

[ON DRAG\_FINISHED block](1991-on-drag-finished-block.md "ON DRAG_FINISHED block")

[ON DRAG\_ENTER block](1992-on-drag-enter-block.md "ON DRAG_ENTER block")

[ON DRAG\_OVER block](1993-on-drag-over-block.md "ON DRAG_OVER block")

[ON DROP block](1994-on-drop-block.md "ON DROP block")
