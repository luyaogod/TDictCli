---
title: "ON FILL BUFFER block"
source: "fgl-topics/c_fgl_dialog_ON_FILL_BUFFER_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG data blocks > ON FILL BUFFER block"
type: "concept"
description: "Syntax ON FILL BUFFER instruction [...] Usage The ON FILL BUFFER block is used to fill a page of rows into the dynamic array, based on an offset and a number of rows. This data block is only used in ..."
---

# ON FILL BUFFER block

## Syntax

```
ON FILL BUFFER
   instruction [...]
```

## Usage

The `ON FILL BUFFER` block is used to fill a page of rows into the dynamic array,
based on an offset and a number of rows.

This data block is only used in `DISPLAY
ARRAY` dialog blocks.

The `ON FILL BUFFER` block is executed when the runtime system needs data rows to
fill the current page of the list dialog. This can happen before a `BEFORE DISPLAY`
of a singular `DISPLAY ARRAY`, or before the `BEFORE DIALOG` block of
a `DIALOG` / `END DIALOG` instruction containing `DISPLAY
ARRAY` sub-dialogs.

The offset can be retrieved with the [`fgl_dialog_getbufferstart()`](../15_library-reference/2750-fgl-dialog-getbufferstart.md "Returns the row offset of the page to feed a paged display array.") built-in function and the number of rows to
provide is defined by the [`fgl_dialog_getbufferlength()`](../15_library-reference/2751-fgl-dialog-getbufferlength.md "Returns the number of rows to feed a paged DISPLAY ARRAY.") built-in function.

When entering the `ON FILL BUFFER` block, the dynamic array is already cleared by
the runtime system.

```
    ON FILL BUFFER
       LET ofs = fgl_dialog_getBufferStart()
       LET len = fgl_dialog_getBufferLength()
       LET row = ofs
       FOR i=1 TO len
          FETCH ABSOLUTE row c1 INTO arr[i].*
          IF sqlca.sqlcode==NOTFOUND THEN
            EXIT FOR
          END IF
          LET row = row + 1
       END FOR
```

For more details about `ON FILL BUFFER` usage, see [Paged mode of DISPLAY ARRAY](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.").
