---
title: "XSDUnsignedLong"
source: "fgl-topics/c_gws_XML_attributes_052.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDUnsignedLong"
type: "concept"
description: "Map BDL DECIMAL to XML Schema unsignedLong . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDUnsignedLong,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDUnsignedLong

Map BDL DECIMAL to XML Schema [unsignedLong](http://www.w3.org/TR/xmlschema-2/#unsignedLong).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDUnsignedLong,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>12678967543233</Val>
</Root>
```
