---
title: "OAuthAPI.GetMyAccessToken()"
source: "fgl-topics/c_gws_oauthapi_GetMyAccessToken.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.GetMyAccessToken()"
type: "concept"
---

# OAuthAPI.GetMyAccessToken()

> Get a valid access token.

## Syntax

```
FUNCTION GetMyAccessToken()
RETURNS STRING
```

`NULL` may be returned if the access token is not found or can not be renewed.

## Usage

Use the `GetMyAccessToken()` function to retrieve the access token provided by the
Identity Provider (IdP) for access to a secure RESTful web service. The GWS engine can provide the
IdP token information once it has been initialized by [OAuthAPI.init()](4940-oauthapi-init.md "To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server."). If
the access token has expired, use this function to request a new token (if a refresh token is
available).

In case of error, a `NULL` value will be returned.

## OAuthAPI.GetMyAccessToken function

```
IMPORT FGL OAuthAPI

DEFINE access_token STRING

MAIN
   IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr" ) THEN
      MESSAGE "Error: unable to initialize OAuth"
      EXIT PROGRAM 1
   ELSE
      LET access_token = OAuthAPI.GetMyAccessToken()
   END IF
   # ... 
  
END MAIN
```
