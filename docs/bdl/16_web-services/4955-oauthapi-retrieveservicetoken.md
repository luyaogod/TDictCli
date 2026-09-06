---
title: "OAuthAPI.RetrieveServiceToken()"
source: "fgl-topics/c_gws_oauthapi_RetrieveServiceToken.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.RetrieveServiceToken()"
type: "concept"
---

# OAuthAPI.RetrieveServiceToken()

> Return the OAuth service access token via client app credentials.

## Syntax

```
FUNCTION RetrieveServiceToken(
   timeout INTEGER,
   TokenServiceURL STRING, 
   client_id STRING, 
   secret_id STRING,
   scope STRING )
RETURNS STRING, INTEGER
```

1. timeout defines the number of seconds.
2. TokenServiceURL is the token endpoint of the
   Identity Provider (IdP) securing the service.
3. client\_id is the application ID assigned to the app when
   registered.
4. client\_secret is the application secret created for the
   app.
5. scope is a space-separated list of scopes defining user
   access.

Returns a valid access token and when it expires in seconds. `NULL` may be
returned if the access token is not available.

## Usage

Use this function to retrieve a valid access token for a client app accessing a RESTful web
service using the client app's own client id and client secret credentials.

In case of error, a `NULL` value will be returned.

## OAuthAPI.RetrieveServiceToken function

```
IMPORT FGL OAuthAPI

MAIN

DEFINE metadata OAuthAPI.OpenIDMetadataType
DEFINE token  STRING
DEFINE expire INTEGER
DEFINE client_id STRING
DEFINE secret_id STRING
DEFINE scope STRING
DEFINE idp_url STRING

TRY
     # ...
     CALL OAuthAPI.FetchOpenIDMetadata(20, idp_url)
          RETURNING metadata.*
     IF metadata.issuer IS NULL THEN
         ERROR "IdP not available" 
         EXIT PROGRAM 1
     ELSE
        CALL OAuthAPI.RetrieveServiceToken(5, metadata.token_endpoint, client_id, secret_id, scope ) 
           RETURNING token, expire
        IF token IS NULL THEN
           DISPLAY "Unable to retrieve token"
           EXIT PROGRAM 1
        ELSE
            DISPLAY "Access token value :",token 
            DISPLAY SFMT("Token expires in %1 seconds",expire)
        END IF
     END IF
CATCH
   DISPLAY "ERROR : ",status,sqlca.sqlerrm
   EXIT PROGRAM 1
END TRY

END MAIN
```
