---
title: "gICAPI.onFocus()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_onfocus.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.onFocus()"
type: "concept"
---

# gICAPI.onFocus()

> The gICAPI.onFocus() function is used to detect if the WEBCOMPONENT field lost or acquired focus.

## Purpose of gICAPI.onFocus()

The `gICAPI.onFocus()` function is called when the runtime system gives/grabs the
focus to/from the `WEBCOMPONENT` field.

The `gICAPI.onFocus()` function is mandatory to implement a gICAPI web
component.

Focus gain or loss is defined by the polarity parameter passed to the
`gICAPI.onFocus()` function:

- When the runtime system gives the focus to the `WEBCOMPONENT` field, the
  polarity parameter is `true`.
- When the runtime system grabs the focus from the `WEBCOMPONENT` field, the
  polarity parameter is `false`.

When polarity is `true` and HTML elements of the web component
can get the focus, the `gICAPI.onFocus()` function must set the focus to the expected
element inside the web component.

This function can also be used to handle an internal flag, to know if the web component has got
the focus. If the web component has the focus, you can call the `gICAPI.SetData()`
function to provide the current value of the `WEBCOMPONENT` field.

## When is gICAPI.onFocus() called?

The `gICAPI.onFocus(polarity)` function is called with
`polarity=true` in the following cases:

- If the current dialog selects the `WEBCOMPONENT` field, as a normal tabbing
  candidate,
- If the user code executes an explicit `NEXT FIELD` instruction (or equivalent
  dialog instruction) to the web component,
- If the web component JavaScript code asked to get the focus with
  `gICAPI.SetFocus()`, and the program has accepted to give the focus to the
  `WEBCOMPONENT` field.

The `gICAPI.onFocus(polarity)` function is called with
`polarity=false` in the following cases:

- If the current dialog gives the focus to another field, as a normal tabbing candidate.
- If the user code executes an explicit `NEXT FIELD` instruction (or equivalent
  dialog instruction) to move to another field.

The `gICAPI.onFocus()` function will not be called with
`polarity=false`, if the web component JavaScript code asked the focus with
`gICAPI.SetFocus()`, and the program refused to give the focus.

## Example

The following code example defines the `gICAPI.onFocus()` function to set an
internal flag and give the focus to the textarea:

```
var has_focus = false;

var onICHostReady = function(version) {

    ...

    gICAPI.onFocus = function(polarity) {
        has_focus = polarity;
        if (has_focus) {
            var field1 = document.getElementById("field1");
            field1.focus();
        }
    };

    ...

};
```

## Related links

**Related concepts**  

[gICAPI.SetFocus()](2401-gicapi-setfocus.md "The gICAPI.SetFocus() function must be used to request the focus to the runtime system.")

[gICAPI.SetData()](2400-gicapi-setdata.md "The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program.")
