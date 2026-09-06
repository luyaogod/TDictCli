---
title: "ui.Dialog.setFieldActive"
source: "fgl-topics/c_fgl_ClassDialog_setFieldActive.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setFieldActive"
type: "concept"
---

# ui.Dialog.setFieldActive

> Enable and disable form fields.

## Syntax

```
setFieldActive(
   formFieldList STRING,
   val BOOLEAN )
```

1. formFieldList is the string with the list of field specification, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).
2. val is a boolean value.

## Usage

The `setFieldActive()` method can be used to enable / disable
form fields.

The formFieldList is a string containing the field qualifier, with an optional
prefix ("`[table.]column`"), or a table prefix followed by a dot and an asterisk
("`table.*`").

```
CALL DIALOG.setFieldActive( "customer.cust_addr",
            (rec.cust_name IS NOT NULL) )
```

Do not disable or hide all fields of a dialog, otherwise the dialog
execution stops: At least one field must be able to get the focus during a dialog execution.

Do not disable/hide the current field having the focus. When
disabling or when hiding the current field, the `AFTER FIELD` block of the current
field and the `BEFORE FIELD` block on the next field in the tabbing order will be
executed. In a multiple dialog, if the next field in the tabbing order is in another sub-dialog, the
`AFTER INPUT` block of the current sub-dialog and the `BEFORE
INPUT/CONSTRUCT/DISPLAY` or the next sub-dialog are executed. In all these control blocks, a
`NEXT FIELD` to the field that is disabled or hidden will result in an endless
loop.
