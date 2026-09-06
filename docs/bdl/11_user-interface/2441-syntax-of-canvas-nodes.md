---
title: "Syntax of canvas nodes"
source: "fgl-topics/c_fgl_Canvas_003.html"
breadcrumb: "User interface > User interface programming > Canvases > Syntax of canvas nodes"
type: "concept"
---

# Syntax of canvas nodes

> The AUI tree contains the XML representation of the canvas, with a variety of elements defining shapes.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

Canvas areas are defined in forms with the following XML syntax:

```
<Canvas colName="name" >
{ <CanvasArc canvasitem-attribute="value" [...] />
| <CanvasCircle canvasitem-attribute="value"[...] />
| <CanvasLine canvasitem-attribute="value"[...] />
| <CanvasOval canvasitem-attribute="value"[...] />
| <CanvasPolygon canvasitem-attribute="value"[...] />
| <CanvasRectangle canvasitem-attribute="value"[...] />
| <CanvasText canvasitem-attribute="value"[...] />
} [...]
</Canvas>
[...]
```

| Name | Description |
| --- | --- |
| `CanvasArc` | Arc defined by the bounding square top left point, a diameter, a start angle, an end angle, and a fill color. |
| `CanvasCircle` | Circle defined by the bounding square top left point, a diameter, and a fill color. |
| `CanvasLine` | Line defined by a start point, an end point, width, and a fill color. |
| `CanvasOval` | Oval defined by rectangle (with start point and endpoint), and a fill color. |
| `CanvasPolygon` | Polygon defined by a list of points, and a fill color. |
| `CanvasRectangle` | Rectangle defined by a start point, an end point, and a fill color. |
| `CanvasText` | Text defined by a start point, an anchor hint, the text, and a fill color. |

| Name | Values | Description |
| --- | --- | --- |
| `startX` | `INTEGER (0->1000)` | X position of starting point. |
| `startY` | `INTEGER (0->1000)` | Y position of starting point. |
| `endX` | `INTEGER (0->1000)` | X position of ending point. |
| `endY` | `INTEGER (0->1000)` | Y position of ending point. |
| `xyList` | `STRING` | Space-separated list of XY coordinates. For example: "23 45 56 78" defines (x=23,y=45) (x=56,y=78). |
| `width` | `INTEGER` | Width of the shape. |
| `height` | `INTEGER` | Height of the shape. |
| `diameter` | `INTEGER` | Diameter for circles and arcs. |
| `startDegrees` | `INTEGER` | Beginning of the angular range occupied by an arc. |
| `extentDegrees` | `INTEGER` | Size of the angular range occupied by an arc. |
| `text` | `STRING` | The text to draw. |
| `anchor` | `"n","e","w","s"` | Anchor hint to give the draw direction for texts. |
| `fillColor` | `STRING` | Name of the color to be used for the element. |
| `acceleratorKey1` | `STRING` | Name of the key associated to a left button click. |
| `acceleratorKey3` | `STRING` | Name of the key associated to a right button click. |
