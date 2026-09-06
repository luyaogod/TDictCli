---
title: "Cordova plugin front calls"
source: "fgl-topics/c_fgl_frontcalls_cordova.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls"
type: "concept"
---

# Cordova plugin front calls

> Genero provides a set of Cordova plugin front calls that make use of the Cordova plugins.

This table shows the functions implemented by the Android™ and iOS front-end in the "`cordova`" module.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("cordova", "call", [plugin-name, function-name [, param1, param2, ... ] ], [result] ) | Calls a function in a Cordova plugin and returns a result. |
| ui.Interface.frontCall("cordova", "callWithoutWaiting", [plugin-name, function-name [, param1, param2, ... ] ], [callback-id]) | Calls a function asynchronously in a Cordova plugin, without waiting for a result. |
| ui.Interface.frontCall("cordova", "getAllCallbackData", [callback-id-filter], [results]) | Returns all results for asynchronous Cordova plugin front calls, based on a callback ID filter. |
| ui.Interface.frontCall("cordova", "getCallbackDataCount", [], [count]) | Returns the number of pending Cordova plugin results. |
| ui.Interface.frontCall("cordova", "getCallbackData", [], [result, callback-id]) | Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue. |
| ui.Interface.frontCall("cordova", "getPluginInfo", [plugin-name], [result] ) | Returns details about a specific Cordova plugin. |
| ui.Interface.frontCall("cordova", "listPlugins", [ ], [plugins] ) | Returns the list of available Cordova plugins. |

## Related links

**Related concepts**  

[Cordova plugins](../17_mobile-applications/5117-cordova-plugins.md "This section describes how to use Cordova plugins.")

## Child topics

- [cordova.call](3473-cordova-call.md): Calls a function in a Cordova plugin and returns a result.
- [cordova.callWithoutWaiting](3474-cordova-callwithoutwaiting.md): Calls a function asynchronously in a Cordova plugin, without waiting for a result.
- [cordova.getAllCallbackData](3475-cordova-getallcallbackdata.md): Returns all results for asynchronous Cordova plugin front calls, based on a callback ID filter.
- [cordova.getCallbackDataCount](3476-cordova-getcallbackdatacount.md): Returns the number of pending Cordova plugin results.
- [cordova.getCallbackData](3477-cordova-getcallbackdata.md): Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue.
- [cordova.getPluginInfo](3478-cordova-getplugininfo.md): Returns details about a specific Cordova plugin.
- [cordova.listPlugins](3479-cordova-listplugins.md): Returns the list of available Cordova plugins.
