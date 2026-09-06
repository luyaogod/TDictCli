---
title: "OAuthAPI.RetrievePasswordToken()"
source: "fgl-topics/c_gws_oauthapi_RetrievePasswordToken.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.RetrievePasswordToken()"
type: "concept"
---

# OAuthAPI.RetrievePasswordToken()

> Return the OAuth service access token via user name and password.

## Syntax

```
FUNCTION RetrievePasswordToken(
   timeout INTEGER,
   TokenServiceURL STRING, 
   usr STRING, 
   pass STRING,
   scope STRING )
RETURNS ( STRING, INTEGER )
```

1. timeout defines the number of seconds.
2. TokenServiceURL is the token endpoint of the
   Identity Provider (IdP) securing the service.
3. usr is the user's login details.
4. pass is the user password.
5. scope is a space-separated list of scopes defining user
   access.

Returns a valid access token and when it expires in seconds. `NULL` may be
returned if the access token is not available.

## Usage

Use the `RetrievePasswordToken()` function to retrieve a valid access token for a
client app accessing a RESTful Web service using a username and password.

In case of error, a `NULL` value will be returned.

## OAuthAPI.RetrievePasswordToken function

```
IMPORT FGL OAuthAPI

DEFINE metadata OAuthAPI.OpenIDMetadataType
DEFINE token  STRING
DEFINE expire INTEGER
DEFINE usr, pass, scope STRING
DEFINE idp_url STRING

MAIN

   TRY
      # ...
      CALL OAuthAPI.FetchOpenIDMetadata(20, idp_url)
            RETURNING metadata.*
      IF metadata.issuer IS NULL THEN
            ERROR "IdP not available" 
            EXIT PROGRAM 1
      ELSE
         CALL OAuthAPI.RetrievePasswordToken(5, metadata.token_endpoint, usr, pass, scope) 
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

## Related links

**Related concepts**  

[OAuthAPI.RetrieveServiceToken()](4955-oauthapi-retrieveservicetoken.md "Return the OAuth service access token via client app credentials.")

[OAuth access without GAS](4957-get-oauth-access-without-gas.md "This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.")
