---
title: "com.HttpServiceRequest.sendXmlResponse"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_sendXmlResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.sendXmlResponse"
type: "concept"
---

# com.HttpServiceRequest.sendXmlResponse

> Sends an HTTP response with data from a XML document object.

## Syntax

```
sendXmlResponse(
   code INTEGER,
   desc STRING,
   doc xml.DomDocument )
```

1. code specifies the
   status code of the response.
2. description
   specifies the description of the response.
3. doc specifies the XML document
   containing the data to be sent.

## Usage

The `sendXmlResponse()` method performs the HTTP response by sending a status
(code) and description (desc), followed by the headers
previously set, and the XML data contained in the passed [`xml.DomDocument`](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") object as body.

The `code` returned ​can be a value between 100 and 999. You can send
any code in this range. This provides support for the status codes returned by HTTP, WebDAV, other
Request for Comments (RFC) codes, etc. The following determines how the `desc` is
handled:

- If the code you send is recognized (as one of the standards) and the description is
  `NULL`, the description corresponding to the code standard is sent to the client.
- If the code is not recognized, the description you provide is added. If you do not provide a
  description, the default description "No Message" is sent to the client.

> **Note:**
>
> It is important for the server to return a status code, following the
> HTTP standards, otherwise the client may fail to interpret the response. For instance, if the
> request is malformed, the server is expected to send an HTTP response with the code of 400 (Bad
> Request).

For more details about HTTP response codes, see [HTTP
status codes (wikipedia)](http://en.wikipedia.org/wiki/List_of_HTTP_status_codes).

The default Content-Type header is `text/xml`, but it can be changed if of the
form `*/xml` or `*/*+xml`. For example:
`application/xhtml+xml`.

In HTTP 1.1, if the body size is greater than 32k, the response will be sent in several chunks of
the same size.

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
