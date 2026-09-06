---
title: "WSErrorHeader"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSErrorHeader.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > Error handling attributes > WSErrorHeader"
type: "concept"
---

# WSErrorHeader

> Defines a custom HTTP header to include in an error response.

## Syntax

```
WSErrorHeader= "{ code } [,...]"
```

1. code defines a comma-separated list of HTTP status error codes may be
   returned when the GWS server generates a response with an error code. Codes between 400 to 599
   represent HTTP error responses.

`WSErrorHeader` is an optional attribute.

## Usage

You use this attribute in a REST web service function to dynamically return a HTTP error code in
a header instead of the body when the GWS server generates a response with an error code. By
default, the REST operation returns HTTP status codes in the body, but this method allows you to set
the return code in a custom header.

The following are conditions of use:

- You must use `WSErrorHeader` on a return parameter only; otherwise, error [-9143](../15_library-reference/4483-genero-bdl-errors.md) is thrown.
- You must use `WSErrorHeader` in conjunction with [WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return."). The `WSErrorHeader`
  attribute requires the `WSThrows` attribute in your function; otherwise, error [-9145](../15_library-reference/4483-genero-bdl-errors.md) is thrown.
- You must use `WSErrorHeader` with `WSThrows` as follows:
  - Specify `WSErrorHeader` without an error code value to send a header for all
    values ​​present in `WSThrows`.
  - Specify `WSErrorHeader` with one or more error code values ​​present in
    `WSThrows` to send a header for only those values.The error code sent can only be for error codes specified in `WSThrows`;
  otherwise, error [-9146](../15_library-reference/4483-genero-bdl-errors.md) is
  thrown.

You can also use `WSErrorHeader` in conjunction with the [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute. For example:

- You define `WSHeader` to send a custom header when the status is a success (codes
  between 200 to 399)
- Whereas you define the `WSErrorHeader` to send a custom header when the status is
  an error code (codes between 400 to 599).

This example illustrates how you may use `WSHeader` on its own and together with
`WSErrorHeader`:

```
# ...
RETURNS(STRING ATTRIBUTE (WSHeader, WSName = "X-Header-Token"), 
   STRING ATTRIBUTE(WSHeader,WSErrorHeader = "400", WSName = "Access-Control-Allow-Methods"))
# ...
```

- The header "`X-Header-Token`" will only be returned when the status code
  indicates a success. The header will not be returned in the case of a response with error 400.
- The "`Access-Control-Allow-Methods`" header will be returned in the case of a 200
  response and also if the response is 400.

In your function, you code to trap an error at runtime with a call to the [SetRestError()](../15_library-reference/3800-com-webserviceengine-setresterror.md "Manages error handling for a REST high-level Web Service function.") method, to return the HTTP
status code and a description to the client.

## Example WSErrorHeader content error

In this sample
REST function status codes are returned in headers. The `WSThrows` attribute is set
to handle errors that may be encountered. It has options to respond to two HTTP errors:

- Error 400 when the resource is not found
- Error 401 when access is unauthorized

The `test` input parameter is checked. Depending on the value, an error is
thrown and returned in the response with an error description provided in the
`WSError` attributed variable referenced in `@myError` in
`WSThrows`.

Response headers are sent depending on the generated status code:

- The "`Access-Control-Allow-Methods`" header is set for both
  `WSHeader` and `WSErrorHeader`; the header will be returned in the
  case of both a 200 response and also if the response is 400.
- The `"Access-Control-Allow-Origin"` header set for `WSHeader` is
  sent in the case of a 200 response.
- The "`ERROR`" header set for `WSErrorHeader` is sent if the
  response is 400 or 401.

A call to [SetRestError()](../15_library-reference/3800-com-webserviceengine-setresterror.md "Manages error handling for a REST high-level Web Service function.")
returns the error message in the specified header:

```
IMPORT com

PUBLIC DEFINE myError RECORD ATTRIBUTE(WSError = "Service Error")
    code INTEGER,
    message STRING
END RECORD

PUBLIC FUNCTION test_NoContentError(
    test INTEGER ATTRIBUTE(WSQuery))
    ATTRIBUTE(WSGet,
        WSPath = "/nocontent/error",
        WSDescription = "test with WSErrorHeader (return 400)",
        WSThrows = "400:@myError,401")
    RETURNS(STRING ATTRIBUTE(WSHeader,
            WSErrorHeader = "400",
            WSName = "Access-Control-Allow-Methods"),
        STRING ATTRIBUTE(WSHeader, WSName = "Access-Control-Allow-Origin"),
        INTEGER ATTRIBUTE(WSErrorHeader,
            WSDescription = "in case error",
            WSName = "ERROR"))

    DEFINE method, origin STRING
    DEFINE err INTEGER

    LET method = "POST, OPTIONS"
    LET origin = "*"

   
    IF test = 1 THEN
        LET err = 1
        CALL com.WebServiceEngine.SetRestError(400, myError)
    ELSE
        CALL com.WebServiceEngine.SetRestError(401, myError)
    END IF

    RETURN method, origin, err
END FUNCTION
```

![In the image the HTTP error response is sent in the header "ERROR"](../_images/rest_response_wserrorheader.png)

*HTTP response using WSErrorHeader*

## Related links

**Related concepts**  

[Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")

[WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.")
