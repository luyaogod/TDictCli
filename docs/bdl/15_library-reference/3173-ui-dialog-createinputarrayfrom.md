---
title: "ui.Dialog.createInputArrayFrom"
source: "fgl-topics/c_fgl_ClassDialog_createInputArrayFrom.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.createInputArrayFrom"
type: "concept"
---

# ui.Dialog.createInputArrayFrom

> Creates an ui.Dialog object to implement a dynamic INPUT ARRAY FROM.

## Syntax

```
ui.Dialog.createInputArrayFrom(
   fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
                    END RECORD,
   screenRecord STRING )
  RETURNS ui.Dialog
```

1. fields is the list of form fields controlled by the dialog.
   This must be a `DYNAMIC ARRAY` of a `RECORD`
   structure, with a `name` and `type` member of type
   `STRING`.
2. screenRecord is the name of the screen array (defined with
   the `SCREEN RECORD` instruction in form files).
3. The method returns a new `ui.Dialog` object.

## Usage

The `ui.Dialog.createInputArrayFrom()` class method creates a dialog object to
implement the equivalent of a static [`INPUT ARRAY
FROM`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") block.

The [current form](../11_user-interface/1562-understanding-windows-and-forms.md "This is an introduction to Genero windows and forms.") will be attached to the
created dialog.

The method takes a list of field definitions as parameter, as described in [Field definition for Dynamic Dialogs](3241-field-definition-for-dynamic-dialogs.md).

A dynamic input array dialog behaves like a static `INPUT ARRAY` using the
`WITHOUT DEFAULTS` option: The values set in the internal rows before starting the
dialog will be used. However, like with a static `INPUT ARRAY`, when adding a new
row, the `DEFAULT` attributes of the form-fields are used.

The second parameter passed to the `createInputArrayFrom()` method is
the name of the screen record which groups the fields together, for the list view of the form.

For example, in the following form definition, the screen record name is
"`sr_custlist`":

```
...
INSTRUCTIONS
SCREEN RECORD sr_custlist
(
 customer.cust_id,
 customer.cust_name,
 ...
);
END
```

For more details, see [Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.").

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
LET d = ui.Dialog.createInputArrayFrom(fields, "sr_custlist")
...
```

## Related links

**Related concepts**  

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
