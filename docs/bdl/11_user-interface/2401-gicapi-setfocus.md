---
title: "gICAPI.SetFocus()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_setfocus.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.SetFocus()"
type: "concept"
---

# gICAPI.SetFocus()

> The gICAPI.SetFocus() function must be used to request the focus to the runtime system.

## Purpose of gICAPI.SetFocus()

The tabbing order and focus management is controlled by the runtime system and the dialog code
(`NEXT FIELD`).

If your `WEBCOMPONENT` field receives an event that requests the focus, you must
first perform a `gICAPI.SetFocus()` call to determine if the program can put the
focus in the field.

## Focus acknowledgment

The `gICAPI.SetFocus()` function is used in conjunction with the
`gICAPI.onFocus()` callback function. If the program can give the focus to the field,
`gICAPI.onFocus()` is called with `true` as parameter. The
`gICAPI.onFocus()` is not called if one of the following conditions is true:

- Another field (the current field) cannot release the focus, because it does not satisfy
  constraints such as `INCLUDE`, or because the text in this field cannot be converted
  to the data type of the corresponding variable.
- The dialog code logic prevents focus change for example when a `NEXT FIELD
  CURRENT` is executed in the `AFTER FIELD` block of another field.

A good practice is to define an internal flag, to know if your `WEBCOMPONENT`
field has gained the focus.

## Example

The following code example uses the `gICAPI.SetFocus()` function to get the focus
from the runtime system:

```
var has_focus = false;
var field1 = {};

var onICHostReady = function(version) {

    ...

    field1 = document.getElementById("field1");

    field1.addEventListener('focus', (event) => {
        if (has_focus === false) {
            gICAPI.SetFocus();
        }
    });

    gICAPI.onFocus = function(polarity) {
        has_focus = polarity;
        if (has_focus === true) {
            field1.focus();
        }
    }

    ...

};
```

## Related links

**Related concepts**  

[gICAPI.onFocus()](2397-gicapi-onfocus.md "The gICAPI.onFocus() function is used to detect if the WEBCOMPONENT field lost or acquired focus.")
