---
title: "Genero Mobile Android front calls"
source: "fgl-topics/c_fgl_frontcalls_android.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile Android™ front calls"
type: "concept"
description: "This section describes front calls specific to the Android platform."
---

# Genero Mobile Android front calls

> This section describes front calls specific to the Android™ platform.

This table shows the functions implemented by the Android front-end in the "`android`" module.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("android","askForPermission", [permission], [result]) | Ask the user to enable a dangerous feature on the Android device. |
| ui.Interface.frontCall("android", "showAbout", [],[]) | Shows the GMA about box displaying version information. |
| ui.Interface.frontCall("android", "showSettings", [], []) | Shows the GMA settings box controlling debug options. |
| ui.Interface.frontCall("android","startActivity", [action, data, category, type, component, extras], []) | Starts an external Android application (activity), and returns to the GMA application immediately. |
| ui.Interface.frontCall("android", "startActivityForResult", [action, data, category, type, component, extras], [outdata, outextras]) | Starts an external application (Android activity) and waits until the activity is closed. |

## Related links

**Related concepts**  

[Genero Mobile common front calls](3441-genero-mobile-common-front-calls.md "This section describes common front calls provided by all mobile front-ends.")
