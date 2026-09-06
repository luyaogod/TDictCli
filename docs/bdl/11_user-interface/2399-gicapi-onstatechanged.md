---
title: "gICAPI.onStateChanged()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_onstatechanged.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.onStateChanged()"
type: "concept"
---

# gICAPI.onStateChanged()

> The gICAPI.onStateChanged() function is executed when the state of the field changes.

## Purpose of gICAPI.onStateChanged()

The `gICAPI.onStateChanged()` function is called when the state of the
`WEBCOMPONENT` field changes. This occurs for example when the field is enabled or
disabled when starting/stopping a dialog instruction, or when using the
`DIALOG.setFieldActive()` method, or when the current dialog type changes (when a
`DISPLAY ARRAY` executes a nested `INPUT` for example).

## Implementing gICAPI.onStateChanged()

When the `gICAPI.onStateChanged()` function is fired, set up your
`WEBCOMPONENT` field accordingly, following the state parameters passed to the
function.

The `gICAPI.onStateChanged()` function gets a string representing a JSON structure
with all parameters set (like the `gICAPI.onProperty()` function). Convert this
string to a JSON object with `JSON.parse()`, then use the object properties as
follows:

1. `params.active`: The active state of the field (`0`: field is
   disabled, `1`: field is enabled)
2. `params.dialogType`: The current dialog type (`Input, DisplayArray,
   InputArray, Construct`)

## Example

This code example defines the `gICAPI.onStateChanged()` function to set up a
textarea element:

```
var onICHostReady = function(version) {

    ...

    gICAPI.onStateChanged = function(ps) {
        var field1 = document.getElementById("field1");
        var params = JSON.parse(ps);
        if ( params.active ) {
            field1.disabled = false;
        } else {
            field1.disabled = true;
        }
        ...
        if ( params.dialogType == 'Input' ) {
           ...
        }
        ...        
    };

    ...

};
```

## Related links

**Related concepts**  

[gICAPI.onFocus()](2397-gicapi-onfocus.md "The gICAPI.onFocus() function is used to detect if the WEBCOMPONENT field lost or acquired focus.")
