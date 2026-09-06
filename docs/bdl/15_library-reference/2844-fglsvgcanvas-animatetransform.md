---
title: "fglsvgcanvas.animateTransform()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_animatetransform.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.animateTransform()"
type: "concept"
---

# fglsvgcanvas.animateTransform()

> Produces an SVG "animateTransform" element.

## Syntax

```
FUNCTION animateTransform(
   attributeName STRING,
   attributeType STRING,
   type STRING,
   from STRING,
   to STRING,
   by STRING,
   begin STRING,
   dur STRING,
   repeatCount STRING )
  RETURNS om.DomNode
```

1. The attributeName, attributeType, type,
   from, to, by, begin,
   dur, and repeatCount parameters are used to set the
   corresponding SVG attributes for the `animateTransform` element. See SVG
   specification for details.

## Usage

This function creates an `"animateTransform"` SVG DOM element from the parameters
passed.

For more details about SVG animation, refer to the SVG specification.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.animateTransform("transform", "XML",
              "rotate", NULL, NULL,
              "360", NULL, "12h", "indefinite")
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
