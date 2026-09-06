---
title: "SAMPLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_SAMPLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > SAMPLE attribute"
type: "concept"
---

# SAMPLE attribute

> The SAMPLE attribute defines the text to be used to compute the width of a form field widget.

> **Important:**
>
> The `SAMPLE` attribute is not yet supported.

## Syntax

```
SAMPLE = "text"
```

1. text is the sample string that will be used to compute the width of the form
   field widget.

## Usage

By default, form fields are rendered by the client with a size determined by the current font and
the number of cells used by the form tag in the layout grid. The field width is computed so that the
largest value can fit in the widget.

Sometimes, the default computed width is too wide for the typical values displayed in the field.
For example, numeric fields usually need less space than alphanumeric fields. If the values are
always smaller, you can use the `SAMPLE` attribute to provide a hint for the
front-end to compute the best width for that form field.

When specifying the `SAMPLE` attribute, you do not have to fill the sample
string up to the width of the corresponding field tag: The front-ends will be able to compute a
physical width by applying a ratio to fit the best visual result. For example, for a sample of
`"XY"` used for a field defined with 10 characters, is equivalent to specifying a
sample of `"XYXYXYXYXY"`.

By default, when no `SAMPLE` attribute is used, the first 6 cells are always
computed with the pixel width of the `"M"` character in the current font. Next cells
are computed with the pixel width of the `"0"` (zero) character. In other words, the
default sample model is `"MMMMMM000000....."`, reduced to the size of the field tag
in the
layout:

```
   -123456789-123456789-         Default SAMPLE
   [f01 ]                        MMMM
   [f02   ]                      MMMMMM
   [f03             ]            MMMMMM0000000000
```

A default sample can be specified for all fields used in the form, by using the `DEFAULT
SAMPLE` option in the [`INSTRUCTIONS`](1751-instructions-section.md "The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties.")
section:

```
INSTRUCTIONS
DEFAULT SAMPLE = "0"
END
```

See also [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Example

```
EDIT cid = customer.custid, SAMPLE="0000";
EDIT ccode = customer.ucode, SAMPLE="MMMMMM";
DATEEDIT be01 = customer.created, SAMPLE="00-00-0000";
```

## Related links

**Related concepts**  

[INSTRUCTIONS section](1751-instructions-section.md "The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties.")
