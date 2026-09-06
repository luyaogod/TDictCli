---
title: "XMLNillable"
source: "fgl-topics/c_gws_XML_attribute_XMLNillable.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLNillable"
type: "concept"
description: "Define an XML element to be explicitly null and serialized with the xsi:nil=\"true\" value. It specifies how a NULL value is set on an XML node when the BDL variable is NULL but not optional. The ..."
---

# XMLNillable

Define an XML element to be explicitly null and serialized with the
`xsi:nil="true"` value. It specifies how a NULL value is set on an XML node
when the BDL variable is NULL but not optional.

The `XMLNillable` attribute cannot be set on a `TYPE`
definition. If you need to set nillable in a record defined from a database schema with
`LIKE` clause, the [XMLElementNillable](5025-xmlelementnillable.md)
attribute supports an implementation that sets all elements in the record to
`XMLNillable`.

The `XMLNillable` attribute is used on its own, or in combination with the
[XMLOptional](5022-xmloptional.md) attribute to specify conditions for the
XML node depending on how you want the BDL variable value to be serialized:

- XML node optional (when NULL)
- XML node nillable (when NULL)
- Option of either XML node optional or nillable, with preference for nillable.

Examples are shown for each use.

## Example optional

The record member "val2" is defined as optional when NULL:

```
DEFINE rec RECORD
     val1 STRING,
     val2 INTEGER ATTRIBUTES(XMLOptional),
     val3 FLOAT
END RECORD
```

When `rec.val2` is NULL, the resulting XML document is:

```
<rec>
  <val1>Hello</val1>
  <val3>3.1415</val3>
</rec>
```

## Example nillable

The record member "val2" is defined as nillable when NULL, but not
optional.

```
DEFINE rec RECORD
     val1 STRING,
     val2 INTEGER ATTRIBUTES(XMLNillable),
     val3 FLOAT
END RECORD
```

When `rec.val2` is NULL, the resulting XML document contains the "val2" element
with the `xsi:nil="true"` attribute:

```
<rec>
  <val1>Hello</val1>
  <val2 xsi:nil="true"/>
  <val3>3.1415</val3>
</rec>
```

## Example optional or nillable

The record member "val2" is defined as either optional or nillable when
NULL:

```
DEFINE rec RECORD
     val1 STRING,
     val2 INTEGER ATTRIBUTES(XMLNillable,XMLOptional),
     val3 FLOAT
END RECORD
```

When `rec.val2` is NULL, the resulting XML document does not contain the "val2"
element, as it is defined by XMLOptional (the default):

```
<rec>
  <val1>Hello</val1>
  <val3>3.1415</val3>
</rec>
```

Or the resulting XML document can contain the "val2" element, with the
`xsi:nil="true"`
attribute:

```
<rec>
  <val1>Hello</val1>
  <val2 xsi:nil="true"/>
  <val3>3.1415</val3>
</rec>
```

If both XMLOptional and XMLNillable are set and the BDL variable is NULL, BDL to XML
serialization uses XMLOptional by default. If you prefer to serialize as XMLNillable,
`xsi:nil="true"`, you must specify the "preferred" value as a parameter of
XMLNillable.

```
DEFINE rec RECORD
     val1 STRING,
     val2 INTEGER ATTRIBUTES(XMLNillable="preferred",XMLOptional),
     val3 FLOAT
END RECORD
```

## XML transformation into BDL

When getting an XML transformation into BDL, the Web Service engine will not
raise an error and the variable is set correctly whether the variable tag is
missing or has `xsi:nil="true"` set. In either case, the BDL
variable is set as NULL.
