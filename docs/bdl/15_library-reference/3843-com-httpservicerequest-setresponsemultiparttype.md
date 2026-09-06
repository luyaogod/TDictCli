---
title: "com.HttpServiceRequest.setResponseMultipartType"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_setResponseMultipartType.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.setResponseMultipartType"
type: "concept"
---

# com.HttpServiceRequest.setResponseMultipartType

> Sets HTTP response in multipart mode of given type.

## Syntax

```
setResponseMultipartType(
   type STRING,
   start STRING,
   boundary STRING )
```

1. type defines one of the following:
   - form-data: Browser Xform with attachment
   - mixed: Parts are independent
   - related: Parts are dependent (Required for SOAP)
   - alternative: Parts are different types of the same document
   - *or any other type*
   - NULL: switch multipart mode off
2. start specifies the Content-ID value
   of root multipart document. Must be ASCII. (optional)
3. boundary specifies the string used as
   multipart boundary. Must be ASCII. (optional)

## Usage

Sets HTTP response in multipart mode of given type. Calling one of the standard request methods
will send the HTTP response as the given multipart type, even if no other part has been set.

The root HTTP part must be handled via the standard HttpServiceRequest methods such as
`sendTextRequest()`, `sendXmlRequest()`,
`sendDataRequest()` and [`BeginXmlResponse()`](3810-com-httpservicerequest-beginxmlresponse.md "Starts an HTTP streaming response.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
