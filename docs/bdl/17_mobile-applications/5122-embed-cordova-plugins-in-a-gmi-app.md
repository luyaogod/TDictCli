---
title: "Embed Cordova plugins in a GMI app"
source: "fgl-topics/c_fgl_cordova_build_gmi.html"
breadcrumb: "Mobile applications > Cordova plugins > Embed Cordova plugins in a GMI app"
type: "concept"
---

# Embed Cordova plugins in a GMI app

> To be included in your app, Cordova plugins need to be specified in the build process.

## Including Cordova plugins in your iOS IPA package

Use the [gmibuildtool](5110-gmibuildtool.md "The gmibuildtool is a utility to create and test applications for an iOS devices.")
`--build-cordova` command-line option, to specify which Cordova plugins to embed in
your iOS IPA
package:

```
$ gmibuildtool --build-cordova=cordova-plugin-media,GeneroTestPlugin
```

## The plugin.xml file

A plugin can contain several assets (images, sounds, native forms) which are listed in the
plugin.xml file that is bundled by gmibuildtool.

## Shipping the Cordova wrapper library with the app package

If the Cordova plugin front calls are called from BDL wrapper functions, the
.42m p-code module of the wrapper library needs to be included in the IPA
package.

When building the API package for an app using Cordova plugins, the
gmibuildtool automatically includes the corresponding .42m
wrapper module(s).

## Related links

**Related concepts**  

[Deploying mobile apps on iOS devices](5107-deploying-mobile-apps-on-ios-devices.md "This section contains information to create a mobile application to be deployed on iOS devices.")
