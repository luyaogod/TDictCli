---
title: "com.HttpRequest.setMultipartType"
source: "fgl-topics/c_gws_ComHTTPRequest_setMultipartType.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setMultipartType"
type: "concept"
---

# com.HttpRequest.setMultipartType

> Switch HttpRequest in multipart mode of a given type.

## Syntax

```
setMultipartType(
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
2. start defines the Content-ID value of root
   multipart document. (optional)
3. boundary defines a string used as multipart
   boundary. (optional)

## Usage

Switch HttpRequest in multipart mode of a given type. Calling one of the standard request methods
will send the HTTP request as the given multipart type, even if no other part has been set.

The root HTTP part is the part handled via the standard HttpRequest methods such as [`doTextRequest()`](3862-com-httprequest-dotextrequest.md "Performs the request by sending an entire string at once."), [`doXmlRequest()`](3863-com-httprequest-doxmlrequest.md "Performs the request by sending an entire XML document at once."), [`doDataRequest()`](3858-com-httprequest-dodatarequest.md "Performs the request by sending binary data.") and [`beginXmlRequest()`](3851-com-httprequest-beginxmlrequest.md "Starts a streaming HTTP request.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
