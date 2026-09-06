---
title: "xml.DomDocument.getElementById"
source: "fgl-topics/c_gws_XmlDomDocument_getElementById.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getElementById"
type: "concept"
---

# xml.DomDocument.getElementById

> Returns the xml.DomNode element that has an attribute of type ID with the given value.

## Syntax

```
getElementById(
   id STRING )
  RETURNS xml.DomNode
```

1. id defines the Id value.

## Usage

Use this method to return the `xml.DomNode` element that has an attribute of type
ID with the given value, or `NULL` if there is none.

Attributes with the name "ID" or "id" are not of type ID unless so defined with [setIdAttribute](4077-xml-domnode-setidattribute.md "Set the XML Attribute of given name to be of type ID. Declare (or undeclare) the ID as user-determined.") or
[setIdAttributeNS](4078-xml-domnode-setidattributens.md "Set the namespace-qualified XML Attribute of given name and namespace to be of type ID. Declare (or undeclare) the ID as user-determined."). However, there is a specific attribute called `xml:id`
belonging to the namespace `http://www.w3.org/XML/1998/namespace` that is always of
type ID even if not set with setIdAttributeNS.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Navigation methods usage examples](4011-navigation-methods-usage-examples.md "Examples using the navigation methods of the xml.DomDocument class.")
