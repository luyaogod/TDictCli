---
title: "XMLAny"
source: "fgl-topics/c_gws_XML_attribute_XMLAny.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLAny"
type: "concept"
description: "Map a Xml.DomDocument object to a wildcard XML element: DEFINE myVar RECORD ATTRIBUTES(XMLName=\"Root\",XMLNamespace=\"http://tempuri.org\") val1 INTEGER ATTRIBUTES(XMLName=\"Val1\"), any Xml.DomDocument ..."
---

# XMLAny

Map a Xml.DomDocument object to a wildcard XML element:

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root",XMLNamespace="http://tempuri.org")
  val1  INTEGER ATTRIBUTES(XMLName="Val1"),
  any   Xml.DomDocument ATTRIBUTES(XMLAny,XMLNamespace="##other"),
  val2  FLOAT   ATTRIBUTES(XMLName="Val2")
END RECORD
```

```
<pre:Root xmlns:pre="http://tempuri.org" >
  <pre:Val1>148</pre:Val1>
  <pre2:Doc xmlns:pre2="http://www.mycompany.com">
     <pre2:Element>how</pre2:Element>
     <pre2:Element>are</pre2:Element>
     <pre2:Element>you</pre2:Element>
  </pre2:Doc>
  <pre:Val2>
0.58</pre:Val2>
</pre:Root>
```

Associated with XMLAny, the [XMLNamespace](5038-xmlnamespace.md) attribute
requires either:

- A list of space-separated URIs to accept each attribute belonging to one of this namespace URI
  as a wildcard attribute.
- The value **##any** to accept any attribute as a wildcard attribute.
- The value **##other** to accept any attribute not in the main schema namespace as a wildcard
  attribute.

For example:

- If XMLNamespace="http://tmpuri.org http://www.mycompany.com", then only
  the XML documents belonging to one of those namespaces will be accepted and serialized (or
  deserialized) into the Xml.DomDocument object.
- If XMLNamespace="##any", then any XML document will be accepted and serialized (or
  deserialized) into the Xml.DomDocument object.
- If XMLNamespace="##other", then any XML document not belonging to the targetNamespace of the
  XML Schema where the any definition is used will be accepted and serialized (or deserialized) into
  the Xml.DomDocument object.
