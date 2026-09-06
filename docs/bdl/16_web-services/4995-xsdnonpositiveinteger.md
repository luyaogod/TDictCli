---
title: "XSDNonPositiveInteger"
source: "fgl-topics/c_gws_XML_attributes_041.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDNonPositiveInteger"
type: "concept"
description: "Map BDL DECIMAL to XML Schema nonPositiveInteger . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDNonPositiveInteger,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDNonPositiveInteger

Map BDL DECIMAL to XML Schema [nonPositiveInteger](http://www.w3.org/TR/xmlschema-2/#nonPositiveInteger).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDNonPositiveInteger,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>-8574</Val>
</Root>
```
