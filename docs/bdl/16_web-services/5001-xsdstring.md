---
title: "XSDString"
source: "fgl-topics/c_gws_XML_attributes_047.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDString"
type: "concept"
description: "Map BDL STRING, CHAR, TEXT or VARCHAR to XML Schema string . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 STRING ATTRIBUTES(XSDString,XMLName=\"Val\") END RECORD <Root> <Val>Hello world, ..."
---

# XSDString

Map BDL STRING, CHAR, TEXT or VARCHAR to XML Schema [string](http://www.w3.org/TR/xmlschema-2/#string).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 STRING ATTRIBUTES(XSDString,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>Hello world, how are you ?</Val>
</Root>
```
