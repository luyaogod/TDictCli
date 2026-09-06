---
title: "com.TcpRequest.setConnectionTimeOut"
source: "fgl-topics/c_gws_ComTCPRequest_setConnectionTimeOut.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.setConnectionTimeOut"
type: "concept"
---

# com.TcpRequest.setConnectionTimeOut

> Defines the connection time out.

## Syntax

```
setConnectionTimeOut(
   timeout INTEGER )
```

1. timeout
   specifies a time out value in seconds.

## Usage

This method defines the time in seconds to wait for a connection, before a break.

If the time out is set to -1, it waits indefinitely.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
