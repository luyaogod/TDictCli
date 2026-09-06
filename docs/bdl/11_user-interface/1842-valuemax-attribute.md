---
title: "VALUEMAX attribute"
source: "fgl-topics/c_fgl_FSFAttributes_VALUEMAX.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > VALUEMAX attribute"
type: "concept"
---

# VALUEMAX attribute

> The VALUEMAX attribute defines a upper limit of values displayed in widgets (such as progress bars).

## Syntax

```
VALUEMAX = integer
```

1. integer is an integer literal.

## Usage

This attribute is typically used to define the upper limit in `PROGRESSBAR`,
`SPINEDIT` and `SLIDER` fields.

Default values for `VALUEMIN`/`VALUEMAX` depend on
the type of the widget:

- For a [`PROGRESSBAR`](1698-progressbar-item-type.md "Defines a progress indicator field."), the
  default is `VALUEMIN=0`, `VALUEMAX=100`.
- For a [`SLIDER`](1701-slider-item-type.md "Defines a slider form item."), the default
  is `VALUEMIN=0`, `VALUEMAX=5`.
- For a [`SPINEDIT`](1702-spinedit-item-type.md "Defines a spin box widget to enter integer values."), there is
  no default.

This attribute is not used by the runtime system to validate the field. Use
the `INCLUDE` attribute to
control value boundaries.

> **Important:**
>
> For maximum security, when the data is stored in a database, define a `CHECK`
> constraint on the SQL column corresponding to the field, in order to allow only values matching the
> `VALUEMIN` / `VALUEMAX` attributes, when an `INSERT` or
> `UPDATE` statement is executed.

## Example

```
SLIDER s01 = FORMONLY.slider01,
                        VALUEMIN=0,
                        VALUEMAX=500;
```

## Related links

**Related concepts**  

[VALUEMIN attribute](1841-valuemin-attribute.md "The VALUEMIN attribute defines a lower limit of values displayed in widgets (such as progress bars).")
