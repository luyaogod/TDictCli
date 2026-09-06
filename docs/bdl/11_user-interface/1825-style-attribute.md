---
title: "STYLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_STYLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > STYLE attribute"
type: "concept"
---

# STYLE attribute

> The STYLE attribute specifies a presentation style for a form element.

## Syntax

```
STYLE = "style-name [...]"
```

1. style-name is a set of user-defined style name. Separator is space.

## Usage

This attribute specifies a presentation style to be applied to a form element.

The presentation style can define decoration attributes such as a background color, a
font type, and so on.

Several style names can be combined, by using a space as separator.

The names used to define the `STYLE` attribute must be a
style-name only, these must not contain the element-type
that is typically used to define the style in a .4st file (as
`CheckBox.important` for example)

For more details about using the STYLE attribute, see [Using presentation styles](1610-using-presentation-styles.md "Use presentation styles to centralize the decoration of your user interface.").

## Example

```
EDIT c01 = item.pkey, STYLE = "important keyval";
EDIT c02 = item.comment, STYLE = "comment";
```

## Related links

**Related concepts**  

[COLOR attribute](1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.")

[COLOR WHERE Attribute](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically.")
