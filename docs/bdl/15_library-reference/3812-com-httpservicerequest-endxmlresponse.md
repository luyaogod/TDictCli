---
title: "com.HttpServiceRequest.endXmlResponse"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_endXmlResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.endXmlResponse"
type: "concept"
---

# com.HttpServiceRequest.endXmlResponse

> Terminates an HTTP streaming response.

## Syntax

```
endXmlResponse(
   stax xml.StaxWriter )
```

1. stax defines
   `xml.StaxWriter` used to write the HTTP body.

## Usage

The `endXmlResponse()` method terminates the HTTP streaming response by closing
the `xml.StaxWriter` object created by [`beginXmlResponse`](3810-com-httpservicerequest-beginxmlresponse.md "Starts an HTTP streaming response.").

The body of the request is discarded.

New incoming requests can be retrieved again with the [`com.WebServiceEngine.GetHTTPServiceRequest()`](3788-com-webserviceengine-gethttpservicerequest.md "Get a handle for an incoming HTTP service request.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
