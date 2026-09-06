---
title: "XSDMinExclusive"
source: "fgl-topics/c_gws_XML_attributes_083.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDMinExclusive"
type: "concept"
description: "Define the minimum exclusive value allowed and depending on the data type where it is set, namely all numeric, date and time data types. The minimum value cannot exceed or be equal to the implicit ..."
---

# XSDMinExclusive

Define the minimum exclusive value allowed and depending on the
data type where it is set, namely all numeric, date and time data
types.

The minimum value cannot exceed or be equal to the implicit minimum value supported by
the data type itself or the compiler will complain. For instance, with XSDShort the
minimum value is -32768.

## Example

```
DEFINE myCode SMALLINT ATTRIBUTES(XSDShort, XSDMinExclusive="-1000",
  XMLName="MyCode")
```

```
DEFINE myRate DECIMAL(4,2) ATTRIBUTES(XSDDecimal, XSDMinExclusive="100.01",
  XMLName="MyRate")
```
