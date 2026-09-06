---
title: "Understanding drag & drop"
source: "fgl-topics/c_fgl_DragDrop_basics.html"
breadcrumb: "User interface > User interface programming > Drag & drop > Understanding drag & drop"
type: "concept"
---

# Understanding drag & drop

> This is an introduction to drag & drop programming.

Drag & drop is a well-known feature of graphical applications, allowing the end user to use
the mouse to drag an element of a window to another window in the same program or into an external
application. The front-end platform/device must support this feature.

> **Important:**
>
> This feature is not supported on mobile platforms.

Drag & drop can be implemented in regular tables and tree-views controlled by a singular
`DISPLAY ARRAY` or a `DISPLAY ARRAY` sub-dialog within a
`DIALOG` instruction. Drag & drop is not supported in other dialog contexts,
such as a singular `INPUT`, `INPUT ARRAY` or
`CONSTRUCT`.

![Form with drag & drop screenshot](../_images/DragAndDrop01_gbc.jpg)

*Form with Drag & Drop*

With drag & drop, end users can:

- Move drag-able objects between lists and tree-views in the same Genero form or program.
- Move drag-able objects between lists and tree-views in different Genero forms and programs.
- Move drag-able objects between other desktop applications and tables / tree-views in Genero
  programs.

Drag & drop control is implemented in a `DISPLAY ARRAY` with specific
interaction blocks, to handle the events related to the drag and drop operation. These specific
blocks will be triggered when drag and drop events arrive from the front-end.

- [`ON DRAG_START`](1990-on-drag-start-block.md)
- `ON DRAG_FINISHED`
- [`ON DRAG_ENTER`](1992-on-drag-enter-block.md)
- [`ON DRAG_OVER`](1993-on-drag-over-block.md)
- [`ON DROP`](1994-on-drop-block.md)

Each of these interaction blocks takes an `ui.DragDrop` object as a parameter. A
reference variable to that object must be declared before the dialog. In the interaction block, the
`ui.DragDrop` object can be used to configure the drag & drop action to take. For
example, a "drag enter" event can be refused.

The `ON DRAG_START` and `ON DRAG_FINISHED` triggers apply to the
source of the drag & drop operation; the dialog where the object was dragged. The other triggers
provide notification to the drop target dialog, used to inform the program when the different drop
events occur and to let the target accept or reject the drop action.

This example illustrates the use of a drag & drop interaction block with the
`ui.DragDrop` control object:

```
DEFINE dnd ui.DragDrop
...
DISPLAY ARRAY arr TO sr.* ...
...
    ON DRAG_ENTER(dnd)
      IF ok_to_drop THEN
        CALL dnd.setOperation("move")
      ELSE
        CALL dnd.setOperation(NULL)
      END IF
...
END DISPLAY
```

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")
