---
title: "Configuring a dynamic dialog"
source: "fgl-topics/c_fgl_dynamic_dialogs_config.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Configuring a dynamic dialog"
type: "concept"
---

# Configuring a dynamic dialog

> The dynamic dialogs can be configured with attributes.

After creating the `ui.Dialog` dynamic dialog object, use the
`setDialogAttribute()` method to specify dialog options.

For singular dialogs, define dialog attributes after using one of the `create*()`
methods:

```
DEFINE d ui.Dialog
...
LET d = ui.Dialog.createInputByName(fields)
CALL d.setDialogAttribute( "UNBUFFERED", TRUE )
```

When creating a multiple dialog, the `setDialogAttribute()` method applies the
attribute value to the last added sub-dialog:

```
DEFINE d ui.Dialog
...
LET d = ui.Dialog.createMultipleDialog()
...
CALL d.addDisplayArrayTo(fields, "sr_cust")
CALL d.setDialogAttribute( "HELP", 1023 )  -- applies to sr_cust
...
CALL d.addInputArrayFrom(fields, "sr_ord")
CALL d.setDialogAttribute( "INSERT ROW", FALSE )  -- applies to sr_ord
...
```

The attribute name passed as first parameter must match the attribute name that can be specified
in the static dialog instruction syntax. For a complete list of attribute names and possible values,
see [ui.Dialog.setDialogAttribute](../15_library-reference/3225-ui-dialog-setdialogattribute.md "Set an attribute to configure a dynamic dialog.").

## Related links

**Related concepts**  

[Instantiate a dynamic dialog](2428-instantiate-a-dynamic-dialog.md "The dynamic dialogs needs to be created with specific ui.Dialog methods.")
