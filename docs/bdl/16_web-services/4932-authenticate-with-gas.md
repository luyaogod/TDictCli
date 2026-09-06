---
title: "Authenticate with GAS"
source: "fgl-topics/c_gws_oauthapi_with_gas.html"
breadcrumb: "Web services > Reference > OAuthAPI library > Overview > Authenticate with GAS"
type: "concept"
---

# Authenticate with GAS

> To authenticate an application that is behind Genero Application Server, OAuth is implemented through delegation.

A Genero delegate service for OpenID Connect is delivered in
$FGLDIR/web\_utilities/services. This must be running to manage all delegated
requests for applications or services run on the GAS. For more information, see the pages on
How to implement delegation in the Genero Application Server User Guide.

You call the `OAuthAPI.init()` method
in your client app. This initializes OAuth by:

- Reading the GAS environment variable OIDC\_ACCESS\_TOKEN that contains the initial access token
  the app got from the Genero Identity Provider (GIP), or the third party Identity Provider.
- Initiating a global variable with that access token.

Then any REST request used in your application will automatically set the access token in a
HTTP request header and get access to the service. See the [Main program code for access to secure service](4788-code-to-access-secure-service.md "Code to get an access token for a secure RESTful Web service from a Genero application also secured.") topic.
