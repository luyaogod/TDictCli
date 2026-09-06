---
title: "Field-anchored windows"
source: "fgl-topics/c_fgl_windows_and_forms_win_pos_field.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Field-anchored windows"
type: "concept"
---

# Field-anchored windows

> The Window style attribute "position" can be set to "field" in order to display the window under the current field.

## The Window style attribute "`position`"

The "`position`" style attribute for Window elements can take several values, to
control where the window will appear on the screen. It can for example be anchored to an existing
window, or centered. For more details, see the [`position`](1655-window-style-attributes-basics.md) style attribute reference topic.

In this topic, we will describe how to use the "`field`" value of the Window
`position` style attribute, to display the window under the current field having the
focus. This will result in a "drop-down" window, typically used to select a record in a list:

![The figure shows a screenshot of a parent window displaying a drop-down window, to select a record from a list.](../_images/win_pos_field_1_gbc.jpg)

*Typical Window.position="field" usage*

## Defining your custom "dropdown" Window style

The "`position`" style attribute for Window elements must be defined under a
`Style` element.

We use the name "`dropdown`" to reference the style element in program source
files.

Create the following styles.4st file:

```
<StyleList>
  <Style name="Window.dropdown">
     <StyleAttribute name="position" value="field" />
     <StyleAttribute name="windowType" value="modal" />
     <StyleAttribute name="border" value="none" />
     <StyleAttribute name="windowSystemMenu" value="no" />
     <StyleAttribute name="actionPanelPosition" value="none" />
     <StyleAttribute name="ringMenuPosition" value="none" />
     <StyleAttribute name="statusBarType" value="none" />
  </Style>
<StyleList>
```

Additional Window style attributes are required to get the expected rendering. For example, we
remove the border, system menu, statusbar, action panel, etc, to get a minimal window frame. Note
also that the `windowType` attribute must be "`modal`".

## The zoom.per form

The zoom.per form defines a `TABLE` container, with the
`STYLE` attribute set to "`dropdown`":

```
LAYOUT
TABLE t1 (STYLE="dropdown")
{
[c1      |c2           ]
[c1      |c2           ]
[c1      |c2           ]
[c1      |c2           ]
[c1      |c2           ]
}
END
END
ATTRIBUTES
c1 = FORMONLY.col1, TITLE="Id";
c2 = FORMONLY.col2, TITLE="Name";
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

## The main.per form

The main.per form defines a `GRID` container with
two `BUTTONEDIT` fields, with an `ACTION`, that will fire
the `ON ACTION` program code to open the zoom window:

```
LAYOUT
GRID
{
[f1             ] [f2             ]
}
END      
END      
ATTRIBUTES
BUTTONEDIT f1 = FORMONLY.field1, ACTION=zoom1, IMAGE="zoom";
BUTTONEDIT f2 = FORMONLY.field2, ACTION=zoom2, IMAGE="zoom";
END
```

## The main.4gl program

The main.4gl module implements a simple `INPUT`
with `ON ACTION` handlers calling a function that opens the zoom window.

Note that the `OPEN WINDOW` instruction uses the `STYLE="dropdown"`
attribute:

```
MAIN
    DEFINE rec RECORD
                field1 STRING,
                field2 STRING
           END RECORD

    CALL ui.Interface.loadStyles("styles")

    OPEN FORM f1 FROM "main"
    DISPLAY FORM f1

    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
        ON ACTION zoom1
           LET rec.field1 = open_zoom(rec.field1)
        ON ACTION zoom2
           LET rec.field2 = open_zoom(rec.field2)
    END INPUT

END MAIN

FUNCTION open_zoom(cv)
    DEFINE cv STRING
    DEFINE arr DYNAMIC ARRAY OF RECORD
               col1 INT,
               col2 STRING
           END RECORD
    DEFINE i INT
    FOR i=1 TO 100
        LET arr[i].col1 = i
        LET arr[i].col2 = SFMT("Item %1",i)
    END FOR
    OPEN WINDOW wz WITH FORM "zoom" ATTRIBUTES(STYLE="dropdown")
    LET int_flag = FALSE
    DISPLAY ARRAY arr TO sr.*
        AFTER DISPLAY
          IF NOT int_flag THEN
             LET cv = arr[arr_curr()].col2
          END IF
    END DISPLAY
    CLOSE WINDOW wz
    RETURN cv
END FUNCTION
```

## Related links

**Related concepts**  

[Presentation styles](1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")
