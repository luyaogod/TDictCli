---
title: "XML-to-BDL conversion options"
source: "fgl-topics/c_gws_XML_serialization_options.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML-to-BDL conversion options"
type: "concept"
---

# XML-to-BDL conversion options

> This topic describes the options that relax XML-to-BDL deserialization by allowing the serializer to ignore unexpected XML attributes or elements.

By default, XML-to-BDL deserialization raises an error when the XML document contains attributes
or elements that are not defined in the BDL variable receiving the data.

To relax the conversion rules, you can set the `xml_ignoreunknownattributes` and
`xml_ignoreunknownelements` options on the [xml.Serializer](../15_library-reference/4170-the-serializer-class.md "The xml.Serializer class provides methods to manage options for the serializer engine, and to use the serializer engine to serialize variables and XML element nodes.") class. These options instruct the serializer to ignore unexpected XML
attributes or elements during deserialization. For a detailed list of option flags, refer to [Serialization option flags](../15_library-reference/4189-serialization-option-flags.md "Serialization option flags for the xml.Serializer class.").

For example, consider the following BDL record:

```
DEFINE data RECORD ATTRIBUTES(XMLName="Demo"),
val1 INTEGER ATTRIBUTES(XMLName="Value1"),
val2 STRING  ATTRIBUTES(XMLName="Value2"),
attr INTEGER ATTRIBUTES(XMLAttribute, XMLName="MyAttr")
END RECORD
```

The next XML fragment contains an undefined attribute and an unexpected element. Without relaxed
options, these will cause a conversion error.

```
<Demo MyAttr="hello" badAttr="BAD">
<Value1>128</Value1>
<Unexpected1>Will be ignored</Unexpected1>
<Value2>Hello</Value2>
</Demo>
```

To allow the document to load without errors, configure the options as follows:

```
xml_ignoreunknownattributes = TRUE
xml_ignoreunknownelements  = TRUE
#...
xml.Serializer.DomToVariable(doc, data)
```

These options cannot relax all structural incompatibilities. In complex cases, the
conversion may still fail when the XML structure does not match the expected BDL definition.

For details on setting options programmatically, see [xml.Serializer.SetOption](../15_library-reference/4182-xml-serializer-setoption.md "Sets a global option value for the serializer engine").

## Related links

**Related reference**  

[Serialization option flags](../15_library-reference/4189-serialization-option-flags.md "Serialization option flags for the xml.Serializer class.")
