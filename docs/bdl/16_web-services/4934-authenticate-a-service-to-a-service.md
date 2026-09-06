---
title: "Authenticate a service to a service"
source: "fgl-topics/c_gws_oauthapi_oauth_service_to_service.html"
breadcrumb: "Web services > Reference > OAuthAPI library > Overview > Authenticate a service to a service"
type: "concept"
---

# Authenticate a service to a service

> To successfully authenticate a service (server side) to connect (as a service client) to another service protected by an access token, you will need to implement OAuth this way.

The access token got from the IdP to check access to the service, is also required to access the
client service.

1. Call the `OAuthAPI.initService()` function to initiate OAuth and to register the access
   token.
2. Then you can call any of the OAuthAPI methods, such as `OAuthAPI.CreateHTTPAuthorizationRequest()`, to perform requests to the other
   service.

See the code sample in `OAuthAPI.initService()`.
