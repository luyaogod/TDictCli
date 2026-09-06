---
title: "MINWIDTH attribute"
source: "fgl-topics/c_fgl_FSFAttributes_MINWIDTH.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > MINWIDTH attribute"
type: "concept"
---

# MINWIDTH attribute

> The MINWIDTH attribute defines the minimum width of a form.

## Syntax

```
MINWIDTH = grid-cells
```

1. grid-cells is an integer that defines the minimum width of the element,
   as a number of grid cells.

## Usage

The `MINWIDTH` attribute is used to define a minimum width for the form/window.
It must be specified in the attributes of the `LAYOUT` section.

On `LAYOUT` element, the `MINWIDTH` and
`MINHEIGHT` attributes apply only to windows created with the
`windowType` style attribute set to `"modal"`.

The unit defaults to a number of grid cells. This is the equivalent of the
`CHARACTERS` in the `WIDTH` attribute specification. For more details
about grid cells, see [Widget position and size in grid](1552-widget-position-and-size-in-grid.md "Form items render as widgets in the window, at a given position and with a given size.")

## Example

```
LAYOUT ( MINWIDTH=60, MINHEIGHT=50 )
GRID
  ...
```

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")

[MINHEIGHT attribute](1799-minheight-attribute.md "The MINHEIGHT attribute defines the minimum height of a form.")
