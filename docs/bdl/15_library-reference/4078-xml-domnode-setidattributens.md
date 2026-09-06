---
title: "xml.DomNode.setIdAttributeNS"
source: "fgl-topics/c_gws_XmlDomNode_setIdAttributeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.setIdAttributeNS"
type: "concept"
---

# xml.DomNode.setIdAttributeNS

> Set the namespace-qualified XML Attribute of given name and namespace to be of type ID. Declare (or undeclare) the ID as user-determined.

## Syntax

```
setIdAttributeNS(
   name STRING,
   ns STRING,
   isId INTEGER )
```

1. name defines the name of the XML Attribute to set.
2. ns defines the namespace URI of the XML Attribute to set.
3. isId declares whether the attribute is a user-determined ID attribute.

## Usage

This method sets the namespace-qualified XML Attribute of given name and namespace to be of type
ID. Use the value `TRUE` for the parameter isID, to declare that
attribute as a user-determined ID attribute, otherwise use `FALSE`.

This affects the behavior of [getElementById](3982-xml-domdocument-getelementbyid.md "Returns the xml.DomNode element that has an attribute of type ID with the given value.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
