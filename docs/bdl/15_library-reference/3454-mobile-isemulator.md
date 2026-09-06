---
title: "mobile.isEmulator"
source: "fgl-topics/c_fgl_frontcall_mobile_isemulator.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.isEmulator"
type: "concept"
---

# mobile.isEmulator

> Indicates if the mobile app runs or displays forms on an emulator/simulator.

## Syntax

```
ui.Interface.frontCall("mobile", "isEmulator",
   [],  [result] )
```

1. result - Is set to `TRUE`, if the app is running in an device
   emulator/simulator, or when the front-end displaying application forms runs in an emulator.

## Usage

The "`isEmulator`" front call checks if the mobile app is running or is displayed
in a device emulator/simulator.

```
DEFINE emul BOOLEAN
CALL ui.Interface.frontCall("mobile", "isEmulator", []. [emul] )
```
