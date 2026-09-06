---
title: "The gICAPI web component interface script"
source: "fgl-topics/c_fgl_webcomponent_gicapi_script.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > The gICAPI web component interface script"
type: "concept"
---

# The gICAPI web component interface script

> The gICAPI web components are controlled on the front-end through a gICAPI interface object, defined in a JavaScript script.

## gICAPI interface basics

The goal of the gICAPI interface is to manage communication between the program and the
web component with a basic API, to handle the interaction events, the focus, and the value
of the web component field.

The interface script is written in JavaScript and bound to the `WEBCOMPOMENT`
form field by using an [HTML document as
container](2392-html-document-and-javascript-for-the-gicapi-object.md "A gICAPI web component is identified by an HTML document containing the JavaScript interface (or a reference to the .js file).").

![gICAPI-based Web Component communication management diagram](../_images/WebCompo1.jpg)

*Web Component communication management*

The gICAPI web component API relies on a published global JavaScript object named
`gICAPI`.

## JavaScript exceptions in gICAPI code

If an exception is thrown in the JavaScript code of your gICAPI interface, the gICAPI framework
will detect it and raise a fatal error that will stop the application.

For ex, when parsing some malformed JSON string with `util.JSON`, in code that is
not protected with a try/catch block, this can will produce a front-end error like
following:

```
"Application ended" The application encountered a problem.
Webcomponent threw an error : "SyntaxError: Unexpected end of JSON input"
```

To avoid such fatal error, make sure that exceptions are caught in a try/catch
block:

```
try {
    // Code that can potentially throw execptions
    ...
}
catch(err)
    // Handle error
    ...
}
```

See [gICAPI.onData()](2395-gicapi-ondata.md "The gICAPI.onData() function is executed when field data is sent by the program.") for more details about JSON
usage.

## gICAPI initialization function

The `onICHostReady()` global function must be implemented, to execute code after
the HTML page has been loaded and the `gICAPI` interface has been initialized.

The `gICAPI` object is ready in the context of
`onICHostReady()`.
> **Important:**
>
> The `onICHostReady()`
> function is only called for the initial HTML page defined by the `WEBCOMPONENT` form
> item. If you implement JavaScript code that loads another HTML page with a set of gICAPI interface
> functions, these will be ignored. For example, using
> `window.location="./new-page.html"` will load another HTML page, making all current
> gICAPI `WEBCOMPONENT` interaction functions invalid.

| Method | Description |
| --- | --- |
| `onICHostReady( version String )` | Called when the gICAPI web component interface is ready. The version passed in the parameter allows you to check that your component is compatible with the API, and initialization code can be execute in this function. |

The programming interface of the gICAPI class is identified by a version number, to make sure
that the user code corresponds to the current gICAPI implementation. Verify that the runtime version
number matches the gICAPI version used during development, by checking the value passed as parameter
to `onICHostReady()`:

```
var onICHostReady = function(version) {
    if ( version != "1.0" ) {
        alert('Invalid API version');
    }
    ...
}
```

The rest of the `onICHostReady()` function body is used to do some initialization
and to assign the `gICAPI.on*()` callback functions as described later in this
topic.

## gICAPI management functions

The `gICAPI` object supports a set of callback functions (like
`onFlushData()`) and control functions (like `Action()`), to handle
field value changes, properties changes and focus requests.

> **Important:**
>
> The `gICAPI` object must be instantiated, before
> defining and assigning these methods. The `gICAPI` object is created and initialized
> by the web component framework before calling the `onICHostReady()` global function.
> Therefore, `on*` callback methods are typically defined and assigned to the
> `gICAPI` object, inside the body of the `onICHostReady()`
> function.

```
var onICHostReady = function(version) {

    if ( version != "1.0" ) {
        alert('Invalid API version');
    }

    current_color = "#000000";

    gICAPI.onProperty = function(properties) {
        ...
    }

    ...
```

The execution order of the `gICAPI.on*` functions is undefined. For example, the
`onData()` function may be fired before the `onProperty()` function
when the form is initialized. Consider writing code that takes this behavior into account.

| Method | Description |
| --- | --- |
| `Action( action String )` | Triggers an action event, which will execute the corresponding `ON ACTION` code.For more details, see [gICAPI.Action()](2394-gicapi-action.md "The gICAPI.Action() function is used to perform an action in the current dialog."). |
| `onData( data String )` | Called when the program value of the form field is updated by the runtime system.For more details, see [gICAPI.onData()](2395-gicapi-ondata.md "The gICAPI.onData() function is executed when field data is sent by the program."). |
| `onFlushData( )` | Called when the gICAPI framework needs to send a value to the runtime system.For more details, see [gICAPI.onFlushData()](2396-gicapi-onflushdata.md "The gICAPI.onFlushData() function is executed when the front-end must send the field value to the program."). |
| `onFocus( polarity Boolean )` | Called when the runtime system / program changes the focus.For more details, see [gICAPI.onFocus()](2397-gicapi-onfocus.md "The gICAPI.onFocus() function is used to detect if the WEBCOMPONENT field lost or acquired focus."). |
| `onProperty( properties String )` | Called to get `WEBCOMPONENT` [`PROPERTIES`](1810-properties-attribute.md "The PROPERTIES attribute is used to define a list of widget-specific characteristics.") attributes.For more details, see [gICAPI.onProperty()](2398-gicapi-onproperty.md "The gICAPI.onProperty() function is executed when web component properties change."). |
| `onStateChanged( params [] )` | Called when the field state changes (for example, when it gets active/inactive)For more details, see [gICAPI.onStateChanged()](2399-gicapi-onstatechanged.md "The gICAPI.onStateChanged() function is executed when the state of the field changes."). |
| `SetData( data String )` | Registers data to be sent to the program, in order to set the form field value in the program.For more details, see [gICAPI.SetData()](2400-gicapi-setdata.md "The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program."). |
| `SetFocus()` | Generates a focus request.For more details, see [gICAPI.SetFocus()](2401-gicapi-setfocus.md "The gICAPI.SetFocus() function must be used to request the focus to the runtime system."). |
| `UseGbcThemeVariables()` | Declares the list of GBC theme variables that can be used in the web component CSS.For more details, see [gICAPI.UseGbcThemeVariables()](2402-gicapi-usegbcthemevariables.md "The gICAPI.UseGbcThemeVariables() function can be used to declare a set of GBC theme variables to be referenced in the CSS of a web component."). |

## Related links

**Related concepts**  

[Web component front calls](../15_library-reference/3422-web-component-front-calls.md "This section describes web component specific front calls.")

## Child topics

- [gICAPI.Action()](2394-gicapi-action.md): The gICAPI.Action() function is used to perform an action in the current dialog.
- [gICAPI.onData()](2395-gicapi-ondata.md): The gICAPI.onData() function is executed when field data is sent by the program.
- [gICAPI.onFlushData()](2396-gicapi-onflushdata.md): The gICAPI.onFlushData() function is executed when the front-end must send the field value to the program.
- [gICAPI.onFocus()](2397-gicapi-onfocus.md): The gICAPI.onFocus() function is used to detect if the WEBCOMPONENT field lost or acquired focus.
- [gICAPI.onProperty()](2398-gicapi-onproperty.md): The gICAPI.onProperty() function is executed when web component properties change.
- [gICAPI.onStateChanged()](2399-gicapi-onstatechanged.md): The gICAPI.onStateChanged() function is executed when the state of the field changes.
- [gICAPI.SetData()](2400-gicapi-setdata.md): The gICAPI.SetData() function registers WEBCOMPONENT field data to be sent to the program.
- [gICAPI.SetFocus()](2401-gicapi-setfocus.md): The gICAPI.SetFocus() function must be used to request the focus to the runtime system.
- [gICAPI.UseGbcThemeVariables()](2402-gicapi-usegbcthemevariables.md): The gICAPI.UseGbcThemeVariables() function can be used to declare a set of GBC theme variables to be referenced in the CSS of a web component.
