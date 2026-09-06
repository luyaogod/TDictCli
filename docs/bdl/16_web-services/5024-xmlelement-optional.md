---
title: "XMLElement (Optional)"
source: "fgl-topics/c_gws_XML_attribute_XMLElement.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLElement (Optional)"
type: "concept"
description: "Map a BDL simple data type to an XML Element. The attribute cannot be set on a type definition. Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 INTEGER ATTRIBUTES(XMLElement, ..."
---

# XMLElement (Optional)

Map a BDL simple data type to an XML Element. The attribute cannot be set on a type
definition.

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 INTEGER ATTRIBUTES(XMLElement, XSDunsignedShort, XMLName="Val1"),
  rec  RECORD ATTRIBUTES(XMLName="Rec")
    val2 FLOAT 	ATTRIBUTES(XMLElement, XMLName="Val2"),
    val3 STRING	ATTRIBUTES(XMLElement, XMLName="Val3")
  END RECORD
END RECORD
```

```
<Root>
  <Val1>148</Val1>
  <Rec1>
    <Val2>25.8</Val2>
    <Val3>Hello world</Val3>
  </Rec1>
</Root>
```
