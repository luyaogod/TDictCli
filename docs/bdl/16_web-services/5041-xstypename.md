---
title: "XSTypename"
source: "fgl-topics/c_gws_XML_attribute_XSTypename.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XSTypename"
type: "concept"
description: "Define the XML Schema name of a BDL type definition. The attribute must be used with the XSTypenamespace attribute; otherwise, the generated source code will not compile. The attribute is only allowed ..."
---

# XSTypename

Define the XML Schema name of a BDL type definition.

1. The attribute must be used with the [XSTypenamespace](5042-xstypenamespace.md) attribute; otherwise, the generated source code
   will not compile.
2. The attribute is only allowed on a [TYPE](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.")
   definition.

## Example

```
TYPE myType RECORD ATTRIBUTES(XMLSequence,
                             XSTypeName="MyFirstType",
                             XSTypeNamespace="http://tempuri.org" )
  val1  FLOAT   ATTRIBUTES(XMLElement,XMLName="Val1"),
  val2  INTEGER ATTRIBUTES(XMLElement,XMLName="Val2",XMLOptional),
  attr  STRING  ATTRIBUTES(XMLAttribute,XMLName="Attr")
END RECORD
```

```
<xsd:schema xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
 targetNamespace="http://tempuri.org" elementFormDefault="qualified" >
  <xsd:complexType name="MyFirstType">
    <xsd:sequence>
      <xsd:element name="Val1" type="xsd:double" />
      <xsd:element name="Val2" type="xsd:int" minOccurs="0" />
    </xsd:sequence>
    <xsd:attribute name="Attr" type="xsd:string" use="required" />
  </xsd:complexType>
</xsd:schema>
```
