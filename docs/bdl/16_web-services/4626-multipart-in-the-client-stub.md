---
title: "Multipart in the client stub"
source: "fgl-topics/c_gws_client_stub_multipart.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Multipart in the client stub"
type: "concept"
---

# Multipart in the client stub

> You can generate a client stub for a Web service that has multiple parts.

If the WSDL for a Web service indicates that the Web service uses multiple parts, the client stub
generated will support multiple parts.

## For the request

There are as many `com.HTTPPart`
input parameters as parts defined for the input request, plus one
`AnyInputParts DYNAMIC ARRAY OF com.HTTPPart` parameter, to manage
the optional parts a user can add to the request.

For example:

```
FUNCTION xxx_g(InputHttpPart_1, ..., InputHttpPart_n, AnyInputParts)
  DEFINE InputHttpPart_1 com.HTTPPart
  ...
  DEFINE InputHttpPart_n com.HTTPPart
  DEFINE AnyInputParts DYNAMIC ARRAY OF com.HTTPPart
  ...
  RETURN wsstatus
END FUNCTION
```

## For the response

There are as many `com.HTTPPart`
variables as are described in the WSDL, plus one `AnyOutputParts DYNAMIC ARRAY OF
com.HTTPPart` to handle the optional parts that may be returned by a service.

For
example:

```
FUNCTION xxx_g()
  DEFINE wsstatus INTEGER
  DEFINE OutputHttpPart_1 com.HTTPPart
  DEFINE AnyOutputParts DYNAMIC ARRAY OF com.HTTPPart
  ...
  RETURN wsstatus, OutputHttpPart_1, AnyOutputParts
END FUNCTION
```

## Related links

**Related concepts**  

[SOAP multipart style requests in GWS](4553-soap-multipart-style-requests-in-gws.md "This topic describes multipart support with SOAP Genero Web Services")
