---
title: "WSScope (function)"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSScope_function.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > WSScope (function)"
type: "concept"
---

# WSScope (function)

> Specifies the security scope required to access this REST function. Use this attribute to restrict access at the function level.

## Syntax

```
WSScope = "{ scope } [,...]"
```

Where `WSScope` is a comma-separated list of scopes and where:

1. scope defines access permission for the resource.

`WSScope` is an optional attribute.

## Usage

The `WSScope` attribute defines the scope required to call this REST
function. The scope value is forwarded from the Genero Application Server and is compared against
the scope names specified on the function.

Testing your services with WSScope:

- When testing your service in standalone mode without a GAS, the `WSScope` is not
  checked. However, when behind a GAS, the appropriate scope is required and you will need to deploy
  and secure the service with the Genero Identity Provider (GIP).
- Alternatively, if you need to integrate Genero REST services security into your own environment
  system, you can also write your own delegate service to validate any kind of token, extract the
  scope from it, and forward it to the REST service.

You set `WSScope` in either the `ATTRIBUTES()` clause of the
function or in the module’s [service
information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."). A scope defined at the function level applies only to that operation and
overrides the module-level scope.

Use function-level scopes when specific operations require different permissions than the rest of
the service.

## Example 1: Setting security with WSScope at function level

In this sample REST function there is an example of a function that requires authentication
to access it. To execute this REST operation requires the request contains an access token
with a scope that matches what is in `WSScope`.

The `WSScope` attribute is set in the `ATTRIBUTES` clause of
the function. In this example the scope is set to "profile" or "profile.me".

Access token errors are automatically handled by the GWS engine. You
do not need to do anything in your code. If the client request does not have the correct
access token, the service will return HTTP 403.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

TYPE profileType RECORD
     id INTEGER,
     name VARCHAR(100),
     email VARCHAR(255)
     # ...
   END RECORD

PUBLIC FUNCTION FetchMyUserProfile( id INTEGER ATTRIBUTES(WSQuery) )
  ATTRIBUTES(
    WSGet,
    WSPath = "/users/profile",
    WSDescription = "Returns a user profile, requires authentication",
    WSThrows = "404:user not found",
    WSScope = "profile, profile.me")
  RETURNS profileType ATTRIBUTES(WSName = "data",
                                 WSMedia = "application/json,application/xml")
    DEFINE p profileType
    TRY
      SELECT * INTO p.* FROM users
             WHERE @id = id
      IF sqlca.sqlcode = NOTFOUND THEN
        CALL com.WebServiceEngine.SetRestError(404,NULL)
      END IF
    CATCH
       CALL com.WebServiceEngine.SetRestError(505,NULL)
    END TRY
    RETURN p
END FUNCTION
```

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
