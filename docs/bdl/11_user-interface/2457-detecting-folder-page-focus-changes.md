---
title: "Detecting folder page focus changes"
source: "fgl-topics/c_fgl_ui_folders_current_page.html"
breadcrumb: "User interface > User interface programming > Folders > Detecting folder page focus changes"
type: "concept"
---

# Detecting folder page focus changes

> Field focus change defines the current folder page.

When a folder page is selected by the user, and the focus is in the context of that folder, the
focus is set to the field with the lowest [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.") in the selected folder `PAGE`.

The dialog code can then detect which page was selected by using `BEFORE INPUT`,
`BEFORE DISPLAY` or `BEFORE CONSTRUCT` of the corresponding sub-dialog
block:

```
DIALOG ...
    INPUT BY NAME rec.* ...
       BEFORE INPUT -- focus goes to a field of this sub-dialog/folder page
           ...
```

If the focus is not in the context of the folder, the user can still select a folder tab to show
up one of its pages, but the focus is not set to an element of that page: The focus remains in the
current field, unless the user explicitly clicks or tabs to a field of the displayed page. This
allows to view folder pages while entering data in another dialog, without leaving the current
field, which would produce input validation and `AFTER FIELD` / `AFTER
INPUT` blocks execution.
