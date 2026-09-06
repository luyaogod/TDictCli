---
title: "gwa.app.getVersion"
source: "fgl-topics/c_gwa_app_api_getVersion.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.getVersion"
type: "concept"
---

# gwa.app.getVersion

> Returns the application version set by gwabuildtool --app-version.

## Syntax

```
FUNCTION getVersion() RETURNS STRING
```

Returns the application version.

## Usage

Use the `gwa.app.getVersion()` method to return the version set in [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.") with the `--app-version` option. If this
option was not set, it returns `NULL`.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

In this example, the action "`app`" displays the GWA app version to the
message field. An app version is only displayed if a version number was passed in by the
--app-version option of the gwabuildtool or
gwarun command line tools; otherwise, it is
empty.

```
IMPORT FGL gwa.location

FUNCTION menu_gwa_info()

  MENU "GWA info"
    ON ACTION app ATTRIBUTES(TEXT="App ver", COMMENT="Shows the app's version")
           MESSAGE SFMT("App version: %1",gwa.app.getVersion())
    ON ACTION cancel ATTRIBUTES(TEXT="Back")
      EXIT MENU
  END MENU

END FUNCTION
```

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
