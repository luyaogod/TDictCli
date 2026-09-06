---
title: "browser.getApplicationState"
source: "fgl-topics/c_fgl_frontcall_browser_getapplicationstate.html"
breadcrumb: "Library reference > Built-in front calls > Browser front calls > browser.getApplicationState"
type: "concept"
---

# browser.getApplicationState

> Gets the # anchor of the current URL in the browser address bar.

## Syntax

```
ui.Interface.frontCall("browser", "getApplicationState",
   [], [anchor])
```

1. anchor - The `#` anchor in the current URL.

## Usage

The "`getApplicationState`" front call allows the application to retrieve the
`#` anchor in the current URL displayed in the browser address bar.

This API must be used in conjonction with the
`applicationstatechanged` special action. For more details, read [Controlling web application state (#anchor)](../11_user-interface/2229-controlling-web-application-state-anchor.md "With GBC in a web browser, the context/state of a program can be managed with URL # anchors.").

## Related links

**Related concepts**  

[browser.setApplicationState](3433-browser-setapplicationstate.md "Sets the # anchor of the URL in the browser address bar.")
