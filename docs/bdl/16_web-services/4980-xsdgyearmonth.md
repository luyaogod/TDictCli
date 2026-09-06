---
title: "XSDGYearMonth"
source: "fgl-topics/c_gws_XML_attributes_026.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDGYearMonth"
type: "concept"
description: "Map BDL DATETIME to XML Schema gYearMonth . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME YEAR TO MONTH ATTRIBUTES(XSDGYearMonth,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDGYearMonth

Map BDL DATETIME to XML Schema [gYearMonth](http://www.w3.org/TR/xmlschema-2/#gYearMonth).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME YEAR TO MONTH ATTRIBUTES(XSDGYearMonth,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>2006-06</Val>
</Root>
```
