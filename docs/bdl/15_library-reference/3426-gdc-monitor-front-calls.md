---
title: "GDC Monitor Front Calls"
source: "fgl-topics/c_fgl_frontcalls_monitor.html"
breadcrumb: "Library reference > Built-in front calls > GDC Monitor Front Calls"
type: "concept"
---

# GDC Monitor Front Calls

> This section describes front calls specific to the GDC monitor.

The GDC monitor is the administration component of the Genero Desktop Front-End. This component
allows you to configure the GDC and do setup tasks.

The following table shows the functions implemented by all front-ends in the
"`monitor`" module.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("monitor", "update", [ path-to-update-file [,warning-text [,elevation-prompt] ] ], [ result ]) | Starts the GDC update. |
