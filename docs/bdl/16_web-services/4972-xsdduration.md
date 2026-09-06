---
title: "XSDDuration"
source: "fgl-topics/c_gws_XML_attributes_018.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDDuration"
type: "concept"
description: "Map BDL INTERVAL to XML Schema duration . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 INTERVAL DAY TO SECOND ATTRIBUTES(XSDDuration,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDDuration

Map BDL INTERVAL to XML Schema [duration](http://www.w3.org/TR/xmlschema-2/#duration).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 INTERVAL DAY TO SECOND ATTRIBUTES(XSDDuration,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>P3DT10H30M45S</Val>
</Root>
```
