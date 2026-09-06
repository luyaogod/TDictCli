---
title: "com.TcpResponse.endXmlResponse"
source: "fgl-topics/c_gws_ComTCPResponse_endXmlResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpResponse class > TcpResponse methods > com.TcpResponse.endXmlResponse"
type: "concept"
---

# com.TcpResponse.endXmlResponse

> Ends a streaming TCP response.

## Syntax

```
endXmlResponse(
   stax xml.StaxReader )
```

1. stax defines the
   `xml.StaxReader` object created with [`beginXmlResponse()`](3948-com-tcpresponse-beginxmlresponse.md "Starts a streaming TCP response.") .

## Usage

Terminates the streaming TCP response identified by the `xml.StaxReader` object
passed as parameter. This object must have been created with the
`beginXmlResponse()` method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
