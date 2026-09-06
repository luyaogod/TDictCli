---
title: "Understanding folder tabs"
source: "fgl-topics/c_fgl_ui_folders_basics.html"
breadcrumb: "User interface > User interface programming > Folders > Understanding folder tabs"
type: "concept"
---

# Understanding folder tabs

> Folders allow to display a set of form elements in dedicated pages (aka folder tabs).

A [`FOLDER`](1691-folder-item-type.md "Defines a layout area to hold folder pages.") element contains
two or more [`PAGE`](1697-page-item-type.md "Defines the content of a folder page.") elements, that
can contain other containers such as `GRID` and `TABLE`. Each page
gets a tab that can be selected by the user to switch to a given page.

In the next screenshots, a folder defines three pages identified by the tab titles "Fruits",
"Vegetables" and "Meat".

In the first screenshot, the "Fruits" folder page is selected:

![Form with foldertabs screenshot (page 1)](../_images/gbc_foldertabs_1.jpg)

*Form with Folder Tabs showing first page*

In the next screenshot, the "Vegetables" folder page is selected:

![Form with foldertabs screenshot (page 2)](../_images/gbc_foldertabs_2.jpg)

*Form with Folder Tabs showing second page*

To control a form with folder pages containers, use a [`DIALOG`](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.") instruction block where
each sub-dialog controls a specific folder page.

A given sub-dialog may control several folder pages, but in this case, it will not be possible to
detect page selection within `BEFORE INPUT`, `BEFORE DISPLAY` or
`BEFORE CONSTRUCT` triggers in `DIALOG/END DIALOG`.

A folder `PAGE` can define an `ACTION` attribute to fire an
`ON ACTION` handler, when the folder tab is selected. This solution must only be used
in specific cases.

## Related links

**Related concepts**  

[Multiple dialogs (DIALOG - inside functions)](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
