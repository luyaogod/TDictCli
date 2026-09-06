---
title: "INPUT control blocks execution order"
source: "fgl-topics/r_fgl_record_input_010.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT control blocks > INPUT control blocks execution order"
type: "reference"
description: "This table shows the order in which the runtime system executes the control blocks in the INPUT instruction, depending on the user action: Table 1. Control Block Execution Order for INPUT Context / ..."
---

# INPUT control blocks execution order

This table shows the order in which the runtime system executes the control
blocks in the `INPUT` instruction, depending on the user action:

| Context / User action | Control Block execution order |
| --- | --- |
| Entering the dialog | `BEFORE INPUT``BEFORE FIELD` (first field) |
| Moving from field A to field B | `ON CHANGE` (if value has changed for field A)`AFTER FIELD` (for field A)`BEFORE FIELD` (for field B) |
| Changing the value of a field with a specific field like checkbox | `ON CHANGE` |
| Validating the dialog | `ON CHANGE` (if value has changed in current field)`AFTER FIELD``AFTER INPUT` |
| Canceling the dialog | `AFTER INPUT` |
