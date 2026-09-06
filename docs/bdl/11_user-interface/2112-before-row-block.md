---
title: "BEFORE ROW block"
source: "fgl-topics/c_fgl_dialog_BEFORE_ROW_3.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > BEFORE ROW block"
type: "concept"
description: "Syntax BEFORE ROW instruction [...] BEFORE ROW block in singular DISPLAY ARRAY, INPUT ARRAY dialogs In a singular DISPLAY ARRAY or INPUT ARRAY instruction, the BEFORE ROW block is executed each time ..."
---

# BEFORE ROW block

## Syntax

```
BEFORE ROW
   instruction [...]
```

## BEFORE ROW block in singular DISPLAY ARRAY, INPUT ARRAY dialogs

In a singular `DISPLAY ARRAY` or `INPUT ARRAY` instruction, the
`BEFORE ROW` block is executed each time the user moves to another row. This trigger
can also be executed in other situations, such as when you delete a row, or when the user tries to
insert a row but the maximum number of rows in the list is reached.

You typically do some dialog setup / message display in the `BEFORE ROW`
block, because it indicates that the user selected a new row or entered in the list.

When the dialog starts, `BEFORE ROW` will be executed for the current row,
but only if there are data rows in the array.

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") / [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") return the
index of the current row.

In this example, the `BEFORE ROW` block gets the new row number and
displays it in a message:

```
DISPLAY ARRAY ...
   ...
   BEFORE ROW
     MESSAGE "We are on row # ", arr_curr()
   ...
```

## BEFORE ROW block in DISPLAY ARRAY and INPUT ARRAY of procedural DIALOG

In an `INPUT` or `INPUT ARRAY` sub-dialog of a procedural
`DIALOG` instruction, the `BEFORE ROW` block is executed when a
`DISPLAY ARRAY` or `INPUT ARRAY` list gets the focus, or when
the user moves to another row inside a list. This trigger can also be executed in other
situations, for example when you delete a row, or when the user tries to insert a row but
the maximum number of rows in the list is reached.

You typically do some dialog setup / message display in the `BEFORE ROW`
block, because it indicates that the user selected a new row. Do not use this trigger to
detect focus changes; Use the `BEFORE DISPLAY` or `BEFORE INPUT`
blocks instead.

In `DISPLAY ARRAY`, `BEFORE ROW` is executed after
the `BEFORE DISPLAY` block. In `INPUT ARRAY`, `BEFORE
ROW` is executed before the `BEFORE INSERT` and `BEFORE
FIELD` blocks and after the `BEFORE INPUT` blocks.

When the procedural dialog starts, `BEFORE ROW` will only be executed if
the list has received the focus and there is a current row (the array is not empty). If you
have other elements in the form which can get the focus before the list, `BEFORE
ROW` will not be triggered when the dialog starts. You must pay attention to this,
because this behavior is different to the behavior of singular [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") or [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."). In singular dialogs,
the `BEFORE ROW` block is always executed when the dialog starts (and when
there are rows in the array).

When called in this block, [`DIALOG.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.")
/ [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.")
return the index of the current row.

In this example the `BEFORE ROW` block displays a message with the current
row number:

```
   DISPLAY ARRAY p_items TO s_items.*
     BEFORE ROW
       MESSAGE "We are in items, on row #", DIALOG.getCurrentRow("s_items")
```

## Related links

**Related concepts**  

[BEFORE INPUT block](1940-before-input-block.md "BEFORE INPUT block")

[BEFORE DISPLAY block](1973-before-display-block.md "BEFORE DISPLAY block")
