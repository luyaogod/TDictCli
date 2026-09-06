---
title: "ui.Dialog.getFieldBuffer"
source: "fgl-topics/c_fgl_ClassDialog_getFieldBuffer.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getFieldBuffer"
type: "concept"
---

# ui.Dialog.getFieldBuffer

> Returns the input buffer of the specified field.

## Syntax

```
getFieldBuffer(
   name STRING )
  RETURNS STRING
```

1. name is the form field name, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

The `getFieldBuffer()` method returns the current input buffer of the specified
field. The input buffer is used by the dialog to synchronize form fields and program variables.
In some situations, especially when using the [buffered mode](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") or in a [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form."), you may want to access the
field input buffer.

Use of the `getFieldBuffer()` method is only recommended in dialogs allowing field
input (`INPUT`, `INPUT ARRAY`, `CONSTRUCT`). The
behavior is undefined when used in `DISPLAY ARRAY`.

The parameter is a field specification, a string containing the field qualifier, with an optional
prefix ("`[table.]column`").

```
LET buff = DIALOG.getFieldBuffer("customer.cust_name")
```

The input buffer can be set with:

- A [`DISPLAY TO`](../11_user-interface/1880-display-to.md "The DISPLAY ... TO instruction displays data to specific form fields.") or
  [`DISPLAY BY NAME`](../11_user-interface/1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.") instruction
- The [`FGL_DIALOG_SETBUFFER()`](2744-fgl-dialog-setbuffer.md "Sets the input buffer of the current field.") function (only for the current field)
