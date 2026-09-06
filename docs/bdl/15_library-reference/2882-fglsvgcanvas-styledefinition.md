---
title: "fglsvgcanvas.styleDefinition()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_styledefinition.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.styleDefinition()"
type: "concept"
---

# fglsvgcanvas.styleDefinition()

> Produces a CSS style definition with a selection and list of attributes.

## Syntax

```
FUNCTION styleDefinition(
   selector STRING,
   attrs om.SaxAttributes )
 RETURNS STRING
```

1. selector is the style selector.
2. attrs is the `om.SaxAttributes` object defining the
   attributes.

## Usage

This function creates a CSS style definition line from the attribute set defined by an
`om.SaxAttributes` object.

The attribute set must be an [`om.SaxAttributes`](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.") object.

The resulting string can be used in a style list created by the [`styleList()`](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.")
function.

Steps to define and use an attribute set with `styleDefinition()`:

1. Create and fill an `om.SaxAttributes` object,
2. Use the `styleDefinition()` function using the SAX attributes.

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
