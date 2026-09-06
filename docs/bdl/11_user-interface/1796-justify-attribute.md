---
title: "JUSTIFY attribute"
source: "fgl-topics/c_fgl_FSFAttributes_JUSTIFY.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > JUSTIFY attribute"
type: "concept"
---

# JUSTIFY attribute

> The JUSTIFY attribute defines the alignment of a text field content, and table column headers.

## Syntax

```
JUSTIFY = { LEFT | CENTER | RIGHT }
```

## Usage

With the `JUSTIFY` attribute, you specify the justification of the content of a
text field as `LEFT`, `CENTER` or `RIGHT` when the
field is in display state.

`JUSTIFY` has no effect on the content of form item types like
`IMAGE`, `CHECKBOX`, `PROGRESSBAR`. It aligns only the
content of text fields such as `EDIT` and `BUTTONEDIT`. However, the
`JUSTIFY` attribute can be used with all form item types: In addition to the text
field data alignment, `JUSTIFY` defines the alignment of table column headers (this
means table column header follows the alignment of field data). However, column header alignment in
tables may not be enabled by default; check the [`headerAlignment`](1648-table-style-attributes.md) presentation style attribute for the Table class.

If the input field has the focus in a dialog allowing user input, the data alignment rules are
front-end specific, and follow either `JUSTIFY` or the data type of the field
variable. When the current dialog is a [`CONSTRUCT`](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form."), criteria input is always left-aligned.

You can also specify the text alignment of static form labels with the `JUSTIFY`
attribute.

`IMAGE` form items rendering can be controled with a style attribute
`"alignment"`, for more details, see [alignment](1641-image-style-attributes.md).

## Example

```
LABEL t01: TEXT="Hello!", JUSTIFY=RIGHT;
EDIT f01 = order.value, JUSTIFY=CENTER;
```

## Related links

**Related concepts**  

[FORMAT attribute](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.")

[Presentation styles](1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")

**Related reference**  

[Table style attributes](1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container.")
