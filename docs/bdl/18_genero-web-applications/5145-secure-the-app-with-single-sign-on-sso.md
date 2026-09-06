---
title: "Secure the app with Single sign-on (SSO)"
source: "fgl-topics/c_fgl_gwa_application_configure_sso.html"
breadcrumb: "Genero Web applications > Secure the app with Single sign-on (SSO)"
type: "concept"
---

# Secure the app with Single sign-on (SSO)

> Enable SSO for your Genero Web Application (GWA) to increase security and makes it easier for users to sign in.

Genero provides SSO based on the OpenID Connect (OIDC) protocol via its Genero Identity Provider
(GIP). For more information on GIP and SSO, refer to the Single Sign-On User Guide. Any identity
provider that supports OIDC will work with GWA and can manage the creation of user accounts as well
as authentication and authorization during sign-in.

## Overview

When your GWA application is registered with an IdP for SSO, the end-user enters credentials once
in a login form on the browser to get authorization to use the application. If your application
performs further request to REST services, an access token is required. Your application must
initiate a request using a password-credential flow to the IdP to obtain an authorization and access
token. An example of the code demonstrating this process is provided in Typical BDL OAuth initialization.

## CORS constraints

If the IdP is on a different server than the GWA, you must address Cross-Origin Resource Sharing
(CORS) issues by adding the `Access-Control-Allow-Origin` HTTP header on the IdP
server. This allows the GWA to contact the IdP on a different host.

- **Implementation on GAS**: Add the `Access-Control-Allow-Origin` header to the
  `SERVICE (for HTTP)` entry in the as.xcf configuration file,
  specifying the hostname and port number of the IdP server. For example, if your IdP runs on a server
  called "cube", the entry should look like this:

  ```
  #...
  <SERVICE>
      <HEADER Name="Access-Control-Allow-Origin">https://cube:6394</HEADER>
  </SERVICE>
  #...
  ```

## OAuth API

In your GWA application, you must call functions of the [OAuthAPI library](../16_web-services/4930-oauthapi-library.md "The OAuthAPI library provides a set of functions for working with access to web services using the OAuth protocol.") in
order to authenticate and retrieve an access token to call additional REST services. See the example
in Typical BDL OAuth initialization.

Of course, you can also use fglrestful --auth yes to generate the stubs to
connect using OAuth protocol to any REST services. However, as GWA has a limited number of GWS
functions available at this time, you must use option `--legacyJSONApi` of the [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.") tool, and only requests with JSON or plain text are supported.
XML is not supported.

## Typical BDL OAuth initialization

In this sample, there are examples of the OAuthAPI calls you must make in your GWA application in
order to authenticate and get an access token to call REST services. You initialize the SSO user
identification with OAuthAPI:

```
IMPORT FGL OAuthAPI

-- IDP constants
CONSTANT idp_issuer = "https://host:port/gas/ws/r/services/GeneroIdentityProvider"
CONSTANT client_id = "XXXXX" -- OAuth client id as registered in your IdP
CONSTANT private_id = "ZZZZZ" -- OAuth private id as registered in your IdP

-- OAuthAPI
DEFINE metadata OAuthAPI.OpenIDMetadataType
DEFINE tokens OAuthAPI.OpenIdCResponseType
DEFINE user_login, user_pswd STRING
DEFINE s INTEGER
DEFINE r BOOLEAN

MAIN
  -- Initialize user_login and user_pswd (from a login dialog for example)
  -- CALL AskForLoginAndPassword() RETURNING s, user_login, user_pswd
  IF s<0 THEN EXIT PROGRAM 0 END IF

  -- Fetch IDP metadata
  CALL OAuthAPI.FetchOpenIDMetadata(5, idp_issuer)
     RETURNING metadata
  IF metadata.issuer IS NULL THEN
    DISPLAY "ERROR: Could not fetch OAuthAPI.OpenIDMetadataType"
    EXIT PROGRAM 1
  END IF

  -- Perform password credential flow 
  CALL OAuthAPI.RetrievePasswordTokenForNativeApp(5, metadata.token_endpoint,
                                                user_login, user_pswd,
                                                client_id,
                                                private_id,
                                                NULL)
     RETURNING tokens 
  -- Initialize OAuthAPI via returned tokens
  LET r = OAuthAPI.InitNativeApp(5, tokens,
                               client_id,
                               private_id,
                               metadata.token_endpoint)
  IF NOT r THEN
    DISPLAY "ERROR: Could not initialize native app with OAuthAPI.InitNativeApp"
    EXIT PROGRAM 1
  END IF
  -- Use any GWS client stub generated with fglrestful --oauth yes --legacyJSONApi
END MAIN
```

## Related links

**Related concepts**  

[OAuthAPI library](../16_web-services/4930-oauthapi-library.md "The OAuthAPI library provides a set of functions for working with access to web services using the OAuth protocol.")

[fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.")
