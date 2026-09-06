---
title: "XSDShort"
source: "fgl-topics/c_gws_XML_attributes_046.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDShort"
type: "concept"
description: "Map BDL SMALLINT or BIGINT to XML Schema short . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 SMALLINT ATTRIBUTES(XSDShort,XMLName=\"Val\") END RECORD <Root> <Val>12678</Val> </Root>"
---

# XSDShort

Map BDL SMALLINT or BIGINT to XML Schema [short](http://www.w3.org/TR/xmlschema-2/#short).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 SMALLINT ATTRIBUTES(XSDShort,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>12678</Val>
</Root>
```
