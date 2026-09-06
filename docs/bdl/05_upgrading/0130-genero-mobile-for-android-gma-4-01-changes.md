---
title: "Genero Mobile for Android (GMA) 4.01 changes"
source: "fgl-topics/c_fgl_Migrate_to_401_gma_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.01 upgrade guide > Genero Mobile for Android™ (GMA) 4.01 changes"
type: "concept"
description: "Modifications to consider when using the Genero Mobile for Android."
---

# Genero Mobile for Android (GMA) 4.01 changes

> Modifications to consider when using the Genero Mobile for Android™.

> **Note:**
>
> This topic describes feature changes in the GMA 4.01 product. See also the Mobile section in
> [Genero BDL 4.01 New Features page](0059-bdl-4-01-new-features.md "Features added in 4.01 releases of the Genero Business Development Language.").

## GMA 4.01 with FGLGWS 4.01

> **Important:**
>
> The GMA version 4.01 is built on FGLGWS 4.01 and therefore strongly tied to this Genero BDL
> version.

## Need for JDK version 17

JDK 17 is required to build Android apps. For the latest information regarding system requirements and
Java support, please refer to the Supported platforms and databases document, available
on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).

For more details, see [Install Genero Mobile for Android](../04_installation/0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA.").

## Cordova / Gradle compatibility

Starting with GMA 4.01.02, [Cordova plugins](../17_mobile-applications/5117-cordova-plugins.md "This section describes how to use Cordova plugins.") need to
be compatible with Gradle plugin version 7.2.1 and Gradle version 7.4.2. All Cordova plugin using
outdated Gradle features need to be reviewed.

## `standard.playSound` gets a wait parameter

Starting with GMA 4.01.02, the `standard.playSound` front call supports a second
input parameter to force the front call to return immediately (`FALSE`), or wait
until the sound has finished playing (`TRUE`).

For more details, see [standard.playSound](../15_library-reference/3411-standard-playsound.md "Plays the sound file passed as parameter on the front-end platform.").

## `standard.playSound` allows http/https URLs

Starting with GMA 4.01.02, the `standard.playSound` front call now supports sound
files provided by URIs that take the form of http or https and that exist on the Web or on a private
HTTP server. It is no longer required to set up the sound file as a resource on the GAS or
programmatically put the file on the client.

For more details, see [standard.playSound](../15_library-reference/3411-standard-playsound.md "Plays the sound file passed as parameter on the front-end platform.").

## gmabuildtool default keystore type is now PKCS12

When the gmabuildtool option `build-jarsigner-keystore` is not
used, or if the provided keystore file is not valid, gmabuildtool generates a
default keystore.

Starting with GMA 4.01.04, the default keystore file is generated with type PKCS12. Before
4.01.04, the default generated keystore was of type JKS.

For more details, see [Generate the keystore file to sign your app](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Using GBC theme to define app colors

Starting with GMA 4.01.04, the `--build-app-colors` option of
gmabuildtool is desupported. As replacement, use [GBC customization and
theming solutions](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Android 13 SDK (API 33) support

Starting with GMA 4.01.04, GMA is based on API 33 / Android 13 SDK.

## Permissions summary in gmabuildtool output

Starting with GMA 4.01.06, the `gmabuildtool` output shows a summary of the Android permissions used to build the
APK/AAB. These can be:

- Permissions embedded by default
- Permissions added by the [`--build-app-permissions/-ba`](../17_mobile-applications/5106-gmabuildtool.md) option
- Permissions added by Cordova plugins or external libs implemented by the user

## Changes to gmabuildtool command arguments

Starting with GMA 4.01.06, the gmabuildtool supports now the @argfile syntax
to specify options files, and some options are now deprecated and have been replaced by either new
or existing options (for more details, see [`gmabuildtool`](../17_mobile-applications/5106-gmabuildtool.md) command reference):

| Deprecated option | Replacement option | Description |
| --- | --- | --- |
| `--build-force-scaffold-update` /`-bfsu` | [`--build-project` / `-bp`](../17_mobile-applications/5106-gmabuildtool.md) | The option to replace the current scaffold can be achieved by cleaning it manually. Changes to the way the `--build-project` unzips the scaffold has enhancements for working with scaffold files. |
| `--build-quietly` /`-bq` | [`--build-project` / `-bp`](../17_mobile-applications/5106-gmabuildtool.md) | The `--build-project` option now extracts scaffold files in silent mode. |
| `--input-options` /`-i` | [`@argfile`](../17_mobile-applications/5106-gmabuildtool.md) | Instead of:--input-options my_optionfile.txtUse now:@my_optionfile.txt |
| `--build-app-name` /`-bn` | [`--build-app-package-name` / `-bpn`](../17_mobile-applications/5106-gmabuildtool.md) | The application name is now constructed from the package name specified with the `--build-app-package-name` option. |
| `--build-app-genero-program` /`-bgp` | [`--root-path` / `-rp`](../17_mobile-applications/5106-gmabuildtool.md) | The `--root-path` is a new option for specifying the path to the default app icons path and for the default directory for build temp files. Default app icons dir is now root-dir/icons, it was root-dir/gma. Go to [Directory structures required for the build](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more details. |
| `--build-app-genero-program-main` /`-bgpm` | [`--main-app-path` / `-map`](../17_mobile-applications/5106-gmabuildtool.md) | The `--main-app-path` is a new option for specifying the absolute path to the main module of the application.This option is mandatory.Go to [Directory structures required for the build](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more details. |
| `--verbose-fine``--verbose-finer``--verbose-finest` | [`-v` / `-vv` / `-vvv`](../17_mobile-applications/5106-gmabuildtool.md) | Use the short options to set the log level. |
| `--build-jarsigner-keystore` / `-bjks` | [`--build-signing-keystore` / `-bsks`](../17_mobile-applications/5106-gmabuildtool.md) | Option renaming. |
| `--build-jarsigner-alias` / `-bja` | [`--build-signing-alias` / `-bsa`](../17_mobile-applications/5106-gmabuildtool.md) | Option renaming, now mandatory when `-bsks` is used. |
| `--build-jarsigner-storepass` / `-bjs` | [`--build-signing-storepass` / `-bss`](../17_mobile-applications/5106-gmabuildtool.md) | Option renaming, now mandatory when `-bsks` is used. |
| `--build-jarsigner-keypass` / `-bjk` | [`--build-signing-keypass` / `-bsk`](../17_mobile-applications/5106-gmabuildtool.md) | Option renaming, now mandatory when `-bsks` is used. |

## Changes to gmabuildtool proxy options

Starting with GMA 4.01.06, the [`--proxy-host`](../17_mobile-applications/5106-gmabuildtool.md) and [`--proxy-port`](../17_mobile-applications/5106-gmabuildtool.md) options for setting host and
port now apply only to the gmabuildtool updatesdk command.

For more details, see [Update the Android SDK with the GMA buildtool](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Package name manifest attribute for user-extensions and cordova builds

Starting with GMA 4.01.06 (since AGP upgrade), for Java user-extensions and Cordova, the
`package` attribute in AndroidManifest.xml is now deprecated and
replaced by specifying the `namespace` attribute in the
build.gradle.template file located in
templates/src/extension/build.gradle.template.

Before GMA
4.01.06:

```
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="com.gma.extension">
```

Since GMA
4.01.06:

```
...
       release {
            minifyEnabled false
            proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
        }
    }   
    namespace 'com.gma.extension'
...
```

The directory path of the scaffold has also changed, it was
extension-example/gma/project and is now
extension-example/project/scaffold.

For more details, see [Custom GMA binary archive build](../14_extending-the-language/2698-custom-gma-binary-archive-build.md "If you are planning to build Genero Mobile for Android extensions for your GMA project, you need to do this using Android Studio. Follow this procedure to extend GMA.").

## Advertizing ID management

If an app package includes the `com.google.android.gms.permission.AD_ID`
permission, in order to upload the AAB package to the Google Play Console, it is mandatory to
declare that the app uses an advertising ID.

Before GMA version 4.01.06, an app package implicitly included the `AD_ID`
permission. Consequently, it was mandatory to declare that the app uses an advertising ID.

Starting with GMA version 4.01.06, the `AD_ID` permission is no longer implicitly
included. As result, the declaration of the advertising ID is no longer required.

However, if the previous app version uploaded to the Play Console had the `AD_ID`
permission, you must keep the declaration of the advertising ID in order to upload the new version,
even if this new version does not include the `AD_ID` permission. During the upload
procedure, an error will be shown, explaining to the user that the manifest does not include the
permission although it is declared as being used. This error can be ignored. After the new version
is uploaded successfully, the advertising ID declaration can be removed, since the last version of
the app no longer includes this permission.

## Push notifications using FCM API V1

Starting with GMA 4.01.07, the push notifications solution on GMA now uses Firebase Cloud
Messaging API V1.

If you have implemented a push notification service using FCM with the former "legacy" Cloud
Messaging API, you must review the server and app code.

The structure of the JSON returned by mobile.getRemoteNotifications front call hase changed, see
[Front calls changes](0128-front-calls-changes.md "Modifications to consider when using front calls.").

For more details, see [Firebase Cloud Messaging (FCM)](../17_mobile-applications/5113-firebase-cloud-messaging-fcm.md "Follow this procedure to implement push notification with FCM.").

## Device identification with `standard.feInfo` front call

Starting with GMA 4.01.08, the `standard.feInfo` property
`"deviceId"` returns now the `ANDROID_ID`, while the
`"iccid"` and `"imei"` properties are now desupported and will produce
a runtime error if used.

For more details, see [standard.feInfo](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.").

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Changes in GMA 4.00 / BDL 4.00](0138-genero-mobile-for-android-gma-4-00-changes.md "Modifications to consider when using the Genero Mobile for Android.").

Notable changes introduced in maintenance releases:

- The `--no-interactions` option of gmabuildtool updatesdk [New --no-interactions option](0159-genero-mobile-for-android-gma-1-40-changes.md), is available in GMA 4.01.01 and GMA 1.40.21.
