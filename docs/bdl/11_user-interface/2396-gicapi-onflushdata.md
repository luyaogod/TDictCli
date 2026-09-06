---
title: "gICAPI.onFlushData()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_onflushdata.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.onFlushData()"
type: "concept"
---

# gICAPI.onFlushData()

> The gICAPI.onFlushData() function is executed when the front-end must send the field value to the program.

## Purpose of gICAPI.onFlushData()

The `gICAPI.onFlushData()` function is called when the front-end must sync the
`WEBCOMPONENT` field content with the program.

This occurs when the `WEBCOMPONENT` field loses the focus, or when calling
`gICAPI.Action()`.

If the `gICAPI.onFlushData()` function is not implemented, the front-end will use
the value set from the last `gICAPI.SetData()` call.

## Sending values with gICAPI.onFlushData()

In the `gICAPI.onFlushData()` function, use `gICAPI.SetData()` to
provide the value to be send to the runtime system.

The `gICAPI.onFlushData()` function is called when the gICAPI framework requires a
field value synchronization, or after `gICAPI.Action()` is called, to let you provide
the value.

After a `gICAPI.onFlushData()`, the value is sent to program for validation. The
runtime system can accept or reject the field value change. In order to detect that the runtime has
accepted the value, the `gICAPI.onData()` function will be called, with the same
value as the value provided by `gICAPI.SetData()`. The web component then receives an
indication that the VM has accepted the value change. Note that `gICAPI.onData()` is
also fired when the web component value is changed by the program.

## Details about the behavior of gICAPI.onFlushData()

Consider the following facts when implementing `gICAPI.onFlushData():`

1. If the `WEBCOMPONENT` field is focused, executing an action will trigger
   `gICAPI.onFlushData()`
2. If the `WEBCOMPONENT` field is focused, and the front-end gives the focus to
   another element, `gICAPI.onFlushData()` will be executed.
3. Perform a `gICAPI.SetData()` call in `gICAPI.onFlushData()` when
   needed (for example, if the value of the `WEBCOMPONENT` field has changed and needs
   to be sent to the program)
4. Do not use other `gICAPI` functions such as `gICAPI.Action()` or
   `gICAPI.SetFocus()` in the `gICAPI.onFlushData()` function.
5. The code inside `gICAPI.onFlushData()` must be non-blocking and should execute
   rapidely.

## Example

The following code example defines the `gICAPI.onFlushData()` function to provide
the content on a textarea element:

```
var onICHostReady = function(version) {

    ...

    gICAPI.onFlushData = function() {
        var field1 = document.getElementById("field1");
        gICAPI.SetData( field1.value );
    }

    ...

};
```

## Related links

**Related concepts**  

[gICAPI.SetData()](2400-gicapi-setdata.md "The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program.")
