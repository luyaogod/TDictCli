---
title: "Canvas drawing area"
source: "fgl-topics/c_fgl_Canvas_005.html"
breadcrumb: "User interface > User interface programming > Canvases > Canvas drawing area"
type: "concept"
---

# Canvas drawing area

> The canvas area defines a two-dimensional coordinate system for drawing elements.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

The canvas area represents an abstract drawing page where you define size and location of shapes
with coordinates from (0,0) to (1000,1000).

The origin point (0,0), is on the bottom-left of the drawing area.

![Canvas area diagram](../_images/CVSFig01.jpg)

*Canvas area diagram*

The drawing area is defined in the form file with a `CANVAS` form item. At
runtime, you draw the content of canvas areas in the Abstract User Interface tree. In a form
defining canvas areas, the Abstract User Interface tree contains empty
`<Canvas>` nodes that you can fill with canvas items.

A canvas node is identified in the program by the `name` attribute. You can
get the canvas node by name with the `Window.getElement(name)`
method.

You cannot drop canvas area nodes, as they are read-only in a form definition.

## Related links

**Related concepts**  

[CANVAS item definition](1732-canvas-item-definition.md "Defines attributes for a CANVAS drawing area.")
