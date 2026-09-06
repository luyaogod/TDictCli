---
title: "cordova.call"
source: "fgl-topics/c_fgl_frontcall_cordova_call.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.call"
type: "concept"
---

# cordova.call

> Calls a function in a Cordova plugin and returns a result.

## Syntax

```
ui.Interface.frontCall("cordova", "call",
  [plugin-name, function-name [, param1, param2, ... ] ],
  [result]
)
```

1. plugin-name - This is the name of the Cordova plugin.
2. function-name - This is the name of plugin function to be called.
3. param1, param2, ... - Optional parameters to be passed to
   the Cordova function.
4. result - Holds the Cordova function return value.

## Usage

The `call` front call executes synchronously the Cordova plugin function
identified by the plugin-name and the function-name.

The other arguments (param1, param2, …) are arguments for
the Cordova plugin function. Each argument may have a different type like `FLOAT`,
`INTEGER`, `STRING`, `RECORD` or `DYNAMIC
ARRAY` (for `RECORD` and `DYNAMIC ARRAY`, the runtime system
will do the BDL to JSON conversion [automatically](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")).

The front call returns one result variable of type `FLOAT`,
`INTEGER`, `STRING`, `RECORD` or `DYNAMIC
ARRAY`, that matches the JSON equivalent of the plugin function result (for
`RECORD` and `DYNAMIC ARRAY`, the runtime system will do the JSON to
BDL conversion [automatically](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")).

The `call` Cordova front call is synchronous: This means that the call will not
return until the plugin returns a result. Some functions may not return a result at all (such as
start/stop functions) and therefore would cause the front call to wait forever. Those functions need
to be called asynchroneously with the [`callWithoutWaiting`](3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result.") front call.

In case of an error, the front call raises a runtime error -6333 that can be
caught with `TRY/CATCH` or `WHENEVER ERROR`.
> **Note:**
>
> Use the `err_get()` function, to identify the reason of the error. For more
> details about front call error handling, see [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.").

## Example

```
DEFINE calendars DYNAMIC ARRAY OF STRING
CALL ui.interface.frontcall( "cordova", "call",
     ["Calendar", "listCalendars"], [calendars] )
```
