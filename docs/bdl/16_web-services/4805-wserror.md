---
title: "WSError"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSError.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSError"
type: "concept"
---

# WSError

> Defines a description for an HTTP status code returned by the service.

## Syntax

```
WSError= "description"
```

1. description defines a short description as the
   `"Reason-Phrase"` for the status-code that is returned in the HTTP response message.

`WSError` is an optional attribute.

## Usage

You use this attribute in the management of application level errors in conjunction with [WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return."). When you use `WSThrows`
with the "code:@variable" option to return a description of
the status-code instead of the standard HTTP status code description, you must reference a variable
that has the `WSError` attribute.

You must set the attribute on a public variable defined at the modular level. You can define it
as any Genero BDL simple or complex type depending on your requirements for the detail about the
error. For example, if defined as a `RECORD` or `ARRAY`, by default,
it will be serialized in `JSON/XML` on the client side. If you define it as a simple
type, it will be serialized in `TEXT/PLAIN`.

You code in your REST function to set the variable at runtime with a specific reason for the HTTP
status code. The response is sent in a call to the [com.WebServiceEngine.SetRestError](../15_library-reference/3800-com-webserviceengine-setresterror.md "Manages error handling for a REST high-level Web Service function.") method referencing the variable.

When the client stub is generated from the OpenAPI documentation of your Web service, the correct
Genero BDL variable corresponding to the HTTP status code is deserialized on the client side.

## Example WSError in record variable

```
PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD
```

## Example WSError in integer variable

```
PUBLIC DEFINE fatalError INTEGER ATTRIBUTES(WSError = "fatal error")
```

## Related links

**Related concepts**  

[Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")
