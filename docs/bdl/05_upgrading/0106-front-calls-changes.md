---
title: "Front calls changes"
source: "fgl-topics/c_fgl_Migrate_to_501_frontcalls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.01 upgrade guide > Front calls changes"
type: "concept"
---

# Front calls changes

> Modifications to consider when using front calls.

## New front calls

Starting with GBC 5.01.04:

- [`table.autofitAllColumns`](../15_library-reference/3420-table-autofitallcolumns.md "Adapts the width of table columns to the displayed data.")
- [`table.fitToViewAllColumns`](../15_library-reference/3421-table-fittoviewallcolumns.md "Adapts the width of table columns to show all columns.")

## `mobile.scanBarCode` output format specification (GMA)

Starting with GMA 5.01.01, the `mobile.scanBarCode` front call accepts an optional
parameter to specify the output format for the bar code value.

Refer to [mobile.scanBarCode](../15_library-reference/3459-mobile-scanbarcode.md "Allow the user to scan a barcode with a mobile device") for more details.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Front
call changes in BDL 5.00](0117-front-calls-changes.md "Modifications to consider when using front calls.").

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
