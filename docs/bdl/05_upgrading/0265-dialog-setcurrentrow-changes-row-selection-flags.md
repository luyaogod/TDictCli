---
title: "DIALOG.setCurrentRow() changes row selection flags"
source: "fgl-topics/c_fgl_Migrate_to_230_dialog_setcurrentrow.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > DIALOG.setCurrentRow() changes row selection flags"
type: "concept"
---

# DIALOG.setCurrentRow() changes row selection flags

> Row selection flags are reset by a call to setCurrentRow().

Before version 2.30, the `DIALOG.setCurrentRow()` method did not modify the row
selection flags.

Starting with version 2.30, the method resets row selection flags to false and marks the
new current row as selected.

## Related links

**Related concepts**  

[Multiple row selection](../11_user-interface/2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.")
