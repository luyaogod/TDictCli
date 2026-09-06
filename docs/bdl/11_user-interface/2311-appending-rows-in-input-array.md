---
title: "Appending rows in INPUT ARRAY"
source: "fgl-topics/c_fgl_prog_dialogs_temp_row.html"
breadcrumb: "User interface > User interface programming > List dialogs > Appending rows in INPUT ARRAY"
type: "concept"
---

# Appending rows in INPUT ARRAY

> Rows appended at the end of an editable list are temporary until they are edited.

In an `INPUT ARRAY`, a new row can be created at the end of the list. This new row
is called "temporary" because it will be automatically removed if the user leaves the row without
entering data. If data is entered by the user or by program (setting the [touched flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.")), the temporary row becomes
permanent.

A temporary row is promoted to a permanent row under certain conditions
described in this topic. We distinguish also *explicit* temporary
row creation from *automatic* temporary row creation.

Temporary row creation is different from adding new rows with the
`DIALOG.appendRow()` method; when appending a row by program, the row is considered
permanent and remains in the list even if the user did not enter data in fields.

## Explicit temporary row creation

Explicit temporary row creation takes place when the user decides to append a new row explicitly
with the `append` action. If the list is empty, an insert action will have the same
effect as an `append` action (a temporary row is created at position 1).

## Automatic temporary row creation

By default, automatic temporary row creation takes place when `AUTO APPEND` is
`TRUE` (default) and one of the following occurs:

- The user tries to move below the last row, with a Down keystroke or with the mouse.
- The user presses the Tab key when in the last field of the last row.
- The last row of the list is deleted by the user.
- The list has the focus and the last row of the list is deleted by program with
  `DIALOG.deleteRow()` or `DIALOG.deleteAllRows()`.
- When the `INPUT ARRAY` is in a `DIALOG` block, the list has no
  rows and gets the focus (A new temporary row is created to let the user enter data immediately)

## When is a temporary row removed?

The temporary row will be automatically removed if none of the fields has been touched (the
modification flag is not set), and when leaving the list (INPUT ARRAY) or by moving upward in the
list.

The row is "touched" when the user enters data in a field, or when the program simulates a user
input with a `DISPLAY TO / BY NAME` instruction or with the
`DIALOG.setFieldTouched()` method.

When the modification flag is set by program, `NOENTRY` fields are ignored.
However, fields dynamically disabled by `DIALOG.setFieldActive()` are taken into
account.

## Deny row append

Temporary row creation is useful because, in most cases, `INPUT ARRAY` is used to
edit existing rows and append new rows at the end of the list. However, you might want to prevent
row addition or at least avoid the automatic temporary row creation when the last row is deleted or
when an empty list gets the focus.

To avoid explicit temporary row creation, prevent `INPUT ARRAY` from defining the
automatic append action by setting the `APPEND ROW` attribute to
`FALSE` in the `ATTRIBUTE` clause:

```
...
   INPUT ARRAY p_items FROM sa.* ATTRIBUTES(APPEND ROW=FALSE)
   ...
   END INPUT
...
```

When `APPEND ROW` or `INSERT ROW` attributes are set to
`FALSE`, automatic temporary row can still occur when the user deletes the last row
of the list or if the list is empty when the `INPUT ARRAY` is entered. To avoid
automatic temporary row creation when only one of `APPEND ROW=FALSE` or
`INSERT ROW=FALSE`, use `AUTO
APPEND=FALSE`:

```
...
   INPUT ARRAY p_items FROM sa.* ATTRIBUTES(INSERT ROW=FALSE, AUTO APPEND=FALSE)
   ...
   END INPUT
...
```

If both `APPEND ROW` and `INSERT ROW` attributes are set to
`FALSE`, the dialog will prevent explicit temporary row creation and also automatic
temporary row creation, as if `AUTO APPEND = FALSE` was used.

## Row creation control blocks for temporary rows

In order to control row creation, use the `BEFORE INSERT` and `AFTER
INSERT` control blocks. The `BEFORE INSERT` trigger is invoked after a new
row was inserted or appended, just before the user gets control to enter data in fields.
Regarding temporary rows, the `AFTER INSERT` block is invoked if data has been
entered and you leave the new row (for example, when the focus moves to another row or leaves the
current list), or if the dialog is validated, for example with [`ACCEPT DIALOG`](2142-accept-dialog-instruction.md) in case of
`DIALOG` (or [`ACCEPT
INPUT`](2034-accept-input-instruction.md) in case of singular `INPUT ARRAY`). No `AFTER
INSERT` block is invoked if the user did not enter data. The temporary row is automatically
deleted.

In the `BEFORE INSERT` control block, you can tell if a row is a temporary
appended one by comparing the current row (`DIALOG.getCurrentRow()` or `ARR_CURR()`) with the total number of
rows (`DIALOG.getArrayLength()` or `ARR_COUNT()`). If the current row index
equals the row count, you are in a temporary row.

## AFTER ROW and temporary rows

When a temporary row is automatically removed, the `AFTER ROW` block will be executed for the
temporary row, but `ui.Dialog.getCurrentRow()` / `ARR_CURR()` will be
one row greater than `DIALOG.getArrayLength()` / `ARR_COUNT()`. In
this case, ignore the `AFTER ROW` event.

## Related links

**Related concepts**  

[INPUT ARRAY ATTRIBUTES clause](2092-input-array-attributes-clause.md "INPUT ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")

[The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")
