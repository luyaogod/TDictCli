---
title: "CONSTRUCT control blocks execution order"
source: "fgl-topics/c_fgl_Construct_012.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control blocks > CONSTRUCT control blocks execution order"
type: "concept"
description: "This table shows the order in which the runtime system executes the control blocks in the CONSTRUCT instruction, depending on the user action: Table 1. Control block execution order for CONSTRUCT ..."
---

# CONSTRUCT control blocks execution order

This table shows the order in which the runtime system executes the control blocks
in the `CONSTRUCT` instruction, depending on the user action:

| Context / User action | Control Block execution order |
| --- | --- |
| Entering the dialog | [`BEFORE CONSTRUCT`](2057-before-construct-block.md)[`BEFORE FIELD`](1942-before-field-block.md) (first field) |
| Moving from field A to field B | [`AFTER FIELD`](1944-after-field-block.md) (for field A)[`BEFORE FIELD`](1942-before-field-block.md) (for field B) |
| Validating the dialog | `AFTER FIELD`[`AFTER CONSTRUCT`](2058-after-construct-block.md) |
| Canceling the dialog | `AFTER CONSTRUCT` |
