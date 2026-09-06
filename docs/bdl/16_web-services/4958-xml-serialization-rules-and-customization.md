---
title: "XML serialization rules and customization"
source: "fgl-topics/c_gws_XML_attributes_001.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization"
type: "concept"
description: "The topics in this section cover the default BDL/XML type mappings, the attributes that modify these mappings, and the options that control XML serialization and deserialization. These attributes and ..."
---

# XML serialization rules and customization

The topics in this section cover the default BDL/XML type mappings, the attributes that modify
these mappings, and the options that control XML serialization and deserialization. These attributes
and options are also referenced from the [`xml.Serializer`](../15_library-reference/4170-the-serializer-class.md "The xml.Serializer class provides methods to manage options for the serializer engine, and to use the serializer engine to serialize variables and XML element nodes.") class page.

## Child topics

- [BDL to/from XML type mappings](4959-bdl-to-from-xml-type-mappings.md): Map BDL variable types to specific XML data types for web service serialization using optional attributes. For example, use XSDBoolean to map a BDL SMALLINT to an XML boolean, and assign custom XML names as needed.
- [Default BDL/XML mapping](4960-default-bdl-xml-mapping.md): By default, Genero Web Services converts BDL variables to standard XML Schema (XSD) data types for use in web service messages.
- [XML facet constraint attributes](5008-xml-facet-constraint-attributes.md)
- [XML serialization attributes](5021-xml-serialization-attributes.md)
- [Serialize XML from a dynamic array](5046-serialize-xml-from-a-dynamic-array.md): The XMLList and XMLName attributes can be used to serialize dynamic arrays to XML and vice versa.
- [XML-to-BDL conversion options](5047-xml-to-bdl-conversion-options.md): This topic describes the options that relax XML-to-BDL deserialization by allowing the serializer to ignore unexpected XML attributes or elements.
