---
title: "ui.Dialog.setFieldTouched"
source: "fgl-topics/c_fgl_ClassDialog_setFieldTouched.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setFieldTouched"
type: "concept"
---

# ui.Dialog.setFieldTouched

> Sets the modification flag of the specified field.

## Syntax

```
setFieldTouched(
   formFieldList STRING,
   val BOOLEAN )
```

1. formFieldList is the string with the list of field specification, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).
2. val is the boolean value to set the modification
   flag.

## Usage

The `setFieldTouched()` method can be used to change the [modification flag](../11_user-interface/2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") of the specified
field(s).

The formFieldList is a string containing the field qualifier, with an optional
prefix ("`[table.]column`"), or a table prefix followed by a dot and an asterisk
("`table.*`").

You typically use this method to set the touched flag when assigning a variable,
to emulate user input. Remember when using the [`UNBUFFERED`](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") mode,
you don't need to display the value to the fields.
The `setFieldTouched()` method is provided as a 3GL replacement for
the [`DISPLAY BY NAME
/ TO`](../11_user-interface/1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.") instructions to set the modification flags.

```
ON ACTION zoom_city 
  LET p_cust.cust_city = zoom_city()
  CALL DIALOG.setFieldTouched("customer.cust_city", TRUE)
    ...
```

If the parameter is a screen record followed by dot-asterisk, the method checks the
modification flags of all the fields that belong to the screen record. You typically
use this to reset the touched flags of a group of fields, after modifications have
been saved to the database, to get back to the initial state of the
dialog:

```
ON ACTION save 
  CALL save_cust_record()
  CALL DIALOG.setFieldTouched("customer.*", FALSE)
    ...
```

The modification flags are reset to false when using an `INPUT ARRAY`
list, every time you leave the modified row.

## Related links

**Related concepts**  

[ui.Dialog.getFieldTouched](3198-ui-dialog-getfieldtouched.md "Returns the modification flag for a field.")
