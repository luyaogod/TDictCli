---
title: "Genero Mobile for Android (GMA) 1.40 changes"
source: "fgl-topics/c_fgl_Migrate_to_320_gma_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Genero Mobile for Android™ (GMA) 1.40 changes"
type: "concept"
description: "Modifications to consider when using the Genero Mobile for Android."
---

# Genero Mobile for Android (GMA) 1.40 changes

> Modifications to consider when using the Genero Mobile for Android™.

> **Note:**
>
> This topic describes feature changes in the GMA 1.40 product. See also the Mobile section in
> [Genero BDL 3.20 New Features page](0062-bdl-3-20-new-features.md "Features added in 3.20 releases of the Genero Business Development Language.").

## GMA 1.40 with FGLGWS 3.20

> **Important:**
>
> The GMA version 1.40 is built on FGLGWS 3.20 and therefore, strongly tied to
> this Genero BDL version.

## Specifying the GBC for Universal Rendering

Genero 3.20 introduces Universal Rendering based on the GBC web front-end.

The GBC to be bundled with the embedded app is specified with the new
gmabuildtool option [`--build-gbc-runtime`](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.").

## `Button.alignment` style attribute

GMA 1.40.03 supports the `alignment` attribute for `BUTTON` form
items.

For the possible values of this attribute, see [Button style attributes](../11_user-interface/1631-button-style-attributes.md "Button presentation style attributes apply to BUTTON elements.").

## Cordova plugins from the FOURJS Cordova Github

Before Genero 3.20, the Cordova plugins wrapper libraries and demos were shipped in the FGLGWS
package (FGLDIR).

Cordova plugins, as well as demos and BDL wrapper libraries and demos, are now available from the
FOURJS Cordova Github.

For more details, see [Installing Cordova plugins](../17_mobile-applications/5119-installing-cordova-plugins.md "Before usage, Cordova plugins need to be installed in the GMA or GMI development environment.").

## New Cordova front calls `listPlugins` / `getPluginInfo`

Since Genero 3.20 / GMA 1.40, the new [`cordova.listPlugins`](../15_library-reference/3479-cordova-listplugins.md "Returns the list of available Cordova plugins.") and [`cordova.getPluginInfo`](../15_library-reference/3478-cordova-getplugininfo.md "Returns details about a specific Cordova plugin.")
front calls can be used to get information about Cordova plugins bundled with the app.

## 64-bit architecture in APK

Google requires that the APK uploaded to the Play Store must contain a 64-bit version of the app,
otherwise it will be rejected after August 1, 2019.

Starting with GMA 1.40.03, the gmabuildtool generates an APK with a 32-bit and
64-bit version of the app.

For more details, see [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.").

## Android Command Line Tools replacing SDK Tools

Starting with GMA 1.40.10, you need to install the [Android
Command Line Tools](https://developer.android.com/studio#cmdline-tools), to install and/or update the Android SDK. For more details about this
requirement, see [the latest Android SDK Tools release note](https://developer.android.com/studio/releases/sdk-tools).

If you have an older Android SDK installed (without the
$ANDROID\_SDK\_TOOLS/cmdline-tools directory), you can install the command line
tools in the existing Android SDK directory, by following the steps described in [Install Genero Mobile for Android](../04_installation/0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA.").

See also [the
prerequisites to build Android Apps with Genero Mobile](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Android App Bundle (.aab) packages

Starting with GMA 1.40.15, in addition to .apk packages, the
gmabuildtool produces now .aab packages.

For new applications, only Android App Bundle .aab archives can be uploaded
to the Android Play Store.

For more details, see [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.") and [gmabuildtool](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.").

## Android 11 (API 30) related changes

In August 2021, Android applications must target Android 11 / API 30, to be published on the
Google Play Store. GMA 1.40.15 is validated with Android 11 / API 30.

With this update come several security rules to apply:

- The `MESSAGE` and `ERROR` instructions are rendered with the new
  ["snackbar" visual](https://material.io/design/components/snackbars.html), and are no longer visible in the background. However, all presentation
  styles are still applicable.
- The sdcard cannot be accessed directly anymore: Program code must use the paths returned by
  front calls such as [`standard.feInfo/datadirectory`](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.") and [`standard.openFile`](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system."). This change
  impacts also front calls such as [`mobile.takePhoto`](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.") and [`mobile.takeVideo`](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier."), that will save the image in the external file directory
  of the application (and no longer in general share directory on sdcard). App code can no longer
  guess the path to sdcard and do some operation on it: One must either use the path returned
  `standard.feinfo/dataDirectory`, or use paths returned from front calls such as
  `standard.openFile`, to access to a file on the device. The paths returned by these
  front calls are now aligned with the scoped storage (see <https://source.android.com/devices/storage/scoped>)

See also [Handling files on Android devices](../17_mobile-applications/5090-handling-files-on-android-devices.md "How to manipulate file resources with GMA?").

## Support for `mobile.newContact` front call

Starting with GMA 1.40.19, the front call [`mobile.newContact`](../15_library-reference/3456-mobile-newcontact.md "Opens contact input form to create a new entry in the contact database.") is now available on GMA.

## New `--no-interactions` option

Starting with GMA 1.40.21, the `--no-interactions` option of gmabuildtool
updatesdk allows you to install extras silently, by answering yes to all questions asked
during the Android SDK installation
process. For more details, see [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.") and [gmabuildtool](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.").

## Controlling Firebase data collection for analytics

Starting with GMA 1.40.22, it is possible to enable Firebase data collection for analytics with
the new gmabuildtool option `-bfac`.

By default, Firebase data collection for analytics is disabled.

For more details, see [gmabuildtool](../17_mobile-applications/5106-gmabuildtool.md).
