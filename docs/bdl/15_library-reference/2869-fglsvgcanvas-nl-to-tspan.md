---
title: "fglsvgcanvas.nl_to_tspan()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_nl_to_tspan.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.nl_to_tspan()"
type: "concept"
---

# fglsvgcanvas.nl_to_tspan()

> Converts a string to an SVG "text" element with "tspan" sub-elements.

## Syntax

```
FUNCTION nl_to_tspan(
   text om.DomNode,
   x STRING,
   y STRING,
   dx STRING,
   dy STRING,
   content STRING )
```

1. text is the DOM node to update with tspan elements.
2. x is used to set the x attribute of each tspan element.
3. y is used to set the y attribute of each tspan element.
4. dx is used to set the dx attribute of each tspan element.
5. dy is used to set the dy attribute of each tspan element.
6. content is the source text with new-line characters.

## Usage

This function creates `"tspan"` elements for each line of text.

The text DOM element passed as parameter gets `"tspan"`
sub-elements for each new-line character found in the original string.

## Example

```
DEFINE root_svg, t, n om.DomNode
...
LET t = fglsvgcanvas.text(NULL,200,NULL)
CALL root_svg.appendChild( t )
CALL fglsvgcanvas.nl_to_tspan(t,120,NULL,NULL,30,"Text using tspan\nLine 2\nLine 3")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
