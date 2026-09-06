---
title: "com.HttpRequest.beginXmlRequest"
source: "fgl-topics/c_gws_ComHTTPRequest_beginXmlRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.beginXmlRequest"
type: "concept"
---

# com.HttpRequest.beginXmlRequest

> Starts a streaming HTTP request.

## Syntax

```
beginXmlRequest()
  RETURNS xml.StaxWriter
```

## Usage

The `beginXmlRequest()` starts a streaming HTTP request and returns an [`xml.StaxWriter`](4090-the-staxwriter-class.md "The xml.StaxWriter class provides methods compatible with Streaming API for XML(StAX) for writing XML documents.") object ready to send to the server.

Supported methods are:
PUT, POST, PATCH, and DELETE.
> **Warning:**
>
> A message body is allowed in a
> DELETE request, but servers may ignore it if they do not support it.

The default Content-Type header is `text/xml`, but it can be changed if of the
form `*/xml` or `*/*+xml`. For example:
`application/xhtml+xml`.

In HTTP 1.1, if the body size is greater than 32 KB, the request will be sent in several chunks
of the same size.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
