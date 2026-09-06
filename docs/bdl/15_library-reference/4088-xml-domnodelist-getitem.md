---
title: "xml.DomNodeList.getItem"
source: "fgl-topics/c_gws_XmlDomNodeList_getItem.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNodeList class > xml.DomNodeList methods > xml.DomNodeList.getItem"
type: "concept"
---

# xml.DomNodeList.getItem

> Returns the DomNode object at a given position in a DomNodeList object.

## Syntax

```
getItem(
   index INTEGER)
  RETURNS xml.DomNode
```

1. index defines the position of the DomNode object to return (index starts at 1).

## Usage

Use this method to return a `DomNode` object at the given position in this
`DomNodeList` object, where *index* is the position of the
`DomNode` object to return (Index starts at 1).

Returns `NULL` when no `DomNode` object is at the given
position.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
