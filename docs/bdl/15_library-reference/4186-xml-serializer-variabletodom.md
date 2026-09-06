---
title: "xml.Serializer.VariableToDom"
source: "fgl-topics/c_gws_XmlSerializer_VariableToDom.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.VariableToDom"
type: "concept"
---

# xml.Serializer.VariableToDom

> Serializes a BDL variable into a XML element node using a DomNode object.

## Syntax

```
xml.Serializer.VariableToDom(
   var fgl-type,
   node xml.DomNode )
```

1. var is any Genero BDL [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."). Or it can be a structured type, like a
   `RECORD` or `ARRAY`, including elements with these data types.
   Optional [XML mapping attributes](../16_web-services/4958-xml-serialization-rules-and-customization.md) can be added to the definition of variables for XML serialization.
2. node is a [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object of type
   `ELEMENT_NODE` or `DOCUMENT_FRAGMENT_NODE`.

## Usage

This method serializes a BDL variable into a XML element node using a DomNode object. The
resulting XML element node of the serialization process will be appended to the last child of the
given node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
