---
title: "XMLAnyAttribute"
source: "fgl-topics/c_gws_XML_attribute_XMLAnyAttribute.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLAnyAttribute"
type: "concept"
description: "Map a one-dimensional dynamic array to wildcard XML attributes. Example DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\", XMLNamespace=\"http://tempuri.org\") val1 INTEGER ATTRIBUTES(XMLName=\"Val1\"), val2 ..."
---

# XMLAnyAttribute

Map a one-dimensional dynamic array to wildcard XML attributes.

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root", XMLNamespace="http://tempuri.org")
  val1  INTEGER ATTRIBUTES(XMLName="Val1"),
  val2  FLOAT   ATTRIBUTES(XMLName="Val2"),
  attr  STRING  ATTRIBUTES(XMLName="Attr", XMLAttribute),
  any   DYNAMIC ARRAY ATTRIBUTES(XMLAnyAttribute, XMLNamespace="##other")OF RECORD
    ns  STRING,
    name  STRING,
    value  STRING
  END RECORD
END RECORD
```

```
<pre:Root xmlns:pre="http://tempuri.org" pre:Attr="10"
 xmlns:pre2="http://www.mycompany.com" pre2:AnyAttr1="10" pre2:AnyAttr2="">
  <pre:Val1>148</pre:Val1>
  <pre:Val2>0.58</pre:Val2>
</pre:Root>
```

1. The attribute **XMLAnyAttribute** is only allowed on a one-dimensional dynamic
   array of a record with three members of type STRING. The **first** member is for the
   **namespace** of the wildcard attribute, the **second** member is for the
   **name** of the wildcard attribute, and the **third** member is for the
   **value** of the wildcard attribute. **The name cannot be null.**
2. Associated with the XMLAnyAttribute, the [XMLNamespace](5038-xmlnamespace.md) attribute requires either:
   - A list of space-separated URIs to accept each attribute belonging to one of the namespace URIs
     as a wildcard attribute.
   - The value **##any** to accept any attribute as a wildcard attribute.
   - The value **##other** to accept any attribute not in the main schema namespace as a wildcard
     attribute.

   For example:
   - If XMLNamespace="http://tmpuri.org http://www.mycompany.com", then only the
     attributes belonging to one of those namespaces will be accepted and serialized (or deserialized)
     into the array.
   - If XMLNamespace="##any", then any attribute will be accepted and serialized (or deserialized)
     into the array.
   - If XMLNamespace="##other", then any attributes not belonging to the targetNamespace of the XML
     Schema where the anyAttribute definition is used will be accepted and serialized (or deserialized)
     into the array.
