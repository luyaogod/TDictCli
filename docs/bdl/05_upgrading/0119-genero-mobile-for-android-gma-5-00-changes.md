---
title: "Genero Mobile for Android (GMA) 5.00 changes"
source: "fgl-topics/c_fgl_Migrate_to_500_gma_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > Genero Mobile for Android™ (GMA) 5.00 changes"
type: "concept"
description: "Modifications to consider when using the Genero Mobile for Android."
---

# Genero Mobile for Android (GMA) 5.00 changes

> Modifications to consider when using the Genero Mobile for Android™.

> **Note:**
>
> This topic describes feature changes in the GMA 5.00 product. See also the Mobile section in
> [Genero BDL 5.00 New Features page](0058-bdl-5-00-new-features.md "Features added in 5.00 releases of the Genero Business Development Language.").

## GMA 5.00 with FGLGWS 5.00

> **Important:**
>
> The GMA version 5.00 is built on FGLGWS 5.00 and therefore, strongly tied to this Genero BDL
> version.

## ANDROID\_HOME replaces ANDROID\_SDK\_ROOT

When building apps with gmabuildtool, the ANDROID\_SDK\_ROOT environment
variable is deprecated.

To define the Android SDK
installation directory, use the ANDROID\_HOME environment variable instead.

To learn more about the ANDROID\_HOME environment variable, go to the [Environment variables](https://developer.android.com/studio/command-line/variables.html) page on the Android Studio Developer site. To learn more about the
gmabuildtool, go to [gmabuildtool](../17_mobile-applications/5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.").

## Android 32bit desupport

Starting with GMA 5.00, the Android
32bit devices are no longer supported.

The GMA 5.00 front end, or an embedded app built with GMA 5.00 can only be deployed on 64bit
devices.

## Run `gmabuildtool updatesdk` on new ANDROID\_HOME directory

When updating the Android SDK
installation to the GMA 5.00.01 version, it is recommended to create a new ANDROID\_HOME directory to
avoid the warning about the "latest-2" directory that may be raised:

```
Warning: Package "Android SDK Command-line Tools (latest)" (cmdline-tools;latest) should be installed in
"C:\Android\android-sdk\cmdline-tools\latest" but it already exists.
Installing in "C:\Android\android-sdk\cmdline-tools\latest-2" instead.
```

While you can
ignore the warning, as the installation now takes place in the "latest"
directory, the recommendation is to either delete the latest-2 directory or
install in a new directory.
> **Important:**
>
> **Applies to version 5.00.01 only**
>
> In
> future versions of GMA, you can install updates in your current ANDROID\_HOME as usual.

To learn more about updating Android SDK, go to [Install Genero Mobile for Android](../04_installation/0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA.").

## Android 14 SDK (API 34) support

Starting with GMA 5.00.01, GMA is based on API 34 / Android 14 SDK.

## Specifying the Android foreground service type

Since Android 14 (API 34), a foreground Android service needs to specify a service type with the
`foregroundServiceType` manifest parameter.

Starting with GMA 5.00.02, the foreground service type can be controlled with the
gmabuildtool options `--custom-foreground-service-type` and
`--background-explanation`.

For more details, see [Defining the Android service type](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Changes in GMA 4.01 / BDL 4.01](0130-genero-mobile-for-android-gma-4-01-changes.md "Modifications to consider when using the Genero Mobile for Android.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
