---
title: "XSDTime"
source: "fgl-topics/c_gws_XML_attributes_048.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDTime"
type: "concept"
description: "Map BDL DATETIME to XML Schema time . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME ATTRIBUTES(XSDTime,XMLName=\"Val\") END RECORD <Root> <Val>23:16:03.589+01:00</Val> </Root>"
---

# XSDTime

Map BDL DATETIME to XML Schema  [time](http://www.w3.org/TR/xmlschema-2/#time).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME ATTRIBUTES(XSDTime,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>23:16:03.589+01:00</Val>
</Root>
```
