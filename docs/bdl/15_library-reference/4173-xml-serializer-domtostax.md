---
title: "xml.Serializer.DomToStax"
source: "fgl-topics/c_gws_XmlSerializer_DomToStax.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.DomToStax"
type: "concept"
---

# xml.Serializer.DomToStax

> Serializes a XML DomNode object to a StaxWriter object.

## Syntax

```
xml.Serializer.DomToStax(
   node xml.DomNode,
   stax xml.StaxWriter )
```

1. node is a [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object.
2. stax is a [StaxWriter](4090-the-staxwriter-class.md "The xml.StaxWriter class provides methods compatible with Streaming API for XML(StAX) for writing XML documents.") object.

## Usage

Use this method to serialize a `xml.DomNode` object to a
`StaxWriter` object.

The resulting XML element node of the serialization process will be added at the current cursor
position of the StaxWriter object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
