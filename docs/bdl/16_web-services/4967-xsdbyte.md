---
title: "XSDByte"
source: "fgl-topics/c_gws_XML_attributes_013.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDByte"
type: "concept"
description: "Map BDL TINYINT, SMALLINT or BIGINT to XML Schema byte . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 SMALLINT ATTRIBUTES(XSDByte,XMLName=\"Val\") END RECORD <Root> <Val>-126</Val> ..."
---

# XSDByte

Map BDL TINYINT, SMALLINT or BIGINT to XML Schema [byte](http://www.w3.org/TR/xmlschema-2/#byte).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 SMALLINT ATTRIBUTES(XSDByte,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>-126</Val>
</Root>
```
