---
title: "ui.DragDrop.getSelectedMimeType"
source: "fgl-topics/c_fgl_ClassDragDrop_getSelectedMimeType.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.getSelectedMimeType"
type: "concept"
---

# ui.DragDrop.getSelectedMimeType

> Get the previously selected MIME type.

## Syntax

```
getSelectedMimeType()
  RETURNS STRING
```

## Usage

Before retrieving data from the drag & drop buffer with [`getBuffer()`](3270-ui-dragdrop-getbuffer.md "Get drag & drop data from the buffer."), first call the
`getSelectedMimeType()` method, to identify the data format that was previously
selected by a [`selectMimeType()`](3275-ui-dragdrop-selectmimetype.md "Select the MIME type before getting the data.") call.

The `getSelectedMimeType()` method is typically called in `ON
DROP` to identity the format of the dropped object.

## Related links

**Related concepts**  

[Handle drag & drop data with MIME types](../11_user-interface/2375-handle-drag-drop-data-with-mime-types.md "How to handle MIME types with drag & drop?")
