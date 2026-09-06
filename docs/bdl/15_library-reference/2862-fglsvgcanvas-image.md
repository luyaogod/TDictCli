---
title: "fglsvgcanvas.image()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_image.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.image()"
type: "concept"
---

# fglsvgcanvas.image()

> Produces an SVG "image" element.

## Syntax

```
FUNCTION image(
   href STRING,
   x STRING,
   y STRING,
   width STRING,
   height STRING,
   preserveAspectRatio STRING )
  RETURNS om.DomNode
```

1. href defines the xlink:href reference to the image.
2. x defines the X coordinate of the image.
3. y defines the Y coordinate of the image.
4. width defines the width of the image.
5. height defines the height of the image.
6. preserveAspectRatio defines a the aspect ratio to preserve.

## Usage

This function creates an `"image"` SVG DOM element from the parameters.

The href parameter defines the image resource or URL.

> **Important:**
>
> Use the [`ui.Interface.filenameToURI()`](3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.")
> method to produce a URL for application images.

## Example

```
DEFINE n om.DomNode
LET n = fglsvgcanvas.image( ui.Interface.filenameToURI("image02.jpg"),
                            100,100,50,50, "xMidYMid meet" )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
