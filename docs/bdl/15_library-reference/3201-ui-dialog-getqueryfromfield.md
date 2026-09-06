---
title: "ui.Dialog.getQueryFromField"
source: "fgl-topics/c_fgl_ClassDialog_getQueryFromField.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getQueryFromField"
type: "concept"
---

# ui.Dialog.getQueryFromField

> Returns the SQL condition of a field used in a query by example dialog.

## Syntax

```
getQueryFromField(
   name STRING )
  RETURNS STRING
```

1. name is the name of the form field, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

The `getQueryFromField()` method generates the SQL condition from the value
entered in the field specified by the field-name parameter.

This method is used in the context of a [construct dynamic dialog](3171-ui-dialog-createconstructbyname.md "Creates an ui.Dialog object to implement a dynamic CONSTRUCT BY NAME.").

The result of this method can be used to build the `WHERE` part of a
`SELECT` statement to find rows in a database.

Collect and concatenate field conditions returned from `getQueryFromField()`, then
add `AND` or `OR` boolean operators to create an executable SQL
query.

The SQL condition is generated based on the current type of database. The SQL syntax may vary
depending on the target database. Therefore it is not recommended to reuse the generated SQL
conditions. However, the user input of a query by example dialog can be reused for different types
of databases (see [ui.Dialog.setFieldValue](3228-ui-dialog-setfieldvalue.md "Sets the value of a field controlled by the dialog object.") and [ui.Dialog.getFieldValue](3199-ui-dialog-getfieldvalue.md "Returns the value of a field controlled by a dynamic dialog."))

## Related links

**Related concepts**  

[ui.Dialog.createConstructByName](3171-ui-dialog-createconstructbyname.md "Creates an ui.Dialog object to implement a dynamic CONSTRUCT BY NAME.")
