---
title: "com.TcpRequest.beginXmlRequest"
source: "fgl-topics/c_gws_ComTCPRequest_beginXmlRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.beginXmlRequest"
type: "concept"
---

# com.TcpRequest.beginXmlRequest

> Starts a streaming XML request.

## Syntax

```
beginXmlRequest()
  RETURNS xml.StaxWriter
```

## Usage

The `beginXmlRequest()` method begins a streaming HTTP request and returns an
`xml.StaxWriter` object ready to send XML to the server.

After sending all the XML data to the server, you must call the [`endXmlRequest()`](3939-com-tcprequest-endxmlrequest.md "Terminates a streaming TCP request.") method with
the `xml.StaxWriter` object created by the `beginXmlRequest()`
method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
