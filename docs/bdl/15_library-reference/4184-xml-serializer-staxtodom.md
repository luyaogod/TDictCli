---
title: "xml.Serializer.StaxToDom"
source: "fgl-topics/c_gws_XmlSerializer_StaxToDom.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.StaxToDom"
type: "concept"
---

# xml.Serializer.StaxToDom

> Serializes an XML element node into a DomNode object using a StaxReader object.

## Syntax

```
xml.Serializer.StaxToDom(
   stax xml.StaxReader,
   node xml.DomNode )
```

1. stax is a [StaxReader](4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") object where the cursor points
   to an XML Element node.
2. node is a [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object of type
   `ELEMENT_NODE` or `DOCUMENT_FRAGMENT_NODE`.

## Usage

This method serializes a XML element node into a DomNode object using a StaxReader object. The
resulting XML element node of the serialization process will be appended to the last child of the
given node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
