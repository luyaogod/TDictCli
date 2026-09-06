---
title: "XMLNamespace"
source: "fgl-topics/c_gws_XML_attribute_XMLNamespace.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLNamespace"
type: "concept"
description: "Define the namespace of a variable in an XML document. If the attribute is set on a Record, by default all members defined as XMLElement of that record are in the same namespace. If the attribute is ..."
---

# XMLNamespace

Define the namespace of a variable in an XML document.

1. If the attribute is set on a Record, by default all members defined as [XMLElement](5024-xmlelement-optional.md) of that record are in the same
   namespace.
2. If the attribute is set on an Array, by default all elements defined as XMLElement of
   that array are in the same namespace.
3. The attribute cannot be set on a type definition.

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root", XMLNamespace="http://tempuri.org")
  attr1 INTEGER ATTRIBUTES(XMLAttribute,XMLName="Attr1"),
  val1  FLOAT   ATTRIBUTES(XMLName="Val1", XMLNamespace="http://www.mycompany.com"),
  val2  INTEGER ATTRIBUTES(XMLName="Val2"),
  attr2 STRING  ATTRIBUTES(XMLAttribute, XMLName="Attr2",
    XMLNamespace="http://anyuri.org")
END RECORD
```

```
<fjs1:Root xmlns:fjs1="http://tempuri.org" Attr1="158"
 xmlns:fjs3="http://anyuri.org" fjs3:Attr2="Hello">
  <fjs2:Val1 xmlns:fjs2="http://www.mycompany.com">0.5</fjs2:Val1>
  <fjs1:Val2>-18547</fjs1:Val2>
</fjs1:Root>
```
