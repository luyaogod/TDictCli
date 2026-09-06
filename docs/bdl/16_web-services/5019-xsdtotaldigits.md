---
title: "XSDTotalDigits"
source: "fgl-topics/c_gws_XML_attributes_085.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDTotalDigits"
type: "concept"
description: "Define the maximum number of digits allowed on a numeric data type, fraction part inclusive if there is one. The total digits value cannot be equal to or lower then 0. On a BDL decimal, the total ..."
---

# XSDTotalDigits

Define the maximum number of digits allowed on a numeric data type, fraction part inclusive if
there is one.

1. The total digits value cannot be equal to or lower then 0.
2. On a BDL decimal, the total digits value cannot be lower than the precision of
   the BDL decimal itself.
3. Notice that a decimal without any precision and scale value is a decimal(16),
   therefore the total digits value must be equal to or greater than 16.

## Example

```
DEFINE myCode SMALLINT ATTRIBUTES(XSDShort, XSDTotalDigits="4",
  XSDMaxExclusive="1000", XMLName="MyCode")
```

```
DEFINE myRate DECIMAL(4,2) ATTRIBUTES(XSDDecimal, XSDTotalDigits="5",
  XSDMaxExclusive="299.99", XMLName="MyRate")
```
