---
title: "fglsvgcanvas.filter()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_filter.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.filter()"
type: "concept"
---

# fglsvgcanvas.filter()

> Produces the SVG "filter" element.

## Syntax

```
FUNCTION filter(
   id STRING,
   x STRING,
   y STRING,
   width STRING,
   height STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. x The X for the filter dimension.
3. y The Y for the filter dimension.
4. width The width for the filter dimension.
5. height The height for the filter dimension.

## Usage

This function creates a `"filter"` SVG DOM element from the parameters.

Filters must be defined in a `"defs"` element, that is created with the [`defs()`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") function.

Create filter effect sub-elements such as `"feOffset"` or
`"feGaussianBlur"` with the [`createElement()`](2852-fglsvgcanvas-createelement.md "Produces an SVG DOM element with the tag name specified as parameter.") function.

A filter can then be applied to an SVG element, by setting the `"filter"`
attribute with an `"url(#filter-name)"` reference.

## Example

```
DEFINE root_svg, defs, f, fe, fse, n om.DomNode

LET defs = fglsvgcanvas.defs( NULL )
CALL root_svg.appendChild( defs )

CALL defs.appendChild( f:=fglsvgcanvas.filter("blur1", -10, -10, 40, 150 ) )

CALL f.appendChild( fe:=fglsvgcanvas.createElement("feOffset",NULL) )
CALL fe.setAttribute("in","SourceAlpha")
CALL fe.setAttribute("dx","3")
CALL fe.setAttribute("dy","3")
CALL fe.setAttribute("result","offset2")

CALL f.appendChild( fe:=fglsvgcanvas.createElement("feGaussianBlur",NULL) )
CALL fe.setAttribute("in","offset2")
CALL fe.setAttribute("stdDeviation","3")
CALL fe.setAttribute("result","blur2")

CALL f.appendChild( fe:=fglsvgcanvas.createElement("feMerge",NULL) )
CALL fe.appendChild( fse:=fglsvgcanvas.createElement("feMergeNode",NULL) )
CALL fse.setAttribute("in","blur2")
CALL fe.appendChild( fse:=fglsvgcanvas.createElement("feMergeNode",NULL) )
CALL fse.setAttribute("in","SourceGraphic")

...
CALL root_svg.appendChild( n:=fglsvgcanvas.rect(40,50,100,200,5,5) )
CALL n.setAttribute("style", "stroke:gray; fill:blue; filter:url(#blur1);" )
...
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
