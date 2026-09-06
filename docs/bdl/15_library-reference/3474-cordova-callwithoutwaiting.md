---
title: "cordova.callWithoutWaiting"
source: "fgl-topics/c_fgl_frontcall_cordova_callWithoutWaiting.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.callWithoutWaiting"
type: "concept"
---

# cordova.callWithoutWaiting

> Calls a function asynchronously in a Cordova plugin, without waiting for a result.

## Syntax

```
ui.Interface.frontCall("cordova", "callWithoutWaiting",
  [plugin-name, function-name [, param1, param2, ... ] ],
  [callback-id])
```

1. plugin-name - This is the name of the Cordova plugin.
2. function-name - This is the name of plugin function to be called.
3. param1, param2, ... - Optional parameters to be passed to
   the Cordova function.
4. callback-id - Holds the callback identifier for this Cordova asynchroneous
   front call.

## Usage

The `callWithoutWaiting` front call executes asynchronously the Cordova plugin
function identified by the plugin-name and the
function-name.

The other arguments (param1, param2, …) are arguments for
the Cordova plugin function. Each argument may have a different type like `FLOAT`,
`INTEGER`, `STRING`, `RECORD` or `DYNAMIC
ARRAY` (for `RECORD` and `DYNAMIC ARRAY`, the runtime system
will do the BDL to JSON conversion [automatically](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")).

The `callWithoutWaiting` front call behaves the same as the [`call`](3473-cordova-call.md "Calls a function in a Cordova plugin and returns a result.")
front call, but does not wait for a result from the plugin. Instead it returns directly a unique
callback ID, to be able to identify the call later on. The program execution can continue, while the
plugin processes the result asynchronously.

When the plugin produces the result, it is stored internally by GMI/GMA in a result queue. Then
result data can be retrieved with [`getCallbackData`](3477-cordova-getcallbackdata.md "Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue.")/[`getAllCallbackData`](3475-cordova-getallcallbackdata.md "Returns all results for asynchronous Cordova plugin front calls, based on a callback ID filter.") front calls in conjunction with the
"`cordovacallback`" action.

When the current dialog contains an `ON ACTION` handler for the [`cordovacallback` predefined
action](../11_user-interface/2257-list-of-predefined-actions.md), this action is triggered by the front-end to notify the Genero program that there is
plugin data to fetch.

The `cordovacallback` action is triggered:

1. If the result queue is empty, and a result is added to the queue.
2. If a new dialog that contains this action is entered or re-entered, and the result queue is non
   empty.

> **Important:**
>
> Some plugin functions may return results repeatedly in a short time intervals
> with one and the same callback-id, examples are for delivering
> motion/audio/bluetooth data.

In case of an error, the front call raises a runtime error -6333 that can be
caught with `TRY/CATCH` or `WHENEVER ERROR`.
> **Note:**
>
> Use the `err_get()` function, to identify the reason of the error. For more
> details about front call error handling, see [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.").

## Example

```
DEFINE id, song STRING
CALL ui.interface.frontcall( "cordova", "callWithoutWaiting",
     ["Media", "play", song], [id] )
...
   ON ACTION cordovacallback ATTRIBUTE(DEFAULTVIEW=NO)
      -- Process results
   ...
```
