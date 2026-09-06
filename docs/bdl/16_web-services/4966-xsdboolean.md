---
title: "XSDBoolean"
source: "fgl-topics/c_gws_XML_attributes_012.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDBoolean"
type: "concept"
description: "Map BDL BOOLEAN, SMALLINT or INTEGER to XML Schema boolean . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 INTEGER ATTRIBUTES(XSDBoolean,XMLName=\"Val\") END RECORD <Root> <Val>true</Val> ..."
---

# XSDBoolean

Map BDL BOOLEAN, SMALLINT or INTEGER to XML Schema [boolean](http://www.w3.org/TR/xmlschema-2/#boolean).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
    val1 INTEGER ATTRIBUTES(XSDBoolean,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>true</Val>
</Root>
```
