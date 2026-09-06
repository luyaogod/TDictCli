---
title: "OAuthAPI.GetIdRoles()"
source: "fgl-topics/c_gws_oauthapi_GetIDRoles.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.GetIdRoles()"
type: "concept"
---

# OAuthAPI.GetIdRoles()

> Get OAuth ID Token authorization roles.

## Syntax

```
FUNCTION GetIdRoles()
RETURNS DYNAMIC ARRAY OF STRING
```

## Usage

Use this function to get the list of authorization OAuth 2.0 roles that the Identity Provider
securing the RESTful web service supports. IdP role information can be retrieved from the GWS engine
once it has been initialized by [OAuthAPI.init()](4940-oauthapi-init.md "To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server.").

In case of error, a `NULL` value will be returned.

## OAuthAPI.GetIdRoles function

```
IMPORT FGL OAuthAPI

DEFINE ind INTEGER
DEFINE roles DYNAMIC ARRAY OF STRING

MAIN

   # ...
   IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr" ) THEN
      MESSAGE "Error: unable to initialize OAuth"
      EXIT PROGRAM 1
   ELSE
      LET roles = OAuthAPI.GetIdRoles()
      FOR ind = 1 TO roles.getLength()
        DISPLAY "Role supported is: ", roles[ind]
      END FOR
   END IF
   # ... 
  
END MAIN
```
