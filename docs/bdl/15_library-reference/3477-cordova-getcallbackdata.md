---
title: "cordova.getCallbackData"
source: "fgl-topics/c_fgl_frontcall_cordova_getCallbackData.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.getCallbackData"
type: "concept"
---

# cordova.getCallbackData

> Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue.

## Syntax

```
ui.Interface.frontCall("cordova", "getCallbackData",
  [], [result, callback-id])
```

1. result - Holds the result returned from the result queue.
2. callback-id - Holds the callback identifier of the Cordova asynchroneous
   front call.

## Usage

When initiating an asynchronous Cordova plugin front call with [`callWithoutWaiting`](3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result."),
results are stored in the result queue when the Cordova function terminates, and a
`cordovacallback` action is fired if the current dialog defines a corresponding
`ON ACTION` handler.

The `getCallbackData` front call returns the first Cordova plugin result from the
result queue, and removes it from the queue. A subsequent `getCallbackData` front
call gives back the next result and so on.

The first value returned by the front call (result) is the actual result. This
can be a variable of type `FLOAT`, `INTEGER`, `STRING`,
`RECORD` or `DYNAMIC ARRAY`, that matches the JSON equivalent of the
plugin function result (for `RECORD` and `DYNAMIC ARRAY`, the runtime
system will do the JSON to BDL conversion [automatically](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")).

The second returned value is callback-id, the identifier returned by the [`callWithoutWaiting`](3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result.")
front call that is causing this result. For example, for a Media plugin front call, the
callback-id can look like:

```
Media-messageChannel:0
```

If the result queue is empty, both result and callback-id
are `NULL`.

In case of an error, the front call raises a runtime error -6333 that can be
caught with `TRY/CATCH` or `WHENEVER ERROR`.
> **Note:**
>
> Use the `err_get()` function, to identify the reason of the error. For more
> details about front call error handling, see [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.").

## Example

```
DEFINE res, id STRING
CALL ui.Interface.frontCall("cordova", "getCallbackData",
        [], [res, id])
```

## Related links

**Related concepts**  

[cordova.getCallbackDataCount](3476-cordova-getcallbackdatacount.md "Returns the number of pending Cordova plugin results.")
