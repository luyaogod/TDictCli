---
title: "ui.DragDrop.setMimeType"
source: "fgl-topics/c_fgl_ClassDragDrop_setMimeType.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.setMimeType"
type: "concept"
---

# ui.DragDrop.setMimeType

> Define the MIME type of the dragged object.

## Syntax

```
setMimeType(
   mimeType STRING )
```

1. mimeType defines the MIME type for the drag & drop buffer.

## Usage

Objects dragged from the program to an external application need to be identified with a MIME
type and the program must provide the data. The MIME type can be specified with the
`setMimeType()` method.

The `setMimeType()` method is typically used in an `ON DRAG_START`
block along with [`setBuffer()`](3276-ui-dragdrop-setbuffer.md "Set the text data of the dragged object.").

By default, the source target will use the text/plain MIME type and copy the data of the
selected rows into the Drag & Drop buffer.

## Related links

**Related concepts**  

[Handle drag & drop data with MIME types](../11_user-interface/2375-handle-drag-drop-data-with-mime-types.md "How to handle MIME types with drag & drop?")
