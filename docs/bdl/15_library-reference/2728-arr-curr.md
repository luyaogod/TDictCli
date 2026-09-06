---
title: "arr_curr()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ARR_CURR.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > arr_curr()"
type: "concept"
---

# arr_curr()

> Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.

## Syntax

```
FUNCTION arr_curr()
  RETURNS INTEGER
```

## Usage

The `arr_curr()` function returns an integer value that identifies the current row
of a list of rows in an `INPUT ARRAY` or `DISPLAY ARRAY`
instruction.

The first row is numbered 1.

> **Important:**
>
> The row index returned by the `arr_curr()` function is
> constant in the context of a dialog block, even when removing rows from the array by program.

The `arr_curr()` and [`scr_line()`](2730-scr-line.md "Returns the index of the current row in the form screen array.") functions can return different values if the program array is
larger than the screen array.

The `arr_curr()` can be used to get the index of the last current row, after the
execution of the dialog.

In multiple dialogs implementing several list controllers, consider using the [`ui.Dialog.getCurrentRow()`](3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") method
instead of `arr_curr()` to get the current row of a specific list identified by its
screen array.

## Related links

**Related concepts**  

[Handling the current row](../11_user-interface/2303-handling-the-current-row.md "Query and control the current row in a read-only or editable list of records.")
