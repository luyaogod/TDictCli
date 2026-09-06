---
title: "XSDHexBinary"
source: "fgl-topics/c_gws_XML_attributes_027.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDHexBinary"
type: "concept"
description: "Map BDL BYTE to XML Schema hexBinary . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 BYTE ATTRIBUTES(XSDHexBinary,XMLName=\"Val\") END RECORD <Root> <Val>0FB6</Val> </Root>"
---

# XSDHexBinary

Map BDL BYTE to XML Schema [hexBinary](http://www.w3.org/TR/xmlschema-2/#hexBinary).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 BYTE ATTRIBUTES(XSDHexBinary,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>0FB6</Val>
</Root>
```
