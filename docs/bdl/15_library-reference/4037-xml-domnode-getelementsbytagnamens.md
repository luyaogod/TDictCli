---
title: "xml.DomNode.getElementsByTagNameNS"
source: "fgl-topics/c_gws_XmlDomNode_getElementsByTagNameNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getElementsByTagNameNS"
type: "concept"
---

# xml.DomNode.getElementsByTagNameNS

> Returns a DomNodeList object containing all namespace-qualified XML Element DomNode objects with the same tag name and namespace.

## Syntax

```
getElementsByTagNameNS(
   tag STRING,
   ns STRING )
  RETURNS xml.DomNodeList
```

1. tag defines the name of the XML Element
   tag to match or "\*" to match all tags.
2. ns defines the namespace URI of the XML
   Element tag to match or "\*" to match any namespace.

## Usage

Returns a `DomNodeList` object containing all namespace-qualified XML Element
DomNode objects with the same tag name and namespace, or NULL. tag is the name of
the XML Element tag to match, or "\*" to match all tags; ns is the namespace URI
of the XML Element tag to match, or "\*" to match any namespace.

The `getElementsByTagNameNS` and [getElementsByTagName](4036-xml-domnode-getelementsbytagname.md "Returns a DomNodeList object containing all XML Element DomNode objects with the same tag name.") methods return a
`DomNodeList` object, unlike the other methods that return a `DomNode`
object. The `DomNodeList` is restricted to contain objects with the same tag name
and/or namespace.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
