---
title: "XML serialization attributes"
source: "fgl-topics/r_gws_XML_attributes_005.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes"
type: "reference"
description: "The following attributes are used to change the default serialization of BDL into XML, and vice versa. Some attributes cannot have values, some may have values, and some attributes have mandatory ..."
---

# XML serialization attributes

The following attributes are used to change the default serialization of BDL into XML, and vice
versa. Some attributes cannot have values, some may have values, and some attributes have mandatory
values.

The following attributes cannot have values:

| Attribute | Definition |
| --- | --- |
| [XMLOptional](5022-xmloptional.md) | Define whether the variable can be missing. |
| [XMLNillable](5023-xmlnillable.md) | Define an XML element to be explicitly null and serialized with the `xsi:nil="true"` value. |
| [XMLElement](5024-xmlelement-optional.md) | Map a BDL simple data type to an XML Element. |
| [XMLElementNillable](5025-xmlelementnillable.md) | Define the default for all element members in a `RECORD` (defined by `TYPE` or `DEFINE`) to be serialized as `xsi:nil="true"` if NULL. |
| [XMLAttribute](5026-xmlattribute.md) | Map a BDL simple data type to an XML Attribute. |
| [XMLBase](5027-xmlbase.md) | Set the base type of an XML Schema simpleContent. |
| [XMLAll](5028-xmlall.md) | Map a BDL Record to an XML Schema all structure. |
| [XMLChoice](5029-xmlchoice.md) | Map a BDL Record to an XML Schema choice structure. |
| [XMLSequence](5030-xmlsequence-optional.md) | Map a BDL Record to an XML Schema sequence structure. |
| [XMLSimpleContent](5031-xmlsimplecontent.md) | Map a BDL Record to an XML Schema simpleContent structure. |
| [XSComplexType](5032-xscomplextype.md) | Map a BDL Record type definition to an XML Schema complexType. |
| [XMLList](5033-xmllist.md) | Map a one-dimensional array to an XML Schema list. |
| [XMLSelector](5034-xmlselector.md) | Define which member of an XMLChoice record is selected. |
| [XMLAny](5035-xmlany.md) | Map a `xml.DomDocument` object to a wildcard XML element node. |
| [XMLAnyAttribute](5036-xmlanyattribute.md) | Map a BDL one-dimensional dynamic array of a record with 3 strings to XML wildcard attributes. |

Values are mandatory for the following attributes: (for
example, `XMLName="myname"`)

| Attribute | Definition |
| --- | --- |
| [XMLName](5037-xmlname.md) | Define the XML Name of a variable in an XML document. |
| [XMLNamespace](5038-xmlnamespace.md) | Define the XML Namespace of a variable in an XML document. |
| [XMLType](5039-xmltype.md) | Force the XML type name of a variable. |
| [XMLTypenamespace](5040-xmltypenamespace.md) | Force the XML type namespace of a variable. |
| [XSTypename](5041-xstypename.md) | Define the XML Type Name of a BDL type definition. |
| [XSTypenamespace](5042-xstypenamespace.md) | Define the XML Type Namespace of a BDL type definition. |
| [XMLElementNamespace](5043-xmlelementnamespace.md) | Define the default XML namespace of all children defined as XMLElement in a Record. |
| [XMLAttributeNamespace](5044-xmlattributenamespace.md) | Define the default XML namespace of all children defined as XMLAttribute in a Record. |

Values may be required for the following attributes: (for example,
`XMLOptimizedContent="image/*"`)

| Attribute | Definition |
| --- | --- |
| [XMLOptimizedContent](5045-xmloptimizedcontent.md "Set on STRING or BYTE data type so that such string content represents a file on disk to be transmitted as base64 binary in SOAP via HTTP attachment.") | Set on STRING or BYTE data type so that such string content represents a file on disk to be transmitted as base64 binary in SOAP via HTTP attachment. |
