---
title: "XMLALL"
source: "fgl-topics/c_gws_XML_attribute_XMLAll.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLALL"
type: "concept"
description: "Map a BDL Record to an XML Schema all structure. The order in which the record members appear in the XML document is not significant. Example DEFINE myall RECORD ATTRIBUTES(XMLAll,XMLName=\"Root\") val1 ..."
---

# XMLALL

Map a BDL Record to an XML Schema [all](http://www.w3.org/TR/xmlschema-1/#element-all) structure.

The order in which the record members appear in the XML document is not significant.

## Example

```
DEFINE myall RECORD ATTRIBUTES(XMLAll,XMLName="Root")
  val1  INTEGER ATTRIBUTES(XMLName="Val1"),
  val2  FLOAT   ATTRIBUTES(XMLAttribute,XMLName="Val2"),
  val3  STRING  ATTRIBUTES(XMLName="Val3")
END RECORD
```

```
<Root Val2="25.8">
  <Val3>Hello world</Val3>
  <Val1>148</Val1>
</Root>
```

```
<Root Val2="25.8">
  <Val1>148</Val1>
  <Val3>Hello world</Val3>
</Root>
```
