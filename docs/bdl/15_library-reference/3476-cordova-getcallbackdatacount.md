---
title: "cordova.getCallbackDataCount"
source: "fgl-topics/c_fgl_frontcall_cordova_getCallbackDataCount.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.getCallbackDataCount"
type: "concept"
---

# cordova.getCallbackDataCount

> Returns the number of pending Cordova plugin results.

## Syntax

```
ui.Interface.frontCall("cordova", "getCallbackDataCount",
  [], [count])
```

1. count - Holds the number of pending results.

## Usage

When initiating an asynchronous Cordova plugin front call with [`callWithoutWaiting`](3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result."),
results are stored in the result queue when the Cordova function terminates, and a
`cordovacallback` action is fired if the current dialog defines a corresponding
`ON ACTION` handler.

The `getCallbackDataCount` front call returns the number of results currently in
the result queue, for all asynchronous Cordova plugin front calls initiated by a [`callWithoutWaiting`](3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result.").

It is then possible to implement a `FOR` loop to retrieve all results with the
[`getCallbackData`](3477-cordova-getcallbackdata.md "Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue.")
front call.

In case of an error, the front call raises a runtime error -6333 that can be
caught with `TRY/CATCH` or `WHENEVER ERROR`.
> **Note:**
>
> Use the `err_get()` function, to identify the reason of the error. For more
> details about front call error handling, see [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.").

## Example

```
DEFINE cnt SMALLINT
CALL ui.Interface.frontCall("cordova", "getCallbackDataCount",
        [], [cnt])
```
