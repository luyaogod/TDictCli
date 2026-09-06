---
title: "webcomponent.call"
source: "fgl-topics/c_fgl_frontcall_webcomponent_call.html"
breadcrumb: "Library reference > Built-in front calls > Web component front calls > webcomponent.call"
type: "concept"
---

# webcomponent.call

> Calls a JavaScript function through the web component.

## Syntax

```
ui.Interface.frontCall("webcomponent", "call",
  [aui-name, function-name
     [, param1, param2, ... ]
  ],
  [result]
)
```

1. aui-name - This is the name of the web component name in the AUI tree.
2. function-name - This is the name of a user-defined JavaScript function
   to be called.
3. param1, param2, ... - Optional parameters to be passed
   to the web component JavaScript function.
4. result - Holds the JavaScript function return value.

## Usage

This front call executes a JavaScript function of the `WEBCOMPONENT` form field
identified by aui-name.

Use the `webcomponent.call` front call only for specific needs:
With desktop and web front ends, a front call implies a network roundtrip and abstract user
interface update that may cause unwanted delays. Prefer using [gICAPI interface](../11_user-interface/2393-the-gicapi-web-component-interface-script.md "The gICAPI web components are controlled on the front-end through a gICAPI interface object, defined in a JavaScript script.") triggers such as
`onData` / `onProperty` instead.

The JavaScript function identified by function-name must be implemented in the
HTML content pointed by the URL-based web component, or in the user-defined JavaScript of a
gICAPI-based web component.

The function specified in the `webcomponent.call` front call must be a user-defined
function. Calling directly JavaScript API methods such as `window.alert()` is not
supported. To call the `window.alert()` method, implement a wrapper that calls the
window method.

The arguments following the function-name argument will be passed to the
JavaScript function.

The result variable will contain the value returned by the JavaScript
function.

When using simple data types in arguments, values are passed as is to the JavaScript function of
the web component. When using `RECORD` or `DYNAMIC ARRAY` types, the
runtime system converts the structured data to a JSON string that can be received in JavaScript as
an object. Similarly, if the JavaScript function returns a complex data structure in JSON notation,
it can directly by assigned to a `RECORD` or `DYNAMIC ARRAY` in the
return arguments of the frontcall.

For more details about JSON notation usage for complex data types, see the [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")
method.

## Example 1: Using a simple string parameter and return value:

```
DEFINE result STRING
CALL ui.Interface.frontCall("webcomponent","call",
        ["formonly.mywebcomp","echoString","abcdef"],
        [result])
```

## Example 2: Using structured and array variables:

```
DEFINE options RECORD
          filter STRING,
          creation DATE
       END RECORD,
       items DYNAMIC ARRAY OF STRING
LET options.filter = "abc*"
LET options.creation = TODAY
CALL ui.Interface.frontCall("webcomponent","call",
        ["formonly.mywebcomp","getItems",options],
        [items])
```

## Example 3: Using a JSON string as parameter:

```
IMPORT util
...
DEFINE options RECORD
          filter STRING,
          creation DATE
       END RECORD,
       js STRING,
       result INTEGER
LET options.filter = "abc*"
LET options.creation = TODAY
LET js = util.JSON.stringify( options )
CALL ui.Interface.frontCall("webcomponent","call",
        ["formonly.mywebcomp","setOptions",js]
        [result])
```

For a complete example, see [Example 1: Calling a JavaScript function](../11_user-interface/2411-example-1-calling-a-javascript-function.md).
