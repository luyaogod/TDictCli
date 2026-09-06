---
title: "Example 1: Simple canvas"
source: "fgl-topics/c_fgl_Canvas_example_1.html"
breadcrumb: "User interface > User interface programming > Canvases > Examples > Example 1: Simple canvas"
type: "concept"
description: "Important: This feature is deprecated, its use is discouraged although not prohibited. This topic describes the steps to draw elements in a canvas. First define a drawing area in the form file with ..."
---

# Example 1: Simple canvas

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

This topic describes the steps to draw elements in a canvas.

First define a drawing area in the form file with the `CANVAS`
form item type. In this example, the name of the canvas field is
'`canvas01`'. This field name identifies the drawing area:

```
DATABASE FORMONLY
LAYOUT
GRID
{
  Canvas example:
  [ca01                         ]
  [                             ]
  [                             ]
  [                             ]
  [                             ]
  [                             ]
}
END
END
ATTRIBUTES
CANVAS ca01: canvas01;
END
```

In programs, you draw canvas shapes by creating canvas nodes in
the abstract user interface tree with the DOM API utilities.

Define a variable to hold the DOM node of the canvas and a second
to handle children created for shapes:

```
DEFINE c, s om.DomNode
```

Define a window object variable; open a window with the form containing
the canvas area; get the current window object, and then get the canvas
DOM node:

```
DEFINE w ui.Window 
OPEN WINDOW w1 WITH FORM "form1"
LET w = ui.Window.getCurrent()
LET c = w.findNode("Canvas","canvas01")
```

Create a child node with a specific type defining the shape:

```
LET s = c.createChild("CanvasRectangle")
```

Set attributes to complete the shape definition:

```
CALL s.setAttribute( "fillColor", "red" )
CALL s.setAttribute( "startX", 10 )
CALL s.setAttribute( "startY", 20 )
CALL s.setAttribute( "endX", 100 )
CALL s.setAttribute( "endY", 150 )
```

It is possible to bind keys / actions to Canvas items in order to let the end user select
elements with a mouse click. You can assign a function key for left-button mouse clicks with the
`acceleratorKey1` attribute, while `acceleratorKey3` is used to detect
right-button mouse clicks. The function keys you can bind are F1 to F255. If the user clicks on a
Canvas item bound to key actions, the corresponding action handler will be executed in the current
dialog. Several canvas items can be bound to the same action keys; in order to identify which items
have been selected by a mouse click, you can use the `drawGetClickedItemId()`
function of fgldraw.4gl. This method will return the AUI tree node id of the
Canvas item that was selected (`s.getId()`).

```
... Create the Canvas item with s node variable ...
CALL s.setAttribute( "acceleratorKey1", "F50" )
MENU "test"
   COMMAND KEY (F50)
      IF drawGetClickedItemId() = s.getId() THEN
        ...
      END IF
...
END MENU
```

To clear a given shape in the canvas, remove the element in the
canvas node:

```
CALL c.removeChild(s)
```

To clear the drawing area completely, remove all children of the
canvas node:

```
LET s=c.getFirstChild()
WHILE s IS NOT NULL
  CALL c.removeChild(s)
  LET s=c.getFirstChild()
END WHILE
```

## Related links

**Related concepts**  

[The DomNode class](../15_library-reference/3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
