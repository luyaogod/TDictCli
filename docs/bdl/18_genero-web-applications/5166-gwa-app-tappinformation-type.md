---
title: "gwa.app.TAppInformation type"
source: "fgl-topics/c_gwa_app_api_type_TAppInformation.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.TAppInformation type"
type: "concept"
---

# gwa.app.TAppInformation type

> The TAppInformation type defines a record for returning application information.

## Syntax

```
TYPE TAppInformation RECORD
   title STRING
   version STRING
   build gwa.app.TBuildDateTime
   gwaVersion STRING
   embeddedVMVersion STRING
   fgl_getversion STRING
   gbcVersion STRING
   server STRING 
END RECORD
```

1. `title` holds the application title.
2. `version` holds the application version.
3. `build` holds the build date-time stamp.
4. `gwaVersion` holds the GWA version.
5. `embeddedVMVersion` holds the VM version.
6. `fgl_getversion` holds the FGL version.
7. `gbcVersion` holds the GBC version.
8. `server` holds the initial HTTP header 'Server' value of the web server.

## Usage

It provides an API to get application information and is used as the return type in [gwa.app.getInformation](5175-gwa-app-getinformation.md "Returns the application version information."). A variable of the type `TAppInformation`
must be defined.

In this example, the action "`gwa`" displays the GWA version – not the app
version, but the GWA version – to the message field:

```
IMPORT FGL gwa.app

FUNCTION menu_gwa_info()

  DEFINE app_i gwa.app.TAppInformation
  CALL gwa.app.getInformation() RETURNING app_i

  MENU "GWA info"
    ON ACTION gwa ATTRIBUTES(TEXT="GWA ver", COMMENT="Shows the GWA version")
           MESSAGE SFMT("GWA version: %1", app_i.gwaVersion)
    ON ACTION cancel ATTRIBUTES(TEXT="Back")
      EXIT MENU
  END MENU

END FUNCTION
```
