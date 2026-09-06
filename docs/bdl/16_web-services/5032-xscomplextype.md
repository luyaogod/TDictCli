---
title: "XSComplexType"
source: "fgl-topics/c_gws_XML_attribute_XSComplexType.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XSComplexType"
type: "concept"
description: "Map a BDL RECORD type definition to an XML Schema complexType . You can have one member as a nested sequence or choice, or as an XMLList array with a nested sequence or choice as the array's elements; ..."
---

# XSComplexType

Map a BDL RECORD type definition to an XML Schema [complexType](http://www.w3.org/TR/xmlschema-1/#Complex_Type_Definitions).

You can have one member as a nested sequence or choice, or as an XMLList array with a nested
sequence or choice as the array's elements; all other members must have an XMLAttribute
attribute. If not, the compiler raises an error.

## Example

```
TYPE mycomplextype RECORD ATTRIBUTES(XSComplexType,
  XSTypeName="MyComplexType",XSTypeNamespace="http://tempuri.org")
         name DYNAMIC ARRAY ATTRIBUTES(XMLList) OF RECORD 
              ATTRIBUTES(XMLSequence="nested")
     firstname STRING ATTRIBUTES(XMLName="FirstName"),
     lastname  STRING ATTRIBUTES(XMLName="LastName")
END RECORD,
  date  DATE ATTRIBUTES(XMLAttribute,XMLName="Date")
END RECORD
```

```
<xsd:schema xmlns:xsd="http://www.w3.org/2001/XMLSchema"
   targetNamespace="http://tempuri.org" elementFormDefault="qualified" >
  <xsd:complexType name="MyComplexType">
    <xsd:sequence maxOccurs="unbounded">
      <xsd:element name="FirstName" type="xsd:string" />
      <xsd:element name="LastName"  type="xsd:string" />
    </xsd:sequence>
    <xsd:attribute name="Date" type="xsd:date" use="required"/>
  </xsd:complexType>
</xsd:schema>
```
