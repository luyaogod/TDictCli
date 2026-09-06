---
title: "XSDWhiteSpace"
source: "fgl-topics/c_gws_XML_attributes_079.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDWhiteSpace"
type: "concept"
description: "Perform a XML string manipulation before serialization or deserialization depending on the following possible values: preserve : the XML string is not modified. replace : the XML string is modified by ..."
---

# XSDWhiteSpace

Perform a XML string manipulation before serialization or deserialization
depending on the following possible values:

- **preserve**: the XML string is not modified.
- **replace**: the XML string is modified by replacing each **\n**, **\t**, **\r** by
  a single space.
- **collapse**: the XML string is modified by replacing each **\n**, **\t**, **\r** by
  a single space, then each sequence of several spaces are replaced
  by one single space. Leading and trailing spaces are removed too.

1. The whiteSpace facet is always performed before any other facet constraints, or
   serialization or deserialization process.
2. For any BDL variable except STRING, CHAR and VARCHAR, only collapse is
   allowed.

## Example

```
DEFINE myStr STRING ATTRIBUTES(XSDString, XSDWhiteSpace="replace",
  XMLName="MyString")
```

```
DEFINE myDec DECIMAL(3,1) ATTRIBUTES(XSDDecimal,
  XSDWhiteSpace="collapse", XMLName="MyDecimal")
```
