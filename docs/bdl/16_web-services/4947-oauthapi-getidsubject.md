---
title: "OAuthAPI.GetIDSubject()"
source: "fgl-topics/c_gws_oauthapi_GetIDSubject.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.GetIDSubject()"
type: "concept"
---

# OAuthAPI.GetIDSubject()

> Get OAuth subject identifier of ID Token.

## Syntax

```
FUNCTION GetIDSubject()
RETURNS STRING
```

`NULL` may be returned if the metadata is not found.

## Usage

Use this function to get the subject identifier of the Identity Provider securing the RESTful Web
service. IdP subject information can be retrieved from the GWS engine once it has been initialized
by [OAuthAPI.init()](4940-oauthapi-init.md "To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server.").

In case of error, a `NULL` value will be returned.

## OAuthAPI.GetIDSubject function

```
IMPORT FGL OAuthAPI

DEFINE my_user_id STRING
MAIN
  # ...
   
  # Init OAuthAPI
  IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr") THEN
    DISPLAY "Error: unable to initialize OAuth"
    EXIT PROGRAM 1
  ELSE
    LET my_user_id = OAuthAPI.getIDSubject()
  END IF

  # ... 
  
END MAIN
```
