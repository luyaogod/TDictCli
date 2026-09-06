---
title: "NOENTRY attribute"
source: "fgl-topics/c_fgl_FSFAttributes_NOENTRY.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > NOENTRY attribute"
type: "concept"
---

# NOENTRY attribute

> The NOENTRY attribute prevents data entry in the field during an input dialog.

## Syntax

```
NOENTRY
```

## Usage

Use the `NOENTRY` attribute to bypass field input during an `INPUT`
or `INPUT ARRAY` statement.

A `NOENTRY` field is like a field permanently disabled with a [`DIALOG.setFieldActive("fieldname",FALSE)`](../15_library-reference/3226-ui-dialog-setfieldactive.md "Enable and disable form fields.") call, it cannot get the focus.

When compiling a form with a field referencing a `SERIAL` or
`BIGSERIAL` column in the database schema, the `NOENTRY` attribute is
automatically set. However, the attribute will not be set if the field is defined with a
`TYPE LIKE` syntax.

When using a `WITHOUT DEFAULTS` dialog option, the content of the corresponding
program variable is displayed in the field.

The `NOENTRY` attribute does not prevent data entry into a field during a
`CONSTRUCT` statement.

## Example

```
EDIT f001 = order.totamount, NOENTRY;
```

## Related links

**Related concepts**  

[Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.")
