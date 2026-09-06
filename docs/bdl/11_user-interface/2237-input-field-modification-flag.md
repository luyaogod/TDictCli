---
title: "Input field modification flag"
source: "fgl-topics/c_fgl_prog_dialogs_touched_flag.html"
breadcrumb: "User interface > User interface programming > Input fields > Input field modification flag"
type: "concept"
---

# Input field modification flag

> Each input field controlled by a dialog instruction has a modification flag.

The modification flag is used to execute form-level validation rules and trigger `ON
CHANGE` blocks. The flag can also be queried to detect if a field was touched during the
`DIALOG` instruction, for example with the [`FIELD_TOUCHED()`](../08_language-basics/0673-field-touched-function.md "The FIELD_TOUCHED() operator checks if fields were modified during the dialog execution.") operator or with
the `ui.Dialog.getFieldTouched()` method.

Both `FIELD_TOUCHED()` and `ui.Dialog.getFieldTouched()` accept a
list of fields and/or the `screen-record.*` notation in order to check the
modification flag of multiple fields in a unique function call. Use a single `*` star
character, to reference all fields used by the dialog.

The modification flag is set to `TRUE` when the user enters data in a field, or
when the program code executes a [`DISPLAY
TO`](1880-display-to.md "The DISPLAY ... TO instruction displays data to specific form fields."), [`DISPLAY BY
NAME`](1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names."), [`CLEAR
field-spec`](1884-clear-field-list.md "The CLEAR field-list instruction clears specific fields in the current form.") or [`CLEAR SCREEN ARRAY
screen-array.*`](1883-clear-screen-array.md "The CLEAR SCREEN ARRAY instruction clears the values of all rows of the form list identified by the specified screen array.") instruction. The modification flags can also be
set by program to `TRUE` or reset to `FALSE` with the `ui.Dialog.setFieldTouched()` method, to
emulate user input, or to reset the modification flags after data is saved in the database.

The modification flags of all fields are automatically reset to
`FALSE` by the interactive instruction in the following
cases:

- When the dialog instruction starts.
- In a `DIALOG` block, when entering a group of fields controlled by an
  `INPUT` or a `CONSTRUCT` sub-dialog.
- When moving to (or creating) a new row in an `INPUT ARRAY`.
- Within a `DISPLAY ARRAY`, the modification flags are always `TRUE`
  for all fields.

When using a `DISPLAY ARRAY`, the modification flags are set to
`TRUE` for all fields. This behavior exists because of backward compatibility.
Since values cannot be modified by the user, the modification flags are not relevant in this
dialog. However, you must pay attention when implementing nested dialogs, because `DISPLAY
ARRAY` will set the modification flags of the fields driven by the parent dialog, for
example when executing a `DISPLAY ARRAY` from an `INPUT ARRAY`.

Query the modification flags with the `ui.Dialog.getFieldTouched()`
method, typically in the context of
[`AFTER INPUT`](1941-after-input-block.md),
[`AFTER CONSTRUCT`](2058-after-construct-block.md),
[`AFTER INSERT`](2021-after-insert-block.md) or
[`AFTER ROW`](1976-after-row-block.md)
control blocks.

When using a list driven by an `INPUT ARRAY` binding, a [temporary row](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") added at the end of the list will be
automatically removed if all fields have the modification flag set to `FALSE`.

For typical `EDIT` fields, the modification flag
is set when leaving the field. If you want to detect data modification
earlier, it is recommended that you use the
[dialogtouched](2239-immediate-detection-of-user-changes.md "This section describes the dialogtouched predefined action.") predefined
action. However, this event is only an indicator that the user started
to modify a field, the value will not be available in the program variables.

## Related links

**Related concepts**  

[Immediate detection of user changes](2239-immediate-detection-of-user-changes.md "This section describes the dialogtouched predefined action.")

[Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.")
