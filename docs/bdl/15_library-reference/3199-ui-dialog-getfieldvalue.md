---
title: "ui.Dialog.getFieldValue"
source: "fgl-topics/c_fgl_ClassDialog_getFieldValue.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getFieldValue"
type: "concept"
---

# ui.Dialog.getFieldValue

> Returns the value of a field controlled by a dynamic dialog.

## Syntax

```
getFieldValue(
   name STRING )
  RETURNS fgl-type
```

1. name is the name of the form field, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).
2. fgl-type is one of the [primitive data
   types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").

## Usage

The `getFieldValue()` method can be used when implementing a dynamic dialog, to
return the value of a field:

```
DISPLAY d.getFieldValue( "customer.cust_addr" )
```

In a dynamic dialog controlling a list of records (`INPUT ARRAY` /
`DISPLAY ARRAY`), this method returns the value for a field in the current row.

During dialog execution, the `getFieldValue()` method must only be used to get the
value of a field for the current row. In a regular dynamic list dialog, calling the [`setCurrentRow()`](3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") method to change
the current row before calling `getFieldValue()` will have no effect. In a
`DISPLAY ARRAY` using the [paged
mode](../11_user-interface/2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.") (`ON FILL BUFFER`), `getFieldValue()` returns
`NULL`, if the current row is not part of the visible page. In a `DISPLAY
ARRAY` using the [full list mode](../11_user-interface/2308-full-list-mode-of-display-array.md "In order to handle short/medium result sets, use the full list mode of DISPLAY ARRAY."), a
`NULL` value is returned, if there is no current row (when the array is empty).

## Related links

**Related concepts**  

[ui.Dialog.setFieldValue](3228-ui-dialog-setfieldvalue.md "Sets the value of a field controlled by the dialog object.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
