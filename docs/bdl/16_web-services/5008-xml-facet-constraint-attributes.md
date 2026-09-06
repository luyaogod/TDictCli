---
title: "XML facet constraint attributes"
source: "fgl-topics/r_gws_XML_attributes_004.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML facet constraint attributes"
type: "reference"
description: "The following attributes are facet constraints depending on the XSD data type used on a simple BDL variable to restrict the allowed value-space. (Notice that some attributes are allowed only on some ..."
---

# XML facet constraint attributes

The following attributes are facet constraints depending on the XSD data type used on a simple
BDL variable to restrict the allowed value-space.

(Notice that some attributes are allowed only on some XSD data types).

Several facet constraints can be set on the same data type, and mandatory values are expected (
for example, `XSDMinLength="8"`).

| Attribute | Definition |
| --- | --- |
| [XSDLength](5009-xsdlength.md) | Define the exact number of XML character or bytes. |
| [XSDMinLength](5010-xsdminlength.md) | Define the minimum number of XML character or bytes. |
| [XSDMaxLength](5011-xsdmaxlength.md) | Define the maximum number of XML character or bytes. |
| [XSDEnumeration](5012-xsdenumeration.md) | Define a list of allowed values separated by the character **\|**. |
| [XSDWhiteSpace](5013-xsdwhitespace.md) | Perform a XML string manipulation before serialization or deserialization. |
| [XSDPattern](5014-xsdpattern.md) | Define the regular expression the value has to match. |
| [XSDMinInclusive](5015-xsdmininclusive.md) | Define the inclusive minimum value based on the data type where it is set. |
| [XSDMaxInclusive](5016-xsdmaxinclusive.md) | Define the inclusive maximum value based on the data type where it is set. |
| [XSDMinExclusive](5017-xsdminexclusive.md) | Define the exclusive minimum value based on the data type where it is set. |
| [XSDMaxExclusive](5018-xsdmaxexclusive.md) | Define the exclusive maximum value based on the data type where it is set. |
| [XSDTotalDigits](5019-xsdtotaldigits.md) | Define the total number of digits. |
| [XSDFractionDigits](5020-xsdfractiondigits.md) | Define the number of digits of the fraction part. |
