---
title: "com.HttpServiceRequest.setResponseHeader"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_setResponseHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.setResponseHeader"
type: "concept"
---

# com.HttpServiceRequest.setResponseHeader

> Defines a header for the HTTP response.

## Syntax

```
setResponseHeader(
   name STRING,
   value STRING )
```

1. name specifies the name of a
   header.
2. value defines the value of a
   header.

## Usage

The `setResponseHeader()` method sets (or replaces) the name and value of an HTTP
response header.

The Content-Length header cannot be set, because it is computed internally according to the body
size.

The method must be called before sending the response with one of [`sendResponse`](3837-com-httpservicerequest-sendresponse.md "Sends an HTTP response without body."), [`sendTextResponse`](3838-com-httpservicerequest-sendtextresponse.md "Sends an HTTP response with data from a plain string."),
[`sendXmlResponse`](3839-com-httpservicerequest-sendxmlresponse.md "Sends an HTTP response with data from a XML document object."), or [`beginXmlResponse`](3810-com-httpservicerequest-beginxmlresponse.md "Starts an HTTP streaming response.")
and [`endXmlResponse`](3812-com-httpservicerequest-endxmlresponse.md "Terminates an HTTP streaming response.") methods.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
