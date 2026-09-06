---
title: "Control block execution order in multiple dialogs"
source: "fgl-topics/c_fgl_multiple_dialogs_ctrlblock_exec_ord.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > Control block execution order in multiple dialogs"
type: "concept"
description: "This table shows the order in which control blocks are executed in a procedural DIALOG instruction, depending on the context and user action: Table 1. Control block execution order for a procedural ..."
---

# Control block execution order in multiple dialogs

This table shows the order in which control blocks are executed in a procedural
`DIALOG` instruction, depending on the context and user action:

| Context / User action | Control Block execution order |
| --- | --- |
| Entering the dialog | `BEFORE DIALOG``BEFORE INPUT`, `BEFORE CONSTRUCT` or `BEFORE DISPLAY` (first sub-dialog getting focus)`BEFORE ROW` (if focus goes to a list)`BEFORE FIELD` (if focus goes to a field) |
| When the focus goes to an `INPUT` or to a `CONSTRUCT` from a different sub-dialog | *Triggers raised by the context of the sub-dialog you leave*`BEFORE INPUT` or `BEFORE CONSTRUCT` (new sub-dialog getting focus)`BEFORE FIELD` |
| When the focus leaves an `INPUT` or a `CONSTRUCT` to a different sub-dialog | `ON CHANGE` (if INPUT and value of current field has changed)`AFTER FIELD` (for the current field)`AFTER INPUT` or `AFTER CONSTRUCT` (current sub-dialog losing focus)*Triggers raised by the context of the sub-dialog you enter* |
| When the focus goes to a `DISPLAY ARRAY` list or to an `INPUT ARRAY` list from a different sub-dialog | *Triggers raised by the context of the sub-dialog you leave*`BEFORE INPUT` or `BEFORE DISPLAY` (new sub-dialog getting focus)`BEFORE ROW` (the row that was selected in the list)`BEFORE FIELD` (if it's an INPUT ARRAY) |
| When the focus leaves a `DISPLAY ARRAY` or `INPUT ARRAY` list to a different sub-dialog | `ON CHANGE` (if INPUT ARRAY and value of current field has changed)`AFTER FIELD` (if it's an INPUT ARRAY)`AFTER INSERT` (if INPUT ARRAY and current row was just created)or`ON ROW CHANGE` (if INPUT ARRAY and a value in the row has changed)`AFTER ROW` (the current row in the list you leave)`AFTER INPUT` or `AFTER DISPLAY` (current sub-dialog losing focus)*Triggers raised by the context of the sub-dialog you enter* |
| Moving from field A to field B in an `INPUT` or `CONSTRUCT` sub-dialog or in the same row of an `INPUT ARRAY` list | `ON CHANGE` ( if value of current field has changed)`AFTER FIELD A``BEFORE FIELD B` |
| Moving from field A of an `INPUT` or `CONSTRUCT` sub-dialog to field B in another `INPUT` or `CONSTRUCT` sub-dialog | `ON CHANGE` (if value of current field has changed)`AFTER FIELD A``AFTER INPUT` or `AFTER CONSTRUCT` (for sub-dialog of field A)`BEFORE INPUT` or `BEFORE CONSTRUCT` (for sub-dialog of field B)`BEFORE FIELD B` |
| Moving to a different row in a `DISPLAY ARRAY` list | `AFTER ROW` (the row you leave)`BEFORE ROW` (the new current row) |
| Moving to a different row in an `INPUT ARRAY` list when current row was newly created | `ON CHANGE` (if value of current field has changed)`AFTER FIELD` (for field A in the row you leave)`AFTER INSERT` (the newly created row)`AFTER ROW` (the newly created row)`BEFORE ROW` (the new current row)`BEFORE FIELD` (field in the new current row) |
| Moving to a different row in an `INPUT ARRAY` list when current row was modified | `ON CHANGE` (if value of current field has changed)`AFTER FIELD` (for field A in the row you leave)`ON ROW CHANGE` (the values in current row have changed)`AFTER ROW` (for the row that was modified)`BEFORE ROW` (the new current row)`BEFORE FIELD` (field in the new current row) |
| Inserting or appending a new row in an `INPUT ARRAY` list | *Triggers raised by the context of the sub-dialog you leave*`BEFORE INSERT` (for the new current row)`BEFORE ROW` (the new current row)`BEFORE FIELD` (field in the new current row) |
| Deleting a row in an `INPUT ARRAY` list | `BEFORE DELETE` (for the current row to be deleted)`AFTER DELETE` (now the deleted row is removed)`AFTER ROW` (for the deleted row)`BEFORE ROW` (the new current row) |
| Firing the *insert* or *append* action for the `ON INSERT` block in a `DISPLAY ARRAY` list | `AFTER ROW``ON INSERT``BEFORE ROW` |
| Firing the *delete* action for the `ON DELETE` block in a `DISPLAY ARRAY` list | `AFTER ROW``ON DELETE``BEFORE ROW` |
| Validating the dialog with `ACCEPT DIALOG` | `ON CHANGE` (if focus is in input field and value has changed)`AFTER FIELD` (if focus is in input field)`AFTER INSERT` (if INPUT ARRAY and current row was just created)or`ON ROW CHANGE` (if INPUT ARRAY and a value in the row has changed)`AFTER ROW` (if focus is in a list)`AFTER INPUT`, `AFTER CONSTRUCT` or `AFTER CONSTRUCT` (current sub-dialog)`AFTER DIALOG` |
| Canceling the dialog with `EXIT DIALOG` | None of the control blocks will be executed; we just leave the dialog instruction. |
