---
title: "Multiple row selection"
source: "fgl-topics/c_fgl_prog_dialogs_multirange.html"
breadcrumb: "User interface > User interface programming > List dialogs > Multiple row selection"
type: "concept"
---

# Multiple row selection

> Multiple row selection allows the end user to select several rows within a list of records.

The `DISPLAY ARRAY` controller supports multiple row selection when the
`ON SELECTION CHANGE` block is defined, or by enabling the feature
with the [`ui.Dialog.setSelectionMode()`](../15_library-reference/3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") method when the dialog starts. The
`setSelectionMode()` method can also be used to enable or disable the multi-row
selection during the dialog execution.

> **Important:**
>
> This feature is not supported on mobile platforms.

When multi-row selection is enabled, the end user can select one or several rows with the
standard keyboard and mouse click combinations. When the end user selects or deselects rows, the
`ON SELECTION CHANGE` block is fired, if defined. The program can then query the
[`DIALOG.isRowSelected()`](../15_library-reference/3206-ui-dialog-isrowselected.md "Queries row selection for a given list and row.")
method to check for selected rows.

```
DISPLAY ARRAY arr TO sr.*
  ...
  ON SELECTION CHANGE
     LET s = "Selected rows:"
     FOR x=1 TO DIALOG.getArrayLength("list1")
         IF DIALOG.isRowSelected("list1",x) THEN
            LET s = s, x USING " <<<&"
         END IF
     END FOR
     DISPLAY s
  ON ACTION enable_mrs
     CALL DIALOG.setSelectionMode( "sr", 1 )
  ON ACTION disable_mrs
     CALL DIALOG.setSelectionMode( "sr", 0 )
  ...
END DISPLAY
```

Multiple row selection is GUI-specific and therefore cannot be used in TUI mode.

With multiple row selection, you must distinguish between two concepts: row selection
and current row. In GUI mode, a selected row usually has a blue background, while the
current row has a dotted focus rectangle. The current row may not be selected, or a selected row
may not be the current row. When the default single-row selection is used, the current row is always
selected automatically.

If the `ON SELECTION CHANGE` block is not required to detect multi-row selection,
use the `ui.Dialog.setSelectionMode()` method to enable multi-row selection for the
dialog:

```
DISPLAY ARRAY arr TO sr.*
  BEFORE DISPLAY
     CALL DIALOG.setSelectionMode( "sr", 1 )
  ...
END DISPLAY
```

Note that without the `ON SELECTION CHANGE` trigger, it is not possible to detect
row selection changes when staying on the current row, since no `BEFORE ROW` /
`AFTER ROW` trigger is fired in this case.

Row selection flags can be changed by program for a range of rows with the [`DIALOG.setSelectionRange()`](../15_library-reference/3232-ui-dialog-setselectionrange.md "Sets the row selection flags for a range of rows.")
method.

The `DISPLAY ARRAY` dialog implements an implicit row-copy feature. The
selected rows can be dragged to another dialog or external program, or the end-user can do an
"editcopy" predefined action (Ctrl-C shortcut), to copy the selected rows to the front-end
clipboard. The row-copy feature works also when multiple row selection is disabled, but only the
current row will be dragged or copied to the front-end clipboard.

If you delete, insert or append rows in the program array with methods such as
`array.deleteElement()`, selection information is not synchronized. To sync the
selection flags with the data rows, use dialog methods like `DIALOG.insertRow()` (or
`DIALOG.insertNode()` for tree-views).

## Behavior of ui.Dialog class methods with multiple row selection

| Dialog class method | Effect on multiple row selection |
| --- | --- |
| [`appendRow()`](../15_library-reference/3185-ui-dialog-appendrow.md "Appends a new row in the specified list.") | Selection flags of existing rows are unchanged.New row is appended at the end of the list with selection flag set to zero. |
| [`appendNode()`](../15_library-reference/3186-ui-dialog-appendnode.md "Appends a new node in the specified tree-view.") | Selection flags of existing rows are unchanged.New node is appended at the end of the tree with selection flag set to zero. |
| [`deleteAllRows()`](../15_library-reference/3190-ui-dialog-deleteallrows.md "Deletes all rows from the specified list.") | Selection flags of all rows are cleared. |
| [`deleteRow()`](../15_library-reference/3192-ui-dialog-deleterow.md "Deletes a row from the specified list.") | Selection flags of existing rows are unchanged.Selection information is synchronized (i.e., shifted up) for all rows after the deleted row. |
| [`deleteNode()`](../15_library-reference/3191-ui-dialog-deletenode.md "Deletes a node from the specified tree-view.") | Selection flags of existing rows are unchanged.Selection information is synchronized (i.e., shifted up) for all nodes after the deleted node. |
| [`insertRow()`](../15_library-reference/3205-ui-dialog-insertrow.md "Inserts a new row in the specified list.") | Selection flags of existing rows are unchanged.Selection information is synchronized (i.e., shifted down) for all rows after the new inserted row. |
| [`insertNode()`](../15_library-reference/3204-ui-dialog-insertnode.md "Inserts a new node in the specified tree.") | Selection flags of existing rows are unchanged.Selection information is synchronized (i.e., shifted down) for all nodes after the new inserted node. |
| [`setArrayLength()`](../15_library-reference/3220-ui-dialog-setarraylength.md "Sets the number of rows in a DISPLAY ARRAY using paged mode.") | Selection flags of existing rows are unchanged.If the new array length is larger than the previous length, selection flags of new rows are not initialized to zero. |
| [`setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") | Selection flags of all rows are reset, and the new current row gets selected. |
| [`setSelectionMode()`](../15_library-reference/3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.") | When you switch off multiple row selection, the selection flags of existing rows are cleared. |

## Related links

**Related concepts**  

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
