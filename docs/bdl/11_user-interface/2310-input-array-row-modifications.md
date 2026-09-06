---
title: "INPUT ARRAY row modifications"
source: "fgl-topics/c_fgl_prog_dialogs_insdel_rows.html"
breadcrumb: "User interface > User interface programming > List dialogs > INPUT ARRAY row modifications"
type: "concept"
---

# INPUT ARRAY row modifications

> Controlling row creation and deletion in an editable record list.

The `INPUT ARRAY` instruction handles record list editing. This controller allows
the user to directly edit existing rows and to create or remove rows with automatic actions.

The following automatic actions are created by default by the `INPUT ARRAY`
dialog:

- `insert`: creates a new row before the current row. If there are no rows in the
  list, the action adds a new one.
- `append`: creates a new row after the last row of the list.
- `delete`: deletes the current row.

To prevent `INPUT ARRAY` from creating the automatic "insert", "append" and
"delete" actions, set respectively the `INSERT ROW`, `APPEND ROW`, or
`DELETE ROW` [control
attributes](2092-input-array-attributes-clause.md "INPUT ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.") to `FALSE`. To prevent row addition, also set the `AUTO
APPEND` attribute to `FALSE`.

```
...
  INPUT ARRAY p_items FROM sa.*
    -- Allow only row append and delete actions.
    ATTRIBUTES(AUTO APPEND=FALSE,
               INSERT ROW=FALSE)
     ...
  END INPUT
...
```

Specific control blocks are available to take control when a row
is created or deleted:

- `BEFORE INSERT` and `AFTER INSERT` control blocks can be used to
  control row creation. Cancel a row creation with [`CANCEL INSERT`](2037-cancel-insert-instruction.md) in
  `BEFORE INSERT` or `AFTER INSERT` blocks.
- `BEFORE DELETE` and `AFTER DELETE` control blocks can be used to
  control row deletion. Cancel row deletion with the [`CANCEL DELETE`](2036-cancel-delete-instruction.md)
  instruction in `BEFORE DELETE`.

Dynamic arrays and the `ui.Dialog` class provide
methods such as `array.deleteElement()` or `ui.Dialog.appendRow()` to
modify the list. When using these methods, the predefined triggers
such as `BEFORE DELETE` or `BEFORE INSERT` are
not executed. While it is safe to use these methods within a `DISPLAY
ARRAY`, you must take care when using an `INPUT
ARRAY`. For example, it is not recommended to call such methods in triggers
like `BEFORE ROW`, `AFTER INSERT`,
`BEFORE DELETE`.

Users can append [temporary rows](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") by moving to
the end of the list, or when executing the append action. Appending temporary rows is different from
inserting a row; an appended row is considered temporary until the user modifies a field while an
inserted row remains in the list even if the user does not modify a field.

By default, when the last row is removed by a delete action, the `INPUT ARRAY`
instruction will automatically create a new temporary row at the same position. The visual effect of
this behavior can be misinterpreted - if no data were entered in the last row, you cannot see any
difference. However, the last row is actually deleted and a new row is created, and the
`BEFORE DELETE` / `AFTER DELETE` / `AFTER ROW` /
`BEFORE ROW` / `BEFORE INSERT` control block sequence is executed. In
order to avoid the creation of a new temporary row when the last row is deleted, set `AUTO
APPEND = FALSE` attribute.

The insert, append or delete actions will be automatically disabled depending on the context. If
the `INPUT ARRAY` is using a static array that becomes full, or if the
`MAXCOUNT` attribute is reached, both insert and append actions will be disabled. The
delete action is automatically disabled when `AUTO APPEND = FALSE` and there are no
more rows in the array.

## Related links

**Related concepts**  

[The INPUT ARRAY sub-dialog](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.")

[INPUT ARRAY ATTRIBUTES clause](2092-input-array-attributes-clause.md "INPUT ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")

[The DIALOG control class](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.")
