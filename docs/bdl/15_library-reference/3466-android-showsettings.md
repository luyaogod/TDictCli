---
title: "android.showSettings"
source: "fgl-topics/c_fgl_frontcall_android_showsettings.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile Android™ front calls > android.showSettings"
type: "concept"
---

# android.showSettings

> Shows the GMA settings box controlling debug options.

## Syntax

```
ui.Interface.frontCall("android", "showSettings", [], [])
```

## Usage

This front call opens the settings box to enable or disable GMA programming options.

> **Important:**
>
> This front call is only available for an
> application running on an Android™ device.

No input parameters are required, and no parameters are returned.

The following features can be controlled with the GMA settings box:

- HTTP debug server on port 6480 (to inspect the [AUI tree](../11_user-interface/1515-inspecting-the-aui-tree-of-a-front-end.md) and show app logs)
- GUI display ([FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.")) and remote debug
  with [fgldb](../13_programming-tools/2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.") on
  port 6400
- Android logcat
  recording
- Managing allowed certificates (SSH connections)
- Cookies cleanup (for SSO authentication tokens)

## Related links

**Related concepts**  

[Debugging a mobile app](../17_mobile-applications/5101-debugging-a-mobile-app.md "Different solutions are available to debug a mobile app.")
