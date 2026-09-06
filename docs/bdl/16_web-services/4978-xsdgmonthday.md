---
title: "XSDGMonthDay"
source: "fgl-topics/c_gws_XML_attributes_024.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDGMonthDay"
type: "concept"
description: "Map BDL DATETIME to XML Schema gMonthDay . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME MONTH TO DAY ATTRIBUTES(XSDGMonthDay,XMLName=\"Val\") END RECORD <Root> <Val>--12-31</Val> ..."
---

# XSDGMonthDay

Map BDL DATETIME to XML Schema [gMonthDay](http://www.w3.org/TR/xmlschema-2/#gMonthDay).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME MONTH TO DAY ATTRIBUTES(XSDGMonthDay,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>--12-31</Val>
</Root>
```
