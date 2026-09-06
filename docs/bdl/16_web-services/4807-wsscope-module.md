---
title: "WSScope (module)"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSScope_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSScope (module)"
type: "concept"
---

# WSScope (module)

> Defines the access scope required to call any REST operation in the module.

## Syntax

```
WSScope = "{ scope } [,...]"
```

Where `WSScope` is a comma-separated list of scopes and where:

1. scope defines access permission for the resource.

`WSScope` is an optional attribute.

## Usage

You use this attribute to specify secure access via scopes forwarded from the Genero
Application Server to the GWS REST service.

Testing your services with WSScope:

- When testing your service in standalone mode without a GAS, the `WSScope` is not
  checked. However, when behind a GAS, the appropriate scope is required and you will need to deploy
  and secure the service with the Genero Identity Provider (GIP).
- Alternatively, if you need to integrate Genero REST services security into your own environment
  system, you can also write your own delegate service to validate any kind of token, extract the
  scope from it, and forward it to the REST service.

You can set the `WSScope` attribute in the [service information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.") of the
module or in the `ATTRIBUTES()` clause of the REST function. See [WSScope (function)](4824-wsscope-function.md "Specifies the security scope required to access this REST function. Use this attribute to restrict access at the function level.") for an example setting access
using scopes in the REST function.

## Example using WSScope in the information record at the service level

When the scope is set in the service information record, all REST functions in the Web
service are executed if, and only if, the request contains a scope definition that matches
the value in the `WSScope` attribute.

```
PUBLIC DEFINE serviceInfo RECORD ATTRIBUTES(WSInfo, WSScope="users.myservice")
  title STRING,
  description STRING,
  termOfService STRING,
  contact RECORD
    name STRING,
    url STRING,
    email STRING
    END RECORD,
    version STRING
  END RECORD = (
    title: "my service", 
    version: "1.0", 
    contact: ( email:"helpdesk@mysite.com") )
```

The main purpose of the service information record (defined with the
`WSInfo` attribute) is to document the REST service. If you are not setting
the scope at the modular level here, the record is still needed to provide service
information.

## How to determine the scope names

When determining the names for the scopes, it is important to understand the role of scopes. You
create a scope in the Identity Provider (IdP) system and assign it to a group or a user, so that the
user or group member will get an access token containing the scope name, and be allowed to access an
operation that specifies the same scope name using the `WSScope` attribute.

The names you choose can be a simple name (such as "`readonly`"), or it can use
dot notation (such as "`readonly.dev`" and "`readonly.user`") to
provide a logical hierarchy. The hierarchy approach is optional; it is purely provided to allow you
to organize your scopes in a logical manner that makes sense to you. Your end user would still need
to belong to groups that have either "`readonly.dev`" or
"`readonly.user`" scope assigned to them, the "`.dev`" and
"`.user`" extensions do not in themselves have any meaning to the IdP.

For example, you could create a service access list with the server name of
"`ReadOnly`", with two scopes defined: "`ReadOnly.dev`" and
"`ReadOnly.user`". You would then set the `ReadOnly.dev` scope to
the resources to be accessible to developers, and the `ReadOnly.user` scope to
the resources to be accessible to users only.

When determining scope names, it may also be helpful to think of the overall solution, which can
be a complex system with many services all working together. You can evaluate the access needs of
the various services and operations, and then identify the list of scopes that would allow you to
provide (or restrict) access to your various groups of users.

In summary, there is no restriction on the names you choose for scopes. You can set them as you
wish, depending on what you want to achieve.

## Related links

**Related concepts**  

[Handling security](4749-handling-security.md "You handle security in GWS high-level REST via scopes.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")
