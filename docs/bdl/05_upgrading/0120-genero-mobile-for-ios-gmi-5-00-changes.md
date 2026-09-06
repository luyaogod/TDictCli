---
title: "Genero Mobile for iOS (GMI) 5.00 changes"
source: "fgl-topics/c_fgl_Migrate_to_500_gmi_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > Genero Mobile for iOS (GMI) 5.00 changes"
type: "concept"
---

# Genero Mobile for iOS (GMI) 5.00 changes

> Modifications to consider when using Genero Mobile for iOS.

> **Note:**
>
> This topic describes features changes in the GMI 5.00 product. See also the Mobile section in
> [Genero BDL 5.00 New Features page](0058-bdl-5-00-new-features.md "Features added in 5.00 releases of the Genero Business Development Language.").

## GMI 5.00 with FGLGWS 5.00

> **Important:**
>
> The GMI version 5.00 is built on FGLGWS 5.00 and therefore, strongly tied to
> this Genero BDL version.

## App Store requirements for privacy

Starting May 1, 2024, apps will be rejected by the App Store Connect, if they do not describe
their use of required reason API in their privacy manifest file.

The privacy manifest file is called PrivacyInfo.xcprivacy and needs to be
placed in the root of the application bundle.

Starting with GMI 5.00.01, the gmibuildtool bundles a default
PrivacyInfo.xcprivacy file which should be sufficient in most cases. This file
can be customized, if Apple rejects an app because informations are not suitable.

For more details, refer to [Customize ./gmi/PrivacyInfo.xcprivacy file](../17_mobile-applications/5109-building-ios-apps-with-genero.md).

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Changes in GMI 4.01 / BDL 4.01](0131-genero-mobile-for-ios-gmi-4-01-changes.md "Modifications to consider when using Genero Mobile for iOS.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
