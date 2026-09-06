---
title: "Handle drag & drop data with MIME types"
source: "fgl-topics/c_fgl_DragDrop_to_from_external_apps.html"
breadcrumb: "User interface > User interface programming > Drag & drop > Handle drag & drop data with MIME types"
type: "concept"
---

# Handle drag & drop data with MIME types

> How to handle MIME types with drag & drop?

If a drag & drop is intended to work only in the same application, data
can be passed with variables in the context of the current program. For example,
in a program using two tables where the user can drag & drop elements between
the two lists, identify the selected rows and update the program arrays accordingly.
When drag & drop is limited to the current application, avoid the drop outside
the current application.

When a drag & drop operation comes from (or goes to) external applications,
data can be of various types/formats: plain text, formatted text, documents, images,
sounds, videos, and so on. In order to handle the drag & drop data, you must
identify the type of data held in the drag & drop buffer. The type of data in
the buffer is identified by the Multipurpose Internet Mail Extensions (MIME) type.
MIME types are a widely used internet standard specification, first introduced to
identify the content of e-mail attachments.

Only text data can be passed with drag & drop; binary data is not supported.
However, you can pass files by using the [`fgl_getfile()`](../15_library-reference/2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.") file transfer function, and identify the
file with a URI (text-uri-list MIME type). For a working example, see the demos in
$FGLDIR/demo/DragAndDrop.

When using URI MIME types (for file paths for example), the string
returned from the [`getBuffer()`](../15_library-reference/3270-ui-dragdrop-getbuffer.md "Get drag & drop data from the buffer.") method can contain URL-encoded characters such as
`%5E`, which represents the `^` caret. Therefore, you must URL-decode
strings returned from `getBuffer()` with [`util.Strings.urlDecode()`](../15_library-reference/3559-util-strings-urldecode.md "Converts the URL-encoded string to a string in the current application locale."). When
setting the drag & drop buffer content, if required by the front-end platform, the string can be
URL-encoded using [`util.Strings.urlEncode()`](../15_library-reference/3560-util-strings-urlencode.md "Converts a string from the current codeset to a URL-encoded string.") for [`setBuffer()`](../15_library-reference/3276-ui-dragdrop-setbuffer.md "Set the text data of the dragged object."). However, URL-encoding
file paths for `setBuffer()` is usually not required.

Example of MIME types:

- text/plain
- text/uri-list
- text/x-vcard

You can also define your own MIME type, as long as it does not conflict with
existing standard MIME types. For example:

- text/my-remote-file
- text/my-customer-record

If you do not specify a MIME type when the drag starts, the type defaults to
text/plain, and the dialog will by default copy the data from selected rows into
the drag & drop buffer. To prevent drag & drop to external applications,
you must pass an application-specific MIME type to the [`ui.DragDrop.setMimeType()`](../15_library-reference/3278-ui-dragdrop-setmimetype.md "Define the MIME type of the dragged object.")
method, to ensure that other applications do not recognize the MIME type and
therefore reject the drop.

## Preparing the dragged object for external targets

If the program implements drag & drop of objects that can be dropped to external
programs, you must specify the MIME type of the object and copy the data to the drag
& drop buffer, so that the external application can identify the data format and
receive it.

In the `ON DRAG_START` block, you must call the
`ui.DragDrop.setMimeType()` method to define the MIME type of the object,
and copy the text data into the buffer with the [`ui.DragDrop.setBuffer()`](../15_library-reference/3276-ui-dragdrop-setbuffer.md "Set the text data of the dragged object.") method.

This example shows a `DISPLAY ARRAY` dialog preparing the drag &
drop buffer to export VCard data from a dragged row:

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
...
  ON DRAG_START(dnd)
    -- Define the MIME type and copy text data to DnD buffer 
    CALL dnd.setMimeType("text/x-vcard")
    CALL dnd.setBuffer( buildVCardData( arr[arr_curr()].cid ) )
    CALL dnd.setOperation("copy")
...
END DISPLAY
```

## Receiving the dragged object from external sources

This describes how to handle the drop action when the target dialog receives an
object dragged from an external source, by identifying the MIME type of the object.

> **Important:**
>
> Drag from external applications is not possible with GBC.

In the `ON DRAG_ENTER` block, you must call the [`ui.DragDrop.selectMimeType()`](../15_library-reference/3275-ui-dragdrop-selectmimetype.md "Select the MIME type before getting the data.")
method to check that data is available in a format identified by the MIME type, passed
as a parameter. If the type of data is available in the buffer, the method returns
`TRUE`. Later, when the dragged object is dropped (`ON DROP`),
you can get the previously selected MIME type with [`ui.DragDrop.getSelectedMimeType()`](../15_library-reference/3273-ui-dragdrop-getselectedmimetype.md "Get the previously selected MIME type.") before calling [`ui.DragDrop.getBuffer()`](../15_library-reference/3270-ui-dragdrop-getbuffer.md "Get drag & drop data from the buffer.")
to retrieve the actual data.

The next example shows the usage of those methods: In `ON DRAG_ENTER`, the program
checks available MIME types, and denies the drop operation if the buffer does not hold any of the
MIME types that can be treated by the program. In `ON DROP`, the program calls
`getSelectedMimeType()` to check what MIME type was selected, retrieves the data with
`getBuffer()`, then inserts a new row and puts the data in dedicated fields depending
on the MIME type:

```
DEFINE dnd ui.DragDrop 
...
DISPLAY ARRAY arr TO sr.* ...
...
  ON DRAG_ENTER(dnd)
    -- Set operation to NULL if unexpected MIME type found 
    CASE
    WHEN dnd.selectMimeType("text/plain")
    WHEN dnd.selectMimeType("text/uri-list")
    OTHERWISE
      CALL dnd.setOperation(NULL)
    END CASE
...
  ON DROP(dnd)
    -- Select MIME type and get data from buffer 
    LET row = dnd.getLocationRow()
    CALL DIALOG.insertRow("sr", row)
    IF dnd.getSelectedMimeType() == "text/plain" THEN
      LET arr[row].text_data = dnd.getBuffer()
    END IF
...
END DISPLAY
```

## Related links

**Related concepts**  

[The DragDrop class](../15_library-reference/3266-the-dragdrop-class.md "The ui.DragDrop class is used to control the events related to drag & drop events.")
