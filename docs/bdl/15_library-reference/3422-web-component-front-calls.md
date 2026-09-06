---
title: "Web component front calls"
source: "fgl-topics/c_fgl_frontcalls_webcomponent.html"
breadcrumb: "Library reference > Built-in front calls > Web component front calls"
type: "concept"
---

# Web component front calls

> This section describes web component specific front calls.

This table shows the functions provided by the "`webcomponent`" module.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("webcomponent", "call", [aui-name, function-name [, param1, param2, ... ] ], [result] ) | Calls a JavaScript function through the web component. |
| ui.Interface.frontCall("webcomponent", "frontCallAPIVersion", [],[result]) | Returns the API version of web component front-end calls. |
| ui.Interface.frontCall("webcomponent", "getTitle", [aui-name], [result] ) | Returns the title of the HTML doc rendered by a web component. |

## Related links

**Related concepts**  

[Web components](../11_user-interface/2378-web-components.md "This section describes how to use web components in your application.")
