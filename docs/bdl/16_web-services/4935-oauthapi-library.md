---
title: "OAuthAPI: library"
source: "fgl-topics/r_gws_oauthapi_apis.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library"
type: "reference"
---

# OAuthAPI: library

> The OAuthAPI library provides functions and types for working with OAuth such as initializing OAuth, getting metadata, and managing requests using access tokens for web services in different situations.

| Types | Description |
| --- | --- |
| TYPE RegisterRequestType RECORD redirect_uris DYNAMIC ARRAY OF STRING, # REQUIRED response_types DYNAMIC ARRAY OF STRING, # OPTIONAL grant_types DYNAMIC ARRAY OF STRING, # OPTIONAL client_name STRING, client_description STRING, scope STRING # OPTIONAL END RECORD | The `RegisterRequestType` record stores authorization credentials for the request. |
| TYPE RegisterResponseType RECORD client_id STRING, # REQUIRED client_secret STRING, # OPTIONAL grant_types DYNAMIC ARRAY OF STRING, # OPTIONAL redirect_uris DYNAMIC ARRAY OF STRING, # OPTIONAL client_name STRING, client_description STRING, scope STRING # OPTIONAL END RECORD | The `RegisterResponseType` record stores authorization credentials for the response. |
| TYPE OpenIDMetadataType RECORD issuer STRING, # REQUIRED authorization_endpoint STRING, # REQUIRED token_endpoint STRING, # REQUIRED userinfo_endpoint STRING, # RECOMMENDED jwks_uri STRING, # REQUIRED registration_endpoint STRING, # RECOMMENDED scopes_supported DYNAMIC ARRAY OF STRING, # RECOMMENDED response_types_supported DYNAMIC ARRAY OF STRING, # REQUIRED response_modes_supported DYNAMIC ARRAY OF STRING, # OPTIONAL grant_types_supported DYNAMIC ARRAY OF STRING, # OPTIONAL acr_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL subject_types_supported DYNAMIC ARRAY OF STRING, # REQUIRED id_token_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # REQUIRED id_token_encryption_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL id_token_encryption_enc_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL userinfo_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL userinfo_encryption_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL userinfo_encryption_enc_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL request_object_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL request_object_encryption_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL request_object_encryption_enc_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL token_endpoint_auth_methods_supported DYNAMIC ARRAY OF STRING, # OPTIONAL token_endpoint_auth_signing_alg_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL display_values_supported DYNAMIC ARRAY OF STRING, # OPTIONAL claim_types_supported DYNAMIC ARRAY OF STRING, # OPTIONAL claims_supported DYNAMIC ARRAY OF STRING, # RECOMMENDED service_documentation STRING, # OPTIONAL claims_locales_supported STRING, # OPTIONAL ui_locales_supported STRING, # OPTIONAL claims_parameter_supported STRING, # OPTIONAL request_parameter_supported STRING, # OPTIONAL request_uri_parameter_supported STRING, # OPTIONAL require_request_uri_registration STRING, # OPTIONAL op_policy_uri STRING, # OPTIONAL op_tos_uri STRING, # OPTIONAL end_session_endpoint STRING # OPTIONAL END RECORD | The `OpenIDMetadataType` record stores metadata retrieved in a request to the IdP. |
| TYPE OpenIdCResponseType RECORD access_token STRING, token_type STRING, expires_in INTEGER, refresh_token STRING END RECORD | The `OpenIdCResponseType` record stores the access token, refresh token, and token expiry date retrieved in a request to the IdP. |

| Function | Description |
| --- | --- |
| FUNCTION Init( cnx_timeout INTEGER, client_id STRING, client_secret STRING ) RETURNS BOOLEAN | To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server. |
| FUNCTION InitService( cnx_timeout INTEGER, access_token STRING ) RETURNS BOOLEAN | To be called in a Genero web service started via OpenID Connect/OAuth2 accessing another secure RESTful web service as a client. |
| FUNCTION InitNativeApp( cnx_timeout INTEGER, tokens OpenIdCResponseType, client_id STRING, client_secret STRING, token_end_point STRING) RETURNS BOOLEAN | To be called in a Genero application accessing a secure RESTful web service directly (not behind a Genero Application Server). |
| FUNCTION GetOpenIDMetadata() RETURNS OAuthAPI.OpenIDMetadataType | Get metadata from the Identity Provider for a service running on a Genero Application Server (GAS). |
| FUNCTION GetIDPIssuer() RETURNS STRING | Get endpoint of the Identity Provider. |
| FUNCTION GetIdRoles() RETURNS DYNAMIC ARRAY OF STRING | Get OAuth ID Token authorization roles. |
| FUNCTION GetIDScopes() RETURNS DYNAMIC ARRAY OF STRING | Get OAuth ID Token authorization scopes. |
| FUNCTION GetIDSubject() RETURNS STRING | Get OAuth subject identifier of ID Token. |
| FUNCTION GetMyAccessToken() RETURNS STRING | Get a valid access token. |
| FUNCTION CreateHTTPAuthorizationRequest( url STRING ) RETURNS com.HttpRequest | Create an HttpRequest with OAuth access token. |
| FUNCTION RetryHTTPRequest( resp com.HttpResponse ) RETURNS BOOLEAN | Retry an HttpRequest with OAuth access token to check if the access token has expired. |
| FUNCTION ExtractTokenFromHTTPRequest( req com.HttpServiceRequest) RETURNS STRING | Return the OAuth access token from a HTTP request service object. |
| FUNCTION FetchOpenIDMetadata( timeout INTEGER, idp STRING ) RETURNS OAuthAPI.OpenIDMetadataType | Fetch metadata from the Identity Provider at the URL provided. |
| FUNCTION RetrievePasswordToken( timeout INTEGER, TokenServiceURL STRING, usr STRING, pass STRING, scope STRING ) RETURNS ( STRING, INTEGER ) | Return the OAuth service access token via user name and password. |
| FUNCTION RetrievePasswordTokenForNativeApp( timeout INTEGER, TokenServiceURL STRING, username STRING, password STRING, client_id STRING, client_secret STRING, scope STRING) RETURNS OpenIdCResponseType | Returns the OAuth service access token via user credentials (username/password) and client credentials (client\_id/secret\_id). A refresh token allows the access token to be refreshed when it expires. |
| FUNCTION RetrieveServiceToken( timeout INTEGER, TokenServiceURL STRING, client_id STRING, secret_id STRING, scope STRING ) RETURNS STRING, INTEGER | Return the OAuth service access token via client app credentials. |
