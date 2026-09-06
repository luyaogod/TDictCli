---
title: "ui.Dialog.addInputArrayFrom"
source: "fgl-topics/c_fgl_ClassDialog_addInputArrayFrom.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.addInputArrayFrom"
type: "concept"
---

# ui.Dialog.addInputArrayFrom

> Adds a sub-dialog of type INPUT ARRAY FROM to an existing ui.Dialog dynamic dialog.

## Syntax

```
ui.Dialog.addInputArrayFrom(
   fields DYNAMIC ARRAY OF RECORD
                        name STRING,
                        type STRING
                    END RECORD,
   screenRecord STRING )
```

1. fields is the list of form fields controlled by the dialog.
   This must be a `DYNAMIC ARRAY` of a `RECORD`
   structure, with a `name` and `type` member of type
   `STRING`.
2. screenRecord is the name of the screen array (defined with the
   `SCREEN RECORD` instruction in form files). This name will also identify
   the sub-dialog in dialog events and API calls.

## Usage

The `ui.Dialog.addInputArrayFrom()` class method adds a sub-dialog equivalent to a
`INPUT ARRAY FROM` block, to the dynamic multiple dialog created with [ui.Dialog.createMultipleDialog](3175-ui-dialog-createmultipledialog.md "Creates an ui.Dialog object to implement a dynamic DIALOG multiple-dialog.").

The method takes a list of field definitions as parameter, as described in [Field definition for Dynamic Dialogs](3241-field-definition-for-dynamic-dialogs.md).

The second parameter passed to the `addInputArrayFrom()` method is the name of the
screen record which groups the fields together, for the list view of the form. This second parameter
will also be used to identify the sub-dialog.

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
LET d = ui.Dialog.createMultipleDialog()
...
CALL d.addInputArrayFrom(fields, "sr_custlist")
...
```

## Related links

**Related concepts**  

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
