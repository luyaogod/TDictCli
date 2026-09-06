---
title: "PHANTOM item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_PHANTOM_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > PHANTOM item definition"
type: "concept"
---

# PHANTOM item definition

> Defines a form field, that must not be displayed to the end user.

## Syntax

```
PHANTOM [field-name] ;
```

1. field-name identifies the name of the screen record field.

## Usage

Define a phantom form field (that will be used by a dialog, but not displayed in the
form layout), with a `PHANTOM` element in the
`ATTRIBUTES` section.

For more details, see [Phantom fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).").

## Example

```
PHANTOM customer.cust_name;
```

## Related links

**Related concepts**  

[Database column fields](1671-database-column-fields.md "Form fields defined with a table and column name get data type from the database schema file.")

[Formonly fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.")
