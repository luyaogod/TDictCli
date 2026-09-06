---
title: "ui.DragDrop.selectMimeType"
source: "fgl-topics/c_fgl_ClassDragDrop_selectMimeType.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.selectMimeType"
type: "concept"
---

# ui.DragDrop.selectMimeType

> Select the MIME type before getting the data.

## Syntax

```
selectMimeType(
   mimeType STRING )
```

1. mimeType defines the MIME type for dragged objects.

## Usage

Call the `selectMimeType()` method to check that data is available in a format
identified by the MIME type passed as parameter.

If this type of data is available in the buffer, the method returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") and you can later get the data with
[`getBuffer()`](3270-ui-dragdrop-getbuffer.md "Get drag & drop data from the buffer.").

The `selectMimeType()` method is typically used in `ON DRAG_ENTER`,
`ON DRAG_OVER` to deny the drag & drop operation if none of the supported MIME
types is available in the buffer.

The [`getSelectedMimeType()`](3273-ui-dragdrop-getselectedmimetype.md "Get the previously selected MIME type.") method can be used to identify which MIME type has
been selected.

## Related links

**Related concepts**  

[Handle drag & drop data with MIME types](../11_user-interface/2375-handle-drag-drop-data-with-mime-types.md "How to handle MIME types with drag & drop?")
