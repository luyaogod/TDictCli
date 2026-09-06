---
title: "WANTCOLUMNSANCHORED is desupported"
source: "fgl-topics/c_fgl_Migrate_to_200_wantcolumnsanchored.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > WANTCOLUMNSANCHORED is desupported"
type: "concept"
---

# WANTCOLUMNSANCHORED is desupported

> Use UNMOVABLECOLUMNS to specify that table columns cannot be moved around by the user.

Before version 2.00, the `WANTCOLUMNSANCHORED` attribute was undocumented but
still supported by the language, to simplify migration from 1.20.

Starting with version 2.00, the `WANTCOLUMNSANCHORED` attribute is desupported;
you must use `UNMOVABLECOLUMNS` to specify that table columns cannot be moved
around by the user.

## Related links

**Related concepts**  

[UNMOVABLECOLUMNS attribute](../11_user-interface/1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table.")
