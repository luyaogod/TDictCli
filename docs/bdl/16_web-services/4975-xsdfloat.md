---
title: "XSDFloat"
source: "fgl-topics/c_gws_XML_attributes_021.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDFloat"
type: "concept"
description: "Map BDL SMALLFLOAT to XML Schema float . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 SMALLFLOAT ATTRIBUTES(XSDFloat,XMLName=\"Val\") END RECORD <Root> <Val>126.435</Val> </Root>"
---

# XSDFloat

Map BDL SMALLFLOAT to XML Schema [float](http://www.w3.org/TR/xmlschema-2/#float).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 SMALLFLOAT ATTRIBUTES(XSDFloat,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>126.435</Val>
</Root>
```
