---
title: "Identifying fields in ui.Dialog methods"
source: "fgl-topics/c_fgl_ClassDialog_identify_field.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Usage > Identifying fields in ui.Dialog methods"
type: "concept"
description: "Form fields are bound to program variables with the binding clause of the dialog instruction ( INPUT variable-list FROM field-list , INPUT BY NAME variable-list , CONSTRUCT BY NAME sql ON column-list ..."
---

# Identifying fields in ui.Dialog methods

Form fields are bound to program variables with the binding clause of the dialog instruction
(`INPUT variable-list FROM field-list`, `INPUT BY NAME
variable-list`, `CONSTRUCT BY NAME sql ON
column-list`,`CONSTRUCT sql ON column-list FROM
field-list`, `INPUT ARRAY array-name FROM
screen-array.*`).

In `ui.Dialog` methods such as [`setFieldActive()`](3226-ui-dialog-setfieldactive.md "Enable and disable form fields."), the first parameter identifies the form field (or, for
some methods, a list of fields) to be processed. This parameter is the form field name specification
(the name of the field in the .per form), not the program variable name used by
the dialog.

If none of the current dialog fields matches the field specification, the error [-1373](4483-genero-bdl-errors.md) is thrown by the method.

The form field names can be fully-qualified with a prefix or partly-qualified by using a single
name without prefix.

The field name specification can be any of the following:

- field-name
- table-name.field-name
- screen-record-name.field-name
- `FORMONLY`.field-name

Here are some examples:

- `"cust_name"`,
- `"customer.cust_name"`,
- `"cust_screen_record.cust_name"`,
- `"item_screen_array.item_label"`,
- `"formonly.total"`,
- `"customer.*"` (only some methods accept the "dot
  asterisk" notation)

When no prefix is used in the field name specification, the first form field matching the name
will be used. Inside a `DIALOG` block, if several sub-dialogs use the same field
name, the method will use the field of the current sub-dialog.

When using a prefix in the field name specification, it must be equal to the field prefix
used in the field binding clause of the dialog.

- When no screen-record was specified in the field binding clause (for example, when using
  `INPUT BY NAME variable-list`), the field prefix must be the database
  table name (or `FORMONLY`) used in the form file, or any valid screen-record using
  that field.
- When the `FROM` or `TO` clause of the dialog specifies an explicit
  screen-record (for example, in `INPUT variable-list FROM screen-record.* /
  field-list-with-screen-record-prefix`, `INPUT ARRAY array-name FROM
  screen-array.*`) or `DISPLAY ARRAY array-name TO
  screen-array.*`, the field prefix must be the screen-record name
  used in the `FROM`/`TO` clause.

The name of the field passed as parameter can use the same letter case as the field definition in
the form file: The lookup is
case-insensitive:

```
-- In the form file:
EDIT f01 = Customer.CustAddr, ... ;
-- In the program code:
CALL DIALOG.setFieldActive("Customer.CustAddr", FALSE)
```

The methods [`validate()`](3233-ui-dialog-validate.md "Checks form level validation rules."),
[`setFieldActive()`](3226-ui-dialog-setfieldactive.md "Enable and disable form fields."), [`setFieldTouched()`](3227-ui-dialog-setfieldtouched.md "Sets the modification flag of the specified field."), [`getFieldTouched()`](3198-ui-dialog-getfieldtouched.md "Returns the modification flag for a field.") can take a
list of fields as parameter, by using the "dot-asterisk " notation
(`screen-record.*`). This way you can check, query or change a complete list of
fields in one method
call:

```
ON ACTION save 
  CALL save_cust_record()
  CALL DIALOG.setFieldTouched("customer.*", FALSE)
    ...
```
