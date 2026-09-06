---
title: "com.TcpRequest.getResponse"
source: "fgl-topics/c_gws_ComTCPRequest_getResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.getResponse"
type: "concept"
---

# com.TcpRequest.getResponse

> Returns the response after performing a TCP request.

## Syntax

```
getResponse()
  RETURNS com.TcpResponse
```

## Usage

The `getResponse()` method returns a TCP response as a [`com.TcpResponse`](3946-the-tcpresponse-class.md "The com.TcpResponse class provides an interface to perform XML and TEXT responses over TCP, with additional XML streaming possibilities.") object, after a call to
[`doRequest()`](3936-com-tcprequest-dorequest.md "Performs a TCP request."), [`doXmlRequest()`](3938-com-tcprequest-doxmlrequest.md "Performs a request with a DOM document."), [`doTextRequest()`](3937-com-tcprequest-dotextrequest.md "Performs a request with a string."), or [`beginXmlRequest() /
endXmlRequest()`](3933-com-tcprequest-beginxmlrequest.md "Starts a streaming XML request.") calls.

A call to this method will stop the program flow until the response is received.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
