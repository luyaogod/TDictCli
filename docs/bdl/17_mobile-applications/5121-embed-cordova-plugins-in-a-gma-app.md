---
title: "Embed Cordova plugins in a GMA app"
source: "fgl-topics/c_fgl_cordova_build_gma.html"
breadcrumb: "Mobile applications > Cordova plugins > Embed Cordova plugins in a GMA app"
type: "concept"
---

# Embed Cordova plugins in a GMA app

> To be included in your app, Cordova plugins need to be specified in the build process.

## Including Cordova plugins in your Android™ APK package

Use the [gmabuildtool](5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.")
`--build-cordova` command-line option, to specify which Cordova plugins to embed in
your Android APK package:

```
gmabuildtool build
    --build-cordova cordova-plugin-contacts,cordova-plugin-device
    ...
```

## The plugin.xml file

A plugin can contain several assets (images, sounds, native forms) which are listed in the
plugin.xml file that is bundled by gmabuildtool.

## Shipping the Cordova wrapper library with the app package

If the Cordova plugin front calls are called from BDL wrapper functions, the
.42m p-code module of the wrapper library needs to be included in the APK
package.

When building the APK package for an app using Cordova plugins, the
gmabuildtool automatically includes the corresponding .42m
wrapper module(s).

## Related links

**Related concepts**  

[Deploying mobile apps on Android devices](5103-deploying-mobile-apps-on-android-devices.md "This section contains information to create a mobile application to be deployed on Android devices.")
