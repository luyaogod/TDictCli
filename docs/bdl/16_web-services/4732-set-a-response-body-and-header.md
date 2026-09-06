---
title: "Set a response body and header"
source: "fgl-topics/c_gws_restful_high_level_response_body.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Set a response body and header"
type: "concept"
---

# Set a response body and header

> You specify a response body in a return parameter without an attribute. Other return values can be sent in headers, using the WSHeader attribute.

> **Important:**
>
> A message body in the response is required when you
> perform an HTTP GET, POST, PUT, DELETE operation on a resource, otherwise the response
> results in the [error-9106](../15_library-reference/4483-genero-bdl-errors.md).

## Example responses in header and body

In this sample REST function data is returned in a header and in
the message body. The function's `RETURNS` clause has two return values:

- An integer is returned in a header. It is specified with the `WSHeader`
  attribute.
- A string is returned in the body. It is specified without an attribute.

```
PUBLIC FUNCTION help()
  ATTRIBUTES (WSGet,
              WSPath = "/help")
  RETURNS (INTEGER ATTRIBUTE(WSHeader, WSDescription = "Reference number"),
    STRING)
    RETURN 3, "Hello world"
END FUNCTION
```

![Sample output of the HTTP response](../_images/rest_response_head_body_output.png)

*Output of the HTTP response*

In the output the header is given a default name, "rv0",
at runtime. You can change default header naming via the [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") attribute.
> **Warning:**
>
> **Standard response headers**
>
> Setting
> a standard [HTTP header](http://www.iana.org/assignments/message-headers/message-headers.xhtml) on a response must be handled with care,
> especially for those that define the response body such as `Content-Type`, or
> `Content-Encoding`. Make sure what you define with `WSName` does not
> conflict with what is specified in the OpenAPI documentation for the service.

## Related links

**Related concepts**  

[HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.")

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")
