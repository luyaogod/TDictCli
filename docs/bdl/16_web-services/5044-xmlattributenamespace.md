---
title: "XMLAttributeNamespace"
source: "fgl-topics/c_gws_XML_attribute_XMLAttributeNamespace.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLAttributeNamespace"
type: "concept"
description: "Define the default namespace of all members of a record also defined as XML attributes. Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\", XMLNamespace=\"http://tempuri.org\", ..."
---

# XMLAttributeNamespace

Define the default namespace of all members of a record also defined as XML attributes.

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root",
    XMLNamespace="http://tempuri.org",
    XMLAttributeNamespace="http://www.mycompany.com")
  val1  FLOAT   ATTRIBUTES(XMLElement,XMLName="Val1"),
  val2  INTEGER ATTRIBUTES(XMLElement,XMLName="Val2"),
  attr1 STRING  ATTRIBUTES(XMLAttribute,XMLName="Attr1"),
  attr2 DATE    ATTRIBUTES(XMLAttribute,
    XMLName="Attr2",XMLNamespace="http://anyuri.org")
END RECORD
```

```
<fjs1:Root xmlns:fjs1="http://tempuri.org"
 fjs2:Attr1="Hello" xmlns:fjs2="http://www.mycompany.com"
 xmlns:fjs3="http://anyuri.org" fjs3:Attr2="2006-06-24">
  <fjs1:Val1>0.5</fjs1:Val1>
  <fjs1:Val2>-18547</fjs1:Val2>
</fjs1:Root>
```
