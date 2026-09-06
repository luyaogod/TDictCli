---
title: "com.TcpRequest.doXmlRequest"
source: "fgl-topics/c_gws_ComTCPRequest_doXmlRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.doXmlRequest"
type: "concept"
---

# com.TcpRequest.doXmlRequest

> Performs a request with a DOM document.

## Syntax

```
doXmlRequest(
   doc xml.DomDocument )
```

1. doc specifies the DOM document
   describing the request.

## Usage

The `doXmlRequest()` method performs the TCP request by using the information
defined in the [`xml.DomDocument`](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") object passed as parameter.

The connection is shutdown for writing, to confirm that no data will be sent.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
