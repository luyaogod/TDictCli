---
title: "XSDUnsignedShort"
source: "fgl-topics/c_gws_XML_attributes_053.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDUnsignedShort"
type: "concept"
description: "Map BDL INTEGER or BIGINT to XML Schema unsignedShort . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 INTEGER ATTRIBUTES(XSDUnsignedShort,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDUnsignedShort

Map BDL INTEGER or BIGINT to XML Schema [unsignedShort](http://www.w3.org/TR/xmlschema-2/#unsignedShort).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 INTEGER ATTRIBUTES(XSDUnsignedShort,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>65535</Val>
</Root>
```
