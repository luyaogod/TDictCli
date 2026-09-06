---
title: "Inspecting the AUI tree of a front-end"
source: "fgl-topics/c_fgl_DynamicUI_AUI_debug.html"
breadcrumb: "User interface > User interface basics > The abstract user interface tree > Inspecting the AUI tree of a front-end"
type: "concept"
description: "AUI tree introspection When executing a program displaying on a front-end, it is possible to inspect the content of the abstract user interface built on the front-end side. The way to show the AUI ..."
---

# Inspecting the AUI tree of a front-end

## AUI tree introspection

When executing a program displaying on a front-end, it is possible to inspect the content of the
abstract user interface built on the front-end side. The way to show the AUI tree depends on the
type of front-end.

## Genero Desktop Client

The GDC must have been started in debug mode (`-aD` option).

In the current window of the running program, do a control-right-click with the mouse. This opens
the AUI Debug Tree window.

You can then browse the AUI debug tree created on the GDC side.

## GAS + Genero Browser Client

The GAS / GBC must have been started with debug option. In the
as.xcf configuration file, add the following
line:

```
<CONFIGURATION ...>
  <APPLICATION_SERVER>
     ...
     <RESOURCE Id="res.uaproxy.param" Source="INTERNAL">--development</RESOURCE>
     ...
```

Start your application in a web browser; a debug icon appears on the right of the
window. Click the icon to display the AUI debug tree. You can then browse the AUI debug tree created
on the GBC side.

Alternatively, it is also possible to specify the `debugMode=1` query string to
the URL starting the
program:

```
http://localhost:6394/ua/r/HelloWorld?debugMode=1
```

## Genero Mobile for Android™

The GMA must execute with debug mode enabled in the settings panel.

Open a web browser and enter the following
URL:

```
http://device-ip-address:6480
```

You can then browse the AUI debug tree created on the GMA side.

## Genero Mobile for iOS

The GMI must have been started in debug mode: the debug option needs to be enabled in
GMI app settings on the device.

Open a web browser and enter the following
URL:

```
http://device-ip-address:6480 (or 6400)
```

You can then browse the AUI debug tree created on the GMI side.
