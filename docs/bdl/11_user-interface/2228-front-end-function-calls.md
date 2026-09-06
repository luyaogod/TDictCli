---
title: "Front-End function calls"
source: "fgl-topics/c_fgl_prog_dialogs_front_calls.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Front-End function calls"
type: "concept"
---

# Front-End function calls

> The language allows to execute specific functions on the front-end platform.

A set of built-in front calls is available in Genero front-ends, to execute a specific task on
the platform where the front-end executes.

In order to perform a front call, use the [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.") method. For example, when using a mobile
front-end, you can instruct the mobile device to take a picture and return the identifier of the
asset containing the image:

```
DEFINE path STRING
CALL ui.Interface.frontCall( "mobile", "takePhoto", [], [path] )
```

For a complete list of available front calls, see [Built-in front calls](../15_library-reference/3385-built-in-front-calls.md "This section contains the description of all built-in front calls.").
