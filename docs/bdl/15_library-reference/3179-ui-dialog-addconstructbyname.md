---
title: "ui.Dialog.addConstructByName"
source: "fgl-topics/c_fgl_ClassDialog_addConstructByName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.addConstructByName"
type: "concept"
---

# ui.Dialog.addConstructByName

> Adds a sub-dialog of type CONSTRUCT BY NAME to an existing ui.Dialog dynamic dialog.

## Syntax

```
ui.Dialog.addConstructByName(
   fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
                    END RECORD,
   name STRING
   )
```

1. fields is the list of form fields controlled by the dialog. This must be a
   `DYNAMIC ARRAY` of `RECORD` structure, with a
   `name` and `type` member of type
   `STRING`.
2. name is the sub-dialog identifier. This name is used to identify the
   sub-dialog in dialog events and API calls.

## Usage

The `ui.Dialog.addConstructByName()` class method adds a sub-dialog equivalent to
a `CONSTRUCT BY NAME` block, to the dynamic multiple dialog created with [ui.Dialog.createMultipleDialog](3175-ui-dialog-createmultipledialog.md "Creates an ui.Dialog object to implement a dynamic DIALOG multiple-dialog.").

The method takes a list of field definitions as parameter, as described in [Field definition for Dynamic Dialogs](3241-field-definition-for-dynamic-dialogs.md).

This second parameter will be used to identify the sub-dialog.

> **Important:**
>
> Static `CONSTRUCT` dialog instructions use the data types of
> the fields defined in the .per form specification file. Unlike static
> CONSTRUCT, the dynamic construct uses the data type provided in the dynamic
> array defining the list of fields.

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
...
CALL d.addConstructByName(fields,"const1")
...
```

## Related links

**Related concepts**  

[ui.Dialog.getQueryFromField](3201-ui-dialog-getqueryfromfield.md "Returns the SQL condition of a field used in a query by example dialog.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
