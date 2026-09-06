---
title: "xml.DomNode.lookupNamespaceURI"
source: "fgl-topics/c_gws_XmlDomNode_lookupNamespaceURI.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.lookupNamespaceURI"
type: "concept"
---

# xml.DomNode.lookupNamespaceURI

> Looks up the namespace URI associated to a prefix, starting from a specified node.

## Syntax

```
lookupNamespaceURI(
   prefix STRING )
  RETURNS STRING
```

1. prefix defines the prefix to look for.

## Usage

Use this method to look up the namespace URI associated with a prefix, starting from this node,
where prefix is the prefix to look for. If `NULL`, the default
namespace URI will be returned. A namespace URI, or `NULL`, is returned.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
