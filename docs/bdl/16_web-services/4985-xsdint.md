---
title: "XSDInt"
source: "fgl-topics/c_gws_XML_attributes_031.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDInt"
type: "concept"
description: "Map BDL INTEGER or BIGINT to XML Schema int . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 INTEGER ATTRIBUTES(XSDInt,XMLName=\"Val\") END RECORD <Root> <Val>-1258</Val> </Root>"
---

# XSDInt

Map BDL INTEGER or BIGINT to XML Schema [int](http://www.w3.org/TR/xmlschema-2/#int).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 INTEGER ATTRIBUTES(XSDInt,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>-1258</Val>
</Root>
```
