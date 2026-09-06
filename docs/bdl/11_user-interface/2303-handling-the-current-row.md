---
title: "Handling the current row"
source: "fgl-topics/c_fgl_prog_dialogs_current_row.html"
breadcrumb: "User interface > User interface programming > List dialogs > Handling the current row"
type: "concept"
---

# Handling the current row

> Query and control the current row in a read-only or editable list of records.

## Get the current row

To query the current row of a list, use either the [`ui.Dialog.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") method
or the [`arr_curr()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") built-in
function, depending on the context.

The `getCurrentRow()` method can be used for a `DISPLAY ARRAY` or
`INPUT ARRAY` dialog. The method takes the name of the screen array as the argument
to identify the list.

For example, when implementing a `DIALOG` block with two `DISPLAY
ARRAY` subdialogs, you can query the current row of a list in the code block of the other
list
controller:

```
DIALOG ...
   DISPLAY ARRAY arr1 TO sa1.*
      ON ACTION check
         IF arr2[DIALOG.getCurrentRow("sa2")].value > 0 THEN
            ...
         END IF
   END DISPLAY
   DISPLAY ARRAY arr2 TO sa2.*
   END DISPLAY
END DIALOG
```

The `arr_curr()` function must be used in the context of the current
`DISPLAY ARRAY` or `INPUT ARRAY` dialog, or just after executing such
a dialog.

For example, when implementing modification triggers in a `DISPLAY ARRAY` dialog,
the current row and the current screen line can be queried respectively with the
`arr_curr()` and `scr_line()`
functions:

```
DISPLAY ARRAY arr TO sa.*
    ON UPDATE
       INPUT arr[arr_curr()].* WITHOUT DEFAULTS FROM sa[scr_line()].* ;
END DISPLAY
```

After the dialog execution, `arr_curr()` returns the current row index for the
last executed dialog, until a new list dialog is started.

The row index returned by the `arr_curr()` function is constant in the context of
a dialog block, even when removing rows from the array by program. A typical mistake is to reuse the
`arr_curr()` index to get data from the new current row, after deleting the last row
of the array.

In the next code example, the reuse of `arr_curr()` without checking for the new
number of rows will automatically create a new program array element when accessing the element in
the `MESSAGE`
instruction:

```
ON ACTION dialog_delete_row
    CALL DIALOG.deleteRow("sr",arr_curr())
    MESSAGE "Current item:", arr[arr_curr()].name
```

The above code works until you reach the last row: When last row is deleted, the
`MESSAGE` instruction is automatically creating a new array element at the same index
returned by `arr_curr()`.

Unlike `arr_curr()`, the `DIALOG.getCurrentRow()` method is
synchronized with the actual number of rows in the array, as long as methods like [`DIALOG.deleteRow()`](../15_library-reference/3192-ui-dialog-deleterow.md "Deletes a row from the specified list.") are
used:

```
ON ACTION dialog_delete_row
    CALL DIALOG.deleteRow("sr",DIALOG.getCurrentRow("sr"))
    MESSAGE "Current item:", arr[DIALOG.getCurrentRow("sr")].name
```

However, the code must also test if there are still rows in the list after deleting a
row:

```
ON ACTION dialog_delete_row
    LET x = DIALOG.getCurrentRow("sr")
    CALL DIALOG.deleteRow("sr",x)
    LET x = DIALOG.getCurrentRow("sr")
    IF x > 0 THEN
       MESSAGE "Current item:", arr[x].name
    ELSE
       MESSAGE "No more rows in the list"
    END IF
```

## Set the current row

To set the current row in a list controlled by a `DISPLAY ARRAY` or
`INPUT ARRAY`, use the [`ui.Dialog.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") method. This method takes the name of the screen
array and the new row index as
parameters:

```
DISPLAY ARRAY p_items TO sa.*
   ...
   ON ACTION next_empty 
      LET row = findEmptyRow(p_items)
      CALL DIALOG.setCurrentRow("sa", row)
   ...
END DISPLAY
```

Calling the `DIALOG.setCurrentRow()` method will not execute control blocks such
as `BEFORE ROW` and `AFTER ROW`, and will not set the focus. If you
want to set the focus to the list, you must use the `NEXT FIELD` instruction. This
works with `DISPLAY ARRAY` as well as `INPUT ARRAY`.
> **Tip:**
>
> Use this method with care. Let the dialog handle normal navigation automatically,
> and jump to a specific row only in the context of an `ON ACTION` block.

The [`fgl_set_arr_curr()`](../15_library-reference/2782-fgl-set-arr-curr.md "Moves to a specific row in a record list.") function can also be used. This function must be called
in the context of the current list having the focus.

`fgl_set_arr_curr()` triggers control blocks such as `BEFORE ROW`,
while `DIALOG.setCurrentRow()` does not trigger any control blocks.

In a `DISPLAY ARRAY` using [paged
mode](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.") with `COUNT=-1`, before calling `DIALOG.setCurrentRow(
screen-array, row-index )`, call [`DIALOG.setArrayLength(
screen-array, count )`](../15_library-reference/3220-ui-dialog-setarraylength.md "Sets the number of rows in a DISPLAY ARRAY using paged mode.") where count >=
row-index. Otherwise, the `setCurrentRow()` call will have no effect, if
the dialog has not yet seen row-index rows through `ON FILL
BUFFER`.

## Converting visual index to/from program array index

When the end user sorts rows in a table, the program array index (`arr_curr()`)
may differ from the visual row index (the row position as seen by the user).

The `ui.Dialog` class provides methods to convert between these contexts:

The [ui.Dialog.arrayToVisualIndex](../15_library-reference/3187-ui-dialog-arraytovisualindex.md "Converts the program array index to the visual index for a given screen array.") method converts a program array
index to a visual index. It can be used, for example, to display a typical list position message
(Row: current-row / total-rows). The current row
(arr\_curr()/getCurrentRow()) is a program array index that
must be converted to a visual index. Note that you need to display such messages in the
`BEFORE ROW` trigger and `ON SORT` trigger:

```
FUNCTION disp_row(d,n)
   DEFINE d ui.DIALOG, n STRING
   MESSAGE SFMT("Row: %1/%2",
                d.arrayToVisualIndex(n,d.getCurrentRow(n)),
                d.getArrayLength(n))
END FUNCTION
...
   DISPLAY ARRAY arr TO sr.*
       ...
       BEFORE ROW
           CALL disp_row(DIALOG,"sr")
       ON SORT
           CALL disp_row(DIALOG,"sr")
       ...
   END DISPLAY
```

The [ui.Dialog.visualToArrayIndex](../15_library-reference/3234-ui-dialog-visualtoarrayindex.md "Converts the visual index to the program array index for a given screen array.") method converts a visual index
to a program array index. It can be used for example to ask the user for a row position (visual
index), and make that row current by using `DIALOG.setCurrentRow()` after
converting to the program array index:

```
DEFINE i INTEGER
...
DISPLAY ARRAY arr TO sr.*
   ...
   ON ACTION move_to
      PROMPT "Enter row index:" FOR i
      CALL DIALOG.setCurrentRow( "sr", DIALOG.visualToArrayIndex("sr", i))
   ...
END DISPLAY
```

## Related links

**Related concepts**  

[Record list (DISPLAY ARRAY)](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")

[Editable record list (INPUT ARRAY)](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")

[The DISPLAY ARRAY sub-dialog](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.")

[The INPUT ARRAY sub-dialog](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.")

[Sorting rows in a list](2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required.")
