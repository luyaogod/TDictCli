---
title: "OAuthAPI.GetIDPIssuer()"
source: "fgl-topics/c_gws_oauthapi_GetIDPIssuer.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.GetIDPIssuer()"
type: "concept"
---

# OAuthAPI.GetIDPIssuer()

> Get endpoint of the Identity Provider.

## Syntax

```
FUNCTION GetIDPIssuer()
RETURNS STRING
```

`NULL` may be returned if the metadata is not found.

## Usage

Use this function to get the endpoint of the Identity Provider securing the RESTful Web service.
IdP issuer information can be retrieved from the GWS engine once it has been initialized by [OAuthAPI.init()](4940-oauthapi-init.md "To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server.").

In case of error, a `NULL` value will be returned.

## OAuthAPI.GetIDPIssuer function

```
IMPORT FGL OAuthAPI

DEFINE issuer STRING

MAIN
   IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr" ) THEN
      MESSAGE "Error: unable to initialize OAuth"
      EXIT PROGRAM 1
   ELSE
      LET issuer = OAuthAPI.GetIDPIssuer()
   END IF
   # ... 
  
END MAIN
```
