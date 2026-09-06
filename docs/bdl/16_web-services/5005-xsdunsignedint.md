---
title: "XSDUnsignedInt"
source: "fgl-topics/c_gws_XML_attributes_051.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDUnsignedInt"
type: "concept"
description: "Map BDL BIGINT or DECIMAL to XML Schema unsignedInt . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDUnsignedInt,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDUnsignedInt

Map BDL BIGINT or DECIMAL to XML Schema [unsignedInt](http://www.w3.org/TR/xmlschema-2/#unsignedInt).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDUnsignedInt,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>1267896754</Val>
</Root>
```
