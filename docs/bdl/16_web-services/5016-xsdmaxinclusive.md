---
title: "XSDMaxInclusive"
source: "fgl-topics/c_gws_XML_attributes_082.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDMaxInclusive"
type: "concept"
description: "Define the maximum inclusive value allowed and depending on the data type where it is set, namely all numeric, date and time data types. The maximum value cannot exceed the implicit maximum value ..."
---

# XSDMaxInclusive

Define the maximum inclusive value allowed and depending on the
data type where it is set, namely all numeric, date and time data
types.

The maximum value cannot exceed the implicit maximum value supported by the data type
itself or the compiler will complain. For instance, with XSDShort the maximum value is
32767.

## Example

```
DEFINE myCode SMALLINT ATTRIBUTES(XSDShort, XSDMaxInclusive="1000",
  XMLName="MyCode")
```

```
DEFINE myRate DECIMAL(4,2) ATTRIBUTES(XSDDecimal, XSDMaxInclusive="299.99",
  XMLName="MyRate")
```
