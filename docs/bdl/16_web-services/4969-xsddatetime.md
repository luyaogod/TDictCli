---
title: "XSDDateTime"
source: "fgl-topics/c_gws_XML_attributes_015.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDDateTime"
type: "concept"
description: "Map BDL DATETIME to XML Schema dateTime . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME ATTRIBUTES(XSDDateTime,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDDateTime

Map BDL DATETIME to XML Schema [dateTime](http://www.w3.org/TR/xmlschema-2/#dateTime).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME ATTRIBUTES(XSDDateTime,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>2006-06-29T09:35:26.13584+01:00</Val>
</Root>
```
