---
title: "XSDUnsignedByte"
source: "fgl-topics/c_gws_XML_attributes_050.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDUnsignedByte"
type: "concept"
description: "Map BDL SMALLINT or BIGINT to XML Schema unsignedByte . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 SMALLINT ATTRIBUTES(XSDUnsignedByte,XMLName=\"Val\") END RECORD <Root> <Val>254</Val> ..."
---

# XSDUnsignedByte

Map BDL SMALLINT or BIGINT to XML Schema [unsignedByte](http://www.w3.org/TR/xmlschema-2/#unsignedByte).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 SMALLINT ATTRIBUTES(XSDUnsignedByte,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>254</Val>
</Root>
```
