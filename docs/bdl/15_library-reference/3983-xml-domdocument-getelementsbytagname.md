---
title: "xml.DomDocument.getElementsByTagName"
source: "fgl-topics/c_gws_XmlDomDocument_getElementsByTagName.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getElementsByTagName"
type: "concept"
---

# xml.DomDocument.getElementsByTagName

> Returns a xml.DomNodeList object containing all XML Element xml.DomNode objects with the same tag name in the document.

## Syntax

```
getElementsByTagName(
   tag STRING )
  RETURNS xml.DomNodeList
```

1. tag defines the name of the XML Element tag to match or "\*" to match all tags.

## Usage

Use this method to return a `xml.DomNodeList` object containing all XML Element
`xml.DomNode` objects with the same tag name in the entire document. The
tag string contains the name of the XML Element tag to match, or use "\*" to match
all tags.

Returns a [DomNodeList](4085-the-domnodelist-class.md "The xml.DomNodeList class provides methods to manipulate a list of DomNode objects.") object, or
`NULL`.

The returned list is ordered using a Depth-First pass algorithm.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Navigation methods usage examples](4011-navigation-methods-usage-examples.md "Examples using the navigation methods of the xml.DomDocument class.")
