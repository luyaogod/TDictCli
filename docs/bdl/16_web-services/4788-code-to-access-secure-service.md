---
title: "Main program code for access to secure service"
source: "fgl-topics/t_gws_client_configure_oauth.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Create the client application > Write the MAIN program code block > Code to access secure service"
type: "task"
---

# Main program code for access to secure service

> Code to get an access token for a secure RESTful Web service from a Genero application also secured.

Genero RESTful Web services and Genero applications may be secured by either a third-party
Identity Provider (IdP) or the Genero Identity Provider (GIP).

When you configure a Genero application protected by an IdP, you must use delegation in the
application configuration. In the `DELEGATE` element you provide
details of the IdP that authenticates access as described in step 1.

If the application is not secured by an IdP, delegation is not required.

In this task as the application accesses (as client) a RESTful Web service that is protected by
an IdP, you must code to get access tokens from the IdP at runtime. The `OAuthAPI.init()` function gets the access tokens for the service and
registers them with the GWS engine.

## Steps

1. Set your secured app to use delegation in its application configuration file. 

   This step is mandatory.

   ```
   <APPLICATION Parent="defaultwa" ...>
     <EXECUTION>
      ...
       <DELEGATE service="services/OpenIDConnectServiceProvider">
          <IDP>IdP_URL</IDP>
          <CLIENT_PUBLIC_ID>XXXXXXXX</CLIENT_PUBLIC_ID>
          <CLIENT_SECRET_ID>XXXXXX-XXXXXX</CLIENT_SECRET_ID> 
       </DELEGATE>
      ...
     </EXECUTION>
   </APPLICATION>
   ```

   Where:
   - The `OpenIDConnectServiceProvider` is the delegation REST Web service in the
     $FGLDIR.
   - The IdP\_URL can have an entry of `localhost` when everything
     runs on the same Genero Application Server. Otherwise, you must provide the Genero Identity Provider
     (GIP) URL. For example:
     http://othermachine.com/ws/r/services/GeneroIdentityProvider
   - The OAuth access tokens for `CLIENT_PUBLIC_ID` and
     `CLIENT_SECRET_ID` are those you get from the IdP.

     For further information, see the
     Configure delegation for application or service page in the Genero Application Server User Guide.
2. In the `MAIN/END MAIN` clause of your client app, call the
   `OAuthAPI.init` function to get the OAuth access tokens at runtime. This must be done
   before calling any other service functions.

   For
   example:

   ```
   IMPORT FGL OAuthAPI

   DEFINE my_user_id STRING
   MAIN
     # ...
      
     # Init OAuthAPI
     IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr") THEN
       DISPLAY "Error: unable to initialize OAuth"
       EXIT PROGRAM 1
     ELSE
       LET my_user_id = OAuthAPI.getIDSubject()
     END IF

     # ... 
     
   END MAIN
   ```

   You
   can get user information coming from the IdP from variables with the prefix `OIDC_`.
   For example,

   ```
   LET userEmail = fgl_getenv("OIDC_EMAIL")
   ```

   For an example of `OAuthAPI` calls, see the `consoleApp` source
   in FGLDIR\web\_utilities\services\gip\src\console.

   When using a third-party IdP, if it supports OpenID Connect, then the `OAuthAPI`
   can be used the same as for Genero Identity Provider. For further information, see the OpenID
   Connect SSO pages in the Genero Application Server User Guide.

## Related links

**Related concepts**  

[OAuthAPI library](4930-oauthapi-library.md "The OAuthAPI library provides a set of functions for working with access to web services using the OAuth protocol.")

[Access to secure web service](4787-access-to-secure-web-service.md "To access a secure RESTful Web service, the client application must have a valid access token.")

[Write the MAIN program code block](4786-write-the-main-program-code-block.md "Create the MAIN program block that imports the client stub and calls its functions.")
