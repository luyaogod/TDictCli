---
title: "XSDGYear"
source: "fgl-topics/c_gws_XML_attributes_025.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDGYear"
type: "concept"
description: "Map BDL DATETIME to XML Schema gYear . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATETIME YEAR TO YEAR ATTRIBUTES(XSDGYear,XMLName=\"Val\") END RECORD <Root> <Val>2006</Val> </Root>"
---

# XSDGYear

Map BDL DATETIME to XML Schema [gYear](http://www.w3.org/TR/xmlschema-2/#gYear).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATETIME YEAR TO YEAR ATTRIBUTES(XSDGYear,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>2006</Val>
</Root>
```
