---
title: "ui.DragDrop.setBuffer"
source: "fgl-topics/c_fgl_ClassDragDrop_setBuffer.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.setBuffer"
type: "concept"
---

# ui.DragDrop.setBuffer

> Set the text data of the dragged object.

## Syntax

```
setBuffer(
   buffer STRING )
```

1. buffer is a string expression containing drag & drop data.

## Usage

Use the `setBuffer()` method to provide the text data of objects dragged from
the program to an external application.

The `setBuffer()` method is typically used in an `ON DRAG_START`
block together with [`setMimeType()`](3278-ui-dragdrop-setmimetype.md "Define the MIME type of the dragged object.").

By default, the dialog will serialize the data of the selected rows as a tab-separated list of
values.

The text/plain MIME type is the default.

When using URI MIME types (for file paths for example), the string
returned from the [`getBuffer()`](3270-ui-dragdrop-getbuffer.md "Get drag & drop data from the buffer.") method can contain URL-encoded characters such as
`%5E`, which represents the `^` caret. Therefore, you must URL-decode
strings returned from `getBuffer()` with [`util.Strings.urlDecode()`](3559-util-strings-urldecode.md "Converts the URL-encoded string to a string in the current application locale."). When
setting the drag & drop buffer content, if required by the front-end platform, the string can be
URL-encoded using [`util.Strings.urlEncode()`](3560-util-strings-urlencode.md "Converts a string from the current codeset to a URL-encoded string.") for [`setBuffer()`](3276-ui-dragdrop-setbuffer.md "Set the text data of the dragged object."). However, URL-encoding
file paths for `setBuffer()` is usually not required.

## Related links

**Related concepts**  

[Handle drag & drop data with MIME types](../11_user-interface/2375-handle-drag-drop-data-with-mime-types.md "How to handle MIME types with drag & drop?")
