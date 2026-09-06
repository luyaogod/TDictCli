---
title: "XMLChoice"
source: "fgl-topics/c_gws_XML_attribute_XMLChoice.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLChoice"
type: "concept"
description: "Map a BDL Record to an XML Schema choice structure. The choice of the record's member is performed at runtime, and changes dynamically depending on a mandatory member. This specific member must be of ..."
---

# XMLChoice

Map a BDL Record to an XML Schema [choice](http://www.w3.org/TR/xmlschema-1/#element-choice)
structure. The choice of the record's member is performed at runtime, and changes dynamically
depending on a mandatory member. This specific member must be of type SMALLINT or INTEGER, and
have an [XMLSelector](5034-xmlselector.md) attribute set.

The XMLChoice attribute also supports a "nested" value that removes the surrounding XML
tag.

1. The compiler indexes all elements of a record with an `XMLChoice`
   attribute but the Genero Web Server (GWS) only serializes the element set by the selector
   value as the member to be serialized. If the selected element has no
   `XMLName` attribute, it will use the Genero BDL name as the XML tag.
2. Nested choice records cannot be defined as main variables; there must always be a
   surrounding variable.

## Example 1: Record with XMLChoice

In the sample record, note the index value of the record element shown in the comment with
the hash symbol ( # ).

```
DEFINE mychoice RECORD ATTRIBUTES(XMLChoice,XMLName="Root")
  val1  INTEGER   ATTRIBUTES(XMLName="Val1"),              #1 
  val2  FLOAT     ATTRIBUTES(XMLAttribute,XMLName="Val2"), #2 
  sel   SMALLINT  ATTRIBUTES(XMLSelector),                 #3 selector node
  val3  STRING    ATTRIBUTES(XMLName="Val3")               #4 
END RECORD
```

Case where "sel" value is set to 4 (`LET mychoice.sel = 4`). The XML output
is:

```
<Root Val2="25.8">
  <Val3>Hello world</Val3>
</Root>
```

Case where "sel" value is set to 1 (`LET mychoice.sel = 1`). The XML output
is:

```
<Root Val2="25.8">
  <Val1>148</Val1>
</Root>
```

## Example 2: Record with nested XMLChoice record

In this example, the `XMLChoice` attribute is set on the choice
record, which is a record nested in the myvar record. It has a
selector node (nestedSel), which is set with the
`XMLSelector` attribute. The index value of the nested record elements is
shown in the comment with the hash symbol ( # ) for clarity. The selector node is also part
of the index.

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1  INTEGER   ATTRIBUTES(XMLName="Val1"),               
  val2  FLOAT     ATTRIBUTES(XMLAttribute,XMLName="Val2"),  
  choice RECORD   ATTRIBUTES(XMLChoice="nested")
    choice1   INTEGER   ATTRIBUTES(XMLName="ChoiceOne"),    #1 
    choice2   FLOAT     ATTRIBUTES(XMLName="ChoiceTwo"),    #2 
    nestedSel SMALLINT  ATTRIBUTES(XMLSelector)             #3 selector node
  END RECORD,
  val3  STRING    ATTRIBUTES(XMLName="Val3")
END RECORD
```

Case where "nestedSel" value is set to 1 (`LET myVar.choice.nestedSel = 1`).
The XML output is:

```
<Root Val2="25.8">
  <Val1>148</Val1>
  <ChoiceOne>6584</ChoiceOne>
  <Val3>Hello world</Val3>
</Root>
```

Case where "nestedSel" value is set to 2 (`LET myVar.choice.nestedSel = 2`).
The XML output is:

```
<Root Val2="25.8">
  <Val1>148</Val1>
  <ChoiceTwo>85.8</ChoiceTwo>
  <Val3>Hello world</Val3>
</Root>
```

## Example 3: Using `XMLChoice="nested"` for substitutionGroup

In this example the
`XMLChoice="nested"` option produces an XML output supporting a
`substitutionGroup`. If you have a WSDL containing a substitutionGroup as
in the following XSD schema example, the "name" element can be represented as a simple
string type (`simple-name`) or as a more complex type
(`full-name`).

```
<xs:element name="name" />
<xs:element name="simple-name" type="xs:string" substitutionGroup="name" />
<xs:element name="full-name" substitutionGroup="name">
  <xs:complexType>
    <xs:all>
      <xs:element name="first" type="xs:string" minOccurs="0"/>
      <xs:element name="middle" type="xs:string" minOccurs="0"/>
      <xs:element name="last" type="xs:string"/>
    </xs:all>
  </xs:complexType>
</xs:element>

<xs:element name="Identity"/>
  <xs:sequence>
    <xs:element ref="name"/>
  </xs:sequence>
</xs:element>
```

In the next example, a BDL `RECORD`
(Identity) is defined to specify the XML representation for the
schema:

```
DEFINE Identity RECORD ATTRIBUTE(XMLSequence,XMLName="Identity")
         name RECORD ATTRIBUTE(XMLChoice="nested")
           _SELECTOR_ SMALLINT ATTRIBUTE(XMLSelector),
           name STRING,                                             #1
           simple_name STRING ATTRIBUTE(XMLName="simple-name"),     #2
           full_name RECORD ATTRIBUTE(XMLAll,XMLName="full-name")   #3
             first STRING ATTRIBUTE(XMLOptional),
             middle STRING ATTRIBUTE(XMLOptional),
             last STRING 
           END RECORD
         END RECORD
       END RECORD
```

Where:

- The name element is the substitutionGroup tag. The
  `XMLChoice="nested"` attribute is set on this variable. This element is
  not serialized in XML.
- The selector node (\_SELECTOR\_) is set with the
  `XMLSelector` attribute.
- name is a `STRING`.
- simple-name is a `STRING` set with an
  `XMLName` attribute.
- full-name is a `RECORD` set with an
  `XMLAll` and `XMLName` attribute. It has two optional
  `STRING` elements, first and
  middle set with `XMLOptional` attributes. Its
  last  element is a `STRING`.

You can choose at runtime one of the available options for name:
name, simple-name, or full-name.
You set the selector node to the index value of the record element, for example `LET
Identity._SELECTOR_= 2`. The index value of record elements is shown in the
comment with the hash symbol ( # ) for clarity.

This is an example of the XML output
when the `_SELECTOR_` value is
3:

```
<Identity>
  <full-name>
     <first>John</first>
     <last>Smith</last>
  </full-name>
</Identity>
```

This is an example of the XML output when the
`_SELECTOR_` value is 2:

```
<Identity>
  <simple-name>John Smith</simple-name>
</Identity>
```
