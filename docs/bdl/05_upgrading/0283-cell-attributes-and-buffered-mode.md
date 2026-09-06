---
title: "Cell attributes and buffered mode"
source: "fgl-topics/c_fgl_Migrate_to_220_cell_attributes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Cell attributes and buffered mode"
type: "concept"
---

# Cell attributes and buffered mode

> Must use the UNBUFFERED mode when setting cell attributes.

Before version 2.20, array cell attributes were synchronized quite often by the runtime system,
and this was not very efficient. As a result, there was not much difference between using
buffered or unbuffered mode; when changing cell attributes, the result was immediate even in
buffered mode.

Starting with version 2.20, it is recommended that you use the UNBUFFERED mode when setting cell
attributes; otherwise, the colors will not be synchronized on the front-end.

## Related links

**Related concepts**  

[Cell color attributes](../11_user-interface/2313-cell-color-attributes.md "List controllers can display every cell in a specific color.")

[The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
