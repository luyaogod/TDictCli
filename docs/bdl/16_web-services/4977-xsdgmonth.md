---
title: "XSDGMonth"
source: "fgl-topics/c_gws_XML_attributes_023.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDGMonth"
type: "concept"
description: "Map BDL DATETIME to XML Schema gMonth . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME MONTH TO MONTH ATTRIBUTES(XSDGMonth,XMLName=\"Val\") END RECORD <Root> <Val>--12</Val> ..."
---

# XSDGMonth

Map BDL DATETIME to XML Schema [gMonth](http://www.w3.org/TR/xmlschema-2/#gMonth).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME MONTH TO MONTH ATTRIBUTES(XSDGMonth,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>--12</Val>
</Root>
```
