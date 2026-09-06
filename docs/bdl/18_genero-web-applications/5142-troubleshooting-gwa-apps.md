---
title: "Troubleshooting GWA apps"
source: "fgl-topics/c_fgl_gwa_trblshoot.html"
breadcrumb: "Genero Web applications > Troubleshooting GWA apps"
type: "concept"
---

# Troubleshooting GWA apps

> What steps can you take if you have trouble with a Genero web application (GWA)?

If a GWA is stuck during the load process, check the browser debug console for errors.

The error is most likely visible in the console, either from the JavaScript side or from the
fglrun side. fglrun VM runtime errors appear as alerts in the GBC UI like in any other Genero app.

If the error is not an fglrun runtime error:

1. First, try reloading the app with the browser reload option with:
   - "Force Reload this page"(Chrome).
   - "Reload page from Origin" (Safari®).
   - Ctrl-F5 (Firefox™).
2. Next, clear the browser cache and then reload.
3. Next, clear website data for the development host and reload.

If all that is unsuccessful, send a copy of the web development console log to support.

## Related links

**Related concepts**  

[Debugging GWA apps](5143-debugging-gwa-apps.md "Different solutions are available to debug a GWA application.")
