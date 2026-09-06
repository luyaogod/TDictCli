---
title: "com.HttpServiceRequest.getRequestPartCount"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_getRequestPartCount.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.getRequestPartCount"
type: "concept"
---

# com.HttpServiceRequest.getRequestPartCount

> Returns the number of additional multipart elements.

## Syntax

```
getRequestPartCount()
  RETURNS INTEGER
```

## Usage

The root multipart is handled via standard [`readTextRequest()`](3833-com-httpservicerequest-readtextrequest.md "Returns the request body as a plain string."),
[`readXmlRequest()`](3834-com-httpservicerequest-readxmlrequest.md "Returns the request body as an XML document."), [`readDataRequest()`](3830-com-httpservicerequest-readdatarequest.md "Returns the body of a request in a BYTE.")
and [`beginXmlRequest()`](3809-com-httpservicerequest-beginxmlrequest.md "Starts an HTTP streaming request.").

The number of parts is only available when the entire request has been read.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
