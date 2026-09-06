---
title: "OAuthAPI.InitNativeApp()"
source: "fgl-topics/c_gws_oauthapi_InitNativeApp.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.InitNativeApp()"
type: "concept"
---

# OAuthAPI.InitNativeApp()

> To be called in a Genero application accessing a secure RESTful web service directly (not behind a Genero Application Server).

## Syntax

```
FUNCTION InitNativeApp(
  cnx_timeout INTEGER,
  tokens OpenIdCResponseType,
  client_id STRING,
  client_secret STRING,
  token_end_point STRING)
RETURNS BOOLEAN
```

1. cnx\_timeout is a connection timeout to the REST service
   with value in seconds.
2. tokens is the record, provided by the `OAuthAPI.RetrievePasswordTokenForNativeApp()` function, containing the
   access token for accessing the RESTful Web service with information about its expiry, and so on,
   that allows it to be refreshed.
3. client\_id is the application ID assigned to the app when
   registered.
4. client\_secret is the application secret created for the
   app.
5. token\_end\_point is the token endpoint of the Identity Provider (IdP) securing
   the service.

Returns `FALSE` if the mandatory access token is null.

## Usage

Use the `InitNativeApp()` function to register the access token required for a
Genero app to connect to a service (server-side) and for that service to connect as a client to
another protected service. This function checks whether the OAuth service has been initiated.

The primary role of `InitNativeApp()` is to register the access token with the
Genero Web Services (GWS). Once registered, you can call any of the OAuthAPI methods, such as
`CreateHTTPAuthorizationRequest`, to perform requests to the service or services. The
GWS will also be able to refresh the token automatically when it expires, eliminating the need to
restart the application.

Before calling `InitNativeApp()`, ensure that the access token and its expiration
date are set by calling [RetrievePasswordTokenForNativeApp](4954-oauthapi-retrievepasswordtokenfornativeapp.md "Returns the OAuth service access token via user credentials (username/password) and client credentials (client_id/secret_id). A refresh token allows the access token to be refreshed when it expires.") to obtain these from the Identity Provider. You must pass
the [OpenIdCResponseType](4939-oauthapi-openidcresponsetype.md "The OpenIdCResponseType record stores the access token, refresh token, and token expiry date retrieved in a request to the IdP.") record, which
contains the tokens, when calling `InitNativeApp()`.

If you need to retrieve metadata, you can call `FetchOpenIDMetadata()`, which will
save the metadata in an `OpenIDMetadataType` record.

In case of error, a `NULL` value will be returned.

## OAuthAPI.InitNativeApp function

```
IMPORT FGL OAuthAPI

DEFINE metadata OAuthAPI.OpenIDMetadataType
DEFINE tokens OAuthAPI.OpenIdCResponseType
DEFINE usr, pass STRING
DEFINE client_id, secret_id STRING
DEFINE idp_url STRING

MAIN

  # Enter the following information:
  LET idp_url=""   # The IdP's issuer URL
  LET usr=""       # A valid username
  LET pass=""      # Password for the above username
  LET client_id="" # The client ID of an application registered in IdP
  LET secret_id="" # The client secret from the same application

  TRY
      CALL OAuthAPI.FetchOpenIDMetadata(5, idp_url)
              RETURNING metadata
      IF metadata.issuer IS NULL THEN
              DISPLAY "Identity provider not available" 
              EXIT PROGRAM 1
      ELSE
          CALL OAuthAPI.RetrievePasswordTokenForNativeApp(5, metadata.token_endpoint, usr, pass, client_id, secret_id, NULL) 
              RETURNING tokens.*
          IF tokens.access_token IS NULL THEN
              DISPLAY "Unable to retrieve token"
              EXIT PROGRAM 1
          ELSE
              IF NOT OAuthAPI.InitNativeApp(5, tokens, client_id, secret_id, metadata.token_endpoint) THEN
                  DISPLAY "Cannot initiate refresh token service"
              END IF
          END IF
      END IF
  CATCH
      DISPLAY "ERROR: ", status, sqlca.sqlerrm
      EXIT PROGRAM 1
  END TRY

END MAIN
```
