---
title: "REQUIRED attribute"
source: "fgl-topics/c_fgl_FSFAttributes_REQUIRED.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > REQUIRED attribute"
type: "concept"
---

# REQUIRED attribute

> The REQUIRED attribute forces the user to modify the content of a field during an input dialog.

## Syntax

```
REQUIRED
```

## Usage

The `REQUIRED` attribute forces the user to modify the content of a field
controlled by an input dialog (`INPUT` or `INPUT ARRAY`), when the
`INPUT` dialog does not use the [`WITHOUT DEFAULTS`](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.") option. Within
`INPUT ARRAY`, the `REQUIRED` attribute always applies to newly
created rows.

If an [`INPUT`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") dialog uses the
`WITHOUT DEFAULTS` clause, the current value of the variable linked to the
`REQUIRED` field is considered as a default value; the runtime system assumes that
the field satisfies the `REQUIRED` attribute, even if the variable value is null.

In an [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") dialog, the
`REQUIRED` attribute always applies to newly created rows, even if `WITHOUT
DEFAULTS` is used. In other words, when creating a new row, `INPUT ARRAY`
behaves like `INPUT` without the `WITHOUT DEFAULTS` clause.

If `REQUIRED` is effective (`WITHOUT DEFAULTS` is not used), and a
[`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attribute is defined
for the field, or the program variable corresponding to the field is initialized with a non-NULL
value in `BEFORE INPUT`, `BEFORE INSERT` or `BEFORE
DIALOG`, the runtime system assigns the default value to the field, and assumes that the
`REQUIRED` attribute is satisfied.

The `REQUIRED` attribute does not prevent fields being null; If the field contains
a value, and the user subsequently erases this value during the same input, the runtime system
considers the `REQUIRED` attribute satisfied. To insist on a non-null entry, use the
[`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") attribute.

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
EDIT f001 = orders.ord_shipcmt, REQUIRED;
```

## Related links

**Related concepts**  

[Form-level validation rules](2240-form-level-validation-rules.md "Form-level validation rules can be defined for each field controlled by a dialog.")

[ITEMS attribute](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.")
