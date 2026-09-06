---
title: "Deploying the gICAPI web component files"
source: "fgl-topics/c_fgl_webcomponent_gicapi_location.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Deploying the gICAPI web component files"
type: "concept"
---

# Deploying the gICAPI web component files

> Deploy web component files to the front-end platform before using gICAPI web components.

## Standard Genero gICAPI Web Components

The Genero BDL package provides a set of common, ready-to-use web components, that can be used in
your application.

The Genero BDL Web Components are located in $FGLDIR/webcomponents and are
found by default with any type of front-end configuration.

As these files are part of the FGLDIR installation directory, no deployment is required.

## Deploying the HTML document and the JavaScript gICAPI interface

The gICAPI web component files (main HTML file, additional JavaScript files and other
potential assets) must be available on the platform where the front-end executes.
Depending on your configuration, Genero supports several solutions to provide the
gICAPI web component files from a single location. In a distributed configuration
with many individual front-end nodes, consider centralizing the gICAPI files on a
server, instead of copying the gICAPI web component files manually to each front-end
device.

> **Important:**
>
> If the main gICAPI HTML document references external JavaScript
> files, put these files in the same directory as the HTML file referencing them.

## Recommended web component directory layout

When using the default settings in any configuration (for example, no FGLIMAGEPATH is defined,
default GAS settings), put the gICAPI web component files under a webcomponents
directory, along with the other program files, for
example:

```
appdir
appdir/main.42m
appdir/form1.42f
appdir/form2.42f
appdir/webcomponents/3DChart
appdir/webcomponents/3DChart/3DChart.html
appdir/webcomponents/3DChart/3DChart.js
appdir/webcomponents/3DChart/3DChart.css
appdir/webcomponents/3DChart/icon_close.png
...
```

## Deployment methods for front-end types

Web components assets are deployed with different solutions, depending on the front-end
configuration type:

- [Deploying gICAPI assets in direct mode](2404-deploying-gicapi-assets-in-direct-mode.md "Using GDC, GMI or GMA front-ends in direct mode (not through GAS)")
- [Deploying gICAPI assets with GAS](2405-deploying-gicapi-assets-with-gas.md "Using gICAPI web components through the GAS.")
- [Deploying gICAPI assets on mobile](2406-deploying-gicapi-assets-on-mobile.md "Using GMI and GMA front-ends, executing app on mobile device")

The GDC front-end supports also local gICAPI file lookup in the GDC installation directory.
However, this solution is deprecated. Consider centralizing the gICAPI web component files on the
application server, as described above in the dedicated topics.

## Defining the gICAPI files search path by program

With older versions, it was possible to use the [standard.setWebComponentPath](../15_library-reference/3416-standard-setwebcomponentpath.md "Defines the base path where web components are located.") front call, to define by program the base
URL to the web component files.

> **Important:**
>
> This front call is deprecated, consider using one of the other mechanisms
> described in this topic.

## Related links

**Related concepts**  

[Using image resources with the gICAPI web component](2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component.")

## Child topics

- [Deploying gICAPI assets in direct mode](2404-deploying-gicapi-assets-in-direct-mode.md): Using GDC, GMI or GMA front-ends in direct mode (not through GAS)
- [Deploying gICAPI assets with GAS](2405-deploying-gicapi-assets-with-gas.md): Using gICAPI web components through the GAS.
- [Deploying gICAPI assets on mobile](2406-deploying-gicapi-assets-on-mobile.md): Using GMI and GMA front-ends, executing app on mobile device
