---
title: "Colors"
source: "fgl-topics/c_fgl_presentation_styles_015.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Predefined attribute values > Colors"
type: "concept"
---

# Colors

> When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.

This section describes how to specify a value for style attributes
defining colors, such as `textColor`.

## Syntax

```
{ generic-color | #rrggbb }
```

1. generic-color is any of the predefined colors supported by the language.
2. #rrggbb is a numerical color defined by a red/green/blue specification.

## Usage

Colors can be used at different places, typically in style attributes like
`backgroundColor`.

The language defines a set of generic color names, interpreted by the front-end
depending on the graphical capability of the workstation.

## RGB notation

When a predefined color is not what you require, you can specify the exact color with
the RGB notation, starting with a # hash
character: `<Style name="Edit.mandatory">
<StyleAttribute name="textColor" value="#50AEFF" />
</Style>`

Each value of the RGB color specification must be provided in hexadecimal, in the range [00-FF].

## GBC System Colors

The next table lists the system colors names that can be used in a color style attribute. Note
that the corresponding GBC theme variable gets the same name as the system color, with the
`"$gbc-genero-"` prefix:

| System color name | GBC Theme Variable | Description |
| --- | --- | --- |
| `appWorkSpace` | `$gbc-genero-appWorkSpace` | Same as `window`. |
| `background` | `$gbc-genero-background` | Background color. |
| `buttonFace` | `$gbc-genero-buttonFace` | Button background color. |
| `buttonText` | `$gbc-genero-buttonText` | Text color of buttons. |
| `grayText` | `$gbc-genero-grayText` | Grayed (disabled) text. |
| `highLight` | `$gbc-genero-highLight` | Background color of row(s) selected in a array. |
| `highLightText` | `$gbc-genero-highLightText` | Text color of row(s) selected in a array. |
| `infoBackground` | `$gbc-genero-infoBackground` | N/A (same as `window`) |
| `infoText` | `$gbc-genero-infoText` | N/A (same as `windowText`) |
| `systemAlternateBackground` | `$gbc-genero-systemAlternateBackground` | Background color of the alternate row in array. |
| `window` | `$gbc-genero-window` | Window background color. |
| `windowText` | `$gbc-genero-windowText` | Text in windows. |

> **Note:**
>
> GBC theme variables with the `"$gbc-genero-"` prefix are created by the system
> color name, and inherited from the theme definitions. Editing variables with this prefix does not
> effect the GBC theme. For information on editing the GBC theme colors, go to Theme
> reference in the Genero Browser Client User Guide.

## Generic color names

Use generic color names, to keep your style definitions portable across several front-end types.
Depending on the configuration of the front end, generic color names may be rendered with different
RGB color value. The use of the `REVERSE` attribute can also affect the resulting RGB
values used.

The generic color names are: `black, blue, cyan, darkBlue, darkCyan, darkGray, darkGreen,
darkMagenta, darkOlive, darkOrange, darkRed, darkTeal, darkYellow, gray, green, lightBlue,
lightCyan, lightGray, lightGreen, lightMagenta, lightOlive, lightOrange, lightRed, lightTeal,
lightYellow, magenta, olive, orange, red, teal, white, yellow`.

## Example

```
<StyleAttribute name="backgroundColor" value="lightBlue" />
<StyleAttribute name="textColor" value="#00FF45" />
```

## Related links

**Related concepts**  

[Font families](1624-font-families.md "Use the fontFamily style attribute to define a font family.")
