---
title: "BEFORE ROW no longer executed when array is empty"
source: "fgl-topics/c_fgl_Migrate_to_220_before_row.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > BEFORE ROW no longer executed when array is empty"
type: "concept"
---

# BEFORE ROW no longer executed when array is empty

> In order to trigger the BEFORE ROW block when entering an array, the array must not be empty.

Before version 2.20, the `BEFORE ROW` block was always executed when
entering a `DISPLAY ARRAY` or `INPUT ARRAY` dialog, even if the number
of real data rows was zero. Starting with 2.20, when using an empty dynamic array or when using a
static array and specifying zero data rows with a `SET_COUNT(0)` call or with the
`COUNT=0` attribute, the `BEFORE ROW` control block is no longer
executed when the dialog starts.

The `BEFORE ROW` block will be executed when a new row is created in `INPUT
ARRAY`. When entering an `INPUT ARRAY` with an empty array, a new temporary
row is created by default, except if you use the `AUTO APPEND = FALSE` attribute.

## Related links

**Related concepts**  

[BEFORE ROW block](../11_user-interface/1975-before-row-block.md "BEFORE ROW block")
