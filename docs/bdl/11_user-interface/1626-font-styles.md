---
title: "Font styles"
source: "fgl-topics/c_fgl_presentation_styles_font_style.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Predefined attribute values > Fonts > Font styles"
type: "concept"
---

# Font styles

> Use the fontStyle style attribute to define the style of a font.

## Syntax

```
{ italic | roman | oblique }
```

## Usage

The style of a font can be specified with a generic name, that is interpreted by the front-end
depending on the graphical capabilities of the platform. For example, on "Android™ devices, `italic` and `oblique`
result in the same font aspect.

| Generic font style name | Definition |
| --- | --- |
| `italic` | Specifies an italic font style, using a typeface that slants slightly to the right. Uses a different glyph as the roman style. |
| `oblique` | Specifies an oblique font style. This style is similar to italic, except that it uses the same glyphs as the roman type, but distorted. |
| `roman` | Specifies a roman font style. This is the typical default font style in Latin-script typography. |

## Example

```
<StyleAttribute name="fontStyle" value="italic" />
```
