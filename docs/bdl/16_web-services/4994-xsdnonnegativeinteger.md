---
title: "XSDNonNegativeInteger"
source: "fgl-topics/c_gws_XML_attributes_040.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDNonNegativeInteger"
type: "concept"
description: "Map BDL DECIMAL to XML Schema nonNegativeInteger . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDNonNegativeInteger,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDNonNegativeInteger

Map BDL DECIMAL to XML Schema [nonNegativeInteger](http://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDNonNegativeInteger,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>1589</Val>
</Root>
```
