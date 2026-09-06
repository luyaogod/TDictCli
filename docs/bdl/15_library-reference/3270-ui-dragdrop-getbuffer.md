---
title: "ui.DragDrop.getBuffer"
source: "fgl-topics/c_fgl_ClassDragDrop_getBuffer.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods > ui.DragDrop.getBuffer"
type: "concept"
---

# ui.DragDrop.getBuffer

> Get drag & drop data from the buffer.

## Syntax

```
getBuffer()
  RETURNS STRING
```

## Usage

After identifying the MIME type of a dropped object with [`getSelectedMimeType()`](3273-ui-dragdrop-getselectedmimetype.md "Get the previously selected MIME type."),
you can call the `getBuffer()` method to get text data from the drag & drop
buffer.

Drag & drop data is only available at `ON DROP` time, therefore, the
`getBuffer()` method must be called in `ON DROP` only.

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
