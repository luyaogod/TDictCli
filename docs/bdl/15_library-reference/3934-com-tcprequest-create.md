---
title: "com.TcpRequest.Create"
source: "fgl-topics/c_gws_ComTCPRequest_Create.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.Create"
type: "concept"
---

# com.TcpRequest.Create

> Creates a new TCP request object.

## Syntax

```
com.TcpRequest.Create(
   uri STRING )
  RETURNS com.TcpRequest
```

1. uri specifies the URL of the TCP
   request.

## Usage

This class method creates a new `com.TcpRequest` object based on the URL passed as
parameter.

The URL must use the TCP or TCPS protocol. Examples of valid URLs include:

- `tcp://localhost:4242/`
- `tcps://localhost:4343/`

The URL can be an identifier of a URL mapping with an optional alias://
prefix. See [FGLPROFILE configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") for more
details about URL mapping with aliases, and for proxy and security configuration.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
