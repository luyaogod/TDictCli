---
title: "STRETCHCOLUMNS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_STRETCHCOLUMNS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > STRETCHCOLUMNS attribute"
type: "concept"
---

# STRETCHCOLUMNS attribute

> The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.

## Syntax

```
STRETCHCOLUMNS[@screen-size]
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when the stretch
   attribute must apply, depending on the size of the screen. Several
   `STRETCHCOLUMNS@screen-size` attributes can be used for the same
   element.

## Usage

The `STRETCHCOLUMNS` attribute makes all columns of a [`TABLE`](1703-table-item-type.md "Defines a list view widget.") or [`TREE`](1706-tree-item-type.md "Defines a tree view widget.") container stretch or shrink, when
the width of the container changes.

The `STRETCHCOLUMNS` attribute at the `TABLE/TREE` level can be
combined with individual [`STRETCH=NONE`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attributes defined at the column level.

To create a responsive layout form, the `STRETCHCOLUMNS` attribute can be combined
with the `@screen-size` selector, to get different stretch
behaviors depending on the screen size.

`STRETCHCOLUMNS` can be combined with field-specific [`STRETCHMIN`](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.") and [`STRETCHMAX`](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.") attributes, to control
the width of a table column.

## Example

```
TABLE mytable ( STRETCHCOLUMNS@LARGE )
```

## Related links

**Related concepts**  

[Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.")

[Responsive Layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.")
