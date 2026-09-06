---
title: "OAuthAPI.GetOpenIDMetadata()"
source: "fgl-topics/c_gws_oauthapi_GetOpenIDMetadata.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.GetOpenIDMetadata()"
type: "concept"
---

# OAuthAPI.GetOpenIDMetadata()

> Get metadata from the Identity Provider for a service running on a Genero Application Server (GAS).

## Syntax

```
FUNCTION GetOpenIDMetadata()
RETURNS OAuthAPI.OpenIDMetadataType
```

`NULL` may be returned if metadata is not found.

## Usage

Use the `GetOpenIDMetadata()` function to retrieve metadata—such as the
registration endpoint, supported scopes, and user information endpoints—provided by the Genero
Identity Provider for accessing a secure RESTful web service. A request is made to obtain this
metadata, which is then stored in an `OpenIDMetadataType` record.

If the service is not started behind the GAS, you should call the [FetchOpenIDMetadata()](4952-oauthapi-fetchopenidmetadata.md "Fetch metadata from the Identity Provider at the URL provided.") function instead.

In case of error, a `NULL` value will be returned.

## OAuthAPI.GetOpenIDMetadata function

```
IMPORT FGL OAuthAPI

DEFINE metadata OAuthAPI.OpenIDMetadataType
DEFINE ind INTEGER

MAIN
  # ...
   CALL OAuthAPI.GetOpenIDMetadata() RETURNING metadata.*
   IF metadata.issuer IS NULL THEN
       ERROR "IdP not available"
   ELSE
     DISPLAY "Registration endpoint is:", metadata.registration_endpoint
     FOR ind = 1 TO metadata.scopes_supported.getLength()
       DISPLAY "Scope is: ", metadata.scopes_supported[ind]
     END FOR
   END IF
   # ... 
  
END MAIN
```

## Related links

**Related concepts**  

[OAuthAPI.FetchOpenIDMetadata()](4952-oauthapi-fetchopenidmetadata.md "Fetch metadata from the Identity Provider at the URL provided.")

[OAuthAPI.OpenIDMetadataType](4938-oauthapi-openidmetadatatype.md "The OpenIDMetadataType record stores metadata retrieved in a request to the IdP.")
