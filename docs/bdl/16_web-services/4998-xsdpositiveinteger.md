---
title: "XSDPositiveInteger"
source: "fgl-topics/c_gws_XML_attributes_044.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDPositiveInteger"
type: "concept"
description: "Map BDL DECIMAL to XML Schema positiveInteger . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(32,0) ATTRIBUTES(XSDPositiveInteger,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDPositiveInteger

Map BDL DECIMAL to XML Schema [positiveInteger](http://www.w3.org/TR/xmlschema-2/#positiveInteger).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(32,0) ATTRIBUTES(XSDPositiveInteger,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>+41893</Val>
</Root>
```
