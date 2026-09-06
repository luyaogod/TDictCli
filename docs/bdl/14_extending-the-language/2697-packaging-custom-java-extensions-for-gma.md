---
title: "Packaging custom Java extensions for GMA"
source: "fgl-topics/c_fgl_JavaBridge_gma_packaging.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Executing Java code with GMA > Packaging custom Java extensions for GMA"
type: "concept"
description: "Custom Java extension must be integrated in the GMA to run on Android devices."
---

# Packaging custom Java extensions for GMA

> Custom Java extension must be integrated in the GMA to run on Android™ devices.

Genero Mobile apps for Android are
created from Genero Studio, or from the command-line with [gmabuildtool](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device."); you need to provide the custom GMA binary archive
containing your Java extensions to Genero Studio or to
gmabuildtool.

- Genero Studio finds the GMA binary archive from the GMADIR variable defined in the
  configuration settings.
- The gmabuildtool requires the Android Studio project directory used to build
  the custom GMA to be specified with the `--build-project` option.

JDK 17 is required to build Android apps. For the latest information regarding system requirements and
Java support, please refer to the Supported platforms and databases document, available
on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).

Along with the GMA binary archive, you must provide the .jar files of your
Java extensions, that will be used to compile Genero application code on the development machine, as
well as the .apk Android packages of GMA, to deploy the front-end part on the device for client/server
development (typically with user-defined front calls).

The original GMA binary archive is a zip file containing several .aar Android libraries. A customized GMA binary
archive contains the .aar files from the GMA core libraries, the Genero runtime
system core libraries, and custom .aar files built from your own Java
libraries. The custom .aar libs are created from [Android
Studio](http://developer.android.com/sdk/installing/studio.html). The minimum Android Studio version is 0.8.9.

To create a new GMA binary archive, the extension.jar file, and the
.apk packages, including your Java extensions, perform the steps described in
[Custom GMA binary archive build](2698-custom-gma-binary-archive-build.md "If you are planning to build Genero Mobile for Android extensions for your GMA project, you need to do this using Android Studio. Follow this procedure to extend GMA.").

After completing these steps:

- When compiling application code, Genero Studio can find your .jar
  libraries to resolve Java symbols.
- When deploying the front-end only for client/server development, Genero Studio will
  find the  .apk packages to be installed on the device.
- When building an Android app in Genero
  Studio, it will be created from the custom GMA binary archive that includes your Java
  extensions.
- When building an Android app with
  gmabuildtool, it can be created by specifying the custom GMA Android project directory with the
  `--build-project` option.

## Related links

**Related concepts**  

[Implement front call modules for GMA](2720-implement-front-call-modules-for-gma.md "Custom front call modules for the Android front-end are implemented by using the API for GMA front calls in Java.")

[Implement Java user extensions in GMA](2695-implement-java-user-extensions-in-gma.md "A GMA app can execute custom Java code.")

## Child topics

- [Custom GMA binary archive build](2698-custom-gma-binary-archive-build.md): If you are planning to build Genero Mobile for Android extensions for your GMA project, you need to do this using Android Studio. Follow this procedure to extend GMA.
