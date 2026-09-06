---
title: "ON DRAG_ENTER block"
source: "fgl-topics/c_fgl_dialog_ON_DRAG_ENTER_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON DRAG_ENTER block"
type: "concept"
description: "Syntax ON DRAG_ENTER ( dnd-object ) instruction [...] Usage When the ON DROP control block is defined, the ON DRAG_ENTER block will be executed when the mouse cursor enters the visual boundaries of ..."
---

# ON DRAG_ENTER block

## Syntax

```
ON DRAG_ENTER (dnd-object)
   instruction [...]
```

## Usage

When the `ON DROP` control block is
defined, the `ON DRAG_ENTER` block will be executed when the mouse cursor enters the
visual boundaries of the drop target dialog. Entering the target dialog is accepted by default if no
`ON DRAG_ENTER` block is defined. However, when `ON DROP` is defined,
it is recommended that you also define `ON DRAG_ENTER` to prevent the drop of objects
with an unsupported MIME type coming from other applications.

The program can decide to disallow or allow a specific drop operation with a call to [`ui.DragDrop.setOperation()`](../15_library-reference/3279-ui-dragdrop-setoperation.md "Define the type of Drag & Drop operation.");
passing a `NULL` to the method will prevent the drop.

To check what MIME type is available in the drag & drop buffer, the program uses the [`ui.DragDrop.selectMimeType()`](../15_library-reference/3275-ui-dragdrop-selectmimetype.md "Select the MIME type before getting the data.")
method. This method takes the MIME type as a parameter and returns `TRUE` if the
passed MIME type is used. You can call this method several times to check the availability of
different MIME types.

You may also define the visual effect when hovering over the target list with [`ui.DragDrop.setFeedback()`](../15_library-reference/3277-ui-dragdrop-setfeedback.md "Define the appearance of the target during Drag & Drop.").

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DRAG_ENTER (dnd)
    IF dnd.selectMimeType("text/plain") THEN
      CALL dnd.setOperation("copy")
      CALL dnd.setFeedback("all")
    ELSE
      CALL dnd.setOperation(NULL)
    END IF
  ON DROP (dnd)
    ...
END DISPLAY
```

Once the mouse has entered the target area, subsequent mouse cursor moves can be detected with
the [`ON DRAG_OVER`](1993-on-drag-over-block.md) trigger.

When using a table or tree-view as drop target, you can control the visual effect when the mouse
moves over the rows, depending on the type of drag & drop you want to achieve.

Basically, a dragged object can be:

1. Inserted in between two rows (visual effect must show where the object will be inserted)
2. Copied/merged to the current row (visual effect must show the row under the mouse)
3. Dropped somewhere on the target widget (the exact location inside the widget does not
   matter)

The visual effect can be defined with the `ui.DragDrop.setFeedback()` method,
typically called in the `ON DRAG_ENTER` block.

The values to pass to the `setFeedback()` method to get the desired visual effects
described are respectively:

1. `insert` (default)
2. `select`
3. `all`

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
  ...
  ON DRAG_ENTER (dnd)
    IF canDrop() THEN
      CALL dnd.setOperation(NULL)
    ELSE
      CALL dnd.setFeedback("select")
    END IF
    ...
END DISPLAY
```

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")

[Drag & drop](2370-drag-drop.md "Explains programming techniques for the drag & drop feature.")
