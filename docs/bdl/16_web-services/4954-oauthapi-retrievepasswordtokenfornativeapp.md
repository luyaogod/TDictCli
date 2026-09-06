---
title: "OAuthAPI.RetrievePasswordTokenForNativeApp()"
source: "fgl-topics/c_gws_oauthapi_RetrievePasswordTokenForNativeApp.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.RetrievePasswordTokenForNativeApp()"
type: "concept"
---

# OAuthAPI.RetrievePasswordTokenForNativeApp()

> Returns the OAuth service access token via user credentials (username/password) and client credentials (client_id/secret_id). A refresh token allows the access token to be refreshed when it expires.

## Syntax

```
FUNCTION RetrievePasswordTokenForNativeApp(
  timeout INTEGER,
  TokenServiceURL STRING,
  username STRING,
  password STRING,
  client_id STRING,
  client_secret STRING,
  scope STRING)
RETURNS OpenIdCResponseType
```

1. timeout defines the number of seconds.
2. TokenServiceURL is the token endpoint of the
   Identity Provider (IdP) securing the service.
3. username. This is the user's login details.
4. password. This is the user password.
5. client\_id is the application ID assigned to the app when
   registered.
6. client\_secret is the application secret created for the
   app.
7. scope is a space-separated list of scopes defining user
   access.

Returns an [OpenIdCResponseType](4939-oauthapi-openidcresponsetype.md "The OpenIdCResponseType record stores the access token, refresh token, and token expiry date retrieved in a request to the IdP.") record with the access token, refresh token, and access token expiration
date. `NULL` may be returned if the access token is not available.

## Usage

Use the `RetrievePasswordTokenForNativeApp()` function to obtain an access token
that will be automatically refreshed when it expires, allowing the user to continue using the
application without needing to restart it. This function is particularly useful for a Genero app
that operates outside of a Genero Application Server while accessing a secure RESTful web
service.

The process of obtaining the token is similar to that of `RetrievePasswordToken`
or `RetrieveServiceToken`. However, in this case, you must provide both user
credentials (username and password) and client credentials (client\_id and secret\_id) as parameters
in the call.

The access token, access token expiration date, and refresh token returned are stored in an
`OpenIdCResponseType` record. This record must be passed to the `InitNativeApp()` function to initiate the
service.

In case of error, a `NULL` value will be returned.

## OAuthAPI.RetrievePasswordTokenForNativeApp function

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

## Related links

**Related concepts**  

[OAuthAPI.RetrieveServiceToken()](4955-oauthapi-retrieveservicetoken.md "Return the OAuth service access token via client app credentials.")

[OAuth access without GAS](4957-get-oauth-access-without-gas.md "This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.")
