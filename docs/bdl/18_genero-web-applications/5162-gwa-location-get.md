---
title: "gwa.location.get"
source: "fgl-topics/c_gwa_api_location_get.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.location module > gwa.location module > gwa.location.get"
type: "concept"
---

# gwa.location.get

> Returns the location description of the application window.

## Syntax

```
FUNCTION get()
  RETURNS gwa.location.TGWALocation
```

Returns a `gwa.location.TGWALocation` record.

## Usage

The `gwa.location.get()` method is a wrapper for the JavaScript API
`window.location.get` object. For more information about the JavaScript API, refer to
[Location](https://developer.mozilla.org/en-US/docs/Web/API/Location) (external link). Use `gwa.location.get()` to return
the location description of the application window.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

```
IMPORT util
IMPORT FGL gwa.location 

MAIN
  CALL fgl_settitle("App location")
  MENU "Main menu"
    COMMAND "Exit"
      EXIT MENU
    COMMAND "Get app location"
      VAR ret=gwa.location.get()
      MESSAGE "location:",util.JSON.stringify(ret)
    COMMAND "Get my page"
      VAR location=gwa.location.get()
      MESSAGE "href:",location.href
    COMMAND "Get my host"
      VAR location=gwa.location.get()
      MESSAGE "host:",location.host
  END MENU
END MAIN
```

In this second example, the action "`url`" displays the GWA installation
details in a message dialog and the GWA installation URL to the message
field:

```
IMPORT FGL gwa.location
IMPORT FGL fgldialog

FUNCTION menu_gwa_info()

  DEFINE loc_i gwa.location.TGWALocation
  DEFINE s_loc STRING

  MENU "GWA info"
    ON ACTION url ATTRIBUTES(TEXT="URL", COMMENT="Shows the GWA installation URL")
      CALL gwa.location.get() RETURNING loc_i
      LET s_loc=SFMT("%1 \nOrigin: %2 \nPath: %3 \nProtocol: %4",
                       loc_i.href, loc_i.origin, loc_i.pathname, loc_i.protocol)  
      CALL fgl_winmessage("Location", s_loc, "information")    
      MESSAGE SFMT("URL: %1",loc_i.href)
    ON ACTION cancel ATTRIBUTES(TEXT="Back")
      EXIT MENU
  END MENU

END FUNCTION
```
