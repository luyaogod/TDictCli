---
title: "xml.DomNode.isDefaultNamespace"
source: "fgl-topics/c_gws_XmlDomNode_isDefaultNamespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.isDefaultNamespace"
type: "concept"
---

# xml.DomNode.isDefaultNamespace

> Checks whether the specified namespace URI is the default namespace.

## Syntax

```
isDefaultNamespace(
   ns STRING )
  RETURNS INTEGER
```

1. ns defines the namespace URI to look for.

## Usage

Use this method to check whether the specified namespace URI is the default namespace, where
ns is the namespace URI to look for. Returns `TRUE` if the given
namespace is the default namespace, `FALSE` otherwise.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
