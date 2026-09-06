---
title: "xml.DomNode.getElementsByTagName"
source: "fgl-topics/c_gws_XmlDomNode_getElementsByTagName.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getElementsByTagName"
type: "concept"
---

# xml.DomNode.getElementsByTagName

> Returns a DomNodeList object containing all XML Element DomNode objects with the same tag name.

## Syntax

```
getElementsByTagName(
   tag STRING )
  RETURNS xml.DomNodeList
```

1. tag defines the name of the XML Element tag to match or "\*" to match all tags.

## Usage

Returns a `DomNodeList` object containing all XML Element DomNode objects with the
same tag name, or NULL; tag is the name of the XML Element tag to match, or "\*"
to match all tags.

The `getElementsByTagName` and [getElementsByTagNameNS](4037-xml-domnode-getelementsbytagnamens.md "Returns a DomNodeList object containing all namespace-qualified XML Element DomNode objects with the same tag name and namespace.") methods return a
`xml.DomNodeList` object, unlike the other methods that return a
`DomNode` object. The `DomNodeList` is restricted to objects with the
same tag name and/or namespace.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
