---
title: "cordova.getAllCallbackData"
source: "fgl-topics/c_fgl_frontcall_cordova_getAllCallbackData.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.getAllCallbackData"
type: "concept"
---

# cordova.getAllCallbackData

> Returns all results for asynchronous Cordova plugin front calls, based on a callback ID filter.

## Syntax

```
ui.Interface.frontCall("cordova", "getAllCallbackData",
  [callback-id-filter], [results])
```

1. callback-id-filter - Provides the callback ID filter for Cordova
   asynchroneous front calls.
2. results - The array of results returned by the Cordova asynchroneous front
   call identified by callback-id.

## Usage

When initiating an asynchronous Cordova plugin front call with [`callWithoutWaiting`](3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result."),
results are stored in the result queue when the Cordova function terminates, and a
`cordovacallback` action is fired if the current dialog defines a corresponding
`ON ACTION` handler.

The purpose of the `getAllCallbackData` front call is to retrieve high traffic
data in bloc (such as motion sensor data) to avoid too many front calls in a short time frame.

The `getAllCallbackData` front call returns all results produced by Cordova
asynchronous front calls, that match the callback-id-filter provided as
parameter. For example, if you specify `"Media-"` as filter, all results related to
callback ids starting with `"Media-"` will be retrieved.

All results in the queue matching the filter are retrieved into the array, and are removed from
the result queue.

The results are returned into a `DYNAMIC ARRAY OF STRING` or a `DYNAMIC
ARRAY OF RECORD`, a structured record to be defined as the JSON equivalent of the plugin
function result (when using a `DYNAMIC ARRAY OF RECORD`, the runtime system can do
the JSON to BDL conversion [automatically](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")).

In case of an error, the front call raises a runtime error -6333 that can be
caught with `TRY/CATCH` or `WHENEVER ERROR`.
> **Note:**
>
> Use the `err_get()` function, to identify the reason of the error. For more
> details about front call error handling, see [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.").

## Example

```
DEFINE events DYNAMIC ARRAY OF RECORD
  x FLOAT,
  y FLOAT,
  z FLOAT,
  timestamp DECIMAL
END RECORD
DEFINE id STRING
DEFINE x INTEGER
...
CALL ui.interface.frontcall("cordova", "callWithoutWaiting",
        ["Accelerometer","start"], [id])
...
   ON IDLE 2
      CALL ui.interface.frontcall("cordova", "getAllCallbackData",
              [id], [events] )
      FOR x=1 TO events.getLength()
          CALL process_motion_event( events[x].* )
      END FOR
```

## Related links

**Related concepts**  

[cordova.getCallbackDataCount](3476-cordova-getcallbackdatacount.md "Returns the number of pending Cordova plugin results.")
