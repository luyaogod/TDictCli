---
title: "XMLSimpleContent"
source: "fgl-topics/c_gws_XML_attribute_XMLSimpleContent.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLSimpleContent"
type: "concept"
description: "Map a BDL RECORD to an XML Schema simpleContent structure. One member must have the XMLBase attribute; all other members must have an XMLAttribute attribute. If not, the compiler complains. Example ..."
---

# XMLSimpleContent

Map a BDL RECORD to an XML Schema [simpleContent](http://www.w3.org/TR/xmlschema-1/#element-simpleContent)
structure.

One member must have the XMLBase attribute; all other members must have an XMLAttribute
attribute. If not, the compiler complains.

## Example

```
DEFINE mysimpletype RECORD ATTRIBUTES(XMLSimpleContent,XMLName="Root")
  base  STRING  ATTRIBUTES(XMLBase),
  val1  INTEGER ATTRIBUTES(XMLAttribute,XMLName="Val1"),
  val2  FLOAT   ATTRIBUTES(XMLAttribute,XMLName="Val2")
END RECORD
```

```
<Root Val1="148" Val2="25.8">
  Hello
</Root>
```
