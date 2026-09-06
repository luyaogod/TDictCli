---
title: "fglsvgcanvas.symbol()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_symbol.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.symbol()"
type: "concept"
---

# fglsvgcanvas.symbol()

> Produces an SVG "symbol" element.

## Syntax

```
FUNCTION symbol(
   id STRING,
   width STRING,
   height STRING,
   viewBox STRING,
   preserveAspectRatio STRING,
   refX STRING,
   refY STRING )
  RETURNS om.DomNode
```

1. id is the SVG object identifier.
2. width defines the width of the symbol.
3. height defines the height of the symbol.
4. viewBox defines the SVG viewbox to draw the symbol.
5. preserveAspectRatio is the aspect ratio to preserve for the symbol.
6. refX defines the X coordinate of the reference point.
7. refY defines the Y coordinate of the reference point.

## Usage

This function creates a `"symbol"` SVG DOM element from the parameters.

SVG symbols define graphical templates to be instantiated by a [`"use"`](2892-fglsvgcanvas-use.md "Produces an SVG \"use\" element.") element.

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
    DEFINE root_svg, defs, s, n, e om.DomNode
    DEFINE cid, x, y SMALLINT

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    CALL fglsvgcanvas.initialize()

    LET cid = fglsvgcanvas.create("formonly.webcomp1")

    LET root_svg = fglsvgcanvas.setRootSVGAttributes(
                      "svg1", "5em", NULL, "0 0 100 100", "xMidYMid meet" )
    CALL root_svg.setAttribute(SVGATT_CLASS,"root_svg")

    LET defs = fglsvgcanvas.defs( NULL )

    LET s = fglsvgcanvas.symbol("s1",10,10,"0 0 100 100","xMidYMid meet",NULL,NULL)
    LET n = fglsvgcanvas.rect(5,5,90,90,NULL,NULL)
    CALL s.appendChild( n )
    CALL n.setAttribute(SVGATT_FILL,"none")
    CALL n.setAttribute(SVGATT_STROKE,"green")
    LET n = fglsvgcanvas.rect(20,20,60,60,NULL,NULL)
    CALL n.setAttribute(SVGATT_FILL,"blue")
    CALL s.appendChild( n )
    CALL defs.appendChild( s )

    CALL root_svg.appendChild( defs )

    FOR x = 0 TO 90 STEP 10
        FOR y = 0 TO 90 STEP 10
            LET e = fglsvgcanvas.use("s1",x,y)
            CALL root_svg.appendChild( e )
        END FOR
    END FOR

    CALL fglsvgcanvas.display(cid)

    MENU "test" COMMAND "Quit" EXIT MENU END MENU

    CALL fglsvgcanvas.finalize()

END MAIN
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
