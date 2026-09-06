---
title: "Font sizes"
source: "fgl-topics/c_fgl_presentation_styles_font_size.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Predefined attribute values > Fonts > Font sizes"
type: "concept"
---

# Font sizes

> Use the fontSize style attribute to influence the size of a font.

## Syntax

```
{ generic-size | pointspt | sizeem }
```

1. generic-size is one of the generic font size names (such as
   '`small`' or '`xx-large`') listed in Table 1.
2. points defines an absolute size in points. Specify a number followed
   immediately by `pt`, for example, `3pt`.
3. size defines relative size. Specify a number followed immediately by
   `em`, for example, `3em`.

## Usage

Specify either a generic font size, an absolute size in points with the "pt" unit, or a relative
size with the "em" unit.

Absolute sizes (using the "pt" suffix) define a font size in physical points. Physical points
are much like pixels, in that they are fixed-size units and cannot scale in size. For
example, on HTML pages using CSS styles, one point is equal to 1/72 of an inch.

Relative sizes (using the "em" suffix) define a font size in a scalable size unit that adapts
to the front-end platform, where one "em" unit results in the same size as the size of the
default font on the platform. For example, if the size of the platform default font is 16
points, `1em = 16pt`, `2em = 32pt`, etc.

Generic font sizes are interpreted by the front-end depending on the graphical capability of the
platform.

Use generic font sizes such as `medium`, `large`,
`small`, or sizes relative to the user-chosen font (using `em` units),
rather than absolute point values. In an HTML browser you can choose two fonts (proportional/fixed),
and a well-designed document does not use more than 2 fonts. This is also valid for
applications.

| Generic font size name | Definition |
| --- | --- |
| `xx-small` | Tiny font size |
| `x-small` | Extra-small font size |
| `small` | Small font size |
| `medium` | Medium font size |
| `large` | Large font size |
| `x-large` | Extra-large font size |
| `xx-large` | Huge font size |

You can also specify an absolute font size, by giving a numeric value
followed by the units such as `pt` or
`em`:

## Example

```
<StyleAttribute name="fontSize" value="medium" />
<StyleAttribute name="fontSize" value="xx-large" />
<StyleAttribute name="fontSize" value="12pt" />
<StyleAttribute name="fontSize" value="1em" />
```

## Related links

**Related concepts**  

[Font families](1624-font-families.md "Use the fontFamily style attribute to define a font family.")
