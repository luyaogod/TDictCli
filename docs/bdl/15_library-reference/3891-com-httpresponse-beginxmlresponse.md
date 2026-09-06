---
title: "com.HttpResponse.beginXmlResponse"
source: "fgl-topics/c_gws_ComHTTPResponse_beginXmlResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.beginXmlResponse"
type: "concept"
---

# com.HttpResponse.beginXmlResponse

> Starts a streaming HTTP response.

## Syntax

```
beginXmlResponse()
  RETURNS xml.StaxReader
```

## Usage

The `beginXmlResponse()` method starts a streaming HTTP response and returns a
[`xml.StaxReader`](4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") object ready
to read XML from the server.

The Content-Type header must be of the form `*/xml` or `*/*+xml`.
For example: `application/xhtml+xml`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")
