---
title: "gICAPI.SetData()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_setdata.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.SetData()"
type: "concept"
---

# gICAPI.SetData()

> The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program.

## Purpose of gICAPI.SetData()

If the content of the `WEBCOMPONENT` field needs to be transmitted to the program,
use the `gICAPI.SetData()` function to register the data to be sent to the runtime
system.

> **Important:**
>
> The call to `gICAPI.SetData()` will be ignored, if the
> `WEBCOMPONENT` field is not the current field and has the focus set by a call to
> [`gICAPI.SetFocus()`](2401-gicapi-setfocus.md "The gICAPI.SetFocus() function must be used to request the focus to the runtime system.") in
> the appropriate conditions. Use the `gICAPI.onFocus()` callback function to detect if
> the `WEBCOMPONENT` field has the focus.

The data must be a string. It is typically serialized as a JSON string.

## When to use gICAPI.SetData()?

Best practice is to use the `gICAPI.onFlushData()` callback function, and call the
`gICAPI.SetData()` function in `onFlushData()` to provide the
`WEBCOMPONENT` field value. When the `WEBCOMPONENT` field loses the
focus, or when the `gICAPI.Action()` function is called, the gICAPI framework will
call the `gICAPI.onFlushData()` function implicitly.

The `gICAPI.SetData()` can also be called outside the context of
`gICAPI.onFlushData()`, or when this callback is not implemented, typically before
calling `gICAPI.Action()`. If no `gICAPI.onFlushData()` is
implemented, the value provided by the last `gICAPI.SetData()` call will be used.
However, in order to register the value with `gICAPI.SetData()`, the
`WEBCOMPONENT` field must have the focus. Keep in mind that setting the focus with
[`gICAPI.SetFocus()`](2401-gicapi-setfocus.md "The gICAPI.SetFocus() function must be used to request the focus to the runtime system.") may
be rejected by the runtime system.

## Handling NULL values

In order to send a `NULL` value, call the `gICAPI.SetData()`
function with JavaScript `null` as parameter:

```
if (value.length==0) {
   gICAPI.SetData(null);
} else {
   gICAPI.SetData(value);
}
```

## Example

The following code example registers data to be sent to the runtime system when the
`gICAPI.onFlushData()` callback function is invoked:

```
var onICHostReady = function(version) {

    ...

    gICAPI.onFlushData = function() {
        var field1 = document.getElementById("field1");
        gICAPI.SetData( field1.value );
    };

    ...

};
```

## Related links

**Related concepts**  

[gICAPI.onData()](2395-gicapi-ondata.md "The gICAPI.onData() function is executed when field data is sent by the program.")

[gICAPI.onFocus()](2397-gicapi-onfocus.md "The gICAPI.onFocus() function is used to detect if the WEBCOMPONENT field lost or acquired focus.")
