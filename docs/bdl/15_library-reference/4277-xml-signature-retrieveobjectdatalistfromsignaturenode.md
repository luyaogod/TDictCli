---
title: "xml.Signature.RetrieveObjectDataListFromSignatureNode"
source: "fgl-topics/c_gws_XmlSignature_RetrieveObjectDataListFromSignatureNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.RetrieveObjectDataListFromSignatureNode"
type: "concept"
---

# xml.Signature.RetrieveObjectDataListFromSignatureNode

> Returns a DomNodeList containing all embedded XML nodes related to the signature object

## Syntax

```
RetrieveObjectDataListFromSignatureNode(
   signNode xml.DomNode,
   index INTEGER )
  RETURNS xml.DomNodeList
```

1. signNode defines the XML Signature [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").
2. index defines the index of the
   signature object.

## Usage

This method returns a [xml.DomNodeList](4085-the-domnodelist-class.md "The xml.DomNodeList class provides methods to manipulate a list of DomNode objects.") containing all embedded XML nodes related to the signature object specified
by the index index within the XML signature node, signNode.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
