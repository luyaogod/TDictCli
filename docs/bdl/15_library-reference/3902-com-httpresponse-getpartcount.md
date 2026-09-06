---
title: "com.HttpResponse.getPartCount"
source: "fgl-topics/c_gws_ComHTTPResponse_getPartCount.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getPartCount"
type: "concept"
---

# com.HttpResponse.getPartCount

> Returns the number of additional parts in the HTTP response.

## Syntax

```
getPartCount()
  RETURNS INTEGER
```

## Usage

Returns the number of additional parts in the HTTP response. The root part element must be
handled via [`getXmlResponse()`](3908-com-httpresponse-getxmlresponse.md "Returns the entire HTTP response in a DOM document."), [`getTextResponse()`](3907-com-httpresponse-gettextresponse.md "Returns the entire HTTP response in a string."), [`getDataResponse()`](3894-com-httpresponse-getdataresponse.md "Returns the entire HTTP response in a BYTE.") and [`beginXmlResponse()`](3891-com-httpresponse-beginxmlresponse.md "Starts a streaming HTTP response."). In
other words, there are `getPartCount()` +1 parts if [`getMultipartType()`](3900-com-httpresponse-getmultiparttype.md "Returns whether a response is multipart or not, and the kind of multipart if any.") does
not return NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
