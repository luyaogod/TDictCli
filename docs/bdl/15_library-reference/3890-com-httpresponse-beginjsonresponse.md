---
title: "com.HttpResponse.beginJSONResponse"
source: "fgl-topics/c_gws_ComHTTPResponse_beginJSONResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.beginJSONResponse"
type: "concept"
---

# com.HttpResponse.beginJSONResponse

> Starts a HTTP response streaming JSON.

## Syntax

```
beginJSONResponse()
RETURNS json.JSONReader
```

## Usage

The `beginJSONResponse()` method starts a streaming HTTP response and returns a
`json.JSONReader` object ready to read JSON from the server.

The Content-Type header must be `application/json`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
