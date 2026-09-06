---
title: "xml.Serializer.VariableToStax"
source: "fgl-topics/c_gws_XmlSerializer_VariableToStax.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.VariableToStax"
type: "concept"
---

# xml.Serializer.VariableToStax

> Serializes a BDL variable into a XML element node using a StaxWriter object.

## Syntax

```
xml.Serializer.VariableToStax(
   var fgl-type,
   stax xml.StaxWriter )
```

1. var is any Genero BDL [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."). Or it can be a structured type, like a
   `RECORD` or `ARRAY`, including elements with these data types.
   Optional [XML mapping attributes](../16_web-services/4958-xml-serialization-rules-and-customization.md) can be added to the definition of variables for XML serialization.
2. stax is a [StaxWriter](4090-the-staxwriter-class.md "The xml.StaxWriter class provides methods compatible with Streaming API for XML(StAX) for writing XML documents.") object.

## Usage

This method serializes a BDL variable into a XML element node using a StaxReader object.

The resulting XML element node of the serialization process will be added at the
current cursor position of the StaxWriter object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
