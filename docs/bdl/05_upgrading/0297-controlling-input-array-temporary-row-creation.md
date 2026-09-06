---
title: "Controlling INPUT ARRAY temporary row creation"
source: "fgl-topics/c_fgl_Migrate_to_220_input_array_tmp_row.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Controlling INPUT ARRAY temporary row creation"
type: "concept"
---

# Controlling INPUT ARRAY temporary row creation

> Down move after last row in INPUT ARRAY creates a new temporary row.

The INPUT ARRAY dialog and sub-dialog provides the APPEND ROW and AUTO APPEND attributes to
control row creation at the end of a list (known as temporary row creation).

`APPEND ROW` controls explicit temporary row creation, while `AUTO
APPEND` controls automatic temporary row creation.

Starting with version 2.20, moving down after the last row (with the mouse or keyboard) or
leaving the last column of the last row with a TAB key are considered events that trigger automatic
temporary row creation.

Before version 2.20, these cases were considered as events for an explicit temporary row
creation. In other words, if you want to deny temporary row creation in such case, it is now done
with `AUTO APPEND = FALSE` while in older versions it was controlled by
`APPEND ROW = FALSE`.

## Related links

**Related concepts**  

[Appending rows in INPUT ARRAY](../11_user-interface/2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.")
