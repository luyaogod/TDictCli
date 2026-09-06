---
title: "XMLOptional"
source: "fgl-topics/c_gws_XML_attribute_XMLOptional.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLOptional"
type: "concept"
description: "Define whether a variable can be omitted or not. It specifies how a NULL value is interpreted in XML if it is optional. For an option to set an XML node nillable when the BDL variable is NULL but not ..."
---

# XMLOptional

Define whether a variable can be omitted or not. It specifies how a NULL value is interpreted in
XML if it is optional. For an option to set an XML node nillable when the BDL variable is NULL
but not optional, see the [XMLNillable](5023-xmlnillable.md) attribute.

1. The attribute cannot be set on a type definition.
2. The attribute cannot be set if the main variable is not a RECORD.

## Example

The variable "ValTwo" is optional when
null.

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root")
  val1 INTEGER ATTRIBUTES(XSDint,XMLName="ValOne"),
  val2 FLOAT ATTRIBUTES(XSDdouble,XMLName="ValTwo",XMLOptional)
END RECORD
```

In the resulting XML document, "ValTwo" is defined when not NULL.
For example, if the output is viewed in an XML viewer, it would display as shown.

```
<Root>
  <ValOne>458</ValOne>
  <ValTwo>58.48</ValTwo>
</Root>
```

In the resulting XML document, "ValTwo" is not defined when NULL.

```
<Root>
  <ValOne>458</ValOne>
</Root>
```
