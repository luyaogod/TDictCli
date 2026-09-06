---
title: "Local storage front calls"
source: "fgl-topics/c_fgl_frontcalls_localstorage.html"
breadcrumb: "Library reference > Built-in front calls > Local storage front calls"
type: "concept"
---

# Local storage front calls

> This section describes front calls to store data on the front-end platform.

Key/Value pairs can be stored locally on the front-end side with the
`localStorage` front calls.

Local storage is supported by all front-ends.

The data is stored on the platform where the front-end executes, and is persistent across
application sessions.

This feature can for example be used to keep a trace of the authenticated users, if the
authentication mechanism is written in Genero.

The table shows the functions implemented by all front-ends in the
"`localStorage`" module.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("localStorage", "clear", [], []) | Removes all local storage key/value pairs. |
| ui.Interface.frontCall("localStorage", "getItem", [key], [value]) | Returns the current value of local storage key. |
| ui.Interface.frontCall("localStorage", "keys", [], [key-list] ) | Returns the list of defined local storage keys. |
| ui.Interface.frontCall("localStorage", "removeItem", [key], []) | Deletes the specified local storage key. |
| ui.Interface.frontCall("localStorage", "setItem", [key,value], []) | Sets a value for local storage key. |
