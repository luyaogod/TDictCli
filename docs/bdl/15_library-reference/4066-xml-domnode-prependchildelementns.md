---
title: "xml.DomNode.prependChildElementNS"
source: "fgl-topics/c_gws_XmlDomNode_prependChildElementNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.prependChildElementNS"
type: "concept"
---

# xml.DomNode.prependChildElementNS

> Creates and adds a child namespace-qualified XML Element node to the beginning of the list of child nodes for an XML Element DomNode object.

## Syntax

```
prependChildElementNS(
   prefix STRING,
   name STRING,
   ns STRING )
  RETURNS xml.DomNode
```

1. prefix defines the prefix of the XML Element to add.
2. name defines the name of the XML Element to add.
3. ns defines the namespace URI of the XML Element to add.

## Usage

Use this method to create and add a child namespace-qualified XML Element node to the beginning
of the list of child nodes for this XML Element `DomNode` object.

It returns the XML Element DomNode object, or `NULL`.

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
