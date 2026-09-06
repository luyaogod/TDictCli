---
title: "xml.DomDocument.getDocumentNodesCount"
source: "fgl-topics/c_gws_XmlDomDocument_getDocumentNodesCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getDocumentNodesCount"
type: "concept"
---

# xml.DomDocument.getDocumentNodesCount

> Returns the number of child xml.DomNode objects for this xml.DomDocument object.

## Syntax

```
getDocumentNodesCount()
  RETURNS INTEGER
```

## Usage

Returns the number of child `xml.DomNode` objects in this
`xml.DomDocument` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Navigation methods usage examples](4011-navigation-methods-usage-examples.md "Examples using the navigation methods of the xml.DomDocument class.")
