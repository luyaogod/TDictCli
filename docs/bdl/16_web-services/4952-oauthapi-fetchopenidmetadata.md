---
title: "OAuthAPI.FetchOpenIDMetadata()"
source: "fgl-topics/c_gws_oauthapi_FetchOpenIDMetadata.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.FetchOpenIDMetadata()"
type: "concept"
---

# OAuthAPI.FetchOpenIDMetadata()

> Fetch metadata from the Identity Provider at the URL provided.

## Syntax

```
FUNCTION FetchOpenIDMetadata(
   timeout INTEGER, 
   idp STRING )
RETURNS OAuthAPI.OpenIDMetadataType
```

1. timeout defines the number of seconds.
2. idp is the URL of the Identity Provider (IdP).

`NULL` may be returned if metadata is not found.

## Usage

Use the `FetchOpenIDMetadata()` function to retrieve the metadata, such as
registration endpoint, supported scopes, user information endpoints, and so on, provided by the
Identity Provider for access to a secure RESTful web service. The information is retrieved in a
request to the IdP via the URL provided and stored in the `OpenIDMetadataType`
record.

If the service is started behind the GAS, you should call the [GetOpenIDMetadata()](4943-oauthapi-getopenidmetadata.md "Get metadata from the Identity Provider for a service running on a Genero Application Server (GAS).") function instead.

In case of error, a `NULL` value will be returned.

## OAuthAPI.FetchOpenIDMetadata function

```
IMPORT FGL OAuthAPI

PRIVATE DEFINE metadata OAuthAPI.OpenIDMetadataType
DEFINE idp_url STRING 

MAIN
  # ...
     CALL OAuthAPI.FetchOpenIDMetadata(20, idp_url)
          RETURNING metadata.*
     IF metadata.issuer IS NULL THEN
         ERROR "IdP not available"
     ELSE
       DISPLAY "Registration endpoint is:", metadata.registration_endpoint
     END IF
  # ... 
  
END MAIN
```

## Related links

**Related concepts**  

[OAuth access without GAS](4957-get-oauth-access-without-gas.md "This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.")
