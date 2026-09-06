---
title: "gICAPI.onData()"
source: "fgl-topics/c_fgl_webcomponent_gicapi_ondata.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script > gICAPI.onData()"
type: "concept"
---

# gICAPI.onData()

> The gICAPI.onData() function is executed when field data is sent by the program.

## Purpose of gICAPI.onData()

The `gICAPI.onData()` function is called each time the
`WEBCOMPONENT` field content modification comes from the program. This occurs for
example when the current dialog sets the field value, or when a `DISPLAY value TO
wc_field` instruction is performed.

The `gICAPI.onData()` function is also used to check that the runtime system has
accepted the web component value change, after a call to the `gICAPI.SetData()`
function, when the value needs to be transmitted from the `WEBCOMPONENT` field to the
program.

## Handling gICAPI.onData() values

When the `gICAPI.onData()` function is fired, assign the data value to the web
component, or check that the runtime system has validated the value provided with
`gICAPI.SetData()`.

The data parameter is a string that contains the field value as provided by the program. It is up
to your JavaScript code to interpret the program value to be rendered on the HTML page. For example,
the value may just be GPS coordinates, that will be used to display a location on a map. The data is
typically serialized as a JSON string.

If the `WEBCOMPONENT` field value can be `NULL`, the
`onData()` function must check for null values as follows:

```
gICAPI.onData = function(value) {
    if (value == null || value.length == 0) {
        // Process null case.
        ...
    }
    ...
```

For structured data coming for example from a BDL `RECORD`, use the JSON notation
to pass values back-and-forth to/from the `WEBCOMPONENT` field. To serialize /
de-serialize structured data in the BDL code, use the [`util.JSON`](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") class. In the JavaScript code, use the
`JSON.parse()` and `JSON.stringify()` methods.

> **Important:**
>
> To avoid a [fatal error](2393-the-gicapi-web-component-interface-script.md) that would stop the program with a front-end error, protect the JS code parsing
> JSON strings.

For example:

```
gICAPI.onData = function(value) {
    try {
        var jo = JSON.parse(value);
        if (jo.current > 0) {
            index = jo.current - 1;
        }
        if (jo.selected) {
            selected = jo.selected;
        }
    } catch (err) {
        console.error("onData(): Invalid JSON string");
        index = -1;
        selected = [];
    }
};
```

## Example

The following code example defines the `onData()` function to set the content on a
textarea element:

```
var onICHostReady = function(version) {

    ...

    gICAPI.onData = function(content) {
        var field1 = document.getElementById("field1");
        field1.value = content;
    };

    ...

};
```

## Related links

**Related concepts**  

[gICAPI.SetData()](2400-gicapi-setdata.md "The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program.")
