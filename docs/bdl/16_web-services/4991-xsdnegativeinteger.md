---
title: "XSDNegativeInteger"
source: "fgl-topics/c_gws_XML_attributes_037.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDNegativeInteger"
type: "concept"
description: "Map BDL DECIMAL to XML Schema negativeInteger . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDNegativeInteger,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDNegativeInteger

Map BDL DECIMAL to XML Schema [negativeInteger](http://www.w3.org/TR/xmlschema-2/#negativeInteger).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDNegativeInteger,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>-4828</Val>
</Root>
```
