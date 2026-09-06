---
title: "XSDDecimal"
source: "fgl-topics/c_gws_XML_attributes_016.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDDecimal"
type: "concept"
description: "Map BDL DECIMAL to XML Schema decimal . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(5,3) ATTRIBUTES(XSDDecimal,XMLName=\"Val\") END RECORD <Root> <Val>12.345</Val> </Root>"
---

# XSDDecimal

Map BDL DECIMAL to XML Schema [decimal](http://www.w3.org/TR/xmlschema-2/#decimal).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(5,3) ATTRIBUTES(XSDDecimal,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>12.345</Val>
</Root>
```
