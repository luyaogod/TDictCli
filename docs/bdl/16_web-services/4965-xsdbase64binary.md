---
title: "XSDBase64binary"
source: "fgl-topics/c_gws_XML_attributes_011.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML type mapping attributes > XSDBase64binary"
type: "concept"
description: "Map BDL BYTE to XML Schema base64binary . Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\") val1 BYTE ATTRIBUTES(XSDBase64binary,XMLName=\"Val\") END RECORD <Root> ..."
---

# XSDBase64binary

Map BDL BYTE to XML Schema [base64binary](http://www.w3.org/TR/xmlschema-2/#base64binary).

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 BYTE ATTRIBUTES(XSDBase64binary,XMLName="Val")
END RECORD
```

```
<Root>
  <Val>F0FFC8D27FF001547FC219E1FFF009F0FFC8D27FF001547D</Val>
</Root>
```
