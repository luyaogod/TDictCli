---
title: "fglsvgcanvas.initialize()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_initialize.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.initialize()"
type: "concept"
---

# fglsvgcanvas.initialize()

> Prepares the fglsvgcanvas library for use.

## Syntax

```
FUNCTION initialize( )
```

## Usage

This function initializes the fglsvgcanvas module usage. When the fglsvgcanvas library is no
longer needed, call the finalization function [`fglsvgcanvas.finalize()`](2858-fglsvgcanvas-finalize.md "Releases the fglsvgcanvas library.").

Initialization and finalization functions can be called several times by different modules using
the fglsvgcanvas library.

## Example

```
IMPORT FGL fglsvgcanvas
FUNCTION show_svg_content()
  ...
  CALL fglsvgcanvas.initialize()
  ...
  CALL fglsvgcanvas.finalize()
END FUNCTION
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
