---
title: "PLACEHOLDER attribute"
source: "fgl-topics/c_fgl_FSFAttributes_PLACEHOLDER.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > PLACEHOLDER attribute"
type: "concept"
---

# PLACEHOLDER attribute

> The PLACEHOLDER attribute defines a hint for the user when the field contains no value.

## Syntax

```
PLACEHOLDER = [%]"string"
```

1. string defines the hint to be displayed when the field is
   empty, with the % prefix it is a localized string.

## Usage

The `PLACEHOLDER` attribute can be used to show hint text in an input field, when
a field contains no value.

The placeholder text is displayed when the field value is `NULL`, during an
`INPUT`, `CONSTRUCT` or `INPUT ARRAY` dialog. In the
case of `INPUT ARRAY`, placeholders are only displayed for the current row where the
end user enters data.

This attribute is typically used for mobile and web applications, to display grayed text inside
empty input fields.

The [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element.") attribute can,
in specific situations, define a default placeholder message.

## Example

```
EDIT f001 = customer.cust_name, PLACEHOLDER = "<Enter customer name>";
```
