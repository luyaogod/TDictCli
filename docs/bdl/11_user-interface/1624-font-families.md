---
title: "Font families"
source: "fgl-topics/c_fgl_presentation_styles_font_family.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Predefined attribute values > Fonts > Font families"
type: "concept"
---

# Font families

> Use the fontFamily style attribute to define a font family.

This section describes the possible values of the `fontFamily` style
attribute.

## Syntax

```
font-family [,...]
```

1. font-family is a specific or generic font name such as
   `sans-serif`, `monospace`.

## Usage

The `fontFamily` style attribute allows to define a single or a set of font names
separated by a comma. These are used as `font-family` property in a CSS style.

Consider using generic font names:

- `serif`
- `sans-serif`
- `cursive`
- `fantasy`
- `monospace`

For more details about possible fonts, see [CSS
font families](https://www.w3.org/TR/CSS2/fonts.html#propdef-font-family).

## Example

```
<StyleAttribute name="fontFamily" value="sans-serif" />
<StyleAttribute name="fontFamily" value="'Courier New'" />
<StyleAttribute name="fontFamily" value="'Times New Roman',Times,serif" />
```

## Related links

**Related concepts**  

[Syntax of presentation styles file (.4st)](1609-syntax-of-presentation-styles-file-4st.md "A .4st presentation styles file is an XML file defining style attributes to be applied by front-ends.")
