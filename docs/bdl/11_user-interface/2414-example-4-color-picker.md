---
title: "Example 4: Color picker"
source: "fgl-topics/c_fgl_webcomponent_gicapi_example_4.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Examples > Example 4: Color picker"
type: "concept"
description: "Introduction This topic describes the different steps to implement a color picker with a gICAPI-based web component. In this example, we will implement a simple color picker, that will allow the user ..."
---

# Example 4: Color picker

## Introduction

This topic describes the different steps to implement a color picker with a gICAPI-based web
component.

In this example, we will implement a simple color picker, that will allow the user the select
a color from a predefined set. Colors are drawn as square boxes using SVG graphics, user can
change the current selected color with a separate `COMBOBOX` field, modify the
title of the HTML body, and query for the color list with a `webcomponent.call`
front call.

The HTML file is described in detail, and complete code example with program and form file is
available at the end of this topic.

## HTML code description

As any HTML source code, the file starts with the typical HTML tags:

```
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html" charset="utf-8" />
<meta name='viewport' content='initial-scale=1.0, maximum-scale=1.0' />
```

> **Note:**
>
> The "viewport" meta is provided to adjust the viewport size for mobile devices.

The JavaScript code needs to the enclosed in a `<script>` element:

```
<script language="JavaScript">
```

Global variables are defined to hold information that must be persistent during the web component
life:

```
var current_color;
var wanted_color;
var has_focus;
```

The global function `onICHostReady()` will be called by the front-end, when the web
component interface is ready. The version passed as parameter allows you to check that your component
code is compatible with the current gICAPI framework, and to define and assign the
`gICAPI.on*` callback methods (these will be defined in the body of the
`onICHostReady()` function:

```
var onICHostReady = function(version) {

    if ( version != "1.0" ) {
        alert('Invalid API version');
    }

    ... some initialization code ...

    gICAPI.onProperty = function(properties) {
        ... see below for function body ...
    }

    gICAPI.onData = function(data) {
        ... see below for function body ...
    }

    gICAPI.onFocus = function(polarity) {
        ... see below for function body ...
    }

}
```

At this point, the gICAPI interface is ready and the `gICAPI` object can be used.

The `gICAPI.onProperty()` method is called when a web component property changes
(properties will be initialized at form creation, or changed during form usage). In this code example,
when the property "title" is changed by the program, the element with id="title" is updated with the new
title:

```
gICAPI.onProperty = function(properties) {
    try{
        var ps = JSON.parse(properties);
        document.getElementById("title").innerHTML = ps.title;
    }
    catch (err){
        console.error("onProperty(): Invalid JSON string");
    }
}
```

> **Note:**
>
> The `ON ACTION change_title` in the dialog code will change the title property
> after the form initialization, to show that the `gICAPI.onProperty()` function can also
> be invoked after the web component field creation.

The `showFocusRectangle()` function shows a border around the specified color item
(SVG element), following the color identifier passed as parameter and the focus status (focus can be
true, false or -1, to keep the current border color and just modify the position of the
border):

```
var showFocusRectangle = function(color, focus) {
    // See complete code example for details
}
```

The `changeColor()` function implements a color change, by registering a field value
change with `gICAPI.SetData()`, and by triggering a specific action with
`gICAPI.Action()`, to inform the program that a color was selected:

```
var changeColor = function(color) {
    current_color = color;
    showFocusRectangle(current_color, true);
    gICAPI.SetData(current_color);
    gICAPI.Action("color_selected");
}
```

Next lines implement the `gICAPI.onFocus()` function, executed when the web component
gets or loses the focus. The code distinguishes the case when the focus is gained (by a mouse click on
a color item), selecting a new color with a call to `changeColor()`, and the case when
the focus is set to the web component by the runtime system. A blue border will be added to the current
color item, when the component gets the focus, and the border color is reset to gray when the focus is
lost:

```
gICAPI.onFocus = function(polarity) {
    if ( polarity == true ) {
        has_focus = true;
        if (wanted_color != undefined) {
            changeColor(wanted_color);
            wanted_color = undefined;
        } else {
            showFocusRectangle(current_color, true);
        }
    } else {
        has_focus = false;
        showFocusRectangle(current_color, false);
    }
}
```

The `gICAPI.onData()` function must be implemented to detect web component value
changes done in the program, and to acknowledge `gICAPI.SetData()` calls. This will
be triggered by assigning the `rec.webcomp` variable in the dialog code, typically
in the `ON CHANGE color` block, when modifying the combobox value.
The `showFocusRectangle()` function moves the focus border to the color item
corresponding to the color identifier passed as parameter.:

```
gICAPI.onData = function(data) {
    current_color = data;
    showFocusRectangle(current_color, -1);
}
```

The `selectColor()` function will be called through the `onclick`
event of the `<rect>` SVG elements representing colors. If the web component does
not have the focus yet, the function will call `gICAPI.SetFocus()`, in order to ask
the runtime system, if the focus can go to the web component field. If the runtime system accepts
to set the focus to the web component field, the `gICAPI.onFocus()` method will be
called with `true` as parameter, and will handle the requested color change (using
`wanted_color`). if the focus cannot be set to the web component, the
`onFocus()` method will not be called:

```
var selectColor = function(color) {   
    if (has_focus) {
        changeColor(color);
    } else {
        wanted_color = color;
        gICAPI.SetFocus();
        // Color item change is done in onFocus(), because
        // VM may refuse to set the focus to the wc field.
    }
}
```

> **Note:**
>
> The only way to detect that the focus was gained by the web component field, is
> when `onFocus(true)` is called.

End the JavaScript element with the `</script>` ending tag:

```
</script>
```

Close the HTML head element with the `</head>` ending tag:

```
</head>
```

The rest of the HTML page defines the graphical elements for the color picker, with a
`<h3>` title and and an `<svg>` element containing
`<rect>` element to show clickable color items. Note the
`<rect>` element with `id="focus_rectangle"`,
used to show a border for the current color item:

```
<body height="100%" width="100%">
<h3 id="title">no-title</h3>

<svg id="svg_container" width="230" height="130">

  <rect x="5" y= "5" rx="5" ry="5" width="30" height="30"
        id="#FFFFCC"
        style="fill:#FFFFCC;stroke:black;stroke-width:1"
        onclick="selectColor('#FFFFCC')" />
  ...

  <rect x="178" y= "73" rx="7" ry="7" width="34" height="34"
        id="focus_rectangle"
        style="fill:none;stroke:gray;stroke-width:3" />

</svg>

</body>
```

## Complete source code

File color\_picker.per:

```
LAYOUT
GRID
{
 Id:    [f1             ]
[f2                     ]
[                       ]
[                       ]
[                       ]
[                       ]
[                       ]
[                       ]
[                       ]
 Color: [f3             ]
[f4                     ]
[                       ]
[                       ]
}
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.id;
WEBCOMPONENT f2 = FORMONLY.webcomp,
   COMPONENTTYPE="color_picker",
   PROPERTIES = (title="My color picker"),
   STRETCH=BOTH;
COMBOBOX f3 = FORMONLY.pgcolor, NOT NULL,
   ITEMS=( "#FFFFCC", "#FFFFAA", "#FFFF00",
           "#FFAD99", "#FF0000", "#990000",
           "#99CCFF", "#0066FF", "#000099",
           "#FF99FF", "#FF00FF", "#990099",
           "#99FF99", "#009933", "#006600",
           "#FFFFFF", "#AAAAAA", "#000000" );
TEXTEDIT f4 = FORMONLY.info, STRETCH=X;
END
```

File color\_picker.4gl:

```
IMPORT util

MAIN
    DEFINE rec RECORD
               id INTEGER,
               webcomp STRING,
               pgcolor STRING,
               info STRING
           END RECORD,
           f ui.Form,
           n om.DomNode,
           tmp STRING,
           colors DYNAMIC ARRAY OF STRING

    OPTIONS INPUT WRAP

    OPEN FORM f1 FROM "color_picker"
    DISPLAY FORM f1

    LET rec.id = 98344
    LET rec.webcomp = "#FF0000"
    LET rec.pgcolor = rec.webcomp

    INPUT BY NAME rec.* WITHOUT DEFAULTS
          ATTRIBUTES(UNBUFFERED)

        ON CHANGE pgcolor
           LET rec.webcomp = rec.pgcolor

        ON ACTION color_selected ATTRIBUTES( DEFAULTVIEW=NO )
           IF rec.webcomp == "#000000" THEN
              LET rec.webcomp = rec.pgcolor
              LET rec.info = NULL
              ERROR "Black color is denied!"
           ELSE
              LET rec.pgcolor = rec.webcomp
              LET rec.info = "Color selected:", rec.pgcolor
           END IF

        ON ACTION change_title ATTRIBUTES(TEXT="Change title")
           LET f = DIALOG.getForm()
           LET n = f.findNode("Property", "title")
           CALL n.setAttribute("value", "New title")
           LET rec.info = "Title changed."

        ON ACTION get_colors ATTRIBUTES(TEXT="Get colors")
           TRY
               CALL ui.Interface.frontCall("webcomponent", "call",
                       ["formonly.webcomp", "getColorList"], [tmp] )
               CALL util.JSON.parse(tmp, colors)
               LET rec.info = "Color list: ", tmp
           CATCH
               ERROR "Front call failed."
           END TRY

    END INPUT

END MAIN
```

File color\_picker.html:

```
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html" charset="utf-8" />
<meta name='viewport' content='initial-scale=1.0, maximum-scale=1.0' />

<script language="JavaScript">

var current_color;
var wanted_color;
var has_focus;

var onICHostReady = function(version) {

    if ( version != "1.0" ) {
        alert('Invalid API version');
    }

    current_color = "#000000";

    gICAPI.onProperty = function(properties) {
        try{
            var ps = JSON.parse(properties);
            document.getElementById("title").innerHTML = ps.title;
        }
        catch (err){
            console.error("onProperty(): Invalid JSON string");
        }
    }

    gICAPI.onFocus = function(polarity) {
        if ( polarity == true ) {
            has_focus = true;
            if (wanted_color != undefined) {
                changeColor(wanted_color);
                wanted_color = undefined;
            } else {
                showFocusRectangle(current_color, true);
            }
        } else {
            has_focus = false;
            showFocusRectangle(current_color, false);
        }
    }

    gICAPI.onData = function(data) {
        current_color = data;
        showFocusRectangle(current_color, -1);
    }

}

var showFocusRectangle = function(color, focus) {
    var f = document.getElementById("focus_rectangle");
    var e = document.getElementById(color);
    if (e == null) {
        e = document.getElementById("#000000");
    }
    var e_x = e.getAttribute("x") - 2;
    var e_y = e.getAttribute("y") - 2;
    f.setAttribute("x", e_x );
    f.setAttribute("y", e_y );
    if (focus == true) {
        f.style.stroke = "blue";
    } else if (focus == false) {
        f.style.stroke = "gray";
    }
}

var changeColor = function(color) {
    current_color = color;
    showFocusRectangle(current_color, true);
    gICAPI.SetData(current_color);
    gICAPI.Action("color_selected");
}

var selectColor = function(color) {
    if (has_focus) {
        changeColor(color);
    } else {
        wanted_color = color;
        gICAPI.SetFocus();
        // Color item change is done in onFocus(), because
        // VM may refuse to set the focus to the wc field.
    }
}

var getColorList = function() {
    var colors = [];
    var p = document.getElementById("svg_container");
    var items = p.getElementsByTagName("rect");
    for (var i = items.length; i--;) {
        var c = items[i].getAttribute("id");
        if (c.indexOf("#") == 0)
            colors.push(c);
    }
    return JSON.stringify(colors);
}

</script>

</head>

<body height="100%" width="100%">
<h3 id="title">no-title</h3>

<svg id="svg_container" width="230" height="130">

  <rect x="5" y= "5" rx="5" ry="5" width="30" height="30"
        id="#FFFFCC"
        style="fill:#FFFFCC;stroke:black;stroke-width:1"
        onclick="selectColor('#FFFFCC')" />
  <rect x="5" y="40" rx="5" ry="5" width="30" height="30"
        id="#FFFFAA"
        style="fill:#FFFFAA;stroke:black;stroke-width:1"
        onclick="selectColor('#FFFFAA')" />
  <rect x="5" y="75" rx="5" ry="5" width="30" height="30"
        id="#FFFF00"
        style="fill:#FFFF00;stroke:black;stroke-width:1"
        onclick="selectColor('#FFFF00')" />

  <rect x="40" y= "5" rx="5" ry="5" width="30" height="30"
        id="#FFAD99"
        style="fill:#FFAD99;stroke:black;stroke-width:1"
        onclick="selectColor('#FFAD99')" />
  <rect x="40" y="40" rx="5" ry="5" width="30" height="30"
        id="#FF0000"
        style="fill:#FF0000;stroke:black;stroke-width:1"
        onclick="selectColor('#FF0000')" />
  <rect x="40" y="75" rx="5" ry="5" width="30" height="30"
        id="#990000"
        style="fill:#990000;stroke:black;stroke-width:1"
        onclick="selectColor('#990000')" />

  <rect x="75" y= "5" rx="5" ry="5" width="30" height="30"
        id="#99CCFF"
        style="fill:#99CCFF;stroke:black;stroke-width:1"
        onclick="selectColor('#99CCFF')" />
  <rect x="75" y="40" rx="5" ry="5" width="30" height="30"
        id="#0066FF"
        style="fill:#0066FF;stroke:black;stroke-width:1"
        onclick="selectColor('#0066FF')" />
  <rect x="75" y="75" rx="5" ry="5" width="30" height="30"
        id="#000099"
        style="fill:#000099;stroke:black;stroke-width:1"
        onclick="selectColor('#000099')" />

  <rect x="110" y= "5" rx="5" ry="5" width="30" height="30"
        id="#FF99FF"
        style="fill:#FF99FF;stroke:black;stroke-width:1"
        onclick="selectColor('#FF99FF')" />
  <rect x="110" y="40" rx="5" ry="5" width="30" height="30"
        id="#FF00FF"
        style="fill:#FF00FF;stroke:black;stroke-width:1"
        onclick="selectColor('#FF00FF')" />
  <rect x="110" y="75" rx="5" ry="5" width="30" height="30"
        id="#990099"
        style="fill:#990099;stroke:black;stroke-width:1"
        onclick="selectColor('#990099')" />

  <rect x="145" y= "5" rx="5" ry="5" width="30" height="30"
        id="#99FF99"
        style="fill:#99FF99;stroke:black;stroke-width:1"
        onclick="selectColor('#99FF99')" />
  <rect x="145" y="40" rx="5" ry="5" width="30" height="30"
        id="#009933"
        style="fill:#009933;stroke:black;stroke-width:1"
        onclick="selectColor('#009933')" />
  <rect x="145" y="75" rx="5" ry="5" width="30" height="30"
        id="#006600"
        style="fill:#006600;stroke:black;stroke-width:1"
        onclick="selectColor('#006600')" />

  <rect x="180" y= "5" rx="5" ry="5" width="30" height="30"
        id="#FFFFFF"
        style="fill:#FFFFFF;stroke:black;stroke-width:1"
        onclick="selectColor('#FFFFFF')" />
  <rect x="180" y="40" rx="5" ry="5" width="30" height="30"
        id="#AAAAAA"
        style="fill:#AAAAAA;stroke:black;stroke-width:1"
        onclick="selectColor('#AAAAAA')" />
  <rect x="180" y="75" rx="5" ry="5" width="30" height="30"
        id="#000000"
        style="fill:#000000;stroke:gray;stroke-width:1"
        onclick="selectColor('#000000')" />

  <rect x="178" y= "73" rx="7" ry="7" width="34" height="34"
        id="focus_rectangle"
        style="fill:none;stroke:gray;stroke-width:3" />

</svg>

</body>
</html>
```
