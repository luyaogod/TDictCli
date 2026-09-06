---
title: "XSDPattern"
source: "fgl-topics/c_gws_XML_attributes_080.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes > XSDPattern"
type: "concept"
description: "Define a regular expression to match. The value has to be serialized or deserialized without any error. The regular expression is defined in the XML Schema Part 2 specification available here . ..."
---

# XSDPattern

Define a regular expression to match. The value has to be serialized or deserialized without any
error.

1. The regular expression is defined in the XML Schema Part 2 specification
   available [here](http://www.w3.org/TR/xmlschema-2/#regexs).
2. Backslash characters '\' in a regular expression must be escaped by duplicating
   it.

## Example

```
DEFINE myStr STRING ATTRIBUTES(XSDString, XSDPattern="A.*Z",
  XMLName="MyString")
```

```
DEFINE myZipCode INTEGER ATTRIBUTES(XSDInt, XSDPattern="[0-9]{5}",
  XMLName="MyZipCode")
```

```
DEFINE myOtherZipCode INTEGER ATTRIBUTES(XSDInt,
 XSDPattern="\\d{5}", XMLName="myOtherZipCode") # regex is \d{5} see note
```
