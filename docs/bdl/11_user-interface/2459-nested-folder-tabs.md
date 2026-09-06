---
title: "Nested folder tabs"
source: "fgl-topics/c_fgl_ui_folders_page_change.html"
breadcrumb: "User interface > User interface programming > Folders > Nested folder tabs"
type: "concept"
---

# Nested folder tabs

> Folder pages can contain other folders. How to deal with this?

When a `FOLDER` `PAGE` contains another `FOLDER`
container, it defines a sub-folder.

When a user selects the tab of the parent folder page containing the sub-folder, it is possible
to show a given sub-folder page, by calling the `ui.Form.ensureFieldVisible()` or
`ui.Form.ensureElementVisible()` method, in the `BEFORE INPUT`,
`BEFORE DISPLAY` or `BEFORE CONSTRUCT` of the sub-dialog controlling
the parent folder page:

```
DIALOG ...
    -- Sub-dialog controlling the parent folder page
    INPUT BY NAME rec_customer.* ...
       BEFORE INPUT -- focus goes to a field of this sub-dialog/folder page
           -- Show the main page of the sub-folder with orders info
           CALL DIALOG.getForm().ensureFieldVisible("orders.order_num")
       ...
    END INPUT
    -- Sub-dialog controlling sub-folder main page
    INPUT BY NAME rec_orders.* ...
       ...
```

Note that the focus will stay in the current field of the parent folder page because this is not
a `NEXT FIELD`.
