---
title: "Controlling the gICAPI web component in programs"
source: "fgl-topics/c_fgl_webcomponent_gicapi_in_progs.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Controlling the gICAPI web component in programs"
type: "concept"
description: "Controlling the gICAPI-based web components with ON ACTION Once a WEBCOMPONENT field is defined in the form file with the COMPONENTTYPE attribute pointing to an HTML content file, it can be used as a ..."
---

# Controlling the gICAPI web component in programs

## Controlling the gICAPI-based web components with ON ACTION

Once a `WEBCOMPONENT` field is defined in the form file with the
`COMPONENTTYPE` attribute pointing to an HTML content file, it can be used as a
regular edit field in program dialogs. The data of the gICAPI web component is transmitted with the
field value, and usually needs to be serialized and deserialized (typically in JSON), when the data
is not a simple scalar value.

When the web component field value is changed in the program, the [`onData()`](2393-the-gicapi-web-component-interface-script.md) method of the `gICAPI` object is fired, and you can
parse the serialized string in your JavaScript.

In order to detect web component value changes in the program, you need to combine the
`gICAPI.SetData()` and `gICAPI.Action()` methods, to transmit the
value and to fire an action, that will be handled by an `ON ACTION` block.

The `ON CHANGE` trigger is not
executed automatically for gICAPI-based web components, just by using
`gICAPI.SetData()`.

This example serializes and deserializes a dynamic array using the JSON format:

```
IMPORT util
...
DEFINE mywc STRING
DEFINE data_array DYNAMIC ARRAY OF RECORD ...
...
...
INPUT BY NAME mywc, ...
      ATTRIBUTES(WITHOUT DEFAULTS, UNBUFFERED)
   ...
   ON ACTION set_wc_values -- Bound to form button
      LET mywc = util.JSON.stringify( data_array )

   ON ACTION wc_data_changed -- Triggered by gICAPI.Action()
      CALL util.JSON.parse( mywc, data_array )
...
```

> **Important:**
>
> All data will be transmitted through the abstract user interface protocol.
> Transmitting a lot of data will not be efficient and is likely to slow down your application.

## Controlling the gICAPI-based web components with properties

Use the [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.")
attribute in the form specification, to define the configuration of the field.

When a property of the web component is modified, the [`onProperty()`](2393-the-gicapi-web-component-interface-script.md) method of the `gICAPI` object in the JavaScript
will be invoked with the list of properties in JSON notation.

Note that the complete property set will be passed, even if a single property is modified.

## Controlling gICAPI-based web components with front calls

The web component can be manipulated with specific front calls. The web component-specific front
calls are provided in the "`webcomponent`" front call module.

Use the `webcomponent.call` front call only for specific needs:
With desktop and web front ends, a front call implies a network roundtrip and abstract user
interface update that may cause unwanted delays. Prefer using [gICAPI interface](2393-the-gicapi-web-component-interface-script.md "The gICAPI web components are controlled on the front-end through a gICAPI interface object, defined in a JavaScript script.") triggers such as
`onData` / `onProperty` instead.

The `webcomponent.call`
front call can be used for general purposes. It takes as parameters the name of the form field, a
JavaScript function to call, and optional parameters to be passed to the specified JavaScript
function. The JavaScript function must be implemented in the HTML content of the gICAPI web
component field. The front call returns the result of the JavaScript function.

```
DEFINE title STRING
CALL ui.Interface.frontCall("webcomponent", "call",
     ["formonly.mychart", "eval", "document.title"],
     [title] )
```

The `getTitle` function is another useful `webcomponent` front call
that can get the title of the HTML document of the web component:

```
DEFINE info STRING
CALL ui.Interface.frontCall("webcomponent", "getTitle",
     ["formonly.url_field"], [info] )
```

Some providers return key information in the title of the HTML document.

## Related links

**Related concepts**  

[Defining a gICAPI web component in forms](2407-defining-a-gicapi-web-component-in-forms.md "When defining a gICAPI web component in a form specification file, you can also provide a sizing policy and define additional properties.")

[webcomponent.getTitle](../15_library-reference/3425-webcomponent-gettitle.md "Returns the title of the HTML doc rendered by a web component.")
