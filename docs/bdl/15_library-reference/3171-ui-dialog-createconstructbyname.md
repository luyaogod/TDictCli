---
title: "ui.Dialog.createConstructByName"
source: "fgl-topics/c_fgl_ClassDialog_createConstructByName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.createConstructByName"
type: "concept"
---

# ui.Dialog.createConstructByName

> Creates an ui.Dialog object to implement a dynamic CONSTRUCT BY NAME.

## Syntax

```
ui.Dialog.createConstructByName(
   fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
                    END RECORD
   )
  RETURNS ui.Dialog
```

1. fields is the list of form fields controlled by the dialog. This must be a
   `DYNAMIC ARRAY` of `RECORD` structure, with a
   `name` and `type` member of type
   `STRING`.
2. The method returns a new `ui.Dialog` object.

## Usage

The `ui.Dialog.createConstructByName()` class method creates a new dialog object
to implement the equivalent of a static [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") block.

The [current form](../11_user-interface/1562-understanding-windows-and-forms.md "This is an introduction to Genero windows and forms.") will be attached to the
created dialog.

The method takes a list of field definitions as parameter, as described in [Field definition for Dynamic Dialogs](3241-field-definition-for-dynamic-dialogs.md).

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
LET d = ui.Dialog.createConstructByName(fields)
...
```

## Related links

**Related concepts**  

[ui.Dialog.getQueryFromField](3201-ui-dialog-getqueryfromfield.md "Returns the SQL condition of a field used in a query by example dialog.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
