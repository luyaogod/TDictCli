---
title: "Font weights"
source: "fgl-topics/c_fgl_presentation_styles_font_weight.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Predefined attribute values > Fonts > Font weights"
type: "concept"
---

# Font weights

> Use the fontWeight style attribute to define the aspect of a font.

## Syntax

```
{ normal
| bold
| bolder
| lighter
| numeric-weight
}
```

1. numeric-weight is an whole number in the range 1 to 1000. See CSS
   `font-weight` for more details.

## Usage

The value of the `fontWeight` style attribute is directly used a CSS
`font-weight` attribute.

The applicable font weights depend on the [`fontFamily`](1624-font-families.md "Use the fontFamily style attribute to define a font family.") style attribute that is used by the element.

For more details about the meaning and possible font weight values, see <https://developer.mozilla.org/en-US/docs/Web/CSS/font-weight>.

## Example

```
<StyleAttribute name="fontWeight" value="bold" />
```
