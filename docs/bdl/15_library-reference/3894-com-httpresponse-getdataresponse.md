---
title: "com.HttpResponse.getDataResponse"
source: "fgl-topics/c_gws_ComHTTPResponse_getDataResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getDataResponse"
type: "concept"
---

# com.HttpResponse.getDataResponse

> Returns the entire HTTP response in a BYTE.

## Syntax

```
getDataResponse(
   b BYTE )
```

1. b defines a `BYTE`
   variable receiving the HTTP response data.

## Usage

The `getDataResponse()` method returns the body of an HTTP response into a [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") variable.

The `BYTE` variable must be located in memory, otherwise operation fails.

Returns binary data as response from a server into a `BYTE`.

Previous content is discarded.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
