---
title: "OAuthAPI.RegisterResponseType"
source: "fgl-topics/c_gws_oauthapi_RegisterResponseType.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.RegisterResponseType"
type: "concept"
---

# OAuthAPI.RegisterResponseType

> The RegisterResponseType record stores authorization credentials for the response.

## Syntax

```
TYPE RegisterResponseType RECORD
  client_id STRING, # REQUIRED
  client_secret STRING, # OPTIONAL
  grant_types DYNAMIC ARRAY OF STRING, # OPTIONAL
  redirect_uris DYNAMIC ARRAY OF STRING, # OPTIONAL
  client_name STRING,
  client_description STRING,
  scope STRING # OPTIONAL
END RECORD
```

The system adds these parameters to the query component of the authorization endpoint when the
resource owner grants the access request. For more information on these parameters, refer to the
[Authorization Response](http://tools.ietf.org/html/rfc6749#section-4.1.2) (external link) section of the OAuth specification.

## Usage

This type defines a record structure to construct the query component of the authorization
endpoint when the resource owner grants the access request.
