---
title: "SOAP multipart style requests in GWS"
source: "fgl-topics/c_gws_server_multipart.html"
breadcrumb: "Web services > Concepts > SOAP multipart style requests in GWS"
type: "concept"
---

# SOAP multipart style requests in GWS

> This topic describes multipart support with SOAP Genero Web Services

## What is multipart style in SOAP?

Multipart style SOAP is the ability to send and receive a SOAP request in multiple pieces.
The sending of attached files in separate parts of the SOAP request is one example of a multipart
style SOAP request.

## Multipart SOAP on the client

When using a WSDL with multipart style, fglwsdl generates a client-side
stub handling multipart requests. For more details, see [Multipart in the client stub](4626-multipart-in-the-client-stub.md "You can generate a client stub for a Web service that has multiple parts.").

## Multipart SOAP on the server

Multipart style is not yet supported with the high-level WS API of Genero.

- It is not possible to write a GWS server handling multipart style SOAP requests with
  the high-level API.
- When generating code from a WSDL using multipart style, the
  fglwsdl will produce a warning message: `WARNING : Unable
  to manage MIME Mutlipart binding on message 'name'`, where
  name is the name of the message in XML.

## Implementing multipart using the low-level APIs

If required, you can implement a WS server handling multipart with the low-level APIs of
Genero Web Services. For more details, see [com.HttpServiceRequest.getRequestMultipartType](../15_library-reference/3819-com-httpservicerequest-getrequestmultiparttype.md "Returns the multipart type of an incoming request.").

## Related links

**Related concepts**  

[SOAP Web services APIs](4510-soap-web-services-apis.md "Genero provides high-level and low-level APIs for creating SOAP Web services.")

[fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).")
