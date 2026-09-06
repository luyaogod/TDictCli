---
title: "Migrating screen arrays to tables"
source: "fgl-topics/c_fgl_Mig0000_018.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > Migrating screen arrays to tables"
type: "concept"
---

# Migrating screen arrays to tables

> A TABLE container in Genero BDL displays using a real table widget, providing a more robust display and interaction than a screen array, while the SCROLLGRID container renders a list of records in separated field cells, providing a replacement for screen arrays using the OPTIONS="-nolist"

This topic also concerns IBM® Informix® 4GL migration, see
the [I4GL Migration](0354-migrating-screen-arrays-to-tables.md "Tables in Genero BDL display using a real table widget, providing a more robust display and interaction than the I4GL screen array.") page
for mode details.

By default with Four Js Business Development Suite (BDS), fields of a screen array were rendered
as listviews:

![Screen shot of a Four Js BDS screen array using listviews.](../_images/BDS_listview.jpg)

*Four Js BDS listview rendering for screen arrays*

To split such listview into individual field cells, Four Js BDS provided the
`OPTIONS="-nolist"` attribute, to be specified for each field of the screen
array:

```
f01 = FORMONLY.cust_fname, OPTIONS="-nolist";
f02 = FORMONLY.cust_lname, OPTIONS="-nolist";
```

![Screen shot of a Four Js BDS screen array with fields using OPTIONS="-nolist".](../_images/BDS_option_nolist.jpg)

*Four Js BDS screen array with fields using OPTIONS="-nolist"*

With Genero BDL, consider using a [`SCROLLGRID` container](../11_user-interface/1723-scrollgrid-container.md "Defines a scrollable grid view widget."), to render a list of records in separated field cells.
The `SCROLLGRID` container allows to position field tags on several grid lines, by
repeating the same layout for each row:

![Screen shot of a Genero BDL SCROLLGRID container.](../_images/scrollgrid_list_gbc.jpg)

*Genero BDL SCROLLGRID container*
