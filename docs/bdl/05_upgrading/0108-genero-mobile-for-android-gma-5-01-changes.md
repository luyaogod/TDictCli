---
title: "Genero Mobile for Android (GMA) 5.01 changes"
source: "fgl-topics/c_fgl_Migrate_to_501_gma_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.01 upgrade guide > Genero Mobile for Android™ (GMA) 5.01 changes"
type: "concept"
description: "Consider these modifications when you use Genero Mobile for Android."
---

# Genero Mobile for Android (GMA) 5.01 changes

> Consider these modifications when you use Genero Mobile for Android™.

> **Note:**
>
> This topic describes feature changes in the GMA 5.01 product. See also the Mobile section in
> [Genero BDL 5.01 New Features page](0057-bdl-5-01-new-features.md "Features added in 5.01 releases of the Genero Business Development Language.").

## GMA 5.01 with FGLGWS 5.01

> **Important:**
>
> The GMA version 5.01 is built on FGLGWS 5.01 and therefore, strongly tied to this Genero BDL
> version.

## Large screen devices no longer support "allowedOrientations"

Starting with GMA 5.01.02, the window style attribute “allowedOrientations” no longer affects
large screen devices. Large screen devices are devices where the smallest width qualifier equals or
exceeds sw600dp; the sw600dp qualifier typically targets 7-inch tablets and larger.

For more details, go to [Screen orientation](../17_mobile-applications/5096-screen-orientation.md "Detecting screen orientation changes with the mobile device.").

To review the announcement from Android, go to [Apps targeting Android 16 - Device form factors](https://developer.android.com/about/versions/16/behavior-changes-16#large-screens-form-factors) (external
link).

## shortService service type apps restart when returned to the foreground

Starting with GMA 5.01.02, if the operating system kills a `shortService` service
type app while it is running in the background, the app will automatically relaunch when you bring
it back to the foreground. You can code the app to reload its state when relaunching.

For more information, see [Defining the Android service type](../17_mobile-applications/5105-building-android-apps-with-genero.md).

## Changes in earlier versions

Check the upgrade notes of earlier versions to avoid missing changes introduced in maintenance
releases. For more details, see [Changes in GMA
5.00 / BDL 5.00](0119-genero-mobile-for-android-gma-5-00-changes.md "Modifications to consider when using the Genero Mobile for Android.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
