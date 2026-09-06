---
title: "com.TcpRequest.endXmlRequest"
source: "fgl-topics/c_gws_ComTCPRequest_endXmlRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.endXmlRequest"
type: "concept"
---

# com.TcpRequest.endXmlRequest

> Terminates a streaming TCP request.

## Syntax

```
endXmlRequest(
  stax xml.StaxWriter )
```

1. stax specifies the
   `xml.StaxWriter` object used for streaming.

## Usage

The `endXmlRequest()` method terminates a streaming TCP request performed with the
`xml.StaxWriter` object that was created by the [`beginXmlRequest()`](3933-com-tcprequest-beginxmlrequest.md "Starts a streaming XML request.")
method.

The connection is shutdown for writing, to confirm that no data will be sent.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
