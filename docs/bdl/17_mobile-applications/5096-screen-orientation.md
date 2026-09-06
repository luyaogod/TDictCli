---
title: "Screen orientation"
source: "fgl-topics/c_fgl_mobile_screen_orientation.html"
breadcrumb: "Mobile applications > Screen orientation"
type: "concept"
---

# Screen orientation

> Detecting screen orientation changes with the mobile device.

## Designing forms adapting to mobile screen orientation

With Responsive Layout features, app forms can adapt to the screen orientation of a mobile
device, by using specific form attributes such as `HBOX` attribute
`ORIENTATION@SMALL=VERTICAL`.

On mobile devices with a small screen (classical smartphones), the portrait mode is consisdered
as a `SMALL` screen, and landscape mode is considered as `MEDIUM`
screen. On a typical tablet, portrait mode will be considered as a `MEDIUM` screen,
while landscape mode is `LARGE`.

For more details about available form attributes, see [Responsive Layout](../11_user-interface/1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.").

## Controlling screen orientation on GMA / Android

Use the `allowedOrientations` Window style attribute to control screen orientation
options on Android™
devices:

```
<Style name="Window.main">
  <StyleAttribute name="allowedOrientations" value="portrait_reverse" />
</Style>
```

Starting with GMA 5.01.02, the allowedOrientations Window style attribute is ignored for large
devices – devices where the smallest width qualifier is equal to or greater than sw600dp.

For the complete list of possible values, see [`allowedOrientations`](../11_user-interface/1656-window-style-attributes-miscellaneous.md) style attribute.

## Controlling screen orientation on GMI / iOS

On iOS devices using GMI, the screen orientation possibilities must be defined in the
Info.plist resource file. Keys such as
`UISupportedInterfaceOrientations` must be defined in the file, for
example:

```
<?xml version="1.0" encoding= "UTF-8"?>
<!DOCTYPE plist PUBLIC  "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd" >
<plist version= "1.0">
<dict>
        <key>UISupportedInterfaceOrientations</key>
        <array>
                <string>UIInterfaceOrientationPortrait</string>
                <string>UIInterfaceOrientationPortraitUpsideDown</string>
        </array>
        <key>UISupportedInterfaceOrientations~ipad</key>
        <array>
                <string>UIInterfaceOrientationPortrait</string>
                <string>UIInterfaceOrientationPortraitUpsideDown</string>
        </array>
</dict>
</plist>
```

For more details about Info.plist usage, see [the section about
defining iOS app properties in Info.plist](5109-building-ios-apps-with-genero.md).
