---
title: "XMLSequence (Optional)"
source: "fgl-topics/c_gws_XML_attribute_XMLSequence.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLSequence (Optional)"
type: "concept"
description: "Map a BDL RECORD to an XML Schema sequence structure. The order in which the record members appear in the XML document must match the order of the BDL RECORD. The XMLSequence attribute also supports a ..."
---

# XMLSequence (Optional)

Map a BDL RECORD to an XML Schema [sequence](http://www.w3.org/TR/xmlschema-1/#element-sequence)
structure. The order in which the record members appear in the XML document must match the order of
the BDL RECORD. The XMLSequence attribute also supports a "nested" value that removes the
surrounding XML tag.

> **Important:**
>
> Nested sequence records cannot be defined as main variables; there must always be a
> surrounding variable.

## Example

```
DEFINE mysequence RECORD ATTRIBUTES(XMLSequence,XMLName="Root")
  val1  INTEGER   ATTRIBUTES(XMLName="Val1"),
  val2  FLOAT     ATTRIBUTES(XMLAttribute,XMLName="Val2"),
  val3  STRING    ATTRIBUTES(XMLName="Val3")
END RECORD
```

```
<Root Val2="25.8">
  <Val1>-859</Val1>
  <Val3>Hello world</Val3>
</Root>
```

Nested example:

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1  INTEGER   ATTRIBUTES(XMLName="Val1"),
  val2  FLOAT     ATTRIBUTES(XMLAttribute,XMLName="Val2"),
  sequence RECORD   ATTRIBUTES(XMLSequence="nested")
    seq1   INTEGER   ATTRIBUTES(XMLName="SeqOne"),
    seq2   FLOAT     ATTRIBUTES(XMLName="SeqTwo")
  END RECORD,
  val3  STRING    ATTRIBUTES(XMLName="Val3")
END RECORD
```

```
<Root Val2="25.8">
  <Val1>148</Val1>
  <SeqOne>6584</SeqOne>
  <SeqTwo>85.597</SeqTwo>
  <Val3>Hello world</Val3>
</Root>
```
