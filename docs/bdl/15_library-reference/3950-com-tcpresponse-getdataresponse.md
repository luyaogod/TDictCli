---
title: "com.TcpResponse.getDataResponse"
source: "fgl-topics/c_gws_ComTCPResponse_getDataResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpResponse class > TcpResponse methods > com.TcpResponse.getDataResponse"
type: "concept"
---

# com.TcpResponse.getDataResponse

> Returns a TCP response in binary format.

## Syntax

```
getDataResponse(
   data BYTE )
```

1. data defines a `BYTE` variable that will hold the response
   data in binary format. The `BYTE` variable must be located `IN MEMORY`.

## Usage

This method retrieves the TCP response in binary format into the `data` variable
passed as parameter. The method will read the TCP stream, until the peer closes the connection.

> **Note:**
>
> The `data` variable must be located `IN MEMORY`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
