---
title: "BEFORE DISPLAY / BEFORE ROW execution order"
source: "fgl-topics/c_fgl_Mig0000_039.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > BEFORE DISPLAY / BEFORE ROW execution order"
type: "concept"
---

# BEFORE DISPLAY / BEFORE ROW execution order

> Genero BDL scenarios define whether the BEFORE DISPLAY or BEFORE ROW block is executed first.

When starting a `DISPLAY ARRAY` dialog with Four Js Business Development Suite
(BDS), the [`BEFORE ROW`](../11_user-interface/1975-before-row-block.md) control
block is executed first, then [`BEFORE
DISPLAY`](../11_user-interface/1973-before-display-block.md) is executed.

This is not logical: The `BEFORE DISPLAY` block is typically used to do some
initializations. Therefore, it should be executed before any other dialog trigger.

With Genero BDL, when a singular `DISPLAY ARRAY` dialog starts, the `BEFORE
DISPLAY` block is executed first, then the `BEFORE ROW` block is executed. In
a `DISPLAY ARRAY` subdialog of a `DIALOG` instruction, the
`BEFORE DISPLAY` is executed when the list gets the focus, then `BEFORE
ROW` is executed.

## Related links

**Related concepts**  

[Record list (DISPLAY ARRAY)](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")

[Multiple dialogs (DIALOG - inside functions)](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
