---
title: "XSDGDay"
source: "fgl-topics/c_gws_XML_attributes_022.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDGDay"
type: "concept"
description: "Map BDL DATETIME to XML Schema gDay . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME DAY TO DAY ATTRIBUTES(XSDGDay,XMLName=\"Val\") END RECORD <Root> <Val>---25</Val> </Root>"
---

# XSDGDay

Map BDL DATETIME to XML Schema [gDay](http://www.w3.org/TR/xmlschema-2/#gDay).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME DAY TO DAY ATTRIBUTES(XSDGDay,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>---25</Val>
</Root>
```
