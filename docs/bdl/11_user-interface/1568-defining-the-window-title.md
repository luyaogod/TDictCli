---
title: "Defining the window title"
source: "fgl-topics/c_fgl_windows_and_forms_011.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Defining the window title"
type: "concept"
---

# Defining the window title

> Use the TEXT attribute to define a title for a window.

The `TEXT` attribute in the `ATTRIBUTE` clause of `OPEN WINDOW` defines the default title of the window.
If the window is opened with a form (`WITH FORM` clause) that defines a
`TEXT` attribute in the [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section, the default is ignored. Subsequent `OPEN
FORM` / `DISPLAY FORM` instructions may change the window title if the new
form defines a different title in the `LAYOUT` section.

It is recommended that you specify the window title in the form
file, instead of using the `TEXT` attribute of
the `OPEN WINDOW` instruction.

If you want to set a window title dynamically, you can use the `setText()` method of the `ui.Window` built-in class.
