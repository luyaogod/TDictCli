---
title: "Handling application level errors"
source: "fgl-topics/c_gws_restful_high_level_error_handling.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling application level errors"
type: "concept"
---

# Handling application level errors

> There are many situations in which you need to notify an error to a client using your REST web service.

Errors in calls to Web services may occur for various reasons; for example:

- the user input is incorrect
- the resource does not exist
- the server response failed
- the client does not have access to the resource
- the database connection failed

In these cases, you normally return an HTTP status code in the range 400 to 599 with its
default message. The most common ones you encounter are as follows:

- client (400-499)
  - 400 Bad Request or No matching Rest operation found
  - 401 Unauthorized
  - 403 Forbidden
  - 404 Not Found
- server (500-599)
  - 500 Internal Server Error
  - 501 Not Implemented

For information on the full range of HTTP status codes see [RFC 2616](http://tools.ietf.org/html/rfc2616).

There are some situations where it is more useful to describe the error specifically, and in your
REST function you can use the error handling attributes for this purpose.

## Handling errors with attributes

In GWS REST you can perform error handling using two attributes:

- with the `WSThrows`
  attribute, you list errors that may be encountered accessing a resource. In the client stub,
  generated from the OpenAPI documentation, code handles errors for the HTTP status codes defined in
  `WSThrows`.
- with the `WSError`
  attribute, you can provide a description of the status-code to replace the standard HTTP status code
  description.

In your function, you code to trap an error at runtime with a call to the `SetRestError()` method, to return
the HTTP status code and a description to the client.

You can define a custom header to return the error response of the web service function. For more
details, see [WSErrorHeader](4831-wserrorheader.md "Defines a custom HTTP header to include in an error response.").

## Publishing errors

When you generate the service description, the errors you declared in the
`WSThrows` attribute, the code reference and the error detail, are published in the
"responses" section for that function in the OpenAPI documentation. The modular variables declared
with `WSError` attribute are found in the "component" section at the end of the
documentation file.

## Related links

**Related concepts**  

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")

[Troubleshooting no matching REST operation](4791-troubleshooting.md "Understanding how the GWS finds an operation to use for a request can help in troubleshooting unexpected errors.")

## Child topics

- [Example: handling expected application errors](4747-handling-expected-errors.md): Error handling is supported with the WSThrows, WSError, and WSErrorHeader attributes.
- [Example: handling an unexpected error](4748-handling-unexpected-errors.md): Example of a method you can use to code for unexpected errors.
