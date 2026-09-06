---
title: "fglsvgcanvas.setRootSVGAttributes()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_setrootsvgattributes.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.setRootSVGAttributes()"
type: "concept"
---

# fglsvgcanvas.setRootSVGAttributes()

> Produces the root SVG element.

## Syntax

```
FUNCTION setRootSVGAttributes(
   id STRING,
   width STRING,
   height STRING,
   viewBox STRING,
   preserveAspectRatio STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. width and height define the SVG viewport.
3. viewBox defines the SVG viewbox (the internal coordinate system).
4. preserveAspectRatio is the aspect ratio to preserve.

## Usage

This function sets the attributes of the root "svg" SVG element from the parameters, and returns
the SVG DOM element.

This method does not create a new DOM node object (it is created when calling the [`create()`](2850-fglsvgcanvas-create.md "Creates a new SVG canvas handler.") function):
You can repeat calls to the `setRootSVGAttributes()` function to change/reset the
root DOM node SVG attributes.

Create the SVG canvas with the following sequence of calls:

1. Initialize the fglsvgcanvas library with [`initialize()`](2863-fglsvgcanvas-initialize.md "Prepares the fglsvgcanvas library for use."),
2. Create a new SVG canvas handler with [`create()`](2850-fglsvgcanvas-create.md "Creates a new SVG canvas handler.") (this
   creates the DOM node for the root SVG),
3. Set root SVG attributes and get the corresponding DOM node object with
   `setRootSVGAttributes()`.
4. Use the root DOM node to append child DOM element created from fglsvgcanvas functions.

The `width` and `height` attributes define the SVG viewport, the
visible area of the SVG image. You want to leave the viewport attributes to `NULL`,
to let the SVG image adapt to its container.

The viewbox is specified with the `viewBox` parameter, to define the coordinate
system to draw SVG objects. For example `"0 0 100 100"` defines a viewbox where 0,0
are the coordinates of the top/left corner, and 100,100 are the coordinates of the bottom/right
corner. If you draw a rectangle with 25,25,50,50 it will be centered and use half of the
viewbox.

The preserveAspectRatio parameter is used in conjunction with the
`viewBox` attribute, to control how the SVG content is stretched, whether or not to
force uniform scaling. In most case you want to use the `"xMidYMid meet"` as aspect
ratio.

See SVG reference for more details about viewports and viewboxes.

## Example

```
DEFINE cid SMALLINT
DEFINE root_svg om.DomNode
CALL fglsvgcanvas.initialize()
LET cid = fglsvgcanvas.create("formonly.canvas")
LET root_svg = fglsvgcanvas.setRootSVGAttributes(
                  "agenda",
                  NULL, NULL,
                  "0 0 1000 1000",
                  "xMidYMin slice"
               )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
