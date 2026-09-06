---
title: "Genero Mobile for Android (GMA) 1.30 changes"
source: "fgl-topics/c_fgl_Migrate_to_310_gma_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Genero Mobile for Android™ (GMA) 1.30 changes"
type: "concept"
description: "Modifications to consider when using the Genero Mobile for Android."
---

# Genero Mobile for Android (GMA) 1.30 changes

> Modifications to consider when using the Genero Mobile for Android™.

> **Note:**
>
> This topic describes feature changes in the GMA 1.30 product. See also the Mobile section in
> [Genero BDL 3.10 New Features page](0063-bdl-3-10-new-features.md "Features added in 3.10 releases of the Genero Business Development Language.").

## GMA 1.30 with FGLGWS 3.10

> **Important:**
>
> The GMA version 1.30 is built on FGLGWS 3.10 and therefore, strongly tied to this Genero BDL
> version.

## Desupport of GMA front-end application on Google Play

The GMA front-end app is no longer available for download on Google Play.

In order to install the GMA front-end on your device, use the gmabuiltool as
described in [Genero Mobile Development Client for Android](../17_mobile-applications/5084-genero-mobile-development-client-for-android.md "Set up a development environment to display app forms on an Android device.").

## ANDROID\_SDK\_ROOT replaces ANDROID\_HOME

When building apps with gmabuildtool, the ANDROID\_HOME environment variable is
deprecated.

To define the Android SDK installation directory, use ANDROID\_SDK\_ROOT instead.

If the [`--android-sdk` option](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.")
is not specified, the gmabuildtool will first use ANDROID\_SDK\_ROOT, then
ANDROID\_HOME as fallback.

Consider changing your environment settings to follow Android SDK specifications.

> **Important:**
>
> The ANDROID\_SDK\_ROOT environment variable has since been deprecated. When building Android apps with Genero, use the
> ANDROID\_HOME environment variable. For more information, go to [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.").

## Default directory for localized strings in GMA apps

The .42s localized string files for the default language can be provided in
the appdir/defaults directory.

Resource files such as .42s provided in
pwd are always loaded by the runtime
system. As GMA, pwd and
appdir are the same, it was not
possible to provide default strings files in
appdir. It was required to provide
localized string files for any language, in the corresponding
appdir/locale-code
directories.

For more details, see [Localized string files on mobile devices](../09_advanced-features/0909-loading-localized-strings-at-runtime.md) and [Deploying mobile apps](../17_mobile-applications/5102-deploying-mobile-apps.md "This section describes how to build and deploy mobile apps with Genero.").

## New --no-install-extras option

By default, when updating the Android
SDK with gmabuildtool updatesdk, the process installs also extra SDK modules.

The `--no-install-extras` option of gmabuildtool updatesdk can
be used to skip installation of extra SDK modules when not needed.

## Unique package for all architectures

The GMA bundle is now provided as a single package, supporting both ARM and x86 device
architectures.

The `--build-types` option of gmabuildtool build is no longer
available.

## Unique scaffolding archive

Before GMA 1.30, two GMA binary archives where provided:

- fjs-gma-\*-android-scaffolding.zip
- fjs-gma-\*-android-extension-project.zip

Starting with GMA 1.30, the scaffolding and extension project have been merged in a single
fjs-gma-\*-android-scaffolding.zip archive.

For more details, see [Executing Java code with GMA](../14_extending-the-language/2693-executing-java-code-with-gma.md).

## New `scaffold` command

The gmabuildtool scaffold command has been added to manage scaffold
archives.

In its initial version, the scaffold command provides the `--list-plugins` option,
to show available plugins, and the `--install-plugins` option, to install plugins in
the scaffold archice (for Cordova support).

## No more gma/temp directory

The Genero Mobile Android project
directory does not longer need the gma/temp directory to build an app.

Since GMA 1.30, the `--build-distribution` option of
gmabuildtool is no longer available.

## gmabuildtool options use `--build-app-genero-program` as base directory by default

Starting with GMA 1.30, the gmabuildtool options listed below use the
application program files directory (`--build-app-genero-program` option) as base
directory for their default values.

In previous versions, the default base directory was the current working directory. Note however,
that the default value for the `--build-app-genero-program` option is still the
current working directory.

- `-bp / --build-project`
- `-bih / --build-app-icon-hdpi`
- `-bim / --build-app-icon-mdpi`
- `-bixh / --build-app-icon-xhdpi`
- `-bixxh / --build-app-icon-xxhdpi`
- `-bsh / --build-status-icon-hdpi`
- `-bsm / --build-status-icon-mdpi`
- `-bsxh / --build-status-icon-xhdpi`

## GMA scaffolding archive usage (**--build-force-scaffold-update** option)

During the manual installation procedure (to build GMA apps from the command line without GST),
it was required to unzip the scaffolding zip archive in a dedicated directory
(gma-scaffold-project), which could be referenced with the
`--build-project` option of the gmabuildtool build command.

Starting with GMA 1.30, if not yet done, the APK build process will automatically unzip the
original GMA scaffold archive (gma-install-dir/artifacts)
into the `--build-project` directory. If needed, you can force a cleanup and update
these scaffold files with the `--build-force-scaffold-update` option.

## New **--build-quietly** option

The `--build-quietly` option of gmabuildtool build allows to
build the APK silently, by answering yes to all questions asked during the build process.

## `Image.alignment` style attribute

GMA 1.30.01 now supports the `alignment` style attribute for
`IMAGE` form items.

For the possible values of this attribute, see [Image style attributes](../11_user-interface/1641-image-style-attributes.md "Image style presentation attributes apply to an IMAGE element.").

## Dangerous permissions no longer set by default

Starting with GMA 1.30, the Android
Dangerous Permissions such as `android.permission.WRITE_EXTERNAL_STORAGE` are no
longer set by default when building an APK. Such permissions must be specified explicitly with the
`--build-app-permissions` option of gmabuildtool, depending on the
frontcalls used by the app.

For more details, see [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Push Notification: Firebase Cloud Messaging replaces Google Cloud Messaging

Starting with GMA 1.30.18, the Firebase Cloud Messaging framework replaces Google Cloud Messaging
to implement push notifications.

What you need to consider:

1. A new Firebase Cloud Messaging project must be created from the [FCM console](https://console.firebase.google.com).
2. The package name of the FCM app must match the package name used to build your APK (with the
   `--build-app-package-name` option of [`gmabuildtool`](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device."))
3. You must download the google-services.json configuration file an put it in
   the [appdir](../17_mobile-applications/5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).").
4. The sender-id parameter of the [`registerForRemoteNotifications`](../15_library-reference/3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications."), [`getRemoteNotifications`](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") and [`unregisterFromRemoteNotifications`](../15_library-reference/3462-mobile-unregisterfromremotenotifications.md "This front call unregisters the mobile device from push notifications.") frontcalls is no longer needed (all
   informations are in the google-services.json parameter file).
5. In the server application sending push notifications, replace the Google API key by the Server
   Key found in the Firebase Cloud Message project parameters.
6. Update your FCM server endpoint to send push requests: Replace the
   `https://gcm-http.googleapis.com/gcm/send` URL by
   `https://fcm.googleapis.com/fcm/send`.

For more details, see [Push notifications](../17_mobile-applications/5112-push-notifications.md "This section describes how to implement push notification with Genero.").

## Setting debug ports with adb

The debug ports 6400 and 6480 need to be forwarded with the adb command, the
telnet/redir solution is no longer supported.

For more details, see [Debugging on a mobile device](../13_programming-tools/2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.").

## Related links

**Related concepts**  

[Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.")
