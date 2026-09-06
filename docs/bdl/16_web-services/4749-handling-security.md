---
title: "Handling security"
source: "fgl-topics/c_gws_restful_high_level_security_handling.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling security"
type: "concept"
---

# Handling security

> You handle security in GWS high-level REST via scopes.

Security is configured based on the recommendations of the OpenAPI specification via scopes
implemented by OAuth, and Bearer token authentication. OAuth allows a user registered with an
Identity Provider to access the protected content from the resource server without sharing their
credentials with the service. Instead access is granted by access tokens forwarded to the GWS REST
service.

To implement security the Genero Application Server delegate service must be running. For more
information, see the How to implement delegation and Genero Identity Provider
(GIP) pages in the Genero Application Server User Guide.

If there is no access token, or there is no delegate service providing the
verification, the security function is not executed. Security is executed if, and only if, the
request contains a scope definition, and its value matches what is in the [WSScope (function)](4824-wsscope-function.md "Specifies the security scope required to access this REST function. Use this attribute to restrict access at the function level.") or [WSScope (module)](4807-wsscope-module.md "Defines the access scope required to call any REST operation in the module.") attributes.

Testing your services with WSScope:

- When testing your service in standalone mode without a GAS, the `WSScope` is not
  checked. However, when behind a GAS, the appropriate scope is required and you will need to deploy
  and secure the service with the Genero Identity Provider (GIP).
- Alternatively, if you need to integrate Genero REST services security into your own environment
  system, you can also write your own delegate service to validate any kind of token, extract the
  scope from it, and forward it to the REST service.

## Handling security errors

Access token errors are automatically handled by the GWS engine. You
do not need to do anything in your code. If the client request does not have the correct
access token, the service will return HTTP 403.

## Publishing scopes

When you generate the [service
description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML."), the scope you declared in the `WSScope` attribute, the user or
group member detail, is published in the "security" section for that function in the [OpenAPI description](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code."). If you declare
scopes with a modular variable with the `WSInfo` and `WSScope`
attributes, the "security" tag at the end of the specification file contains the security details.

## Related links

**Related concepts**  

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")

[OAuthAPI overview](4931-overview.md "The OAuthAPI library supports the OAuth protocol that authenticates user access and issues access tokens.")

## Child topics

- [Example: set security with WSScope](4750-set-security-with-wsscope.md): You can set security using the WSScope attribute either at the function level or at the service level.
