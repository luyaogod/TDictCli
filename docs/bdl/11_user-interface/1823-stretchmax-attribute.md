---
title: "STRETCHMAX attribute"
source: "fgl-topics/c_fgl_FSFAttributes_STRETCHMAX.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > STRETCHMAX attribute"
type: "concept"
---

# STRETCHMAX attribute

> The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.

## Syntax

```
STRETCHMAX[@screen-size] = max-width
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when the maximum stretch
   width must apply, depending on the size of the screen. Several
   `STRETCHMAX@screen-size` attributes can be used for the same
   element.
2. max-width is an integer that defines the maximum width of the element. This
   maximum width is specified as a number of cells, as in the [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") of the column.

## Usage

The `STRETCHMAX` attribute specifies the maximum width for a
`TABLE/TREE` column, when the [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute is used by the element, or when the parent container
defines the [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.") attribute.

By default, when the column is defined as horizontally stretchable, there is no maximum stretch
width.

The `STRETCHMAX` attribute can be combined with the
`@screen-size` selector, to get different stretch behaviors
depending on the screen size.

## Example

```
EDIT f01 = customer.cust_name, STRETCH=X,
   STRETCHMAX@SMALL=20, STRETCHMAX@MEDIUM=30;
```

## Related links

**Related concepts**  

[STRETCHMIN attribute](1824-stretchmin-attribute.md "The STRETCHMIN attribute defines the minimum stretching width.")

[Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.")

[Responsive Layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.")
