---
title: "Deploying gICAPI assets in direct mode"
source: "fgl-topics/c_fgl_webcomponent_gicapi_location_direct.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Deploying the gICAPI web component files > Deploying gICAPI assets in direct mode"
type: "concept"
---

# Deploying gICAPI assets in direct mode

> Using GDC, GMI or GMA front-ends in direct mode (not through GAS)

When using a front-end with a direct connection (not through the GAS), web component files can be
automatically transferred to the front-end.

Providing gICAPI web component files with the direct mode mechanism simplifies the development
process for mobile applications, as you do not have to copy the files to the device.

In direct mode, gICAPI web component assets are typically located in the directory where the
`MAIN` program module resides, under a webcomponents
sub-directory. If web component files are located in a different base directory, add the search path
to the [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.") environment variable. However, if your
application is intended for different front-ends, consider using [the recommended gICAPI web component directory layout](2403-deploying-the-gicapi-web-component-files.md), to avoid FGLIMAGEPATH for web
component assets.

Genero BDL provides a set of standard web components in
$FGLDIR/webcomponents. These standard web component assets will be implicitly
found.

The gICAPI web component files are searched in the following directories:

1. `$FGLDIR/webcomponents/component-type/component-type.html`
2. `appdir/webcomponents/component-type/component-type.html`
3. `fglimagepath-dir/webcomponents/component-type/component-type.html`
4. `fglimagepath-dir/component-type.html`

Where:

- $FGLDIR is the runtime installation directory on the application
  server.
- appdir is the directory where the application program resides.
- fglimagepath-dir is one of the base directories defined in FGLIMAGEPATH.
- component-type is the name defined by the `COMPONENTTYPE`
  attribute in the form definition file.

If assets such as .js, .css, .png
files are referenced by a relative path name in the HTML content, the resources are also transferred
via the direct-mode mechanism. If the assets use an absolute path with a concrete URL scheme (such
as `http://something` ), the HTML viewer will try to get the resource from the URL
location.

For example, if you define the gICAPI web component field as
follows:

```
WEBCOMPONENT wc = FORMOMLY.mychart,
   COMPONENTTYPE = "3DChart";
```

If the FGLIMAGEPATH search path contains "/opt/myapp", and the gICAPI files
are located under "/opt/myapp/webcomponents/3DChart", the gICAPI web component
HTML document will be found on the server at:

- /opt/myapp/webcomponents/3DChart/3DChart.html

For backward compatibility, the GDC front-end is able to find web components locally on the
workstation where it executes. However, this solution is deprecated. Consider centralizing your
gICAPI web components on the application server.
