---
title: "com.HttpServiceRequest.setResponseCharset"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_setResponseCharset.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.setResponseCharset"
type: "concept"
---

# com.HttpServiceRequest.setResponseCharset

> Defines the HTTP response character set.

## Syntax

```
setResponseCharset(
   charset STRING )
```

1. charset defines the HTTP response
   character set.

## Usage

The `setResponseCharset()` method defines the character set to use when sending an
HTTP response.

The server must send a response in a character set that the client understands.

If the response character set is not defined by `setResponseCharset()`, the same
character set as the client request is used, or the implicit ISO-8859-1 charset is used if the
character is not defined by the client request. If the data to be return is application/json, the
response will be in UTF-8.

The method must be called before sending the response with one of [`sendResponse`](3837-com-httpservicerequest-sendresponse.md "Sends an HTTP response without body."), [`sendTextResponse`](3838-com-httpservicerequest-sendtextresponse.md "Sends an HTTP response with data from a plain string."),
[`sendXmlResponse`](3839-com-httpservicerequest-sendxmlresponse.md "Sends an HTTP response with data from a XML document object."), or [`beginXmlResponse`](3810-com-httpservicerequest-beginxmlresponse.md "Starts an HTTP streaming response.")
and [`endXmlResponse`](3812-com-httpservicerequest-endxmlresponse.md "Terminates an HTTP streaming response.") methods.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")
