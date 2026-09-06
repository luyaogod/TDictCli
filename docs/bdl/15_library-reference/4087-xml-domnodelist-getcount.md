---
title: "xml.DomNodeList.getCount"
source: "fgl-topics/c_gws_XmlDomNodeList_getCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNodeList class > xml.DomNodeList methods > xml.DomNodeList.getCount"
type: "concept"
---

# xml.DomNodeList.getCount

> Returns the number of DomNode objects in a DomNodeList object.

## Syntax

```
getCount()
  RETURNS INTEGER
```

## Usage

Use this method to return the number of [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.")
objects in this `DomNodeList` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
