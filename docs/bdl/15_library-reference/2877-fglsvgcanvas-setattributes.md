---
title: "fglsvgcanvas.setAttributes()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_setattributes.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.setAttributes()"
type: "concept"
---

# fglsvgcanvas.setAttributes()

> Sets the SVG attributes from an attribute set.

## Syntax

```
FUNCTION setAttributes(
    node om.DomNode,
    attrs om.SaxAttributes )
```

1. node is the DOM node to be updated.
2. attrs is the `om.SaxAttributes` object defining the
   attributes.

## Usage

This function updates the DOM node passed as first parameter with the [`om.SaxAttributes`](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.") object.

If a SAX attribute value is `NULL`, the attribute is removed from the node.

The `setAttributes()` function should only be used to set explicitly node-level
attributes. Consider using CSS style definitions with the [`styleList()`](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.")
function, or inline styling with the [`styleAttributeList()`](2881-fglsvgcanvas-styleattributelist.md "Builds a string with a list of attributes to be used in a style attribute.") function.

Steps to define and use an attribute set with `setAttributes()`:

1. Create and fill an `om.SaxAttributes` object,
2. Create the SVG DOM element with one of the fglsvgcanvas functions,
3. Call the `setAttributes()` function with the node and the SAX attributes.

## Example

```
CONSTANT COLORS_OCEAN = 1
DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes,
       n om.DomNode
...
LET attr[COLORS_OCEAN] = om.SaxAttributes.create()
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_FILL,           "cyan" )
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_FILL_OPACITY,   "0.3" )
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_STROKE,         "blue" )
...
LET n = fglsvgcanvas.polygon("10,10 10,20 20,20")
CALL fglsvgcanvas.setAttributes( n, attr[COLORS_OCEAN] )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
