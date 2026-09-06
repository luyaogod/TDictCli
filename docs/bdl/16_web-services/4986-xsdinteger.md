---
title: "XSDInteger"
source: "fgl-topics/c_gws_XML_attributes_032.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDInteger"
type: "concept"
description: "Map BDL DECIMAL to XML Schema integer . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDInteger,XMLName=\"Val\") END RECORD <Root> <Val>12678</Val> </Root>"
---

# XSDInteger

Map BDL DECIMAL to XML Schema [integer](http://www.w3.org/TR/xmlschema-2/#integer).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDInteger,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>12678</Val>
</Root>
```
