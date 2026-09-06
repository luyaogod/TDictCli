---
title: "XMLOptimizedContent"
source: "fgl-topics/c_gws_XML_attribute_XMLOptimizedContent.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLOptimizedContent"
type: "concept"
---

# XMLOptimizedContent

> Set on STRING or BYTE data type so that such string content represents a file on disk to be transmitted as base64 binary in SOAP via HTTP attachment.

The `XMLOptimizedContent` attribute is allowed on BYTE or STRING data type.

The `XMLOptimizedContent` attribute adds value when used in one of these
scenarios:

1. When MTOM is enabled and the `XMLOptimizedContent` attribute is set on a BYTE, it
   sets the variable that must be transmitted transparently as an HTTP part on the wire. See [Message Transmission Optimization Mechanism (MTOM)](4554-mtom.md "Use MTOM to efficiently send binary data to and from SOAP Web services.") for more information.
2. When MTOM is enabled and the `XMLOptimizedContent` attribute is set on a STRING,
   it contains the name of the file to be transmitted as a Base64 binary, with an optional mime-type
   hint, in order to avoid loading a file in a BYTE. See [Message Transmission Optimization Mechanism (MTOM)](4554-mtom.md "Use MTOM to efficiently send binary data to and from SOAP Web services.") for more
   information.
3. If the `XMLOptimizedContent` attribute is set to "`swaRef`"
   (`XMLOptimizedContent="swaRef"`) and set on a STRING, the contents of the file are
   transmitted according to the SoapWithAttachmentRef specification. See [swaRef (SOAP with attachments using wsi:swaRef)](4555-swaref.md "swaRef is a specific way for sending and receiving attachments in SOAP. It is used when you have to transfer files as attachment and locate on disk.") for more information.

If the `XMLOptimizedContent` attribute is set on a BDL STRING data type, it is
expected to contain the name of the file that the Genero Web Service will load based on the current
directory during transport of the MTOM request.

If the `XMLOptimizedContent` attribute is set on a BDL STRING data type, it will
contain the absolute path of the file GWS has received during the transport of a MTOM response. The
absolute path is based on the Genero temp directory settings and contains a UUID generated name. The
programmer is in charge of moving the file or removing the file from disk.

The optional value is a hint to specify the mime-type of the data. The hint may be used by Java
or .NET to generate different kinds of objects, depending on the information provided by the hint.
For instance, on Java a mime type of "image" will generate a `java.awt.Image` object
instead of a byte[] object.

## Example

The Genero
code:

```
DEFINE rec RECORD
  data1 BYTE ATTRIBUTES(XMLOptimizedContent,XMLName="MyData"),
  date2 STRING ATTRIBUTES(XMLOptimizedContent="image/*",XMLName="MyImage"),
  data3 BYTE
END RECORD
```

The
XSD
representation:

```
<xsd:complexType name="rec">
  <xsd:sequence>
    <xsd:element type="xsd:base64Binary" name="MyData"/>
    <xsd:element xmime:expectedContentTypes="image/*" type="xsd:base64Binary" name="MyImage"/>
    <xsd:element type="xsd:base64Binary" name="data3"/>
  </xsd:sequence>
</xsd:complexType>
```

Things
to observe:

1. With the BYTE data type, regardless of whether the XMLOptimizedContent attribute is set, the
   same type (xsd:base64Binary) appears in the XSD.
2. The STRING holds the path to the document to transfer and is viewed as a BYTE in the XML
   exchange.
3. With the STRING data type, a mime type is specified.
