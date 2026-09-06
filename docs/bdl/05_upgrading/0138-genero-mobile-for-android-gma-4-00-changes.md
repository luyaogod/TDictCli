---
title: "Genero Mobile for Android (GMA) 4.00 changes"
source: "fgl-topics/c_fgl_Migrate_to_400_gma_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Genero Mobile for Android™ (GMA) 4.00 changes"
type: "concept"
description: "Modifications to consider when using the Genero Mobile for Android."
---

# Genero Mobile for Android (GMA) 4.00 changes

> Modifications to consider when using the Genero Mobile for Android™.

> **Note:**
>
> This topic describes feature changes in the GMA 4.00 product. See also the Mobile section in
> [Genero BDL 4.00 New Features page](0060-bdl-4-00-new-features.md "Features added in 4.00 releases of the Genero Business Development Language.").

## GMA 4.00 with FGLGWS 4.00

> **Important:**
>
> The GMA version 4.00 is built on FGLGWS 4.00 and therefore, strongly tied to
> this Genero BDL version.

## Floating Action Button (FAB) desupport

Native Android Floating Action Button
(FAB) is no longer supported with [4.00
Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). The `materialFABActionList` and
`materialFABType` style attributes are desupported: See [Presentation styles changes](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles.").

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Changes in GMA 1.40 / BDL 3.20](0159-genero-mobile-for-android-gma-1-40-changes.md "Modifications to consider when using the Genero Mobile for Android.").

Notable changes introduced in maintenance releases:

- Requirement of [64-bit architecture in APK](0159-genero-mobile-for-android-gma-1-40-changes.md), also available in GMA 4.00.01.
- Using [Android command line tools](0159-genero-mobile-for-android-gma-1-40-changes.md), also
  available since GMA 4.00.01.
- Building [Android App Bundles
  (.aab)](0159-genero-mobile-for-android-gma-1-40-changes.md), also available since GMA 4.00.02.
- Android 11 (API 30) [Android 11 (API 30) related changes](0159-genero-mobile-for-android-gma-1-40-changes.md), also available since GMA 4.00.02.
