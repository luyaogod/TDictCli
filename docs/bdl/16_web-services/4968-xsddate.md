---
title: "XSDDate"
source: "fgl-topics/c_gws_XML_attributes_014.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDDate"
type: "concept"
description: "Map BDL DATE or DATETIME to XML Schema date . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 DATE ATTRIBUTES(XSDDate,XMLName=\"Val\") END RECORD <Root> <Val>2006-06-29+01:00</Val> </Root>"
---

# XSDDate

Map BDL DATE or DATETIME to XML Schema [date](http://www.w3.org/TR/xmlschema-2/#date).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 DATE ATTRIBUTES(XSDDate,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>2006-06-29+01:00</Val>
</Root>
```
