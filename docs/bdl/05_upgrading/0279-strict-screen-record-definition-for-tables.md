---
title: "Strict screen record definition for tables"
source: "fgl-topics/c_fgl_Migrate_to_221_strict_screen_record.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide > Strict screen record definition for tables"
type: "concept"
---

# Strict screen record definition for tables

> The fglform compiler of version 2.21.00 now makes a strict checking of the fields used in the screen record definition for table containers.

It generates error [-6819](../15_library-reference/4483-genero-bdl-errors.md) if
the screen record do not use all columns used in the table. The order can be different, however.

## Related links

**Related concepts**  

[TABLE container](../11_user-interface/1724-table-container.md "Defines a re-sizable table designed to display a list of records.")
