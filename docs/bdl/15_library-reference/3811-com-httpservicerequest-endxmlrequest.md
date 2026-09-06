---
title: "com.HttpServiceRequest.endXmlRequest"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_endXmlRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.endXmlRequest"
type: "concept"
---

# com.HttpServiceRequest.endXmlRequest

> Terminates an HTTP streaming request.

## Syntax

```
endXmlRequest(
   stax xml.StaxReader )
```

1. stax defines the
   `xml.StaxReader` object used for streaming.

## Usage

The `endXmlRequest()` method ends the streaming HTTP request by closing the [`xml.StaxReader`](4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") object passed as parameter.

The stax object must be created with the [`beginXmlRequest()`](3809-com-httpservicerequest-beginxmlrequest.md "Starts an HTTP streaming request.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
