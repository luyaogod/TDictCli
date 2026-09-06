---
title: "gwa.app.getServerVersionAndBuild"
source: "fgl-topics/c_gwa_app_api_getServerVersionAndBuild.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.getServerVersionAndBuild"
type: "concept"
---

# gwa.app.getServerVersionAndBuild

> Returns server application version information.

## Syntax

```
FUNCTION getServerVersionAndBuild()
  RETURNS gwa.app.TServerVersionAndBuildResult
```

Returns a [gwa.app.TServerVersionAndBuildResult](5170-gwa-app-tserverversionandbuildresult-type.md "The TServerVersionAndBuildResult type defines a record for checking the up-to-date status of the application on the server.") record with the server build version, date time stamp,
and error.

## Usage

Use the `gwa.app.getServerVersionAndBuild()` method to make a network request to
the server for index.html and extract the version information. On success, it
returns the server application version and the server build date. On error, the error string is
set.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
