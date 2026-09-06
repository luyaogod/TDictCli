---
title: "INPUT ARRAY control blocks execution order"
source: "fgl-topics/c_fgl_InputArray_012.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY control blocks > INPUT ARRAY control blocks execution order"
type: "concept"
description: "This table shows the order in which the runtime system executes the control blocks in the INPUT ARRAY instruction, based on the user action: Table 1. Control block execution order for INPUT ARRAY ..."
---

# INPUT ARRAY control blocks execution order

This table shows the order in which the runtime system executes the control blocks in the
`INPUT ARRAY` instruction, based on the user action:

| Context / User action | Control Block execution order |
| --- | --- |
| Entering the dialog | [`BEFORE INPUT`](1940-before-input-block.md)[`BEFORE ROW`](1975-before-row-block.md)[`BEFORE FIELD`](1942-before-field-block.md) |
| Moving to a different row from field A to field B | [`ON CHANGE`](1943-on-change-block.md) (if value has changed for field A)[`AFTER FIELD`](1944-after-field-block.md) (for field A in the row you leave)[`AFTER INSERT`](2021-after-insert-block.md) (if the row you leave was inserted or appended)or[`ON ROW CHANGE`](2018-on-row-change-block.md) (if values have changed in the row you leave)[`AFTER ROW`](1976-after-row-block.md) (for the row you leave)`BEFORE ROW` (the new current row)`BEFORE FIELD` (for field B in the new current row) |
| Moving from field A to field B in the same row | `ON CHANGE` (if value has changed for field A)`AFTER FIELD` (for field A)`BEFORE FIELD` (for field B) |
| Deleting a row | [`BEFORE DELETE`](2022-before-delete-block.md) (for the row to be deleted)[`AFTER DELETE`](2023-after-delete-block.md) (for the deleted row)`AFTER ROW` (for the deleted row)`BEFORE ROW` (for the new current row)`BEFORE FIELD` (field in the new current row) |
| Inserting a new row between rows | `ON CHANGE` (if value has changed in the field you leave)`AFTER FIELD` (for the row you leave)`AFTER INSERT` (if the row you leave was inserted or appended)or`ON ROW CHANGE` (if values have changed in the row you leave)`AFTER ROW` (for the row you leave)`BEFORE INSERT` (for the new created row)`BEFORE FIELD` (for the new created row) |
| Appending a new row at the end | `ON CHANGE` (if value has changed in the current field)`AFTER FIELD` (for the row you leave)`AFTER INSERT` (if the row you leave was inserted or appended)or`ON ROW CHANGE` (if values have changed in the row you leave)`AFTER ROW` (for the row you leave)`BEFORE ROW` (for the new created row)`BEFORE INSERT` (for the new created row)`BEFORE FIELD` (for the new created row) |
| Validating the dialog | `ON CHANGE``AFTER FIELD``AFTER INSERT` (if the current row was inserted or appended)or`ON ROW CHANGE` (if values have changed in the current row)`AFTER ROW``AFTER INPUT` |
| Canceling the dialog | `AFTER ROW``AFTER INPUT` |
