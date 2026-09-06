---
title: "standard.feInfo"
source: "fgl-topics/c_fgl_frontcall_standard_feinfo.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.feInfo"
type: "concept"
---

# standard.feInfo

> Queries general front-end properties.

## Syntax

```
ui.Interface.frontCall("standard", "feInfo",
 [name], [result])
```

1. name - The name of the property.
2. result - The value of the property.

## Usage

The `feInfo` front call returns a front-end property value depending on the
property name passed in as the parameter.

Some `feInfo` options take an optional parameter, such as
`screenResolution`:

```
CALL ui.Interface.frontCall("standard", "feInfo", ["screenResolution", 2], [resolution])
```

| Property name | Description |
| --- | --- |
| `browserName` | Returns the web browser type used by the front-end.Possible returned values are: `"Firefox"`, `"Chrome"`, `"Edge"`, `"Opera"`, `"Safari"`" and `"Unknown"` |
| `colorScheme` | Returns the current brightness mode of the display device / user agent.The possible values returned are: `"light"` or `"dark"`. |
| `dataDirectory` | Returns the directory name that can be used for temporary files on the front-end side. This directory is cleaned at front-end start-up and end, and is common to all front-end instances.The possible values returned are:With Genero Application Server (GAS), this is not applicable.With Genero Desktop Client (GDC), the local cache directory. For example, "/home/username/.cache/Four Js/Genero Desktop".With Genero Mobile for Android™ (GMA), this is the GMA application cache directory. Content may be erased, once the app is closed.With Genero Mobile for iOS (GMI), this is the temporary directory in the application sandbox (iOS `NSTemporaryDirectory()` system call). Content may be erased, once the app is closed. |
| `deviceId` | Identifies the device with a unique ID (only for mobile):With Genero Mobile for Android (GMI), returns the [`ANDROID_ID`](https://developer.android.com/reference/android/provider/Settings.Secure.html#ANDROID_ID).With Genero Mobile for iOS (GMI), returns the `identifierForVendor`.**Important:***Security warning*: The `deviceId` should not be used to identify users: Mobile devices expose unique identifiers that can be used to identify users across applications or devices. These identifiers put user privacy at risk, as they might allow the tracking of user activity without consent, while making it difficult or impossible for users to reset them. Privacy violations can cause apps to be removed from app stores and can result in legal action or loss of trust from users. |
| `deviceModel` | Returns the name of the device, for example "iPad4,5".This property is only supported with Genero Mobile for Android (GMA) and Genero Mobile for iOS (GMI). |
| `feName` | The code identifying the type of front-end component.The possible values returned are:`"Genero Desktop Client"` for Genero Desktop Client.`"GBC"` for Genero Browser Client.`"GMA"` for Genero Mobile for Android.`"GMI"` for Genero Mobile for iOS.**Tip:**To save a network round trip, it is recommended to use [`ui.Interface.getFrontEndName()`](3103-ui-interface-getfrontendname.md "Returns the type of the front-end currently in use.") instead. |
| `fePath` | The installation directory of the front-end executable.This property has no meaning with Genero Application Server (GAS).With Genero Desktop Client (GDC), it returns the path to the installation directory of the GDC.With Genero Mobile for Android (GMA), it returns the installation directory. For example, "/data/data/com.fourjs.gma/fgl".With Genero Mobile for iOS (GMI), it returns the installation directory. For example: "/private/var/mobile/Applications/B3E6-C48A-ED4EFA". Below the installation directory are the "Documents" (which is by default pwd), "GMI.app" (deployed p-code resides in GMI.app/app/) and "tmp" directories.**Important:**The installation path returned by this front call may change in future versions, do not base application code on this. On mobile devices, consider using the [os.Path.pwd](3735-os-path-pwd.md "Returns the current working directory.") utility function to get the application working directory when executing programs. |
| `freeStorageSpace` | Returns the number of bytes available on the mobile device.This property is only valid in the context of Genero Mobile for Android (GMA) and Genero Mobile for iOS (GMI) front-ends. |
| `ip` | Returns the IP address of the network interface used for the GUI connection.This property has no meaning with Genero Application Server (GAS).With Genero Desktop Client (GDC), this is the IP address of the computer where GDC is running.For mobile platforms (GMA, GMI), this is the preferred IP address of the device. If there is WIFI, either the IPv4 address is returned (for example: `"192.168.0.12"`) or if there is no IPv4 address, the IPv6 address is returned (for example: `"2a02:810a:82c0:478:d462:e334:6a1d:fb78"`). If there is no WIFI, either the cellular IPv4 or IPv6 address is returned. If there is no network, `NULL` is returned. |
| `numScreens` | Number of screens available on the front-end platform.With Genero Application Server (GAS), the number of screens is always 1.With Genero Desktop Client (GDC), returns the number of screens plugged to the computer where the GDC is running. |
| `osType` | The operating system type where the front-end is running.Possible return values include `"WINDOWS"`, `"LINUX"`, `"OSX"`, `"ANDROID"`, `"IOS"`. |
| `osVersion` | The version of the operating system.Example of returned values: `"4.3"`, `"5.10.15"`.With Genero Application Server (GAS), the returned OS version is always `"unknown"`. |
| `ppi` | Returns the screen pixel density of the front-end platform (Pixels Per Inch). This front call takes an optional screen number as parameter (1 is the default).With Genero Mobile for iOS, it returns the `PixelsPerInch` of an iOS device.With Genero Mobile for Android, it returns the DPI (ppi == dpi) |
| `screenResolution` | Returns the screen resolution of the front-end platform. This front call takes an optional screen number as parameter (1 is the default).Example of returned values: `"1200x1824"`, `"1920x1104"`.**Note:**For mobile devices, the value can change depending on the device orientation. |
| `target` | Returns the build platform target code name, identifying the operating system the front-end binary was compiled.This front call is provided for debugging purpose: Do not base code on the returned value. Use the `osType` property instead.With Genero Application Server (GAS), the returned value is always `"web"`.Example of returned values:`"w64v110"` = Windows® 64 bits, Visual C++ 11.`"d64a050"` = Android 4.0 ARM 64 bits.`"d64x050"` = Android 4.0 x86 64 bits.`"i32a070"` = iOS 7.0 ARM 32 bits.`"i32x070"` = iOS 7.0 x86 32 bits. |
| `windowSize` | Returns the current size of the front-end view-port.For mobile front-ends, this is the size of the mobile screen.For Genero Desktop Client, this is the size of the window container.For Genero Browser Client, this is the size of the browser webview.Example of returned values: `"1200x1824"`, `"1920x1104"`. |
| `userPreferredLang` | Returns the language and territory of the locale defined on the front-end platform, in the `language_territory` format.With Genero Desktop Client (GDC) and mobile front-ends (GMA, GMI), the front-end locale is defined by the operating system.With Genero Application Server (GAS/GBC), the front-end locale is defined in the web browser preferences. |
