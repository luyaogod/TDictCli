---
title: "XSDMaxLength"
source: "fgl-topics/c_gws_XML_attributes_077.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDMaxLength"
type: "concept"
description: "Restrict the length of the data to the maximum number of XML characters allowed when set on a BDL STRING, VARCHAR, CHAR or TEXT, or the number of bytes allowed when set on a BDL BYTE. XSDMinLength and ..."
---

# XSDMaxLength

Restrict the length of the data to the maximum number of XML characters allowed when set on a BDL
STRING, VARCHAR, CHAR or TEXT, or the number of bytes allowed when set on a BDL BYTE.

1. XSDMinLength and XSDMaxLength can be used together, but XSDMaxLength value must
   be greater then XSDMinLength
2. XSDMaxLength cannot be used with XSDLength

## Example

```
DEFINE myStr STRING ATTRIBUTES(XSDString, XSDMaxLength="12",
  XMLName="MyString")
```

```
DEFINE myByte BYTE ATTRIBUTES(XSDBase64Binary, XSDMaxLength="8000",
  XMLName="MyPicture")
```
