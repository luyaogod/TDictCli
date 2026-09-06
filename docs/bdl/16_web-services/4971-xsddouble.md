---
title: "XSDDouble"
source: "fgl-topics/c_gws_XML_attributes_017.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDDouble"
type: "concept"
description: "Map BDL FLOAT to XML Schema double . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 FLOAT ATTRIBUTES(XSDDouble,XMLName=\"Val\") END RECORD <Root> <Val>12.78e-2</Val> </Root>"
---

# XSDDouble

Map BDL FLOAT to XML Schema [double](http://www.w3.org/TR/xmlschema-2/#double).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 FLOAT ATTRIBUTES(XSDDouble,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>12.78e-2</Val>
</Root>
```
