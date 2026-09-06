---
title: "Serialize XML from a dynamic array"
source: "fgl-topics/c_gws_XML_attributes_serialize_dynamic_array.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > Serialize XML from a dynamic array"
type: "concept"
---

# Serialize XML from a dynamic array

> The XMLList and XMLName attributes can be used to serialize dynamic arrays to XML and vice versa.

The `XMLList` and `XMLName` attributes attributes have distinct
uses depending on the output you want:

- The [XMLList](5033-xmllist.md) attribute is used to
  output a list of the same elements from an array, without a surrounding
  tag. Typically, you use it inside another XML element to map a one
  dimensional array to an XML schema element that has more than one
  occurrence.

  However, the output from `XMLList` on its own is not a
  valid XML document because there is no parent node. To output a valid
  XML document, the attribute must be used inside another variable, such
  as a record. Compare the output in Example 1: Array of the same elements with XMLList and Example 3: Array of the same elements with XMLList inside a record.

  On the other hand, an array without `XMLList` always
  produces a surrounding tag and thus a valid XML document. See Example 2: Array of the same elements with XMLName but without XMLList and Example 4: Array of the same elements with XMLName inside a record.
- The [XMLName](5037-xmlname.md) attribute's function is
  to provide a name for XML elements. If you do not use the `XMLName` attribute, the tags used in the
  XML output are the names of the Genero BDL variables. Or if there are no variables, they are named
  "<element>".

> **Tip:**
>
> The variables with attributes you use to output to XML, may also be used to
> load an XML document into `DYNAMIC ARRAY`.

## Example 1: Array of the same elements with XMLList

This example shows how an array defined with `XMLList` outputs a list of the same
elements. Typically, this array is defined inside another variable (see Example 3: Array of the same elements with XMLList inside a record), because the resulting XML document is not valid otherwise. The
`XMLName` attribute is set to name the elements within the array.

```
DEFINE a RECORD
    list DYNAMIC ARRAY ATTRIBUTES(XMLList) OF STRING ATTRIBUTE(XMLName="MyElt")
END RECORD
```

This would produce an XML output like this, which is not a valid XML document because there is no
parent
node:

```
<MyElt>One</MyElt><MyElt>Two</MyElt>
```

When
viewed in an XML viewer, it would display as shown.

```
<MyElt>One</MyElt>
<MyElt>Two</MyElt>
```

## Example 2: Array of the same elements with `XMLName` but without `XMLList`

In this example an array of the same elements is defined without `XMLList`. This
produces output that is a valid XML document because the `XMLName`
attribute is set to tag the array with a surrounding tag, and to tag the elements
within the array.

If you do not use the `XMLName` attribute, the tags used in the
XML output are the names of the Genero BDL variables. Or if there are no variables, they are named
"<element>".

```
DEFINE arr DYNAMIC ARRAY ATTRIBUTE(XMLName="SurroundingTag") OF STRING ATTRIBUTE(XMLName="MyElt")
```

The
XML output may look like this, which is a valid XML document because there is a
parent
node:

```
<SurroundingTag><MyElt>One</MyElt><MyElt>Two</MyElt></SurroundingTag>
```

When
viewed in an XML viewer, it would display as shown.

```
<SurroundingTag>
   <MyElt>One</MyElt>
   <MyElt>Two</MyElt>
</SurroundingTag>
```

## Example 3: Array of the same elements with XMLList inside a record

In this example an array of the same elements is defined with `XMLList` inside a
record, in order to produce output that is a valid XML document. Several lists may
be included in the record.

```
DEFINE myVar1 RECORD ATTRIBUTES(XMLName="Root")
   val1  INTEGER ATTRIBUTES(XMLName="Val1"),
   list  DYNAMIC ARRAY ATTRIBUTES(XMLList) OF STRING ATTRIBUTES(XMLName="MyElt"),
   val2  FLOAT   ATTRIBUTES(XMLName="Val2")
END RECORD
```

This would produce XML output like this, which is a valid XML document because
there is a parent
node:

```
<Root><Val1>148</Val1><MyElt>hello</MyElt><MyElt>there</MyElt><Val2>0.58</Val2></Root>
```

When
viewed in an XML viewer, it would display as shown.

```
<Root>
  <Val1>148</Val1>
  <MyElt>hello</MyElt>
  <MyElt>there</MyElt>
  <Val2>0.58</Val2>
</Root>
```

## Example 4: Array of the same elements with `XMLName` inside a record

In this example an array of the same elements is defined without `XMLList` inside
a record. The `XMLName` attribute is set to tag the array with a
surrounding tag, and to tag the elements within the array.

If you do not use the `XMLName` attribute, the tags used in the
XML output are the names of the Genero BDL variables. Or if there are no variables, they are named
"<element>".

```
DEFINE myVar2 RECORD ATTRIBUTES(XMLName="Root")
   val1  INTEGER ATTRIBUTES(XMLName="Val1"),
   arr  DYNAMIC ARRAY ATTRIBUTES(XMLName="Sample") OF STRING ATTRIBUTES(XMLName="MyElt"),
   val2  FLOAT   ATTRIBUTES(XMLName="Val2")
END RECORD
```

This would produce XML output like this, which is a valid XML document because
there is a parent
node:

```
<Root><Val1>148</Val1><Sample><MyElt>hello</MyElt><MyElt>there</MyElt></Sample><Val2>0.58</Val2></Root>
```

When
viewed in an XML viewer, it would display as shown.

```
<Root>
   <Val1>148</Val1>
   <Sample>
	<MyElt>hello</MyElt>
	<MyElt>there</MyElt>
   </Sample>
   <Val2>0.58</Val2>
</Root>
```
