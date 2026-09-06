---
title: "Example 2: Basic clickable SVG shapes with fglsvgcanvas"
source: "fgl-topics/c_fgl_fglsvgcanvas_example_2.html"
breadcrumb: "User interface > User interface programming > Web components > Built-in web components > Built-in web components reference > The fglsvgcanvas web component > Examples > Example 2: Basic clickable SVG shapes with fglsvgcanvas"
type: "concept"
description: "Example using the fglsvgcanvas web component to draw basic shapes and detect mouse events. Figure: fglsvgcanvas web component - basics Form definition file svgbasics.per : LAYOUT GRID { [cv |fi ] [ | ..."
---

# Example 2: Basic clickable SVG shapes with fglsvgcanvas

Example using the fglsvgcanvas web component to draw basic shapes and detect mouse events.

![Screenshot of a program using the fglsvgcanvas web component](../_images/fglsvgcanvas_example_basics_gbc.jpg)

*fglsvgcanvas web component - basics*

Form definition file
svgbasics.per:

```
LAYOUT
GRID
{
[cv                               |fi              ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
[                                 |                ]
}
END
END

ATTRIBUTES

WEBCOMPONENT cv=FORMONLY.canvas,
   COMPONENTTYPE="fglsvgcanvas",
   PROPERTIES=(
     selection = "item_selection",
     selection2 = "item_selection2",
     selection3 = "item_selection3",
     mouse_over = "item_mouse_over",
     mouse_out  = "item_mouse_out",
     mouse_event_timeout = 600,
     mouse_event_focus   = false
   ),
   SIZEPOLICY=FIXED,
   STRETCH=BOTH,
   SCROLLBARS=NONE
;
TEXTEDIT fi = FORMONLY.info,
   STRETCH=BOTH;

END
```

Program file
svgbasics.4gl:

```
IMPORT util
IMPORT FGL fglsvgcanvas

DEFINE rec RECORD
           canvas STRING,
           info STRING
       END RECORD

DEFINE canvas_id SMALLINT

MAIN

    OPEN FORM f1 FROM "svgbasics"
    DISPLAY FORM f1

    CALL fglsvgcanvas.initialize()
    LET canvas_id = create_svg()
    CALL fglsvgcanvas.display(canvas_id)

    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
        ON ACTION clear_info ATTRIBUTES(TEXT = "Clear info")
           LET rec.info = NULL
        ON ACTION item_selection ATTRIBUTES(DEFAULTVIEW=NO)
           CALL add_info_line("selection")
        ON ACTION item_selection2 ATTRIBUTES(DEFAULTVIEW=NO)
           CALL add_info_line("selection2")
        ON ACTION item_selection3 ATTRIBUTES(DEFAULTVIEW=NO)
           CALL add_info_line("selection3")
        ON ACTION item_mouse_over ATTRIBUTES(DEFAULTVIEW = NO)
           CALL add_info_line("mouse over")
        ON ACTION item_mouse_out ATTRIBUTES(DEFAULTVIEW = NO)
           CALL add_info_line("mouse out")
    END INPUT

    CALL fglsvgcanvas.destroy(canvas_id)
    CALL fglsvgcanvas.finalize()

END MAIN

FUNCTION create_svg() RETURNS SMALLINT

    CONSTANT CB = 1
    CONSTANT CY = 2
    CONSTANT CG = 3
    CONSTANT T1 = 4
    DEFINE root_svg om.DomNode,
           attr DYNAMIC ARRAY OF om.SaxAttributes,
           defs, g, n om.DomNode
    DEFINE cid SMALLINT

    LET cid = fglsvgcanvas.create("formonly.canvas")
    LET root_svg = fglsvgcanvas.setRootSVGAttributes( NULL,
                                   NULL, NULL,
                                   "0 0 500 500",
                                   "xMidYMid meet"
                                )
    CALL root_svg.setAttribute(SVGATT_CLASS,"root_svg")

    LET attr[CB] = om.SaxAttributes.create()
    CALL attr[CB].addAttribute(SVGATT_FILL,           "cyan" )
    CALL attr[CB].addAttribute(SVGATT_FILL_OPACITY,   "0.3" )
    CALL attr[CB].addAttribute(SVGATT_STROKE,         "blue" )
    CALL attr[CB].addAttribute(SVGATT_STROKE_WIDTH,   "5" )
    CALL attr[CB].addAttribute(SVGATT_STROKE_OPACITY, "0.3" )

    LET attr[CY] = om.SaxAttributes.create()
    CALL attr[CY].addAttribute(SVGATT_FILL,           "yellow" )
    CALL attr[CY].addAttribute(SVGATT_FILL_OPACITY,   "0.8" )
    CALL attr[CY].addAttribute(SVGATT_STROKE,         "orange" )
    CALL attr[CY].addAttribute(SVGATT_STROKE_WIDTH,   "4" )
    CALL attr[CY].addAttribute(SVGATT_STROKE_OPACITY, "0.8" )

    LET attr[CG] = om.SaxAttributes.create()
    CALL attr[CG].addAttribute(SVGATT_FILL,           "green" )
    CALL attr[CG].addAttribute(SVGATT_FILL_OPACITY,   "0.8" )
    CALL attr[CG].addAttribute(SVGATT_STROKE,         "darkGreen" )
    CALL attr[CG].addAttribute(SVGATT_STROKE_WIDTH,   "4" )
    CALL attr[CG].addAttribute(SVGATT_STROKE_OPACITY, "0.8" )

    LET attr[T1] = om.SaxAttributes.create()
    CALL attr[T1].addAttribute(SVGATT_STROKE,         "gray" )
    CALL attr[T1].addAttribute(SVGATT_STROKE_WIDTH,   "1" )
    CALL attr[T1].addAttribute(SVGATT_STROKE_LINECAP, "round" )
    CALL attr[T1].addAttribute(SVGATT_FILL,           "blue" )
    CALL attr[T1].addAttribute(SVGATT_FONT_FAMILY,    "Sans" )
    CALL attr[T1].addAttribute(SVGATT_FONT_SIZE,      "24px" )

    LET defs = fglsvgcanvas.defs( NULL )
    CALL defs.appendChild( fglsvgcanvas.styleList(
                              fglsvgcanvas.styleDefinition(".style_1",attr[CB])
                           || fglsvgcanvas.styleDefinition(".style_2",attr[CY])
                           || fglsvgcanvas.styleDefinition(".style_3",attr[CG])
                           || fglsvgcanvas.styleDefinition(".style_4",attr[T1])
                           )
                         )
    CALL root_svg.appendChild( defs )

    CALL root_svg.appendChild( fglsvgcanvas.text(30, 40,
           "Basic fglsvgcanvas example...", "style_4" ) )

    CALL root_svg.appendChild( n:=fglsvgcanvas.rect(200,200,350,150,10,10) )
    CALL n.setAttribute("id","R1")
    CALL n.appendChild(fglsvgcanvas.title("Rectangle 1"))
    CALL n.setAttribute(SVGATT_STYLE,'stroke:black;fill:orange' )

    CALL root_svg.appendChild( n:=fglsvgcanvas.circle(40,120,50) )
    CALL n.setAttribute("id","C1")
    CALL n.appendChild( fglsvgcanvas.title("Circle 1 (with mouse hovering effects)") )
    CALL n.setAttribute(SVGATT_CLASS, "style_1")
    CALL n.setAttribute(SVGATT_ONMOUSEOVER, "evt.target.setAttribute('opacity', '0.5');")
    CALL n.setAttribute(SVGATT_ONMOUSEOUT,  "evt.target.setAttribute('opacity', '1.0');")

    CALL root_svg.appendChild( n:=fglsvgcanvas.ellipse(50,400,50,30) )
    CALL n.setAttribute("id","E1")
    CALL n.appendChild( fglsvgcanvas.title("Ellipse 1 (clickable)") )
    CALL n.setAttribute(SVGATT_CLASS, "style_3")
    CALL n.setAttribute(SVGATT_ONCLICK, SVGVAL_ELEM_CLICKED)

    CALL root_svg.appendChild( n:=fglsvgcanvas.rect(400,100,150,50,5,10) )
    CALL n.setAttribute("id","R2")
    CALL n.appendChild( fglsvgcanvas.title("Rectangle 2 (with mouse hovering actions)") )
    CALL n.setAttribute(SVGATT_CLASS, "style_2")
    CALL n.setAttribute(SVGATT_ONMOUSEOVER, SVGVAL_ELEM_MOUSE_OVER)
    CALL n.setAttribute(SVGATT_ONMOUSEOUT, SVGVAL_ELEM_MOUSE_OUT)

    CALL root_svg.appendChild( g:=fglsvgcanvas.g( "G1" ) )
    CALL g.appendChild( fglsvgcanvas.title(
           "Group of elements 1 (clickable and mouse hovering actions)") )
    CALL g.setAttribute(SVGATT_ONCLICK, SVGVAL_ELEM_CLICKED)
    CALL g.setAttribute(SVGATT_ONMOUSEOVER, SVGVAL_ELEM_MOUSE_OVER)
    CALL g.setAttribute(SVGATT_ONMOUSEOUT, SVGVAL_ELEM_MOUSE_OUT)
    CALL g.appendChild( n:=fglsvgcanvas.rect(380,390,100,100,NULL,NULL) )
    CALL n.setAttribute(SVGATT_STYLE,'stroke:black;fill:none' )
    CALL g.appendChild( n:=fglsvgcanvas.rect(400,400,70,50,5,10) )
    CALL n.setAttribute(SVGATT_CLASS, "style_2" )
    CALL g.appendChild( n:=fglsvgcanvas.circle(420,450,30) )
    CALL n.setAttribute(SVGATT_CLASS, "style_3" )

    CALL root_svg.appendChild( n:=fglsvgcanvas.rect(100,150,50,80,20,20) )
    CALL n.setAttribute("id","R3")
    CALL n.appendChild( fglsvgcanvas.title(
           "Rectangle 3 (with animation, clickable and mouse hovering actions)") )
    CALL n.setAttribute(SVGATT_STYLE,'stroke:black;fill:#AA2244' )
    CALL n.setAttribute(SVGATT_ONCLICK, SVGVAL_ELEM_CLICKED)
    CALL n.setAttribute(SVGATT_ONMOUSEOVER, SVGVAL_ELEM_MOUSE_OVER)
    CALL n.setAttribute(SVGATT_ONMOUSEOUT, SVGVAL_ELEM_MOUSE_OUT)
    CALL n.appendChild( fglsvgcanvas.animateTransform("transform", "XML", "rotate",
                                                      "0 100 200", "360 200 150", NULL,
                                                      "0s", "20s", "indefinite" ) )

    RETURN cid

END FUNCTION

FUNCTION add_info_line(event STRING) RETURNS ()
    LET rec.info = rec.info, SFMT( "%1: %2 - %3",
                       CURRENT HOUR TO FRACTION(3),
                       fglsvgcanvas.getItemId(canvas_id),
                       event), "\n"
END FUNCTION
```
