---
title: "SCROLL attribute"
source: "fgl-topics/c_fgl_FSFAttributes_SCROLL.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > SCROLL attribute"
type: "concept"
---

# SCROLL attribute

> The SCROLL attribute can be used to enable horizontal scrolling in a character field.

## Syntax

```
SCROLL
```

## Usage

By default, the maximum data input length is defined by the width of the item-tag of the field.
For example, a form field defined with an item tag of 3 cells like `[f1 ]`, users can
only enter a maximum of 3 latin characters (`abc`, `ééé`), or 1
Chinese ideogram + 1 Latin char, even if the program variable used by the input dialog is a
`VARCHAR(20)`.

To let the user input more characters than the width of the item-tag of the field, use the
`SCROLL` attribute.

The `SCROLL` attribute applies only to fields with character data input.

Use the `SCROLL` attribute only when the layout of the form does not allow for
defining an item tag that is large enough to hold all possible character string data that fits in
the corresponding program variable. Understand that the end user can miss a part of the displayed
data when the field is too small. Therefore, times when you would need to use the
`SCROLL` attribute are rare.

Special consideration needs to be taken regarding full-width characters like Chinese logograms: A
form field defined with 4 cells will by default only allow 2 full-width characters, if the
`SCROLL` attribute is not used.

Using character or byte length semantics with (FGL\_LENGTH\_SEMANTICS) also impacts the maximum
field input length.

For more details, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Related links

**Related concepts**  

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
