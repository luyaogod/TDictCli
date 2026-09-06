---
title: "fglsvgcanvas.marker()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_marker.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.marker()"
type: "concept"
---

# fglsvgcanvas.marker()

> Produces an SVG "marker" element.

## Syntax

```
FUNCTION marker(
   id STRING,
   markerUnits STRING,
   refX STRING,
   refY STRING,
   markerWidth STRING,
   markerHeight STRING,
   orient STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. markerUnits defines how the marker scales ("strokeWidth" or "userSpaceOnUse").
3. refX defines the X coordinate of the reference point.
4. refY defines the Y coordinate of the reference point.
5. markerWidth defines the width of the marker.
6. markerHeight defines the height of the marker.
7. orient defines the orientation of the marker (like "auto") .

## Usage

This function creates a `"marker"` SVG DOM element from the parameters.

SVG markers define the rendering of start, mid and end of a line or path.

Markers are typically defined in a `"defs"` SVG element, and then used in a style
definition with the `marker-start`, `marker-mid` and
`marker-end` attributes, using the [`styleList()`](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.") and
[`styleDefinition()`](2882-fglsvgcanvas-styledefinition.md "Produces a CSS style definition with a selection and list of attributes.") functions (see example below).

## Example

Form file form.per:

```
LAYOUT
GRID
{
[wc1                    ]
[                       ]
[                       ]
}
END
END
ATTRIBUTES
WEBCOMPONENT wc1=FORMONLY.webcomp1,
   COMPONENTTYPE="fglsvgcanvas",
   SIZEPOLICY=FIXED,
   STRETCH=BOTH,
   SCROLLBARS=NONE
;
END
```

Program file main.4gl:

```
IMPORT FGL fglsvgcanvas

MAIN
    DEFINE webcomp1 STRING
    DEFINE root_svg, defs, m, n, p om.DomNode
    DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes
    DEFINE cid SMALLINT

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    CALL fglsvgcanvas.initialize()

    LET cid = fglsvgcanvas.create("formonly.webcomp1")

    LET root_svg = fglsvgcanvas.setRootSVGAttributes(
                      "svg1", "5em", NULL, "0 0 1000 1000", "xMidYMid meet" )
    CALL root_svg.setAttribute(SVGATT_CLASS,"root_svg")

    LET defs = fglsvgcanvas.defs( NULL )

    LET m = fglsvgcanvas.marker("m1", NULL, 5,5,10,10, "auto")
    CALL defs.appendChild( m )
    CALL m.appendChild( n:=fglsvgcanvas.circle(5,5,3) )
    CALL n.setAttribute(SVGATT_STYLE,'stroke:gray; fill:blue;')

    LET m = fglsvgcanvas.marker("m2", NULL, 2,6,15,15, "auto")
    CALL defs.appendChild( m )
    CALL m.appendChild( n:=fglsvgcanvas.path("M2,2 L2,11 L10,6 L2,2") )
    CALL n.setAttribute(SVGATT_STYLE,'stroke:gray; fill:blue;')

    LET attr[1] = om.SaxAttributes.create()
    CALL attr[1].addAttribute(SVGATT_STROKE, "blue" )
    CALL attr[1].addAttribute(SVGATT_STROKE_WIDTH, 1.5 )
    CALL attr[1].addAttribute(SVGATT_FILL, "none" )
    CALL attr[1].addAttribute(SVGATT_MARKER_START, fglsvgcanvas.url("m1") )
    CALL attr[1].addAttribute(SVGATT_MARKER_END,   fglsvgcanvas.url("m2") )
    CALL defs.appendChild( fglsvgcanvas.styleList(
                              fglsvgcanvas.styleDefinition(".style_1",attr[1])
                           )
                         )

    CALL root_svg.appendChild( defs )

    LET p = fglsvgcanvas.path("M100,100 L150,150 L200,150")
    CALL p.setAttribute(SVGATT_CLASS,"style_1")
    CALL root_svg.appendChild( p )

    CALL fglsvgcanvas.display(cid)

    MENU "test" COMMAND "Quit" EXIT MENU END MENU

    CALL fglsvgcanvas.finalize()

END MAIN
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
