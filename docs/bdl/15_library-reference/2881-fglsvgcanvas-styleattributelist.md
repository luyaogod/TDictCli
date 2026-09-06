---
title: "fglsvgcanvas.styleAttributeList()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_styleattributelist.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.styleAttributeList()"
type: "concept"
---

# fglsvgcanvas.styleAttributeList()

> Builds a string with a list of attributes to be used in a style attribute.

## Syntax

```
FUNCTION styleAttributeList(
   attrs om.SaxAttributes )
 RETURNS STRING
```

1. attrs is the `om.SaxAttributes` object defining the
   attributes.

## Usage

This function builds string containing a list of `name:value;` pairs, from the
attribute set defined by an `om.SaxAttributes` object.

The resulting string can then be used to define a CSS style, or it can be defined in an SVG
element using `style="attribute-list"`.

Steps to define and use an attribute set with `styleAttributeList()`:

1. Create and fill an `om.SaxAttributes` object,
2. Create the SVG DOM element with one of the fglsvgcanvas functions,
3. Set the "style" attribute with value produced from `styleAttributeList()`.

## Example

```
CONSTANT COLORS_FIRE = 1
DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes,
       n om.DomNode
...
LET attr[COLORS_FIRE] = om.SaxAttributes.create()
CALL attr[COLORS_FIRE].setAttribute(SVGATT_STROKE,"red")
...
LET n = fglsvgcanvas.circle(100,100,50)
CALL n.setAttribute(SVGATT_STYLE, fglsvgcanvas.styleAttributeList(attr[COLORS_FIRE]) )
...
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
