---
title: "STRETCHMIN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_STRETCHMIN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > STRETCHMIN attribute"
type: "concept"
---

# STRETCHMIN attribute

> The STRETCHMIN attribute defines the minimum stretching width.

## Syntax

```
STRETCHMIN[@screen-size] = min-width
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when the minimum stretch
   width must apply, depending on the size of the screen. Several
   `STRETCHMIN@screen-size` attributes can be used for the same
   element.
2. min-width is an integer that defines the minimum width of the element.
   Default is the number of cells used by the [item
   tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") of the column.

## Usage

The `STRETCHMIN` attribute specifies the minimum width of form elements, when the
[`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute is used by
the element (or by the root [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") element), or when the `TABLE/TREE` parent container
defines the [`STRETCHCOLUMNS`](1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.") attribute.

By default, the minimum width is the number of cells used by the [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") in the `LAYOUT`
section.

Note that with form elements inside a `GRID` or `SCROLLGRID`
container, the [items alignment and grid
dependency rules](1553-form-item-dependencies-in-grids.md "Form items interact with each other in terms of width, depending on the front-end widget size.") can hide the effect of the `STRETCHMIN` attribute.

The [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") attribute
has a lower priority than `STRETCHMIN`.

The `STRETCHMIN` attribute can be combined with the
`@screen-size` selector, to get different stretch behaviors
depending on the screen size.

## Example

```
EDIT f01 = customer.cust_name, STRETCH=X,
   STRETCHMIN@MEDIUM=10, STRETCHMIN@LARGE=15;
```

## Related links

**Related concepts**  

[STRETCHMAX attribute](1823-stretchmax-attribute.md "The STRETCHMAX attribute defines the maximum stretching width for a TABLE/TREE column.")

[Controlling table rendering](2319-controlling-table-rendering.md "Table rendering can be controlled by the use of presentation styles and table attributes.")

[Responsive Layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.")
