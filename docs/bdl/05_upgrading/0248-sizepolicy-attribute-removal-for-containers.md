---
title: "SIZEPOLICY attribute removal for containers"
source: "fgl-topics/c_fgl_Migrate_to_240_sizepolicy.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > SIZEPOLICY attribute removal for containers"
type: "concept"
---

# SIZEPOLICY attribute removal for containers

> The SIZEPOLICY attribute is no longer available for layout containers like TABLE / GRID.

Before version 2.40 it was possible to specify a SIZEPOLICY attribute for several types of form
elements, including containers such as TABLE, GRID. The SIZEPOLICY attribute makes no sense in
containers and is only meaningful for leaf nodes (widgets such as EDIT, COMBOBOX). The form compiler
will now report a syntax error if the SIZEPOLICY attribute is used in the definition of elements
that are not widgets.

## Related links

**Related concepts**  

[SIZEPOLICY attribute](../11_user-interface/1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.")
