---
title: "gwa.app.TServerVersionAndBuildResult type"
source: "fgl-topics/c_gwa_app_api_type_TServerVersionAndBuildResult.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.TServerVersionAndBuildResult type"
type: "concept"
---

# gwa.app.TServerVersionAndBuildResult type

> The TServerVersionAndBuildResult type defines a record for checking the up-to-date status of the application on the server.

## Syntax

```
TYPE TServerVersionAndBuildResult RECORD
  serverVersion STRING
  serverBuild gwa.app.TBuildDateTime
  error STRING
END RECORD
```

1. `serverVersion` holds the server version
2. `serverBuild` holds the server version's build date-time stamp.
3. `error` holds an error message.

## Usage

It provides a data type for checking the current status of the application on the server and is
used as the return type in [gwa.app.getServerVersionAndBuild](5176-gwa-app-getserverversionandbuild.md "Returns server application version information.").

In this example, the action "`svr`" displays the server version to the
message field:

```
IMPORT FGL gwa.app

FUNCTION menu_gwa_info()

  DEFINE svr_i gwa.app.TServerVersionAndBuildResult
  CALL gwa.app.getServerVersionAndBuild() RETURNING svr_i

  MENU "GWA info"
    ON ACTION svr ATTRIBUTES(TEXT="Server ver", COMMENT="Shows the Server version")
           MESSAGE SFMT("Server version: %1", svr_i.serverVersion)
    ON ACTION cancel ATTRIBUTES(TEXT="Back")
      EXIT MENU
  END MENU

END FUNCTION
```
