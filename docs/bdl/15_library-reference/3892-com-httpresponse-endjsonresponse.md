---
title: "com.HttpResponse.endJSONResponse"
source: "fgl-topics/c_gws_ComHTTPResponse_endJSONResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.endJSONResponse"
type: "concept"
---

# com.HttpResponse.endJSONResponse

> Performs the HTTP request.

## Syntax

```
endJSONResponse(
   json json.JSONReader )
```

1. json defines a `json.JSONReader` object used to
   read the HTTP response.

## Usage

The `endJSONResponse()` method ends the streaming HTTP response by closing the
`json.JSONReader` object that was created with the [beginJSONResponse](3890-com-httpresponse-beginjsonresponse.md "Starts a HTTP response streaming JSON.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
