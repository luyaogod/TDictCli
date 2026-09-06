---
title: "OAuthAPI overview"
source: "fgl-topics/c_gws_oauthapi_overview.html"
breadcrumb: "Web services > Reference > OAuthAPI library > Overview"
type: "concept"
---

# OAuthAPI overview

> The OAuthAPI library supports the OAuth protocol that authenticates user access and issues access tokens.

The OAuth protocol is widely-used as a means of securing access to web services. It allows client
access by verifying the identity of the end user. Third party Identity Providers (IdP) usually
provide this service, or you can secure your RESTful services using the Genero Identity Provider
service that is delivered in $FGLDIR/web\_utilities/services/gip. For more
information, see the Genero Identity Provider (GIP) section in the Single Sign-On User Guide

The OAuthAPI library provides functions that enable you to retrieve metadata from the Identity
Provider, including endpoints, access tokens, scopes, user profiles, and more. It supports web
services that are started in the following contexts:

- **Behind a Genero Application Server (GAS).**
- **Not behind a GAS**, such as applications that include mobile, desktop, and Text User
  Interface (TUI) apps.
- **For a service (server-side) to connect as a client** to another service.

An application running behind a GAS implements authentication slightly differently than an
application not behind a GAS or in a service-to-service scenario. However, you will find that the
overall OAuth implementation generally follows the same pattern.

This process involves the following steps:

- The client calls the [OAuthAPI.init()](4940-oauthapi-init.md "To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server.") function to register the token for
  accessing the secure RESTful web service.
- Once OAuth is initialized, you can use functions to:
  - Create HTTP requests with the access token (for example, [CreateHTTPAuthorizationRequest()](4949-oauthapi-createhttpauthorizationrequest.md "Create an HttpRequest with OAuth access token.")).
  - Manage metadata, such as retrieving user information and supported scopes (for example, [GetOpenIDMetadata()](4943-oauthapi-getopenidmetadata.md "Get metadata from the Identity Provider for a service running on a Genero Application Server (GAS).")).

## Child topics

- [Authenticate with GAS](4932-authenticate-with-gas.md): To authenticate an application that is behind Genero Application Server, OAuth is implemented through delegation.
- [Authenticate without GAS](4933-authenticate-without-gas.md): To authenticate an application that is not behind GAS, for example applications that include mobile, desktop, and Text User Interface (TUI) , you need to implement OAuth in a specific way.
- [Authenticate a service to a service](4934-authenticate-a-service-to-a-service.md): To successfully authenticate a service (server side) to connect (as a service client) to another service protected by an access token, you will need to implement OAuth this way.
