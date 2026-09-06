---
title: "VALUEUNCHECKED attribute"
source: "fgl-topics/c_fgl_FSFAttributes_VALUEUNCHECKED.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > VALUEUNCHECKED attribute"
type: "concept"
---

# VALUEUNCHECKED attribute

> The VALUEUNCHECKED attribute defines the value associated with a checkbox item when it is not checked.

## Syntax

```
VALUEUNCHECKED = value
```

1. value is a numeric or string literal, or one
   of the following keywords: `NULL`, `TRUE`, `FALSE`.

## Usage

This attribute is used together with the `VALUECHECKED` attribute to define the
values corresponding to the states of a [`CHECKBOX`](1686-checkbox-item-type.md "Defines a boolean or three-state checkbox field.").

This attribute is not used by the runtime system to validate the field. Use
the `INCLUDE` attribute to
control value boundaries.

> **Important:**
>
> For maximum security, when the data is stored in a database, define a `CHECK`
> constraint on the SQL column corresponding to the field, in order to allow only values matching the
> `VALUECHECKED` / `VALUEUNCHECKED` attributes, when an
> `INSERT` or `UPDATE` statement is executed.

## Example

```
CHECKBOX cb01 = FORMONLY.checkbox01,
                TEXT="OK",
                VALUECHECKED="Y",
                VALUEUNCHECKED="N";
```

## Related links

**Related concepts**  

[VALUECHECKED attribute](1843-valuechecked-attribute.md "The VALUECHECKED attribute defines the value associated with a checkbox item when it is checked.")
