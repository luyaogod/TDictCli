---
title: "ui.Dialog.getFieldTouched"
source: "fgl-topics/c_fgl_ClassDialog_getFieldTouched.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getFieldTouched"
type: "concept"
---

# ui.Dialog.getFieldTouched

> Returns the modification flag for a field.

## Syntax

```
getFieldTouched(
   formFieldList STRING )
  RETURNS BOOLEAN
```

1. formFieldList is the string with the list of field specification, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

The `getFieldTouched()` method returns `TRUE` if
the [modification flag](../11_user-interface/2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") of
the specified field(s) is set.

The formFieldList parameter is a string containing the field qualifier, with
an optional prefix ("`[table.]column`"), a table prefix followed by a dot and an
asterisk ("`table.*`"), or a simple asterisk ("`*`").

This code checks if a specific field has been touched:

```
AFTER FIELD cust_name 
  IF DIALOG.getFieldTouched("customer.cust_address") THEN
     ...
```

If the parameter is a screen record following by dot-asterisk, the method checks
the touched flags of all the fields that belong to the screen record:

```
ON ACTION quit 
  IF DIALOG.getFieldTouched("customer.*") THEN
    ...
```

When passing a simple asterisk (\*) to the method, the runtime system will check
all fields used by the dialog:

```
ON ACTION quit 
  IF DIALOG.getFieldTouched("*") THEN
    ...
```

## Related links

**Related concepts**  

[ui.Dialog.setFieldTouched](3227-ui-dialog-setfieldtouched.md "Sets the modification flag of the specified field.")
