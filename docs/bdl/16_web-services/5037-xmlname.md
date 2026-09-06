---
title: "XMLName"
source: "fgl-topics/c_gws_XML_attribute_XMLName.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLName"
type: "concept"
description: "Define the name of a variable in an XML document. The attribute can be set on a RECORD and any of its elements. The attribute cannot be set on a TYPE definition but it can be set on elements of the ..."
---

# XMLName

Define the name of a variable in an XML document.

1. The attribute can be set on a `RECORD` and any of its elements.
2. The attribute cannot be set on a [TYPE](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") definition but it can
   be set on elements of the type.

## Example 1: Record

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1  INTEGER ATTRIBUTES(XMLName="Val1"),
  val2  FLOAT,
  val3  INTEGER ATTRIBUTES(XMLName="Val3")
END RECORD
```

If the output is viewed in an XML viewer, it would display as shown.

```
<Root>
  <Val1>148</Val1>
  <val2>0.5</val2>
  <Val3>-18547</Val3>
</Root>
```

## Example 2: Type

As the `XMLName` attribute cannot be set on the root of a [TYPE](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") definition, you must assign the type to a variable and the
`XMLName` attribute can be set on the
variable.

```
TYPE Type1 RECORD ATTRIBUTE(XSTypename="MyType")  
  f1, f2 INT
END RECORD

DEFINE var Type1 ATTRIBUTE(XMLName="MyName")
```

This generates the following in the XML schema describing the structure of the XML document:

```
<?xml version="1.0" encoding="UTF-8" ?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
  <xs:complexType name="MyType"> 
   ...
  </xs:complexType>
  <xs:element name="MyName" type="xs:MyType"/>
</xs:schema>
```

## Related links

**Related concepts**  

[Serialize XML from a dynamic array](5046-serialize-xml-from-a-dynamic-array.md "The XMLList and XMLName attributes can be used to serialize dynamic arrays to XML and vice versa.")
