---
title: "browser.setApplicationState"
source: "fgl-topics/c_fgl_frontcall_browser_setapplicationstate.html"
breadcrumb: "Library reference > Built-in front calls > Browser front calls > browser.setApplicationState"
type: "concept"
---

# browser.setApplicationState

> Sets the # anchor of the URL in the browser address bar.

## Syntax

```
ui.Interface.frontCall("browser", "setApplicationState",
   [anchor], [])
```

1. anchor - The `#` anchor to be set in the URL.

## Usage

The "`setApplicationState`" front call allows the application to set the
`#` anchor in the URL displayed in the browser address bar.

This API must be used in conjonction with the
`applicationstatechanged` special action. For more details, read [Controlling web application state (#anchor)](../11_user-interface/2229-controlling-web-application-state-anchor.md "With GBC in a web browser, the context/state of a program can be managed with URL # anchors.").

## Related links

**Related concepts**  

[browser.getApplicationState](3434-browser-getapplicationstate.md "Gets the # anchor of the current URL in the browser address bar.")
