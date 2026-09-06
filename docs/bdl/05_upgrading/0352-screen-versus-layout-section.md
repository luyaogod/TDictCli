---
title: "SCREEN versus LAYOUT section"
source: "fgl-topics/c_fgl_MigI4GL_012.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > SCREEN versus LAYOUT section"
type: "concept"
---

# SCREEN versus LAYOUT section

> When writing new programs for GUI applications, it is recommended that you use a LAYOUT section instead of SCREEN. However, the SCREEN section is still supported to be used to design TUI mode forms.

To design a form with IBM®
Informix® 4GL (I4GL), you organize labels and fields in
the `SCREEN` section of a .per form file. Genero Business
Development Language introduced a new `LAYOUT` section to hold form elements. The new
`LAYOUT` section allows for a more sophisticated form design than the
`SCREEN` section.

![Form using SCREEN section in TUI mode screenshot](../_images/TextMode1.jpg)

*Form using a SCREEN section in TUI mode*

![Form using LAYOUT section in GUI mode](../_images/CustOrders1.jpg)

*Form using a LAYOUT section in GUI mode*

## Related links

**Related concepts**  

[LAYOUT section](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
