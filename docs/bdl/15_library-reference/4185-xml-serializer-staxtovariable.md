---
title: "xml.Serializer.StaxToVariable"
source: "fgl-topics/c_gws_XmlSerializer_StaxToVariable.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.StaxToVariable"
type: "concept"
---

# xml.Serializer.StaxToVariable

> Serializes an XML element node into a BDL variable using a StaxReader object.

## Syntax

```
xml.Serializer.StaxToVariable(
   stax xml.StaxReader,
   var fgl-type )
```

1. stax is a [StaxReader](4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") object where the cursor points
   to an XML Element node.
2. var is any Genero BDL [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."). Or it can be a structured type, like a
   `RECORD` or `ARRAY`, including elements with these data types.
   Optional [XML mapping attributes](../16_web-services/4958-xml-serialization-rules-and-customization.md) can be added to the definition of variables for XML serialization.

## Usage

This method serializes a XML element node into a BDL variable using a StaxReader object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
