---
title: "Authenticate without GAS"
source: "fgl-topics/c_gws_oauthapi_without_gas.html"
breadcrumb: "Web services > Reference > OAuthAPI library > Overview > Authenticate without GAS"
type: "concept"
---

# Authenticate without GAS

> To authenticate an application that is not behind GAS, for example applications that include mobile, desktop, and Text User Interface (TUI) , you need to implement OAuth in a specific way.

In your application code, these functions need to be called in the following order:

1. `OAuthAPI.FetchOpenIDMetadata()` to retrieve the metadata from the Identity Provider
   from the required URL.
2. `OAuthAPI.RetrievePasswordToken()` to return the access token.
3. `OAuthAPI.initService()` to
   register the access token.

See the code sample in [OAuth access without GAS](4957-get-oauth-access-without-gas.md "This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.").
