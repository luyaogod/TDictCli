---
title: "Color and theming"
source: "fgl-topics/c_fgl_mobile_theming.html"
breadcrumb: "Mobile applications > Color and theming"
type: "concept"
---

# Color and theming

> Mobile applications must follow the platform colors and theming.

## User interface design on mobile devices

Genero BDL provides several ways to define colors and styles for a mobile app. This section
introduces features that can be used to customize your mobile app and adapt to the target platform
user interface design. As a general rule, avoid using non-standard ergonomics and decorations, use
defaults to let Genero render your forms depending on the platform standards. For example, the GMA
front-end will use Google Material Design on Android™ devices.

## Defining TTF icon colors

By default, TTF icons get the color of the platform theme. A default color can be defined for all
TTF icons of a window with the `defaultTTFColor` style attribute. In order to define
a color for a specific icon, add an #RGB color specification in your image to font glyph mapping
file.

For more details, see [Using a simple image name (centralized TTF icons)](../11_user-interface/1586-providing-the-image-resource.md)

## App color theme on Android

Android apps can be created
with a specific color theme following the Google Material Design.

When building the APK with the gmabuildtool, you can specify the general app colors with the
`--build-app-colors` option.

For more details, see [Define app's color theme](5105-building-android-apps-with-genero.md).
