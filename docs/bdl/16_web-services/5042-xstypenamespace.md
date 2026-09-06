---
title: "XSTypenamespace"
source: "fgl-topics/c_gws_XML_attribute_XSTypenamespace.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XSTypenamespace"
type: "concept"
description: "Define the XML Schema namespace of a BDL type definition. The attribute must be used with the XSType attribute; otherwise, the generated source code will not compile. The attribute is only allowed on ..."
---

# XSTypenamespace

Define the XML Schema namespace of a BDL type definition.

1. The attribute must be used with the XSType attribute; otherwise, the generated source
   code will not compile.
2. The attribute is only allowed on a type definition.

## Example

```
TYPE myType RECORD ATTRIBUTES(XMLChoice,
                             XSTypeName="MyFirstChoice",
                             XSTypeNamespace="http://tempuri.org" )
  val1  FLOAT   ATTRIBUTES(XMLElement,XMLName="Val1"),
  val2  INTEGER ATTRIBUTES(XMLElement,XMLName="Val2",XMLOptional),
  attr  STRING  ATTRIBUTES(XMLAttribute,XMLName="Attr",XMLOptional),
  set   INTEGER ATTRIBUTES(XMLSelector)
END RECORD
```

```
<xsd:schema xmlns:xsd="http://www.w3.org/2001/XMLSchema"
 targetNamespace="http://tempuri.org"  elementFormDefault="qualified" >
  <xsd:complexType name="MyFirstChoice">
    <xsd:choice>
      <xsd:element name="Val1" type="xsd:double" />
      <xsd:element name="Val2" type="xsd:int" minOccurs="0" />
    </xsd:choice>
    <xsd:attribute name="Attr" type="xsd:string" />
  </xsd:complexType>
</xsd:schema>
```
