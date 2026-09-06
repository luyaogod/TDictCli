---
title: "DISPLAY ARRAY control blocks execution order"
source: "fgl-topics/c_fgl_DisplayArray_020.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY control blocks > DISPLAY ARRAY control blocks execution order"
type: "concept"
description: "This table shows the order in which the runtime system executes the control blocks in the DISPLAY ARRAY instruction, based on the user action: Table 1. Control blocks execution order in DISPLAY ARRAY ..."
---

# DISPLAY ARRAY control blocks execution order

This table shows the order in which the runtime system executes the control blocks in the
`DISPLAY ARRAY` instruction, based on the user action:

| Context / User action | Control Block execution order |
| --- | --- |
| Entering the dialog | `BEFORE DISPLAY``BEFORE ROW` |
| Moving to a different row | `AFTER ROW` (the current row)`BEFORE ROW` (the new row) |
| Validating the dialog | `AFTER ROW``AFTER DISPLAY` |
| Canceling the dialog | `AFTER ROW``AFTER DISPLAY` |
| Firing the insert or append action for the `ON INSERT` block | `AFTER ROW``ON INSERT``BEFORE ROW` |
| Firing the delete action for the `ON DELETE` block | `AFTER ROW``ON DELETE``BEFORE ROW` |

## Related links

**Related concepts**  

[ON APPEND block](1984-on-append-block.md "ON APPEND block")

[ON UPDATE block](1986-on-update-block.md "ON UPDATE block")
