---
title: "OAuthAPI.OpenIDMetadataType"
source: "fgl-topics/c_gws_oauthapi_OpenIDMetadataType.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.OpenIDMetadataType"
type: "concept"
---

# OAuthAPI.OpenIDMetadataType

> The OpenIDMetadataType record stores metadata retrieved in a request to the IdP.

## Syntax

```
TYPE OpenIDMetadataType RECORD
    issuer STRING, # REQUIRED
    authorization_endpoint STRING, # REQUIRED
    token_endpoint STRING, # REQUIRED
    userinfo_endpoint STRING,  # RECOMMENDED
    jwks_uri STRING, # REQUIRED
    registration_endpoint STRING, # RECOMMENDED
    scopes_supported DYNAMIC ARRAY OF STRING, # RECOMMENDED
    response_types_supported DYNAMIC ARRAY OF STRING, # REQUIRED
    response_modes_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    grant_types_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    acr_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    subject_types_supported DYNAMIC ARRAY OF STRING, # REQUIRED
    id_token_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # REQUIRED
    id_token_encryption_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    id_token_encryption_enc_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    userinfo_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    userinfo_encryption_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    userinfo_encryption_enc_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    request_object_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    request_object_encryption_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    request_object_encryption_enc_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    token_endpoint_auth_methods_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    token_endpoint_auth_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    display_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    claim_types_supported DYNAMIC ARRAY OF STRING, # OPTIONAL
    claims_supported DYNAMIC ARRAY OF STRING, # RECOMMENDED
    service_documentation STRING, # OPTIONAL
    claims_locales_supported STRING, # OPTIONAL
    ui_locales_supported STRING, # OPTIONAL
    claims_parameter_supported STRING, # OPTIONAL
    request_parameter_supported STRING, # OPTIONAL
    request_uri_parameter_supported STRING, # OPTIONAL
    require_request_uri_registration STRING, # OPTIONAL
    op_policy_uri STRING, # OPTIONAL
    op_tos_uri STRING, # OPTIONAL
    end_session_endpoint STRING # OPTIONAL
END RECORD
```

The `OpenIDMetadataType` record stores metadata retrieved from the IdP. For more
information on these parameters, refer to the [OpenID Connect Core specification](http://openid.net/specs/openid-connect-core-1_0-final.html) (external link).

## Usage

This type defines a record structure for the metadata retrieved in a requested to the IdP. For a
usage example, go to [OAuthAPI.GetOpenIDMetadata()](4943-oauthapi-getopenidmetadata.md "Get metadata from the Identity Provider for a service running on a Genero Application Server (GAS).").
