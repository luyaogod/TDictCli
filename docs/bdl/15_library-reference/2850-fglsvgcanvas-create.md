---
title: "fglsvgcanvas.create()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_create.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.create()"
type: "concept"
---

# fglsvgcanvas.create()

> Creates a new SVG canvas handler.

## Syntax

```
FUNCTION create( name STRING )
 RETURNS SMALLINT
```

1. name is the name of the form field defined as a [`WEBCOMPONENT`](../11_user-interface/1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.").

## Usage

This function creates a new SVG canvas handle by using the form field name passed as
parameter.

The window/form containing the fglsvgcanvas web component field must be created before calling
this function.

The name is used to bind the `WEBCOMPONENT` form field with the
SVG canvas, to display SVG content.

The function returns the id of the newly-created SVG canvas. This id must be used in subsequent
calls to fglsvgcanvas functions.

When the SVG canvas is no longer needed, free the allocated resources with the [`destroy()`](2854-fglsvgcanvas-destroy.md "Releases resources allocated for the SVG canvas.")
function.

## Example

Form file:

```
...
ATTRIBUTES
WEBCOMPONENT cv = FORMONLY.canvas,
   COMPONENTTYPE = "fglsvgcanvas",
   PROPERTIES = (selection="item_selection"),
   SIZEPOLICY = FIXED,
   STRETCH = BOTH,
   SCROLLBARS = NONE
;
...
```

Program code:

```
DEFINE cid SMALLINT
LET cid = fglsvgcanvas.create("formonly.canvas")
...
CALL fglsvgcanvas.destroy( cid )
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
