---
title: "XSDLong"
source: "fgl-topics/c_gws_XML_attributes_034.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDLong"
type: "concept"
description: "Map BDL BIGINT or DECIMAL to XML Schema long . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DECIMAL(19,0) ATTRIBUTES(XSDLong,XMLName=\"Val\") END RECORD <Root> <Val>1267488</Val> </Root>"
---

# XSDLong

Map BDL BIGINT or DECIMAL to XML Schema [long](http://www.w3.org/TR/xmlschema-2/#long).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DECIMAL(19,0) ATTRIBUTES(XSDLong,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>1267488</Val>
</Root>
```
