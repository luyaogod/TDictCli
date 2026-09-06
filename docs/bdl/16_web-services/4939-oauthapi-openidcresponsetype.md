---
title: "OAuthAPI.OpenIdCResponseType"
source: "fgl-topics/c_gws_oauthapi_OpenIdCResponseType.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.OpenIdCResponseType"
type: "concept"
---

# OAuthAPI.OpenIdCResponseType

> The OpenIdCResponseType record stores the access token, refresh token, and token expiry date retrieved in a request to the IdP.

## Syntax

```
TYPE OpenIdCResponseType RECORD
  access_token  STRING,
  token_type    STRING,
  expires_in    INTEGER,
  refresh_token STRING
END RECORD
```

The `OpenIdCResponseType` record stores the access token, refresh token, and token
expiry date retrieved in a request to the IdP using [OAuthAPI.RetrievePasswordTokenForNativeApp()](4954-oauthapi-retrievepasswordtokenfornativeapp.md "Returns the OAuth service access token via user credentials (username/password) and client credentials (client_id/secret_id). A refresh token allows the access token to be refreshed when it expires.").

## Usage

This type defines a record structure to store tokens. If your Genero app is not behind a Genero
Application Server, you can use this record when getting tokens to access a secure RESTful web
service.
