---
title: "DomDocument Features"
source: "fgl-topics/r_gws_XML_DomDocument_class_DomDocument_features.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > DomDocument Features"
type: "reference"
---

# DomDocument Features

> A list of features for the xml.DomDocument class.

## DomDocument features

| Name | Description |
| --- | --- |
| `format-pretty-print` | Formats the output by adding whitespace to produce a pretty-printed, indented, human-readable form.Possible values are TRUE or FALSE.Default value is FALSE. |
| `comments` | Defines whether the XML comments are kept during the loading of a document into an `xml.DomDocument` object.Possible values are TRUE or FALSE.Default value is TRUE. |
| `whitespace-in-element-content` | Defines whether XML Text nodes that can be considered "Ignorable" are kept during the loading of an XML document into an `xml.DomDocument` object.Possible values are TRUE or FALSE.Default value is TRUE. |
| `cdata-sections` | Defines whether XML CData nodes are kept or replaced by XML Text nodes during the loading of an XML document into an `xml.DomDocument` object.Possible values are TRUE or FALSE.Default value is TRUE. |
| `expand-entity-references` | Defines whether XML EntityReference nodes are kept or replaced during the loading of an XML document into an `xml.DomDocument` object.Possible values are TRUE or FALSE.Default value is FALSE.See [security issues with expand-entity-references](4016-domdocument-features.md). |
| `validation-type` | Defines what kind of validation is performed.Possible values are: DTD, Schema.Default is Schema. |
| `external-schemaLocation` | Defines a list of namespace-qualified XML schemas to use for validation on an `xml.DomDocument` object.Value is a space-separated string of one or several pairs of strings representing the namespace URI of the schema, followed by its location.Example:"http://tempuri/org/NS mySchema1.xsd http://www.mycompany.com mySchema2.xsd" |
| `external-noNamespaceSchemaLocation` | Defines a list of XML schemas to use for validation on an `xml.DomDocument` object.Value is a space-separated string of one or several strings representing the location of a schema.Example:"mySchema1.xsd mySchema2.xsd" |
| `schema-uriRecovery` | Changes the schema location of an XML schema referenced by import tags in other schemas.Value is a space-separated string of one or several pairs of strings representing the original schema location followed by the new schema locationExample:"http://www.w3.org/2001/xml.xsd myXML.xsd http://www.mycompany.com/GWS.xsd myGWS.xsd" |
| `load-save-base64-string` | Changes methods [loadFromString()](3998-xml-domdocument-loadfromstring.md "Loads a XML Document into a xml.DomDocument object from a string.") and [saveToString()](4004-xml-domdocument-savetostring.md "Saves this xml.DomDocument object as a XML Document to a string.") to handle Base64 strings.Parsing an XML document is done from a BASE64 encoded string, and saving an XML document results in a BASE64 encoded string.Possible values are TRUE or FALSE.Default is FALSE. |
| `auto-id-attribute` | Changes the parsing of an XML document in order to set all unqualified attributes named ID, Id, iD or id to be of type ID.They can then be retrieved with method [getElementById()](3982-xml-domdocument-getelementbyid.md "Returns the xml.DomNode element that has an attribute of type ID with the given value.") or with an XPath expression without calling [setIdAttribute()](4077-xml-domnode-setidattribute.md "Set the XML Attribute of given name to be of type ID. Declare (or undeclare) the ID as user-determined.").Possible values are TRUE or FALSE.Default is FALSE. |
| `auto-id-qualified-attribute` | Changes the parsing of an XML document in order to set all qualified attributes named ID, Id, iD or id to be of type ID.They can then be retrieved with method [getElementById()](3982-xml-domdocument-getelementbyid.md "Returns the xml.DomNode element that has an attribute of type ID with the given value.") or with an XPath expression without calling [setIdAttributeNS()](4078-xml-domnode-setidattributens.md "Set the namespace-qualified XML Attribute of given name and namespace to be of type ID. Declare (or undeclare) the ID as user-determined.").Possible values are TRUE or FALSE.Default is FALSE. |
| `enable-html-compliancy` | Changes methods to parse, normalize and save HTML document via the `xml.DomDocument` object.Possible values are TRUE or FALSE.Default value is FALSE.The HTML parsing isn't namespace qualified, and document is considered as an XML document after loading.**Note:**This feature works only for HTML 4, it is not supported for HTML 5. |

## Security issues with expand-entity-references

When the `expand-entity-references` document feature is set to TRUE, XML entities
referencing sensitive data may be included when loading the XML document with [xml.DomDocument.load](3996-xml-domdocument-load.md "Loads a XML Document into a DomDocument object from a file or an URL."), [xml.DomDocument.loadFromPipe](3997-xml-domdocument-loadfrompipe.md "Loads a XML Document into this xml.DomDocument object from a PIPE."), [xml.DomDocument.loadFromString](3998-xml-domdocument-loadfromstring.md "Loads a XML Document into a xml.DomDocument object from a string."), or [xml.DomDocument.normalize](3999-xml-domdocument-normalize.md "Normalizes the entire Document.").

For example, in its DTD, the following XML file defines the myref
ENTITY element referencing the /etc/passwd
file:

```
<?xml version="1.0" encoding="ISO-8859-1"?>
 <!DOCTYPE foo [  
      <!ELEMENT foo ANY >
      <!ENTITY myref SYSTEM "file:///etc/passwd" >
   ]>
<foo>&myref;</foo>
```

When
loading this XML file with `expand-entity-references` set to TRUE,
the resulting DOM document will have a `<foo>` node containing a
text node with the content of /etc/passwd.
