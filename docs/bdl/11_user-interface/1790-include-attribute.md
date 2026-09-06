---
title: "INCLUDE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_INCLUDE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > INCLUDE attribute"
type: "concept"
---

# INCLUDE attribute

> The INCLUDE attribute defines a list of possible values for a field.

## Syntax

```
INCLUDE = (
  { NULL
  | literal [ TO literal]
  } [,...]
 )
```

1. literal can be any literal expression supported by the form compiler.

## Usage

The `INCLUDE` attribute specifies acceptable values for a field and causes the
runtime system to check the data before accepting an input value.
> **Important:**
>
> For maximum security, when the data is stored in a database, define a `CHECK`
> constraint on the SQL column corresponding to the field, in order to deny values that are not in the
> `INCLUDE` list, when an `INSERT` or `UPDATE` statement
> is executed.

If the field is `FORMONLY`, you must also specify a data type when you assign the
`INCLUDE` attribute to a field.

Include the `NULL` keyword in the value list to specify that it is acceptable for
the user to leave the field without entering any value.
> **Tip:**
>
> When initializing variables before an new record input with an
> `INPUT` or `INPUT ARRAY` dialog, if a field must have no default
> value, set the corresponding variable to `NULL`. If the value contains only space
> characters (ASCII 32), or any other non-visible whitespace characters such as TAB (ASCII 9), this
> will be considered as a non-null value, when the field defines an [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") attribute with
> `NULL` in the possible values:
>
> ```
> EDIT f02 = order.ord_valid TYPE CHAR, INCLUDE=("Y","N",NULL)
> ```
>
> With a space or
> TAB character, the field will appear empty to the user, but the current value will not match the
> `INCLUDE` constraint. Note that fields with a boolean value such as
> `"Y"/"N"` should be initialized with one of these values and deny nulls with [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."), except in some rare cases
> where the status is undefined.

Use the `TO` keyword to specify an inclusive range of acceptable values. When
specifying a range of values, the lower value must appear first. The field value is accepted if it
is greater or equal to the first literal, and lower or equal to the second literal: `INCLUDE
= (1 TO 999)` is equivalent to `(value>=1 AND
value<=999)`. Special consideration must be taken for character string
fields when using a `TO` range: `INCLUDE = ("AAA" TO "ZZZ")` is
`(value = "AAA" AND value<="ZZZ")`. The
value `"ABC"` is accepted, but values such as `"A!!"` or
`"Zaa"` are denied. When combining several ranges and single values, the value
entered by the user is verified for each element of the `INCLUDE` attribute:
`INCLUDE = (1 TO 999, -1, NULL)` is equivalent to
`(value>=1 AND value<=999) OR
(value==-1) OR (value IS NULL)`.

## Example

```
EDIT f001 = compute.rate, INCLUDE = ( 1 TO 100, 200, NULL);
EDIT f002 = customer.state, INCLUDE = ( "AL" TO "GA", "IA" TO "WY" );
EDIT f003 = FORMONLY.valid TYPE CHAR, INCLUDE = ("Y","N");
```

## Related links

**Related concepts**  

[Form-level validation rules](2240-form-level-validation-rules.md "Form-level validation rules can be defined for each field controlled by a dialog.")

[Formonly fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.")

[ITEMS attribute](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.")
