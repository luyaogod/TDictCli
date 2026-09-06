---
title: "OAuthAPI.GetIDScopes()"
source: "fgl-topics/c_gws_oauthapi_GetIDScopes.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.GetIDScopes()"
type: "concept"
---

# OAuthAPI.GetIDScopes()

> Get OAuth ID Token authorization scopes.

## Syntax

```
FUNCTION GetIDScopes()
RETURNS DYNAMIC ARRAY OF STRING
```

## Usage

Use this function to get the list of authorization OAuth 2.0 scopes that the Identity Provider
securing the RESTful web service supports. IdP scope information can be retrieved from the GWS
engine once it has been initialized by [OAuthAPI.init()](4940-oauthapi-init.md "To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server.").

In case of error, a `NULL` value will be returned.

## OAuthAPI.GetIDScopes function

```
IMPORT FGL OAuthAPI

DEFINE ind INTEGER
DEFINE scopes DYNAMIC ARRAY OF STRING

MAIN

   # ...
   IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr" ) THEN
      MESSAGE "Error: unable to initialize OAuth"
      EXIT PROGRAM 1
   ELSE
      LET scopes = OAuthAPI.GetIDScopes()
      FOR ind = 1 TO scopes.getLength()
        DISPLAY "Scope supported is: ", scopes[ind]
      END FOR
   END IF
   # ... 
  
END MAIN
```
