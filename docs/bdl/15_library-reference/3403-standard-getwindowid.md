---
title: "standard.getWindowId"
source: "fgl-topics/c_fgl_frontcall_standard_getwindowid.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.getWindowId"
type: "concept"
---

# standard.getWindowId

> Returns the local window manager identifier of the window corresponding to the AUI window id passed as parameter.

> **Important:**
>
> This feature is desupported. It may still be
> available, but can be removed in a next version.

## Syntax

```
ui.Interface.frontCall("standard", "getWindowId",
 [aui-win-id], [loc-win-id])
```

1. aui-win-id - The id of the window node in the AUI tree.
2. loc-win-id - The id of the window in the window manager where the front-end
   is running.

## Usage

Returns the window identifier corresponding to the AUI window id passed as parameter, in the
window manager where the front-end is displaying the application forms.

The node id must reference a `Window` node, otherwise "0" is returned. In
traditional mode, window widgets are simple frames. Use "0" as `aui-win-id` parameter
to get the top level window id in the local window system.

This front call is provided for development and debugging purpose only.
