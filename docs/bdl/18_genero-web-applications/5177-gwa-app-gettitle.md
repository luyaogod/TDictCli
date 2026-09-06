---
title: "gwa.app.getTitle"
source: "fgl-topics/c_gwa_app_api_getTitle.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.getTitle"
type: "concept"
---

# gwa.app.getTitle

> Returns the application title set by gwabuildtool --title.

## Syntax

```
FUNCTION getTitle() RETURNS STRING
```

Returns the application title.

## Usage

Use the `gwa.app.getTitle()` method to return the title set in [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.") with the `--title` option. If this option
was not set, it returns the name of the module containing `MAIN`.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

In this example, the action "`tit`" displays the GWA app title to the
message field:

```
IMPORT FGL gwa.location

FUNCTION menu_gwa_info()

  MENU "GWA info"
    ON ACTION tit ATTRIBUTES(TEXT="Title", COMMENT="Shows the app's Title")
           MESSAGE SFMT("App title: %1",gwa.app.getTitle())
    ON ACTION cancel ATTRIBUTES(TEXT="Back")
      EXIT MENU
  END MENU

END FUNCTION
```

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
