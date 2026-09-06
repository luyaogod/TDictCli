---
title: "xml.DomNode.lookupPrefix"
source: "fgl-topics/c_gws_XmlDomNode_lookupPrefix.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.lookupPrefix"
type: "concept"
---

# xml.DomNode.lookupPrefix

> Looks up the prefix associated to a namespace URI, starting from the specified node.

## Syntax

```
lookupPrefix(
   ns STRING )
  RETURNS STRING
```

1. ns defines the namespace URI to look for.

## Usage

Use this method to look up the prefix associated with a namespace URI, starting from this node,
where ns is the namespace URI to look for. Returns the prefix associated with
this namespace URI, or `NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
