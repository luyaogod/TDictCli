---
title: "mobile.getGeolocation"
source: "fgl-topics/c_fgl_frontcall_mobile_getgeolocation.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.getGeolocation"
type: "concept"
---

# mobile.getGeolocation

> Returns the Global Positioning System (GPS) location of a mobile device.

## Syntax

```
ui.Interface.frontCall("mobile", "getGeolocation",
   [], [status, latitude, longitude] )
```

1. status - Holds the status of the front call execution.
2. latitude - Holds the current latitude.
3. longitude - Holds the current longitude.

## Usage

The "`mobile.getGeolocation`" front call is a synonym for [standard.getGeolocation](3401-standard-getgeolocation.md "Returns the Global Positioning System (GPS) location of a device.").
