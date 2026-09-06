---
title: "xml.DomDocument.declareNamespace"
source: "fgl-topics/c_gws_XmlDomDocument_declareNamespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.declareNamespace"
type: "concept"
---

# xml.DomDocument.declareNamespace

> Forces namespace declaration to a XML Element xml.DomNode for this xml.DomDocument object.

## Syntax

```
declareNamespace(
   node xml.DomNode,
   alias STRING,
   ns STRING )
```

1. node defines the XML Element DomNode that carries the namespace definition.
2. alias defines the alias of the namespace to declare.
3. ns defines the URI of the namespace to declare.

## Usage

Forces namespace declaration to a XML Element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") for this `xml.DomDocument` object; node is the XML
Element `xml.DomNode` that carries the namespace definition; alias
is the alias of the namespace to declare, or `NULL` to declare the default namespace;
ns is the URI of the namespace to declare (can only be `NULL` if
alias is `NULL`).

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
