---
title: "Example 2: Simple text input"
source: "fgl-topics/c_fgl_webcomponent_gicapi_example_2.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Examples > Example 2: Simple text input"
type: "concept"
description: "Introduction This topic describes the different steps to implement a simple gICAPI-based web component. In this example, we will implement a simple text editor based on a textarea HTML element. The ..."
---

# Example 2: Simple text input

## Introduction

This topic describes the different steps to implement a simple gICAPI-based web
component.

In this example, we will implement a simple text editor based on a
`textarea` HTML element.

The dialog code implements a couple of triggers to show how the `WEBCOMPONENT`
field interacts with the program.

The HTML file is described in detail, and complete code example with program and form
file is available at the end of this topic.

## HTML code description

As any HTML source code, the file starts with the typical HTML
tags:

```
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html" charset="utf-8" />
<meta name='viewport' content='initial-scale=1.0, maximum-scale=1.0' />
```

> **Note:**
>
> The "viewport" meta is provided to adjust the viewport size for mobile
> devices.

A bunch of CSS is added to

```
<style>
html,body {
  height:100%;
  padding:0;
  margin:0;
  border:0;
  overflow:hidden;
}

textarea#field1 {
  font-weight: bold;
}

textarea {
  display: block;
  font-family: fixed;
  font-size: 10px;
  padding:0;
  margin:0;
  width: 99%;
}
</style>
```

The HTML head is then ended with the typical ending tag:

```
</head>
```

The body of the HTML file defines a `<textarea>` element and references the
external JavaScript file in a `<script>` element:

```
<body>
<textarea id="field1"></textarea>
<script type="text/javascript" src="js/wc_simple.js"></script>
</body>
```

Finally, we end the HTML page with the final tag:

```
</html>
```

## The wc\_simple.js file

The JavaScript code implementing the gICAPI web component starts with some global variables.
These variables will hold information that must be persistent during the web component life:

```
var has_focus;
var field1 = {};
```

The global function `onICHostReady()` will be called by the front-end, when the
web component interface is ready. The version passed as parameter allows you to check that your
component code is compatible with the current gICAPI framework, and to define and assign the
`gICAPI.on*` callback methods (these will be defined in the body of the
`onICHostReady()` function:

```
var onICHostReady = function(version) {

    if ( version != "1.0" ) {
        alert('Invalid API version');
    }

    ... see below for gICAPI interface functions ...

}
```

At this point, the gICAPI interface is ready and the `gICAPI` object can be
used.

After the version checking, the `field1` variable is initialized to reference the
textarea element:

```
    field1 = document.getElementById("field1");
```

The next lines implement the event handler when the `textarea` element gets the
focus, to ask for a focus change with a `gICAPI.setFocus()`
call:

```
    field1.addEventListener('focus', (event) => {
        if (has_focus === false) {
            gICAPI.SetFocus();
        }
    });
```

The `gICAPI.onData()` function must be implemented to detect web component value
changes done in the program, and to acknowledge `SetData()` calls:

```
    gICAPI.onData = function(content) {
        field1.value = content;
    };
```

The `onFocus()` function is used to detect that the web component has got or lost
the focus. If the focus is gained, we need to explicitly set the focus to the expected web component
element:

```
    gICAPI.onFocus = function(polarity) {
        has_focus = polarity;
        if (has_focus) {
            field1.focus();
        }
    }
```

> **Note:**
>
> The only way to detect that the focus was gained by the web compoment field, is when
> `onFocus(true)` is called.

Implement the `gICAPI.onFlushData()` function, to provide
`textarea` content, when the front-end needs to send the field value to the runtime
system:

```
    gICAPI.onFlushData = function() {
        gICAPI.SetData( field1.value );
    };
```

Setup the web component when the field state changes by implementing the
`gICAPI.onStateChanged()`
function:

```
    gICAPI.onStateChanged = function(ps) {
        var params = JSON.parse(ps);
        if ( params.active ) {
            field1.disabled = false;
        } else {
            field1.disabled = true;
        }
    };
```

## Complete source code

File webcomp.per:

```
LAYOUT (TEXT="Simple web component")
GRID
{
Id:                       [id               ]
[wc                                         ]
[                                           ]
[                                           ]
[                                           ]
[                                           ]
[tx                                         ]
[                                           ]
[                                           ]
[                                           ]
}
END
END

ATTRIBUTES
EDIT id = FORMONLY.id;
WEBCOMPONENT wc = FORMONLY.wc,
  COMPONENTTYPE = "wc_simple",
  SCROLLBARS = NONE,
  STRETCH = BOTH;
TEXTEDIT tx = FORMONLY.info;
END
```

File webcomp.4gl:

```
MAIN
    DEFINE rec
        RECORD
            id INTEGER,
            wc STRING,
            info STRING
        END RECORD

    OPTIONS INPUT WRAP, FIELD ORDER FORM

    OPEN FORM f1 FROM "webcomp"
    DISPLAY FORM f1

    LET rec.id = 123
    LET rec.wc = "Hello, world!"

    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED, WITHOUT DEFAULTS)
        BEFORE FIELD wc
           LET rec.info = "BEFORE FIELD wc=\n", rec.wc
        ON CHANGE wc
           LET rec.info = "ON CHANGE wc =\n", rec.wc
        AFTER FIELD wc
           LET rec.info = "AFTER FIELD wc =\n", rec.wc
        ON ACTION show_value
           LET rec.info = "ON ACTION show wc=\n", rec.wc
        ON ACTION set_value
           LET rec.wc = "A new value ", CURRENT HOUR TO FRACTION
        ON ACTION disable
           CALL DIALOG.setFieldActive("wc", FALSE)
        ON ACTION enable
           CALL DIALOG.setFieldActive("wc", TRUE)
    END INPUT

END MAIN
```

File wc\_simple.html:

```
<!DOCTYPE html>
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content-type" content="text/html" charset="utf-8" />
<meta name='viewport' content='initial-scale=1.0, maximum-scale=1.0' />

<style>
html,body {
  height:100%;
  padding:0;
  margin:0;
  border:0;
  overflow:hidden;
}

textarea#field1 {
  font-weight: bold;
}

textarea {
  display: block;
  font-family: fixed;
  font-size: 10px;
  padding:0;
  margin:0;
  height: 99%;
  width: 99%;
}
</style>
</head>

<body>
<textarea id="field1"></textarea>
<script type="text/javascript" src="js/wc_simple.js"></script>
</body>

</html>
```

File wc\_simple.js:

```
var has_focus = false;
var field1 = {};

var onICHostReady = function(version) {

    if ( version != "1.0" ) {
        alert('Invalid API version');
    }

    field1 = document.getElementById("field1");

    field1.addEventListener('focus', (event) => {
        if (has_focus === false) {
            gICAPI.SetFocus();
        }
    });

    gICAPI.onData = function(content) {
        field1.value = content;
    };

    gICAPI.onFocus = function(polarity) {
        has_focus = polarity;
        if (has_focus === true) {
            field1.focus();
        }
    };

    gICAPI.onFlushData = function() {
        gICAPI.SetData( field1.value );
    };

    gICAPI.onStateChanged = function(ps) {
        var params = JSON.parse(ps);
        if ( params.active ) {
            field1.disabled = false;
        } else {
            field1.disabled = true;
        }
    };

}
```
