---
title: "XMLAttribute"
source: "fgl-topics/c_gws_XML_attribute_XMLAttribute.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLAttribute"
type: "concept"
description: "Map a BDL simple data type to an XML Attribute. The attribute cannot be set on a type definition. The attribute can only be set on a RECORD element. Example DEFINE myVar RECORD ..."
---

# XMLAttribute

Map a BDL simple data type to an XML Attribute.

1. The attribute cannot be set on a type definition.
2. The attribute can only be set on a `RECORD` element.

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 INTEGER ATTRIBUTES(XMLAttribute,XSDunsignedShort,XMLName="Val1"),
  rec  RECORD ATTRIBUTES(XMLName="Rec1")
    val2  FLOAT ATTRIBUTES(XMLAttribute,XMLName="Val2"),
    val3  STRING ATTRIBUTES(XMLElement,XMLName="Val3")
  END RECORD
END RECORD
```

```
<Root Val1="148">
  <Rec1 Val2="25.8">
    <Val3>Hello world</Val3>
  </Rec1>
</Root>
```
