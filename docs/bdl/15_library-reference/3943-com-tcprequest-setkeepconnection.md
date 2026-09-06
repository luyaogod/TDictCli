---
title: "com.TcpRequest.setKeepConnection"
source: "fgl-topics/c_gws_ComTCPRequest_setKeepConnection.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.setKeepConnection"
type: "concept"
---

# com.TcpRequest.setKeepConnection

> Defines if the TCP connection is kept open after sending a request.

## Syntax

```
setKeepConnection(
   keep INTEGER )
```

1. keep indicates if the TCP connection
   must be kept open.

## Usage

This method can be used to force the TCP socket to remain open after a send operation, in order
to perform subsequent `do*Request()` calls, without closing the connection (in write
mode).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
