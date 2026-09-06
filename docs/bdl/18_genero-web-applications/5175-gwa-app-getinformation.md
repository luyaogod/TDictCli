---
title: "gwa.app.getInformation"
source: "fgl-topics/c_gwa_app_api_getInformation.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.getInformation"
type: "concept"
---

# gwa.app.getInformation

> Returns the application version information.

## Syntax

```
FUNCTION getInformation()
  RETURNS gwa.app.TAppInformation
```

Returns a `gwa.app.TAppInformation` record with detailed version information about
the application build.

## Usage

Use the `gwa.app.getInformation()` method to return application version
information.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
