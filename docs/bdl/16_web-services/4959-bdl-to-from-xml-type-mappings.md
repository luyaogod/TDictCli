---
title: "BDL to/from XML type mappings"
source: "fgl-topics/c_gws_XML_attributes_007.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > BDL to/from XML type mappings"
type: "concept"
---

# BDL to/from XML type mappings

> Map BDL variable types to specific XML data types for web service serialization using optional attributes. For example, use XSDBoolean to map a BDL SMALLINT to an XML boolean, and assign custom XML names as needed.

Starting with Genero 2.0, you can add optional attributes to the definition of program variables
to be used for XML serialization. These attributes can be used to map a BDL data type used in
the input or output message of a Genero Web Service application to a specific XML data type,
rather than using the [default](4960-default-bdl-xml-mapping.md "By default, Genero Web Services converts BDL variables to standard XML Schema (XSD) data types for use in web service messages.").

For example, if an XML Schema boolean data type is required for an application, and the
corresponding BDL type is a SMALLINT, you can use an attribute to map the BDL SMALLINT
variable to the XML boolean.

The following example uses the [XSDBoolean](4966-xsdboolean.md) attribute
to map a BDL SMALLINT variable to an XML Schema boolean type, and
assigns an uppercase name as the XMLName attribute:

```
GLOBALS
DEFINE invoice_out RECORD
  ok SMALLINT ATTRIBUTES(XSDBoolean,XMLName="OK")
END RECORD
	   
END GLOBALS
```

If you assign your own `XMLName`
attributes, be sure to respect the conventions when using the RPC Service Style.

See the [Tutorial:
Writing a GWS Server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .") for additional information
about input and output messages.
