---
title: "Deploying gICAPI assets on mobile"
source: "fgl-topics/c_fgl_webcomponent_gicapi_location_mobile.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Deploying the gICAPI web component files > Deploying gICAPI assets on mobile"
type: "concept"
---

# Deploying gICAPI assets on mobile

> Using GMI and GMA front-ends, executing app on mobile device

When running the application on mobile (in embedded mode), the gICAPI web component files (along
with other assets) must be deployed on the device: The files will be found locally on the
device.

In mobile development mode (in direct mode, when the application runs on a computer and forms
display on the GMA or GMI on the mobile device), the gICAPI web component assets are transferred to
the device with [the direct-mode
solution](2404-deploying-gicapi-assets-in-direct-mode.md "Using GDC, GMI or GMA front-ends in direct mode (not through GAS)"). Web component assets must be located on the computer when the program executes.

If your application is intended for different front-ends, consider using [the recommended gICAPI web component directory layout](2403-deploying-the-gicapi-web-component-files.md).

Genero BDL provides a set of standard web components in
$FGLDIR/webcomponents. When executing on a mobile device, the standard web
component assets will be implicitly found.

In embedded mode, mobile front-ends make a local search for gICAPI web component files in the
following order:

1. `$FGLDIR/webcomponents/component-type/component-type.html`
2. `appdir/webcomponents/component-type/component-type.html`
3. `appdir/component-type.html`

Where:

- $FGLDIR is the runtime system directory inside the deployed app.
- appdir is the application directory where program files are located.
- component-type is the name defined by the `COMPONENTTYPE`
  attribute in the form definition file.

For example, if you define a custom gICAPI web component field as
follows:

```
WEBCOMPONENT wc = FORMOMLY.mychart,
   COMPONENTTYPE = "3DChart";
```

The gICAPI web component HTML document will be found on the mobile device at:

- appdir/webcomponents/3DChart/3DChart.html

or:

- appdir/3DChart.html

However, using the second location is not recommended (always use a
webcomponents directory).

For more details about appdir on mobile devices, see [Deploying mobile apps on Android devices](../17_mobile-applications/5103-deploying-mobile-apps-on-android-devices.md "This section contains information to create a mobile application to be deployed on Android devices.") and [Deploying mobile apps on iOS devices](../17_mobile-applications/5107-deploying-mobile-apps-on-ios-devices.md "This section contains information to create a mobile application to be deployed on iOS devices.").
