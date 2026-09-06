---
title: "com.TcpRequest.setMaximumResponseLength"
source: "fgl-topics/c_gws_ComTCPRequest_setMaximumResponseLength.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.setMaximumResponseLength"
type: "concept"
---

# com.TcpRequest.setMaximumResponseLength

> Defines the maximum size in Kbyte of the response.

## Syntax

```
setMaximumResponseLength(
   length INTEGER )
```

1. length specifies the max size of a
   response in Kbytes.

## Usage

This method sets the maximum authorized size in Kbyte of the whole response, before a break.

A length of -1 defines no limit.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
