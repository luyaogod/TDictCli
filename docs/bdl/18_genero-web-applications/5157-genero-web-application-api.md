---
title: "GWA API"
source: "fgl-topics/c_gwa_api_functions_overview.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API"
type: "concept"
---

# GWA API

> The Genero Web Application API is a package included with GWA installations, offering modules like gwa.location and gwa.app to assist with application development.

In this example, various methods from the GWA API are used to return information about the GWA
application.

```
IMPORT FGL gwa.app
IMPORT FGL gwa.location

MAIN
  CALL menu_gwa_info()
END MAIN

FUNCTION menu_gwa_info()

  DEFINE loc_i gwa.location.TGWALocation
  DEFINE svr_i gwa.app.TServerVersionAndBuildResult
  DEFINE app_i gwa.app.TAppInformation

  CALL gwa.app.getServerVersionAndBuild() RETURNING svr_i
  CALL gwa.app.getInformation() RETURNING app_i

  MENU "GWA info"
    ON ACTION url
        ATTRIBUTES(TEXT = "URL", COMMENT = "Shows the GWA installation URL")
      CALL gwa.location.get() RETURNING loc_i
      MESSAGE SFMT("URL: %1", loc_i.href)

    ON ACTION appver
        ATTRIBUTES(TEXT = "App ver", COMMENT = "Shows the app's version")
      MESSAGE SFMT("App version: %1", gwa.app.getVersion())

    ON ACTION appbuild
        ATTRIBUTES(TEXT = "App build", COMMENT = "Shows the app's version")
      MESSAGE SFMT("App build: %1", gwa.app.getBuild())

    ON ACTION tit ATTRIBUTES(TEXT = "Title", COMMENT = "Shows the app's Title")
      MESSAGE SFMT("App title: %1", gwa.app.getTitle())

    ON ACTION gwa
        ATTRIBUTES(TEXT = "GWA ver", COMMENT = "Shows the GWA version")
      MESSAGE SFMT("GWA version: %1", app_i.gwaVersion)

    ON ACTION svr
        ATTRIBUTES(TEXT = "Server ver", COMMENT = "Shows the Server version")
      MESSAGE SFMT("Server version: %1", svr_i.serverVersion)

    ON ACTION gbc
        ATTRIBUTES(TEXT = "GBC ver", COMMENT = "Shows the GBC version")
      MESSAGE SFMT("GBC version: %1", gwa.app.getGBCVersion())

    ON ACTION cancel ATTRIBUTES(TEXT = "Back")
      EXIT MENU

  END MENU
END FUNCTION
```

## Child topics

- [Finding modules with FGLLDPATH](5158-finding-modules-with-fglldpath.md): For compilation, you may need to set the FGLLDPATH environment variable to find the modules in the GWA API package.
- [The gwa.location module](5159-the-gwa-location-module.md): The gwa.location module provides methods and types for interfacing GWA applications with HTML DOM API location object.
- [The gwa.app module](5164-the-gwa-app-module.md): The gwa.app module provides methods and types for querying version numbers and for working with application updates.
- [The base.Application.isGWA method](5184-the-base-application-isgwa-method.md): The base.Application.isGWA method can be called to check if the program code is running in a browser as a GWA application.
