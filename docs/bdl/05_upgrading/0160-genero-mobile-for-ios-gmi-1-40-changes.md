---
title: "Genero Mobile for iOS (GMI) 1.40 changes"
source: "fgl-topics/c_fgl_Migrate_to_320_gmi_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Genero Mobile for iOS (GMI) 1.40 changes"
type: "concept"
---

# Genero Mobile for iOS (GMI) 1.40 changes

> Modifications to consider when using Genero Mobile for iOS.

> **Note:**
>
> This topic describes features changes in the GMI 1.40 product. See also the Mobile section in
> [Genero BDL 3.20 New Features page](0062-bdl-3-20-new-features.md "Features added in 3.20 releases of the Genero Business Development Language.").

## GMI 1.40 with FGLGWS 3.20

> **Important:**
>
> The GMI version 1.40 is built on FGLGWS 3.20 and therefore, strongly tied to
> this Genero BDL version.

## Specifying the GBC for Universal Rendering

Genero 3.20 introduces Universal Rendering based on the GBC web front-end.

The GBC to be bundled with the embedded app is specified with the new
gmibuildtool option [`--gbc`](../17_mobile-applications/5110-gmibuildtool.md "The gmibuildtool is a utility to create and test applications for an iOS devices.").

## Minimum required iOS and Xcode® versions

Starting with GMI 1.40, the minimum version of the mobile OS must be iOS 11.

With GMI 1.40.02, Xcode 10 is required to
build apps.

With GMI 1.40.05, Xcode 11 is required to
build apps (for iOS 11, 12 and 13).

## iOS 13 and HTTPS X509 certificates

Since iOS 13, if you need to perform HTTPS connections with GWS or web components, the server
X509 certificates have to follow apple rules defined here: <https://support.apple.com/en-us/HT210176>.

## iOS 13 dark mode and presentation styles

iOS 13 introduced the dark mode. By default, GMI 1.40.05 supports both light and dark modes with
the correct system colors. For example, the rendering of `RED`,
`GREEN`, `BLUE` and `REVERSE` FGL color attributes is
adapted when the device uses dark mode.

However, if custom presentation style background and foreground colors are used, one may want to
force the light mode. This can be achieved by setting the `UIUserInterfaceStyle`
entry in your Info.plist file, in the `<dict/>` element:

```
<dict>
 ...
 <key>UIUserInterfaceStyle</key>
 <string>Light</string>
</dict>
```

## Cordova plugins from the FOURJS Cordova Github

Before Genero 3.20, the Cordova plugins wrapper libraries and demos were shipped in the FGLGWS
package (FGLDIR).

Cordova plugins, as well as demos and BDL wrapper libraries and demos, are now available from the
FOURJS Cordova Github.

For more details, see [Installing Cordova plugins](../17_mobile-applications/5119-installing-cordova-plugins.md "Before usage, Cordova plugins need to be installed in the GMA or GMI development environment.").

## New Cordova front calls `listPlugins` / `getPluginInfo`

Since Genero 3.20 / GMI 1.40, the new [`cordova.listPlugins`](../15_library-reference/3479-cordova-listplugins.md "Returns the list of available Cordova plugins.") and [`cordova.getPluginInfo`](../15_library-reference/3478-cordova-getplugininfo.md "Returns details about a specific Cordova plugin.")
front calls can be used to get information about Cordova plugins bundled with the app.

## Returned barcode scan codes with `mobile.scanBarCode` front call

Since Genero 3.20 / GMI 1.40, the barcode reader used by GMI is a native built-in. Since then it
matches also as close as possible with the GMA barcode types (using underscores). For more details
on supported barcodes types, see [mobile.scanBarCode](../15_library-reference/3459-mobile-scanbarcode.md "Allow the user to scan a barcode with a mobile device").

## New option `--update` for gmibuildtool

Starting with GMI 1.40.14, the gmibuildtool command provides a new
`--update` option, to be used when installing a package on the simulator, in order to
re-install the app without losing the user data/settings.

For more details, see [Installing the app on a device](../17_mobile-applications/5109-building-ios-apps-with-genero.md).
