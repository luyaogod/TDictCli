---
title: "Defining the window icon"
source: "fgl-topics/c_fgl_windows_and_forms_012.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Defining the window icon"
type: "concept"
---

# Defining the window icon

> Use a IMAGE attribute to define the icon for a window.

If the window is opened with `OPEN WINDOW
WITH FORM`, by using a form file that defines an `IMAGE` attribute in
the [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section, the
window will use this image as icon.

Subsequent `OPEN FORM` / `DISPLAY FORM` instructions may change the
window icon if the new form defines a different image in the `LAYOUT` section.

If you want to set a window icon dynamically, you can use the `setImage()` method of the `ui.Window`
built-in class.
