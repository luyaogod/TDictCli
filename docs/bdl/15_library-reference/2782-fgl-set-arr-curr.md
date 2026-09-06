---
title: "fgl_set_arr_curr()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SET_ARR_CURR.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_set_arr_curr()"
type: "concept"
---

# fgl_set_arr_curr()

> Moves to a specific row in a record list.

## Syntax

```
FUNCTION fgl_set_arr_curr(
   row INTEGER )
```

1. row is the row number is the program array variable.

## Usage

This function is used to control navigation in a `DISPLAY ARRAY` or `INPUT
ARRAY`, within an `ON ACTION` or `ON KEY` block. The function
can also be used inside `BEFORE DISPLAY` or `BEFORE INPUT` blocks, to
jump to a specific row when the dialog starts. It is not recommended to use this function in an
other context.

Control blocks like `BEFORE ROW` and field/row validation in `INPUT
ARRAY` are executed, as if the user moved to another row, except when the function is called
in `BEFORE DISPLAY` or `BEFORE INPUT`.

When a new row is reached using this function, the first editable field gets the focus.

> **Tip:**
>
> Since `fgl_set_arr_curr()` triggers control blocks, this can make the dialog more
> difficult to implement, as it chains several dialog blocks executing additional code which can in
> turn again call this function, resulting in an infinite loop. To set the current row from the
> program code, consider using [`DIALOG.setCurrentRow()`](3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") instead of `fgl_set_arr_curr()`, to
> avoid the implicit execution of control blocks.

## Related links

**Related concepts**  

[Editable record list (INPUT ARRAY)](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")

[Record list (DISPLAY ARRAY)](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")

[arr\_curr()](2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.")
