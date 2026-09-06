---
title: "gwa.app.getBuild"
source: "fgl-topics/c_gwa_app_api_getBuild.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.getBuild"
type: "concept"
---

# gwa.app.getBuild

> Returns the date time stamp of when the application was built.

## Syntax

```
FUNCTION getBuild()
  RETURNS gwa.app.TBuildDateTime
```

Returns a `gwa.app.TBuildDateTime` record with the date of the build.

## Usage

Use the `gwa.app.getBuild()` method to get the date time stamp of when the
application was built.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
