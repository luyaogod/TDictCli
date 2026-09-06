---
title: "xml.Serializer.DomToVariable"
source: "fgl-topics/c_gws_XmlSerializer_DomToVariable.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.DomToVariable"
type: "concept"
---

# xml.Serializer.DomToVariable

> Serializes a XML element node into a BDL variable using a DomNode object.

## Syntax

```
xml.Serializer.DomToVariable(
   node xml.DomNode,
   var fgl-type )
```

1. node is a [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object of type
   `ELEMENT_NODE`.
2. var is any Genero BDL [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."). Or it can be a structured type, like a
   `RECORD` or `ARRAY`, including elements with these data types.
   Optional [XML mapping attributes](../16_web-services/4958-xml-serialization-rules-and-customization.md) can be added to the definition of variables for XML serialization.

## Usage

Use this method to serialize a XML element node defined by node as a DomNode
object into a BDL variable.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
