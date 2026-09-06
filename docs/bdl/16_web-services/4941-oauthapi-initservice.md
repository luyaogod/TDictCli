---
title: "OAuthAPI.InitService()"
source: "fgl-topics/c_gws_oauthapi_InitService.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.InitService()"
type: "concept"
---

# OAuthAPI.InitService()

> To be called in a Genero web service started via OpenID Connect/OAuth2 accessing another secure RESTful web service as a client.

## Syntax

```
FUNCTION InitService( 
   cnx_timeout INTEGER, 
   access_token STRING )
  RETURNS BOOLEAN
```

1. cnx\_timeout is a connection timeout to the REST service
   with value in seconds.
2. access\_token is the valid access token for accessing the RESTful web
   service.

Returns `FALSE` if the mandatory access token is null.

## Usage

Use the `InitService()` function to register the access token needed when a
service (server-side) must connect as a client to another service that is protected by an access
token. This function checks whether the OAuth service has been initiated.

The primary role of `InitService()` is to register the access token, allowing you
to call any of the OAuthAPI methods, such as `CreateHTTPAuthorizationRequest`, to
perform requests to the other service.

In the provided code sample, a `WSContext` dictionary is defined to
retrieve information from the service configuration file (.xcf). The access
token set during token validation is required to access the client service, so you must call
`InitService()` with the same token.

Typically, only the access token is required. If you need to retrieve metadata, you can call
`FetchOpenIDMetadata()`,
which will save the metadata in an `OAuthAPI.OpenIDMetadataType` record.

In case of error, a `NULL` value will be returned.

## OAuthAPI.InitService function

```
IMPORT FGL OAuthAPI

PRIVATE DEFINE ctx DICTIONARY ATTRIBUTES(WSContext) OF STRING

MAIN
  DEFINE metadata     OAuthAPI.OpenIDMetadataType
  DEFINE access_token STRING
  DEFINE idp          STRING
 
  # retrieve access_token 
  LET access_token = ctx["OIDC_ACCESS_TOKEN"]
  
  # retrieve IDP URL
  LET idp = ctx["Parameter-IDP"]

  # retrieve metadata
  CALL OAuthAPI.FetchOpenIDMetadata(5,idp) RETURNING metadata.*
  
  # Init OAuth service
  IF NOT OAuthAPI.InitService(5, access_token) THEN
    DISPLAY "Cannot initiate OAuth service"
  END IF

  CALL MyOtherService(access_token)
  
  # ... 
  
END MAIN
```
