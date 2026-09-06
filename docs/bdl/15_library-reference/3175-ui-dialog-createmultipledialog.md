---
title: "ui.Dialog.createMultipleDialog"
source: "fgl-topics/c_fgl_ClassDialog_createMultipleDialog.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.createMultipleDialog"
type: "concept"
---

# ui.Dialog.createMultipleDialog

> Creates an ui.Dialog object to implement a dynamic DIALOG multiple-dialog.

## Syntax

```
ui.Dialog.createMultipleDialog( )
  RETURNS ui.Dialog
```

## Usage

The `ui.Dialog.createMultipleDialog()` class method creates a dialog object to
implement the equivalent of a static [`DIALOG
/ END DIALOG`](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.") block.

The [current form](../11_user-interface/1562-understanding-windows-and-forms.md "This is an introduction to Genero windows and forms.") will be attached to the
created dialog.

The method takes no parameters.

After creating the multiple dialog object, you must add sub-dialogs with the following
methods:

- [ui.Dialog.addConstructByName](3179-ui-dialog-addconstructbyname.md "Adds a sub-dialog of type CONSTRUCT BY NAME to an existing ui.Dialog dynamic dialog.")
- [ui.Dialog.addDisplayArrayTo](3180-ui-dialog-adddisplayarrayto.md "Adds a sub-dialog of type DISPLAY ARRAY TO to an existing ui.Dialog dynamic dialog.")
- [ui.Dialog.addInputArrayFrom](3182-ui-dialog-addinputarrayfrom.md "Adds a sub-dialog of type INPUT ARRAY FROM to an existing ui.Dialog dynamic dialog.")
- [ui.Dialog.addInputByName](3183-ui-dialog-addinputbyname.md "Adds a sub-dialog of type INPUT BY NAME to an existing ui.Dialog dynamic dialog.")

Use the `addTrigger()` method, to add global or sub-dialog triggers. The scope of
the trigger is defined by the `addTrigger()` call order. See [`addTrigger()`](3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog") for more details.

A dynamic dialog can be configured with the [`setDialogAttribute()`](3225-ui-dialog-setdialogattribute.md "Set an attribute to configure a dynamic dialog.")
method.

## Example

```
DEFINE fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
              END RECORD
DEFINE d ui.Dialog
...
LET d = ui.Dialog.createMultipleDialog()
CALL d.addTrigger("ON ACTION close") -- Adds a global trigger for the dialog
...
CALL d.addDisplayArrayTo(fields, "sr_custlist") -- Adds a DISPLAY ARRAY sub-dialog
CALL d.addTrigger("ON ACTION refresh") -- Adds a trigger for the DISPLAY ARRAY sub-dialog
...
```

## Related links

**Related concepts**  

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
