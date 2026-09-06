---
title: "XMLList"
source: "fgl-topics/c_gws_XML_attribute_XMLList.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLList"
type: "concept"
description: "Map a one dimensional array to an XML Schema element that has more than one occurrence. To map a one-dimensional array to an XML Schema element that allows multiple occurrences, the XMLList attribute ..."
---

# XMLList

Map a one dimensional array to an XML Schema element that has more than one occurrence.

To map a one-dimensional array to an XML Schema element that allows multiple occurrences, the
`XMLList` attribute must be used correctly. It is important to note that you cannot
define an `XMLList` attribute directly on a main array; instead, the array must be
placed within a record, which provides the necessary parent element to create a valid XML
document.

The `XMLList` attribute enables the generation of correctly tagged elements from a
dynamic array. The `XMLName` attribute must
be set on the array element to define the tags in the output.

```
list DYNAMIC ARRAY ATTRIBUTES(XMLList) OF STRING ATTRIBUTE(XMLName="MyElt")
```

This
definition would produce the following XML output when viewed in an XML viewer:

```
<MyElt>One</MyElt>
<MyElt>Two</MyElt>
...
```

However, this output is not a valid XML document because it lacks a parent node. To
ensure the output is a valid XML document, the array must be encapsulated within a record, as
demonstrated in the following example:

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1  INTEGER ATTRIBUTES(XMLName="Val1"),
  list  DYNAMIC ARRAY ATTRIBUTES(XMLList) OF STRING ATTRIBUTES(XMLName="Element"),
  val2  FLOAT   ATTRIBUTES(XMLName="Val2")
END RECORD
```

For example, if the output is viewed in an XML viewer, it would display as
shown.

```
<Root>
  <Val1>148</Val1>
  <Element>hello</Element>
  <Element>how</Element>
  <Element>are</Element>
  <Element>you</Element>
  <Val2>0.58</Val2>
</Root>
```

## Related links

**Related concepts**  

[Serialize XML from a dynamic array](5046-serialize-xml-from-a-dynamic-array.md "The XMLList and XMLName attributes can be used to serialize dynamic arrays to XML and vice versa.")
