---
title: "gwa.app.getGBCVersion"
source: "fgl-topics/c_gwa_app_api_getGBCVersion.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.getGBCVersion"
type: "concept"
---

# gwa.app.getGBCVersion

> Returns the version of the GBC embedded in the application.

## Syntax

```
FUNCTION getGBCVersion()
  RETURNS STRING
```

Returns a GBC version.

## Usage

Use the `gwa.app.getGBCVersion()` method to return the version of the GBC embedded
in the application.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

In this example, the action "`tit`" displays the GWA app title to the
message field:

```
IMPORT FGL gwa.location

FUNCTION menu_gwa_info()

  MENU "GWA info"
    ON ACTION gbc ATTRIBUTES(TEXT="GBC ver", COMMENT="Shows the GBC version")
           MESSAGE SFMT("GBC version: %1",gwa.app.getGBCVersion())
    ON ACTION cancel ATTRIBUTES(TEXT="Back")
      EXIT MENU
  END MENU

END FUNCTION
```

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
