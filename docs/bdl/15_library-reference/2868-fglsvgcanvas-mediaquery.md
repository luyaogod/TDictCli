---
title: "fglsvgcanvas.mediaQuery()"
source: "fgl-topics/c_fgl_utility_functions_fglsvgcanvas_mediaQuery.html"
breadcrumb: "Library reference > Utility modules > fglsvgcanvas: SVG drawing module > fglsvgcanvas.mediaQuery()"
type: "concept"
---

# fglsvgcanvas.mediaQuery()

> Builds an SVG "@media" query for CSS.

## Syntax

```
FUNCTION mediaQuery(
   query STRING,
   content STRING )
  RETURNS STRING
```

1. query is a CSS media query list that must match for styles content to
   apply.
2. content defines the styles to apply when the media query matches, build from
   a `styleDefinition()` call.

## Usage

This function creates a CSS `"@media"` entry, using the query string passed as
parameter, that defines the condition for the content styles to apply.

Use `@media` definitions to implement responsible SVG that adapts to the size of
the viewport.

Combine `mediaQuery()` with [`styleList()`](2883-fglsvgcanvas-stylelist.md "Produces a CSS style list.") and
[`styleDefinition()`](2882-fglsvgcanvas-styledefinition.md "Produces a CSS style definition with a selection and list of attributes.") to build the CSS `@media` entry:

```
DEFINE sl om.DomNode
DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes
... attr[] definitions ...
LET sl = fglsvgcanvas.styleList(
            fglsvgcanvas.styleDefinition(".detail-1",attr[1])
         || fglsvgcanvas.mediaQuery("(min-width: 800px)",
               fglsvgcanvas.styleDefinition(".detail-1",attr[2])
            )
         )
```

The resulting DomNode can then be included as child of a [`"defs"`](2853-fglsvgcanvas-defs.md "Produces an SVG \"defs\" element.") element.

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
    DEFINE root_svg om.DomNode
    DEFINE attr DYNAMIC ARRAY OF om.SaxAttributes
    DEFINE defs, e om.DomNode
    DEFINE cid SMALLINT

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    CALL fglsvgcanvas.initialize()

    LET cid = fglsvgcanvas.create("formonly.webcomp1")

    LET root_svg = fglsvgcanvas.setRootSVGAttributes(
                      "svg1","5em",NULL,"0 0 100 100","xMidYMid meet")
    CALL root_svg.setAttribute(SVGATT_CLASS,"root_svg")

    LET attr[1] = om.SaxAttributes.create()
    CALL attr[1].addAttribute("visibility","hidden")
    LET attr[2] = om.SaxAttributes.create()
    CALL attr[2].addAttribute("visibility","visible")
    CALL attr[2].addAttribute("font-size","0.4em")

    LET defs = fglsvgcanvas.defs( NULL )
    CALL defs.appendChild( fglsvgcanvas.styleList(
                              fglsvgcanvas.styleDefinition(".detail-1",attr[1])
                           || fglsvgcanvas.mediaQuery("(min-width: 800px)",
                                 fglsvgcanvas.styleDefinition(".detail-1",attr[2])
                              )
                           )
                         )

    CALL root_svg.appendChild( defs )

    LET e = fglsvgcanvas.circle(50,50,15)
    CALL e.setAttribute(SVGATT_FILL,"none")
    CALL e.setAttribute(SVGATT_STROKE,"red")
    CALL root_svg.appendChild( e )

    LET e = fglsvgcanvas.text(50,50,"Label 1","detail-1")
    CALL e.setAttribute(SVGATT_TEXT_ANCHOR,"middle")
    CALL e.setAttribute("dominant-baseline","middle")
    CALL root_svg.appendChild( e )

    LET e = fglsvgcanvas.rect(0,0,100,100,NULL,NULL)
    CALL e.setAttribute(SVGATT_FILL,"none")
    CALL e.setAttribute(SVGATT_STROKE,"blue")
    CALL root_svg.appendChild( e )

    CALL fglsvgcanvas.display(cid)

    MENU "test" COMMAND "Quit" EXIT MENU END MENU

    CALL fglsvgcanvas.finalize()

END MAIN
```

## Related links

**Related concepts**  

[The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.")
