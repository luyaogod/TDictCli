---
title: "The fglsvgcanvas web component"
source: "fgl-topics/c_fgl_fgl_web_components_fglsvgcanvas.html"
breadcrumb: "User interface > User interface programming > Web components > Built-in web components > Built-in web components reference > The fglsvgcanvas web component"
type: "concept"
---

# The fglsvgcanvas web component

> The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.

> **Note:**
>
> This documentation does not explain SVG drawing principles. Before starting with
> `fglsvgcanvas`, consider learning SVG, with tutorials found on the internet. It is
> especially mandatory to properly understand the root SVG viewport / viewBox / preserveAspectRatio
> concepts.

The `fglsvgcanvas` built-in web component is a [gICAPI web component](2391-using-a-gicapi-web-component.md "This section describes how to add a gICAPI-based web component to your application.").
> **Important:**
>
> Since SVG allows you to specify coordinates and sizes with units such as `"10em"`,
> and including the % percentage unit like in `"50%"`, fglsvgcanvas functions use the
> `STRING` type for parameters such as x, y, width, height. In SVG, the decimal
> separator for numeric values must always be a dot. When computing coordinates and sizes with numeric
> types such as `DECIMAL` or when passing decimal values directly to fglsvgcanvas
> functions, pay attention to numeric to string conversion. By default, the DBMONEY/DBFORMAT settings
> apply and can produce a decimal separator different from the dot. The fglsvgcanvas library converts
> automatically decimal coordinates to the ISO format. However, if possible, define a large SVG
> viewBox (like `"0 0 1000 1000"`), in order to use only integer numbers for
> coordinates and sizes, or `%` percentage units.

The `fglsvgcanvas` web component HTML page is basically a simple HTML container.
It is delivered with the utility library [$FGLDIR/src/webcomponents/fglsvgcanvas/fglsvgcanvas.4gl](../15_library-reference/2841-fglsvgcanvas-svg-drawing-module.md), that can be
used to produce SVG content.

The programming pattern is based on the [built-in
`om.*` API](../15_library-reference/3280-the-om-package.md "These topics cover the built-in classes of the om package"). Create DOM nodes with the utility functions, and construct the
root `<svg/>` element by adding child nodes created from the fglsvgcanvas
functions.

![Screenshot of a program using the fglsvgcanvas web component](../_images/fglsvgcanvas01_gbc.jpg)

*fglsvgcanvas sample from Fourjs Github*

See <https://github.com/FourjsGenero/fgl_svg_chart>

## Defining the `fglsvgcanvas` web component in the form file

In the .per form definition file, define the SVG container as a [`WEBCOMPONENT`](1750-webcomponent-item-definition.md "Defines attributes for a generic form field that can receive an external widget.") form
item with the [`COMPONENTTYPE`](1771-componenttype-attribute.md "The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields.") attribute set to the `"fglsvgcanvas"`
value.

> **Important:**
>
> The form field name will be used in front calls to identify the SVG
> canvas.

Since the SVG canvas web component provides built-in scrollbars, the [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item.") attribute can be set
to `NONE`.

Use [`SIZEPOLICY=FIXED`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") and
[`STRETCH=BOTH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), to get an SVG
canvas that resizes with the parent window.

Additional fglsvgcanvas web component configuration options can be defined with the [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.") attribute (details
will be discussed later in this topic).

For example:

```
LAYOUT
GRID
{
[cv                     ]
[                       ]
[                       ]
...
ATTRIBUTES
WEBCOMPONENT cv = FORMONLY.canvas,
   COMPONENTTYPE = "fglsvgcanvas",
   PROPERTIES = ( selection="item_selection" ),
   SIZEPOLICY = FIXED,
   STRETCH = BOTH,
   SCROLLBARS = NONE;
...
```

## The fglsvgcanvas.4gl library

SVG can be used to draw complex content such as an agenda or a graph, with advanced SVG concepts
such as CSS styles, patterns, nested `<svg/>` elements, etc.

To simplify SVG programming, Genero BDL provides the
$FGLDIR/src/webcomponents/fglsvgcanvas/fglsvgcanvas.4gl utility library. This
library implements a set of functions that produce common SVG elements.

The fglsvgcanvas.4gl library supports following SVG features:

- Attribute sets
- CSS styles
- Patterns
- Masks
- Filters
- Gradients
- Shapes (rect, circle, polygon, etc)
- Simple text, text on path, text tspan
- Animation
- Clickable elements
- Clipping paths
- RGB color utilities (shade, tint)

## fglsvgcanvas library initialization and finalization

Library initialization and finalization functions are provided to prepare the library before
usage, and free resources when the library is no longer needed:

```
CALL fglsvgcanvas.initialize()
...
CALL fglsvgcanvas.finalize()
```

For more details see [fglsvgcanvas.initialize()](../15_library-reference/2863-fglsvgcanvas-initialize.md "Prepares the fglsvgcanvas library for use.") and
[fglsvgcanvas.finalize()](../15_library-reference/2858-fglsvgcanvas-finalize.md "Releases the fglsvgcanvas library.").

## Creating an SVG canvas handler

Before creating new SVG element, you need to create the SVG canvas handler, to get an id that
will be used in subsequent fglsvgcanvas calls. Define a `SMALLINT` variable to hold
the SVG canvas handler id that is returned by the [`create()`](../15_library-reference/2850-fglsvgcanvas-create.md "Creates a new SVG canvas handler.") function.
This function takes the WEBCOMPONENT field name as attribute, to bind the form field to the SVG
canvas handle:

```
DEFINE cid SMALLINT
...
LET cid = fglsvgcanvas.create("formonly.canvas")
```

## Selecting an SVG canvas handler

The SVG canvas is identified by the id returned by the `create()` function. After
creating an SVG canvas, it is automatically defined as the current canvas, and any subsequent calls
to an fglsvgcanvas function will apply to that current canvas handler. If you want to manipulate
several SVG canvases, select the current canvas with the [`setCurrent()`](../15_library-reference/2878-fglsvgcanvas-setcurrent.md "Selects the SVG canvas handler for subsequent SVG canvas API calls.")
function:

```
CALL fglsvgcanvas.setCurrent(cid)
```

## The root SVG node

The root SVG DOM node is created when calling the [`create()`](../15_library-reference/2850-fglsvgcanvas-create.md "Creates a new SVG canvas handler.") function.
However, before drawing your SVG, you need to define essential root SVG attributes.

To define the root SVG attributes, use the [`setRootSVGAttributes()`](../15_library-reference/2879-fglsvgcanvas-setrootsvgattributes.md "Produces the root SVG element.") function.

This function returns the root `om.DomNode` of the SVG tree:

```
DEFINE root_svg, n om.DomNode
...
LET root_svg = fglsvgcanvas.setRootSVGAttributes(
                      "myrootsvg",
                      "10em", "5em", -- viewport
                      "0 0 500 200",  -- viewbox
                      "xMidYMid meet" -- preserveAspectRatio
                   )
...
```

Specify the following properties to define your root SVG element:

- the viewport defines the viewing area for the SVG image (use `NULL,NULL` for
  auto-resize; default unit is `px`, consider using `em` unit),
- the viewBox defines the internal coordinate system (with a (0,0,500,200) viewBox, point
  (250,100) is the middle),
- the parameters to preserve the aspect ratio.

For more details about the `<svg/>` element attributes, see the W3 SVG
specification.

## Destroying an SVG canvas handler

When the SVG canvas handler is no longer needed (for example, before closing the form/window
displaying the corresponding `WEBCOMPONENT`), you can release resources allocated for
the SVG handle by calling the [`destroy()`](../15_library-reference/2854-fglsvgcanvas-destroy.md "Releases resources allocated for the SVG canvas.")
function:

```
DEFINE cid SMALLINT
...
LET cid = fglsvgcanvas.create("formonly.canvas")
...
CALL fglsvgcanvas.destroy( cid )
```

After calling the `destroy()` function, the SVG canvas handler that was created
before the destroyed handler, will be set as the new current SVG
canvas:

```
DEFINE cid1, cid2 SMALLINT
...
LET cid1 = fglsvgcanvas.create("formonly.canvas1") -- current canvas is cid1
...
LET cid2 = fglsvgcanvas.create("formonly.canvas2") -- current canvas is cid2
...
CALL fglsvgcanvas.destroy( cid2 ) -- current canvas is cid1
...
CALL fglsvgcanvas.destroy( cid1 ) -- no current canvas
```

## Building the SVG DOM tree

After creating the root SVG DOM node with the [`setRootSVGAttributes()`](../15_library-reference/2879-fglsvgcanvas-setrootsvgattributes.md "Produces the root SVG element.") function, create other DOM element with fglsvgcanvas
functions, and append the child nodes to the root element or
sub-elements:

```
DEFINE root_svg, n, g om.DomNode
...
LET n = fglsvgcanvas.svg( ... )  -- creates an <svg/> sub-node.
CALL root_svg.appendChild( n )
...
LET g = fglsvgcanvas.g( ... )  -- creates a <g/> sub-node.
CALL n.appendChild( g )
...
```

## Cleaning the SVG canvas

To clean up the SVG canvas content, use the [`clean()`](../15_library-reference/2846-fglsvgcanvas-clean.md "Deletes all SVG elements inside the SVG canvas.")
function:

```
CALL fglsvgcanvas.clean( cid )
```

## Defining CSS styles

SVG supports CSS styling. Styles must be defined in the `<defs/>` element, and
can then be referenced in SVG drawing elements by using the `class` attribute.

To create styles, start by defining a set of SVG attributes with an
`om.SaxAttributes` object. For attribute names, use the predefined [`SVGATT_*` constants](../15_library-reference/2843-fglsvgcanvas-svgatt-constants.md "List of predefined SVG attributes.")
available in the fglsvgcanvas library.
> **Tip:**
>
> Define a dynamic array of
> `om.SaxAttributes`, to define several reusable attribute
> sets.

```
CONSTANT COLORS_OCEAN = 1
CONSTANT COLORS_SAHARA = 2

DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes
...
LET attr[COLORS_OCEAN] = om.SaxAttributes.create()
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_FILL,           "cyan" )
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_FILL_OPACITY,   "0.3" )
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_STROKE,         "blue" )
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_STROKE_WIDTH,   "5" )
CALL attr[COLORS_OCEAN].addAttribute(SVGATT_STROKE_OPACITY, "0.3" )

LET attr[COLORS_SAHARA] = om.SaxAttributes.create()
CALL attr[COLORS_SAHARA].addAttribute(SVGATT_FILL,          "yellow" )
...
```

An SVG attribute set defined in an `om.SaxAttributes` object can be used in
different manners:

1. To explicitly set attributes in an SVG DOM node, with the [`setAttributes()`](../15_library-reference/2877-fglsvgcanvas-setattributes.md "Sets the SVG attributes from an attribute set.") function.
2. To define a CSS style with a selector and a list of `name:value;` pairs, with the
   [`styleDefinition()`](../15_library-reference/2882-fglsvgcanvas-styledefinition.md "Produces a CSS style definition with a selection and list of attributes.") function.
3. To set an inline-style in an DOM node, defining a list of `name:value;` pairs,
   with the [`styleAttributeList()`](../15_library-reference/2881-fglsvgcanvas-styleattributelist.md "Builds a string with a list of attributes to be used in a style attribute.") function.

To create a set of CSS styles, create a `<defs/>` SVG element, containing a
`<style/>` element including your attributes sets.

The `<style/>` element is created with the [`styleList()`](../15_library-reference/2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.")
function, and each CSS style string is created with the [`styleDefinition()`](../15_library-reference/2882-fglsvgcanvas-styledefinition.md "Produces a CSS style definition with a selection and list of attributes.") function, from the `om.SaxAttributes`
objects:

```
DEFINE defs om.DomNode
...
LET defs = fglsvgcanvas.defs( NULL )
CALL defs.appendChild( fglsvgcanvas.styleList(
                          fglsvgcanvas.styleDefinition(".style_ocean",attr[COLORS_OCEAN])
                       || fglsvgcanvas.styleDefinition(".style_sahara",attr[COLORS_SAHARA])
                       )
                     )
CALL root_svg.appendChild( defs )
```

Then, you can, for example, define a `<rect/>` element that references the
style defined with in the `<defs/>` node:

```
DEFINE r om.DomNode
LET r = fglsvgcanvas.rect(10,10,30,40,NULL,NULL)
CALL r.setAttribute(SVGATT_CLASS, "style_ocean")
```

## Displaying the SVG content

After creating the SVG content, it must be sent to the front-end for rendering.

To display the SVG content to the `WEBCOMPONENT` field associated with the SVG
canvas handle, use the [`display()`](../15_library-reference/2855-fglsvgcanvas-display.md "Displays the SVG canvas.") function:

```
CALL fglsvgcanvas.display( cid )
```

## Detecting SVG element selection

To enable clickable SVG elements, define the `WEBCOMPONENT` form field with a
`"selection"` property in the `PROPERTIES` attribute. This property
defines the action that will be fired when the user clicks on an SVG element. Use lowercase action
names:

```
WEBCOMPONENT cv = FORMONLY.canvas,
   ...
   PROPERTIES = (selection="item_selection"),
   ...
```

Define clickable SVG elements by using the `<g/>` SVG grouping element contain
child elements.

The child elements of the clickable group must not use the `fill:none` style,
otherwise they are not clickable.

The group element must have the `SVGATT_ONCLICK/"onclick"` attribute set to
`SVGVAL_ELEM_CLICKED`: The JavaScript function defined by `SVGVAL_ELEM_CLICKED`
is predefined in fglsvgcanvas.html, and will trigger the action defined by the
`selection` property defined by the `WEBCOMPONENT` form field.

The `<g/>` group element must be defined with an `"id"` attribute,
that will be used to identify the clicked elements:

```
DEFINE root_svg, g, c om.DomNode
...
CALL root_svg.appendChild( g := fglsvgcanvas.g("shape_1") )
CALL g.setAttribute(SVGATT_ONCLICK,SVGVAL_ELEM_CLICKED)
CALL g.appendChild( c := fglsvgcanvas.circle(10,20,3) )
CALL c.setAttribute(SVGATT_STYLE, 'stroke:gray;fill:blue;fill-opacity:0.3' )
CALL g.appendChild( c := fglsvgcanvas.circle(14,12,2) )
CALL c.setAttribute(SVGATT_STYLE, 'stroke:gray;fill:yellow;fill-opacity:0.8' )
```

> **Tip:**
>
> Note that child elements can be created, assigned to a variable with the
> `:=` operator and the resulting expression can be directly passed as
> `appendChild()`
> parameter.
>
> ```
> CALL root_svg.appendChild( g := fglsvgcanvas.g("shape_1") )
> ```

In the program code, detect SVG element selection with an `ON ACTION` handler using
the action name that matches the "selection" property of the `WEBCOMPONENT` field. When
the action is fired, the id of the selected SVG element is provided in the
`WEBCOMPONENT` field value:

```
DEFINE rec RECORD
          canvas STRING, -- Variable bound to the WEBCOMPONENT form field
          ...
       END RECORD
...
INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
   ON ACTION item_selection ATTRIBUTES(DEFAULTVIEW = NO)
      DISPLAY "Clicked element:", rec.canvas
   ...
```

Information about the clicked SVG element is provided in JSON notation, for example:

`{"id":"shape_1","source":"action","action":"item_selection"}`

It is also possible to use the [`fglsvgcanvas.getItemId()`](../15_library-reference/2861-fglsvgcanvas-getitemid.md "Returns SVG element id after a user action.") function to get the id of the clicked SVG element.
However, this produces a front call that can be avoided, since the element id is available in the
`WEBCOMPONENT` field value. The `getItemId()` function should only be
used when the `WEBCOMPONENT` field does not have the focus, for example to handle
mouse hovering
events with `mouse_event_focus = false`.

For a complete example, see [Example 2: Basic clickable SVG shapes with fglsvgcanvas](2424-example-2-basic-clickable-svg-shapes-with-fglsvgcanvas.md).

## Detecting mouse double-clicks on SVG elements

To detect a double-click on an SVG element, create SVG elements with the
`SVGATT_ONCLICK` attribute as described in Detecting SVG element selection, and define the
`selection2` property in the `WEBCOMPONENT` field:

```
   PROPERTIES = (
       ...
       selection2="item_selection2"
   )
```

And define the corresponding action handler in the code:

```
   ON ACTION item_selection2
```

Simple click and double-click actions can be used individually (when only double-click is
required, there is no need to define the `selection` property for single-clicks)
> **Important:**
>
> When both single-click (`selection` property) and
> double-click (`selection2` property) actions are used, the single-click
> (`selection`) action will be fired after a delay, because this is the only why to
> distinguish from double-click mouse events in SVG).

For a complete example, see [Example 2: Basic clickable SVG shapes with fglsvgcanvas](2424-example-2-basic-clickable-svg-shapes-with-fglsvgcanvas.md).

## Detecting mouse right-clicks on SVG elements

To detect a right-click (for context-menu) on an SVG element, define the
`selection3` property in the `WEBCOMPONENT` field:

```
   PROPERTIES = (
       ...
       selection3="item_selection3"
   )
```

And define the corresponding action handler in the code:

```
   ON ACTION item_selection3
```

For right-click/context-menu events, no `SVGATT_ONCLICK` attribute needs to be
defined for the SVG elements: The action defined by the `selection3` property will
automatically be fired for any SVG elements having an `id` attribute.

For a complete example, see [Example 2: Basic clickable SVG shapes with fglsvgcanvas](2424-example-2-basic-clickable-svg-shapes-with-fglsvgcanvas.md).

## Detecting mouse hovering on SVG elements

SVG elements can be defined with the mouse hovering events "`onmouseover`" and
"`onmouseout`", respectively defined as the `SVGATT_ONMOUSEOVER` and
`SVGATT_ONMOUSEOUT` constants in fglsvgcanvas.4gl for
convenience and code readability.

> **Tip:**
>
> If you want to implement display effects only, consider using well-known SVG +
> JavaScript techniques to modify element attributes directly in the web browser context, and thus
> avoid unnecessary network round-tips. For example, to produce some responsive rendering effect when
> the mouse goes over an SVG element, you can do
> following:
>
> ```
> DEFINE n om.DomNode
> LET n = fglsvgcanvas.rect(50,50,20,30,2,2)
> CALL n.setAttribute(SVGATT_ONMOUSEOVER, "evt.target.setAttribute('opacity', '0.5');")
> CALL n.setAttribute(SVGATT_ONMOUSEOUT,  "evt.target.setAttribute('opacity', '1.0');")
> ```

The fglsvgcanvas web component can be configured to bind the
`onmouseover/onmouseout` SVG mouse hovering events to `ON ACTION`
handlers and trigger code, when the mouse goes over the SVG elements.

> **Important:**
>
> When SVG element selection is enabled, the corresponding action has a higher
> importance than mouse hovering actions: Depending on the mouse hovering timeout
> (`mouse_event_timeout`), a click on an SVG element may cancel the mouse over / mouse
> out actions, even if corresponding SVG `onmouseover` and `onmouseout`
> events occurred respectively before and after the `onclick` SVG event.

The next code example defines the web component field properties to handle mouse hovering SVG
events:

```
WEBCOMPONENT cv=FORMONLY.canvas,
   ...
   PROPERTIES=(
       ...
       mouse_over = "item_mouse_over",
       mouse_out  = "item_mouse_out",
       mouse_event_timeout    = 600,
       mouse_event_focus      = false
   ),
```

Mouse hovering properties description:

1. The `mouse_over = "action-name"` property identifies the
   `ON ACTION action-name` handler to be triggered, when the
   `onmouseover` SVG event occurs. Use lowercase to define mouse hovering action names.
   This property is mandatory to catch mouse hovering events.
2. The `mouse_out = "action-name"` property identifies the
   `ON ACTION action-name` handler to be triggered, when the
   `onmouseout` SVG event occurs. Use lowercase to define mouse hovering action names.
   The property is typically used to reset mouse hovering management in the program code (the mouse
   goes away from that SVG element)
3. `mouse_event_timeout = milliseconds` is used to avoid network
   clogging: This property defines the number of milliseconds to wait before sending the dialog action,
   after an SVG mouse event has occurred. A JavaScript `setTimeout()` timer is created
   with this value. The timer is re-initialized each time the event occurs. If not defined, default is
   500 milliseconds.
4. `mouse_event_focus = {true|false}` specifies if the web
   component field must have the focus, in order to fire the mouse hovering actions. If not defined,
   default is `true`.

When creating your SVG elements, bind mouse hovering SVG events to predefined JavaScript function
corresponding to the SVG event. The functions take the current SVG element (`this`)
as parameter.

- The `SVGATT_ONMOUSEOVER` SVG event must be bound to
  `SVGVAL_ELEM_MOUSE_OVER`.
- The `SVGATT_ONMOUSEOUT` SVG event must be bound to
  `SVGVAL_ELEM_MOUSE_OUT`.

For
example:

```
DEFINE n om.DomNode
...
LET n = fglsvgcanvas.rect(x, y, w, h, NULL, NULL)
CALL n.setAttribute(SVGATT_ONMOUSEOVER, SVGVAL_ELEM_MOUSE_OVER)
CALL n.setAttribute(SVGATT_ONMOUSEOUT,  SVGVAL_ELEM_MOUSE_OUT)
```

In the dialog code, implement the `ON ACTION` handlers to execute code, when the
actions are fired.

If the mouse hovering action requires field focus (`mouse_event_focus=true`), the
related SVG element id is available in the field value, like for the item selection action. If the
mouse hovering action can be fired when the web component field does not have the focus, use the
[`fglsvgcanvas.getItemId()`](../15_library-reference/2861-fglsvgcanvas-getitemid.md "Returns SVG element id after a user action.") function to get the id of the SVG element.

```
...
        ON ACTION item_mouse_over ATTRIBUTES(DEFAULTVIEW = NO)
           MESSAGE SFMT("%1 : mouse over item : %2",
                        CURRENT HOUR TO FRACTION(5),
                        fglsvgcanvas.getItemId(cid) )

        ON ACTION item_mouse_out ATTRIBUTES(DEFAULTVIEW = NO)
           MESSAGE ""
...
```

For a complete example, see [Example 2: Basic clickable SVG shapes with fglsvgcanvas](2424-example-2-basic-clickable-svg-shapes-with-fglsvgcanvas.md).

## Getting the bounding box of an SVG element

After rendering your SVG, it is possible to get the bounding box of an element.

With elements such as SVG text using a specific font, it is difficult to compute the bounding box
of the text, until it has been rendered.

First define a variable with the [`fglsvgcanvas.t_svg_rect`](../15_library-reference/2842-fglsvgcanvas-t-svg-rect-type.md "The t_svg_rect type defines the position and dimensions of a rectangle.") type, then (after displaying the SVG with
`fglsvgcanvas.display(cid)`), call the [`fglsvgcanvas.getBBox(cid,
element-id)`](../15_library-reference/2860-fglsvgcanvas-getbbox.md "Returns the bounding box of an SVG element.") function, to get the bounding box of the element
identified by
element-id:

```
DEFINE rect fglsvgcanvas.t_svg_rect
...
ON ACTION get_bbox
   CALL fglsvgcanvas.getBBox(cid, "label_23") RETURNING rect.*
   DISPLAY rect.x, rect.y, rect.width, rect.height
```

The bounding box coordinates and size are returned in the current user space.

If the bounding box is required to place and size SVG element based on the position and size of
other elements, consider using SVG scripting instead of the `getBBox()` function:
this will avoid several network roundtrips to render the final SVG image.

For example:

```
<svg version="1.1" baseProfile="full"
     xmlns="http://www.w3.org/2000/svg" 
     width="500px" height="500px" viewBox="0 0 2000 2000"
     onload="setup()">

<script type="text/ecmascript">
// <![CDATA[

function setup_rect(box,bbox){
   box.setAttributeNS(null, "x", bbox.x - 2);
   box.setAttributeNS(null, "y", bbox.y - 2);
   box.setAttributeNS(null, "width", bbox.width + 4);
   box.setAttributeNS(null, "height", bbox.height + 4);
}

function setup(evt){
   setup_rect( document.getElementById("label1Box"), label1.getBBox() );
}

// ]]>
</script>
<g id="label1">
<text x="150" y="250" font-family="Verdana" font-size="55">Hello everybody!</text>
</g>
<rect id="label1Box" stroke="red" stroke-width="3px" fill="none"/>
</svg>
```

## Related links

**Related reference**  

[fglsvgcanvas: SVG drawing module](../15_library-reference/2841-fglsvgcanvas-svg-drawing-module.md "fglsvgcanvas: SVG drawing module")

## Child topics

- [Examples](2422-examples.md): fglsvgcanvas built-in web component usage examples.
