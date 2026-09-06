---
title: "xml.Signature.CreateFromNode"
source: "fgl-topics/c_gws_XmlSignature_CreateFromNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.CreateFromNode"
type: "concept"
---

# xml.Signature.CreateFromNode

> Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification.

## Syntax

```
xml.Signature.CreateFromNode(
   node xml.DomNode )
  RETURNS xml.Signature
```

1. node defines the XML Signature
   node.

## Usage

Returns a [XML
Signature object](4255-the-signature-class.md "The xml.Signature class provides methods to create detached, enveloped or enveloping XML signatures of one or more references of XML documents or document fragments, and to determine whether a signed referenced document has been modified afterwards.") or
NULL.

The node must be an element node with Signature as the local name, and it must
belong to the XML-Signature namespace http://www.w3.org/2000/09/xmldsig#, as
defined in [www.w3.org](http://www.w3.org/TR/xmldsig-core/#sec-Signature).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
