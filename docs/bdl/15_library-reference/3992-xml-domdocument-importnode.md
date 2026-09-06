---
title: "xml.DomDocument.importNode"
source: "fgl-topics/c_gws_XmlDomDocument_importNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.importNode"
type: "concept"
---

# xml.DomDocument.importNode

> Imports a xml.DomNode from a xml.DomDocument object into its new context (attached to a xml.DomDocument object).

## Syntax

```
importNode(
   n xml.DomNode
   deep INTEGER )
  RETURNS xml.DomNode
```

1. n defines the node to import.
2. deep defines a boolean identifying whether to import the node
   only or the node and all its child nodes.

Returns an `xml.DomNode` object.

## Usage

Imports an `xml.DomNode` object from an `xml.DomDocument` object
into its new context (attached to this `xml.DomDocument` object), where
n is the node to import. When deep is FALSE only the node is
imported; when `TRUE` the node and all its child nodes are imported.

Returns the `xml.DomNode` object that has been imported to this
`xml.DomDocument`.

Document
and Document Type nodes cannot be imported.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
