---
title: "OAuthAPI.RegisterRequestType"
source: "fgl-topics/c_gws_oauthapi_RegisterRequestType.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.RegisterRequestType"
type: "concept"
---

# OAuthAPI.RegisterRequestType

> The RegisterRequestType record stores authorization credentials for the request.

## Syntax

```
TYPE RegisterRequestType RECORD
  redirect_uris DYNAMIC ARRAY OF STRING, # REQUIRED 
  response_types DYNAMIC ARRAY OF STRING, # OPTIONAL 
  grant_types DYNAMIC ARRAY OF STRING, # OPTIONAL
  client_name STRING,
  client_description STRING,
  scope STRING # OPTIONAL
END RECORD
```

These parameters are added to the query component of the authorization endpoint. For more
information on these parameters, refer to the [Authorization
Request](http://tools.ietf.org/html/rfc6749#section-4.1.1) (external link) section of the OAuth specification.

## Usage

This type defines a record structure to construct the request URI by adding the parameters to the
query component of the authorization endpoint URI.
