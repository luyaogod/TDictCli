---
title: "gwa.app.versionsAndBuildsEqual"
source: "fgl-topics/c_gwa_app_api_versionsAndBuildsEqual.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.versionsAndBuildsEqual"
type: "concept"
---

# gwa.app.versionsAndBuildsEqual

> Checks if the versions and build date-time stamps are equal. If versions are NULL, only build date-time stamps are compared for equality.

## Syntax

```
FUNCTION versionsAndBuildsEqual(
     version1 STRING, 
     version2 STRING, 
     build1 gwa.app.TBuildDateTime,
     build2 gwa.app.TBuildDateTime) 
  RETURNS BOOLEAN
```

1. version1 holds the version number.
2. version2 holds the comparison version number.
3. build1 holds a `gwa.app.TBuildDateTime` record with the build
   date-time stamp.
4. build2 holds a `gwa.app.TBuildDateTime` record with the
   comparison build date-time stamp.

Returns `TRUE` if the versions/build date-time stamps match.

## Usage

Use the `gwa.app.versionsAndBuildsEqual()` method  to check if the versions
and build date-time stamps are equal. If versions are `NULL`, only build date-time
stamps are compared for equality.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
