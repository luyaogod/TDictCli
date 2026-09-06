---
title: "XSDEnumeration"
source: "fgl-topics/c_gws_XML_attributes_078.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDEnumeration"
type: "concept"
description: "Restrict the allowed value-space to a list of values separated by the characters | . To escape the separator character, simply double it like the following || This attribute can be set on any simple ..."
---

# XSDEnumeration

Restrict the allowed value-space to a list of values separated by the characters **|**.

1. To escape the separator character, simply double it like the following
   **||**
2. This attribute can be set on any simple BDL variable excepted on
   XSDBoolean.

## Example

```
DEFINE myStr STRING ATTRIBUTES(XSDString,
  XSDEnumeration="one|two|three|four", XMLName="MyString")
```

```
DEFINE myDec DECIMAL(3,1) ATTRIBUTES(XSDDecimal,
  XSDEnumeration="12.1|11.8|-24.7", XMLName="MyDecimal")
```
