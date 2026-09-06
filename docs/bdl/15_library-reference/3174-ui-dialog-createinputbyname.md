---
title: "ui.Dialog.createInputByName"
source: "fgl-topics/c_fgl_ClassDialog_createInputByName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.createInputByName"
type: "concept"
---

# ui.Dialog.createInputByName

> Creates an ui.Dialog object to implement a dynamic INPUT BY NAME.

## Syntax

```
ui.Dialog.createInputByName(
   fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
                    END RECORD
   )
  RETURNS ui.Dialog
```

1. fields is the list of form fields controlled by the dialog. This must be a
   `DYNAMIC ARRAY` of a `RECORD` structure, with a `name`
   and `type` member of type `STRING`.
2. The method returns a new `ui.Dialog` object.

## Usage

The `ui.Dialog.createInputByName()` class method creates a dialog object to
implement the equivalent of a static [`INPUT BY
NAME`](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") block.

The [current form](../11_user-interface/1562-understanding-windows-and-forms.md "This is an introduction to Genero windows and forms.") will be attached to the
created dialog.

The method takes a list of field definitions as parameter, as described in [Field definition for Dynamic Dialogs](3241-field-definition-for-dynamic-dialogs.md).

A dynamic input dialog behaves like a static `INPUT` dialog using the
`WITHOUT DEFAULTS` option: The `DEFAULT` attribute of the form-field
is not used.

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
LET d = ui.Dialog.createInputByName(fields)
...
```

## Related links

**Related concepts**  

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
