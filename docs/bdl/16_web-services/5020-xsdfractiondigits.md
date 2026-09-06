---
title: "XSDFractionDigits"
source: "fgl-topics/c_gws_XML_attributes_086.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDFractionDigits"
type: "concept"
description: "Define the maximum number of digits allowed on the fraction part of a numeric data type. The fraction digits value set on a BDL data type without XSDDecimal set can only be 0. On a BDL DECIMAL, the ..."
---

# XSDFractionDigits

Define the maximum number of digits allowed on the fraction part of a numeric data type.

1. The fraction digits value set on a BDL data type without XSDDecimal set can only
   be 0.
2. On a BDL DECIMAL, the fraction digits value cannot be lower than the scale of
   the BDL DECIMAL itself, and must be lower than the XSDTotalDigits value if
   set.

## Example

```
DEFINE myCode SMALLINT ATTRIBUTES(XSDShort, XSDFractionDigits="0",
  XSDMaxExclusive="1000", XMLName="MyCode")
```

```
DEFINE myRate DECIMAL(4,2) ATTRIBUTES(XSDDecimal, XSDTotalDigits="5",
  XSDFractionDigits="3", XSDMaxExclusive="299.99", XMLName="MyRate")
```
