---
title: "xml.DomNode.isAttached"
source: "fgl-topics/c_gws_XmlDomNode_isAttached.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.isAttached"
type: "concept"
---

# xml.DomNode.isAttached

> Returns whether the node is attached to the XML document.

## Syntax

```
isAttached()
  RETURNS INTEGER
```

## Usage

The returned integer indicates whether the node is attached to the XML document or not.

> **Important:**
>
> This method is not part of W3C standard
> API.

This method returns `TRUE` if the [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object is attached to a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") object as a child; otherwise it returns `FALSE`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
