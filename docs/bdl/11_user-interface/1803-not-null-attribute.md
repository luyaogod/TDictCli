---
title: "NOT NULL attribute"
source: "fgl-topics/c_fgl_FSFAttributes_NOT_NULL.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > NOT NULL attribute"
type: "concept"
---

# NOT NULL attribute

> The NOT NULL attribute specifies that the field does not accept NULL values.

## Syntax

```
NOT NULL
```

## Usage

The `NOT NULL` attribute requires that the field contains a non-null value. It can
be specified explicitly in the form field definition, or in the corresponding column definition in
the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."). If no column is
associated to the field, the `NOT NULL` attribute can also be used in the type
definition of [`FORMONLY`
fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.").

> **Important:**
>
> For maximum security, when the data is stored in a database, define a `NOT NULL`
> constraint on the SQL column corresponding to the field, in order to deny `NULL`
> values, when an `INSERT` or `UPDATE` statement is executed.

The `NOT NULL` attribute is effective only when the field name appears in the
list of screen fields of an `INPUT` or `INPUT ARRAY`
statement.

If a `DEFAULT` attribute is used for the field and the input dialog does not
use the `WITHOUT DEFAULTS` option, the runtime system assumes that
the default value satisfies the `NOT NULL` attribute.

Unlike the `REQUIRED` attribute which has no effect when the
`INPUT` dialog uses the `WITHOUT DEFAULTS` option,
the `NOT NULL` attribute is always checked when validating a
dialog.

> **Tip:**
>
> For [accessibility
> compliance](1589-accessibility-guidelines.md "This section describes the best practices to make your application accessible to disabled people."), when a form field is defined with the [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") attribute, the GBC
> front-end automatically sets the HTML attribute [WAI-ARIA `aria-required`](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-required) to
> `true`. Conversely, the `aria-required` attribute is not set when
> using the [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") form field
> attribute, as its functional purpose differs from standard ARIA requirement states.

## Example

```
EDIT f001 = customer.city, NOT NULL;
```

## Related links

**Related concepts**  

[Form-level validation rules](2240-form-level-validation-rules.md "Form-level validation rules can be defined for each field controlled by a dialog.")

[Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.")

[DEFAULT attribute](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.")

[REQUIRED attribute](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.")
