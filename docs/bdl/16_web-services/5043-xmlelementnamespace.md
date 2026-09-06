---
title: "XMLElementNamespace"
source: "fgl-topics/c_gws_XML_attribute_XMLElementNamespace.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLElementNamespace"
type: "concept"
description: "Define the default namespace of all members of a record also defined as XML elements. Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\", XMLNamespace=\"http://tempuri.org\", ..."
---

# XMLElementNamespace

Define the default namespace of all members of a record also defined as XML elements.

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root",
    XMLNamespace="http://tempuri.org",
    XMLElementNamespace="http://www.mycompany.com")
  val1  FLOAT   ATTRIBUTES(XMLElement,XMLName="Val1"),
  val2  INTEGER ATTRIBUTES(XMLElement,XMLName="Val2"),
  attr  STRING  ATTRIBUTES(XMLAttribute,XMLName="Attr")
               END RECORD
```

```
<fjs1:Root xmlns:fjs1="http://tempuri.org" Attr="Hello"
 xmlns:fjs2="http://www.mycompany.com">
  <fjs2:Val1>0.5</fjs2:Val1>
  <fjs2:Val2>-18547</fjs2:Val2>
</fjs1:Root>
```
