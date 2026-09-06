---
title: "Understanding Cordova plugins"
source: "fgl-topics/c_fgl_cordova_intro.html"
breadcrumb: "Mobile applications > Cordova plugins > Understanding Cordova plugins"
type: "concept"
---

# Understanding Cordova plugins

> Cordova plugins allow you to access specific mobile device functionalities.

## What are Cordova plugins in Genero?

Cordova plugins are based on the Apache Cordova cross-platform mobile development framework, to
access mobile device functionality such as the accelerometer, the camera, the compass, the
microphone, and more.

GMI and GMA act as a plugin container and provide a plugin API widely compatible with the
original Cordova API. The implementation is Genero-specific and makes the native plugin interfaces
available in Genero BDL through Cordova plugin front calls.

## Cordova plugins with device simulator or in client/server GUI mode

Depending on the configuration used when executing your app, some Cordova plugins or plugin APIs
may not be available and may cause a malfunction.

- When using a device simulator, some Cordova plugin will not work because the feature is not
  available on the simulator.
- When executing the app on a server in development mode or in "runOnServer" mode, the front calls
  are slowed down by the network roundtrips.

## Installing Cordova plugins into the development environment

In order to be used, Cordova plugins have to be downloaded from the [FOURJS Cordova
GitHub](https://github.com/FourjsGenero-Cordova-Plugins) or from the [Apache Cordova GitHub](https://cordova.apache.org/plugins/), and must be installed by using the
`--install-plugin` option of GMA or GMI build tools.

A Cordova plugin installation into GMA/GMI installation directories must be done by using a local
clone of the GitHub repository.

For more details, see [Installing Cordova plugins](5119-installing-cordova-plugins.md "Before usage, Cordova plugins need to be installed in the GMA or GMI development environment.").

## Building apps with Cordova plugins

Installed Cordova plugins need to be specified in the build process when creating your app.

The Cordova plugins to be bundled with the app are specified with the
`--cordova-plugin` option of GMA and GMI build tools:

- [Embed Cordova plugins in a GMA app](5121-embed-cordova-plugins-in-a-gma-app.md "To be included in your app, Cordova plugins need to be specified in the build process.")
- [Embed Cordova plugins in a GMI app](5122-embed-cordova-plugins-in-a-gmi-app.md "To be included in your app, Cordova plugins need to be specified in the build process.")

## Using Cordova plugin APIs (wrapper functions)

Access to the native code can be achieved using native Cordova APIs through [Cordova plugin front calls](../15_library-reference/3472-cordova-plugin-front-calls.md "Genero provides a set of Cordova plugin front calls that make use of the Cordova plugins.").

To ease the usage of a Cordova plugin, a library of BDL functions can encapsulate the Cordova
front calls. Most Cordova plugins provided by FOURJS are shipped with a BDL wrapper library,
available from the [FOURJS Cordova GitHub](https://github.com/FourjsGenero-Cordova-Plugins).
