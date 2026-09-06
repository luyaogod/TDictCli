---
title: "Browser front calls"
source: "fgl-topics/c_fgl_frontcalls_browser.html"
breadcrumb: "Library reference > Built-in front calls > Browser front calls"
type: "concept"
---

# Browser front calls

> This section describes browser handling front calls.

This table shows the functions implemented in the "`browser`" module, to control
elements of the web browser, when using the GAS/GBC. The front calls related to web browser control
cannot be used in another front-end context such as GDC on desktop.

These front calls were made available starting with GBC 5.01.08

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("browser", "setApplicationState", [anchor], []) | Sets the `#` anchor of the URL in the browser address bar. |
| ui.Interface.frontCall("browser", "getApplicationState", [], [anchor]) | Gets the `#` anchor of the current URL in the browser address bar. |

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")

## Child topics

- [browser.setApplicationState](3433-browser-setapplicationstate.md): Sets the # anchor of the URL in the browser address bar.
- [browser.getApplicationState](3434-browser-getapplicationstate.md): Gets the # anchor of the current URL in the browser address bar.
