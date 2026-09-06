---
title: "gICAPI.Action()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_action.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.Action()"
type: "concept"
---

# gICAPI.Action()

> The gICAPI.Action() function is used to perform an action in the current dialog.

## Purpose of gICAPI.Action()

Use the `gICAPI.Action("action-name")` function in order to
execute an action in the context of the current dialog.

This function takes the name of the action as parameter. The corresponding `ON
ACTION` handler will be called.

Make sure that the parameter passed to `gICAPI.action()` matches the action name
in lowercase letters.

If the named action is not available (if it does not exist, or if it is disabled), the
`gICAPI.Action()` function has no effect.

After calling `gICAPI.Action("action-name")`, the gICAPI
framework will perform a `gICAPI.onFlushData()` callback (if defined), to let you
provide the web component field value with `gICAPI.SetData()`.

However, if you do not use the `gICAPI.onFlushData()` callback, it is also
possible to provide the value with `gICAPI.SetData()`, just before calling with
`gICAPI.Action()`.

## Example 1: Using gICAPI.SetData() directly before gICAPI.Action()

The following code executes the "color\_selected" action, after setting the value of the
`WEBCOMPONENT` field with `gICAPI.SetData()`:

```
var selectColor = function(c) {
    gICAPI.SetData(c);
    gICAPI.Action("color_selected");
}

// In the HTML code, an element defines the onclick handler as follows:
 ... onclick="selectColor('#FFFFCC')" ...
```

## Example 2: Using gICAPI.Action() and gICAPI.onFlushData()

The following code executes the "color\_selected" action, and sets the value of the
`WEBCOMPONENT` field with `gICAPI.SetData()` in the
`gICAPI.onFlushData()` callback:

```
var color = null;

var selectColor = function(c) {
    color = c;
    gICAPI.Action("color_selected");
}

gICAPI.onFlushData = function() {
    gICAPI.SetData(color);
}

// In the HTML code, an element defines the onclick handler as follows:
 ... onclick="selectColor('#FFFFCC')" ...
```

## Action handler in program code

In the program code, define an [`ON
ACTION` block](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.") in a dialog instruction, to execute code when the corresponding web
component action is fired.

`WEBCOMPONENT` actions are not known by the front-end. A default action view will
be created for each `ON ACTION` handler. Consider using the [`DEFAULTVIEW=NO` action
attribute](2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action."), in order to avoid default action views being displayed for your web component
actions.

```
  ON ACTION color_selected ATTRIBUTES( DEFAULTVIEW=NO )
     IF rec.webcomp == "#000000" THEN
        LET rec.webcomp = rec.pgcolor
        LET rec.info = NULL
        ERROR "Black color is denied!"
     ELSE
        LET rec.pgcolor = rec.webcomp
        LET rec.info = "Color selected:", rec.pgcolor
     END IF
```

## Related links

**Related concepts**  

[gICAPI.SetData()](2400-gicapi-setdata.md "The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program.")
