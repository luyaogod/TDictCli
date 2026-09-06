---
title: "The base.Application.isGWA method"
source: "fgl-topics/c_gwa_base_application_isgwa.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The base.Application.isGWA method"
type: "concept"
---

# The base.Application.isGWA method

> The base.Application.isGWA method can be called to check if the program code is running in a browser as a GWA application.

For example, this method can be used in the following way as part of a dialog
statement:

```
IMPORT FGL gwatest

MAIN

    MENU
        ON ACTION gwa_info
            ATTRIBUTES(TEXT = "GWA info",
                COMMENT = "Tells whether or not this is a GWA app.")
            IF base.Application.isGWA() THEN
                CALL gwatest.menu_gwa_info()
            ELSE
                ERROR "This is not a GWA application."
            END IF
        ON ACTION EXIT
            EXIT MENU
    END MENU
    
END MAIN
```

To view the code for gwatest, go to [Genero Web
Application API](5157-genero-web-application-api.md "The Genero Web Application API is a package included with GWA installations, offering modules like gwa.location and gwa.app to assist with application development.")

## Related links

**Related concepts**  

[base.Application.isGWA](../15_library-reference/2977-base-application-isgwa.md "Indicates if the application runs in a browser.")
