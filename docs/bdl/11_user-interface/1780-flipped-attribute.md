---
title: "FLIPPED attribute"
source: "fgl-topics/c_fgl_FSFAttributes_FLIPPED.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > FLIPPED attribute"
type: "concept"
---

# FLIPPED attribute

> The FLIPPED attribute flips TABLE columns into rows.

## Syntax

```
FLIPPED[@screen-size]
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when the "flip columns"
   attribute must apply, depending on the size of the screen. Several
   `FLIPPED@screen-size` attributes can be used for the same
   element.

## Usage

The `FLIPPED` attribute forces the `TABLE` container to render in a
way that all cells of a given row are packed vertically, to get a thinner table. The table is still
scrollable, with each row displayed stacked as a group of fields.

To create a responsive layout form, the `FLIPPED` attribute can be combined with
the `@screen-size` selector, to get different table renderings
depending on the screen size. On a large screen, tables can be rendered in the classic way. On small
screens, the "flip columns" mode rendering is more appropriate.

When the `FLIPPED` attribute applies, the widgets used to render form items may
not keep the same size (especially the height) as in classical table display, as the goal is to pack
cells vertically to display rows on small screens. Some type of form items such as [`RADIOGROUP`](1699-radiogroup-item-type.md "Defines a mutual exclusive set of options field."), [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.") and [`TEXTEDIT`](1704-textedit-item-type.md "Defines a multi-line edit field.") will typically get a
different height and therefore should be used with care in a table with the `FLIPPED`
attribute.

## Example

```
TABLE mytable ( FLIPPED@SMALL )
```

## Related links

**Related concepts**  

[Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.")

[Responsive Layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.")
