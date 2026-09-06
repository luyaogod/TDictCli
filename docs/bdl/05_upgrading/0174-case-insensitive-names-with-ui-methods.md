---
title: "Case insensitive names with UI methods"
source: "fgl-topics/c_fgl_Migrate_to_320_case_insens_names.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Case insensitive names with UI methods"
type: "concept"
---

# Case insensitive names with UI methods

> Methods of built-in classes using user interface object names are now case insensitive.

Before Genero BDL version 3.20, built-in class methods such as
`ui.Dialog.setActionActive()` were case sensitive. Since compilers convert names
to lowercase, it was required to use lowercase names in the built-in class
methods:

```
ON ACTION MyAction   -- converted to myaction
...
CALL DIALOG.setActionActive("MyAction",FALSE)   -- produced a runtime error!
CALL DIALOG.setActionActive("myaction",FALSE)   -- works
```

Starting with version 3.20, these methods are now case insensitive, so you can
write:

```
ON ACTION MyAction   -- converted to myaction
...
CALL DIALOG.setActionActive("MyAction",FALSE)
```

> **Note:**
>
> This change is backward compatible, there is no need to modify the existing code.

User interface object names defined in the form files can also be referenced with the exact
name.

For example, in the .per form
file:

```
EDIT f01 = Customer.CustAddr, ... ;
GROUP g1 : Group1, ... ;
...
SCREEN RECORD CustRec (...);
...
```

In the .4gl source
file:

```
CALL DIALOG.setFieldActive("Customer.CustAddr",FALSE)
...
CALL DIALOG.getForm().setElementHidden("Group1",1)
...
LET r = DIALOG.getCurrentRow("CustRec")
...
```

> **Note:**
>
> The methods to build [dynamic dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.") are still
> case sensitive. For example, to define an `ON ACTION` trigger: `CALL
> mydlg.addTrigger("ON ACTION accept")`.

## Related links

**Related concepts**  

[Identifying actions in ui.Dialog methods](../15_library-reference/3238-identifying-actions-in-ui-dialog-methods.md "Identifying actions in ui.Dialog methods")

[Identifying fields in ui.Dialog methods](../15_library-reference/3239-identifying-fields-in-ui-dialog-methods.md "Identifying fields in ui.Dialog methods")

[Identifying screen-arrays in ui.Dialog methods](../15_library-reference/3240-identifying-screen-arrays-in-ui-dialog-methods.md "Identifying screen-arrays in ui.Dialog methods")

[Identifying elements in ui.Form methods](../15_library-reference/3164-identifying-elements-in-ui-form-methods.md "Identifying elements in ui.Form methods")

[ui.Window.forName](../15_library-reference/3130-ui-window-forname.md "Get a window object by name.")

[ui.ComboBox.forName](../15_library-reference/3251-ui-combobox-forname.md "Search for a combobox in the current form.")
