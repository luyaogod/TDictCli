---
title: "Understanding canvases"
source: "fgl-topics/c_fgl_Canvas_002.html"
breadcrumb: "User interface > User interface programming > Canvases > Understanding canvases"
type: "concept"
---

# Understanding canvases

> This is an introduction to CANVAS drawing.

A canvas element defines a drawing area in a form, to show basic colored shapes.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

Canvas can draw lines, rectangles, ovals, circles, texts, arcs, and polygons.
Keys can be bound to graphical elements for selection with a right or left
mouse click.

In programs, you select a given canvas area by name and you create the shapes
in the abstract user interface tree by using the built-in DOM API, or helper
functions.

The painted canvas is automatically displayed on the front-end when an
interactive instruction is executed, such as `MENU` or
`INPUT`.

Each canvas element is identified by a unique number (id). You can use this
identifier to bind mouse clicks to canvas elements.

> **Note:**
>
> Consider using  [Web Components](2378-web-components.md "This section describes how to use web components in your application.")
> for specific drawing needs (charts, graphics). For example, Genero BDL provides the build-in Web
> Component [fglsvgcanvas](2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.").

## Related links

**Related concepts**  

[CANVAS item definition](1732-canvas-item-definition.md "Defines attributes for a CANVAS drawing area.")

[User interface basics](1508-user-interface-basics.md "This section introduces to the foundation of the Genero user interface.")
