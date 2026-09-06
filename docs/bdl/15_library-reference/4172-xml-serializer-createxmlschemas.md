---
title: "xml.Serializer.CreateXmlSchemas"
source: "fgl-topics/c_gws_XmlSerializer_CreateXmlSchemas.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.CreateXmlSchemas"
type: "concept"
---

# xml.Serializer.CreateXmlSchemas

> Creates XML schemas corresponding to the given variable, and fills a dynamic array with xml.DomDocument objects each representing a XML schema.

## Syntax

```
xml.Serializer.CreateXmlSchemas(
   var fgl-type,
   schemas RECORD )
```

1. var is any Genero BDL [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."). Or it can be a structured type, like a
   `RECORD` or `ARRAY`, including elements with these data types.
   Optional [XML mapping attributes](../16_web-services/4958-xml-serialization-rules-and-customization.md) can be added to the definition of variables for XML serialization.
2. schemas is a dynamic array of [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")
   objects, each representing an XML schema.

## Usage

Use this method to create XML schemas corresponding to the given variable var,
and fill the dynamic array schemas with `xml.DomDocument` objects
each representing a XML schema.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
