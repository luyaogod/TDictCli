---
title: "xml.StaxWriter methods"
source: "fgl-topics/c_gws_XmlStaxWriter_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods"
type: "concept"
---

# xml.StaxWriter methods

> Methods for the xml.StaxWriter class.

| Name | Description |
| --- | --- |
| xml.StaxWriter.Create() RETURNS xml.StaxWriter | Constructor of a StaxWriter object. |

| Name | Description |
| --- | --- |
| getFeature( property STRING ) RETURNS STRING | Gets a feature of a StaxWriter object. |
| setFeature( property STRING, value STRING ) | Sets a feature of a StaxWriter object. |

| Name | Description |
| --- | --- |
| close() | Closes the StaxWriter streaming, and releases all associated resources. |
| writeTo( name STRING ) | Sets the output stream of the StaxWriter object to a file or an URL, and starts the streaming. |
| writeToDocument( doc xml.DomDocument ) | Sets the output stream of the StaxWriter object to an xml.DomDocument object, and starts the streaming. |
| writeToPipe( name STRING ) | Sets the output stream of the StaxWriter object to a PIPE, and starts the streaming. |
| writeToText( txt TEXT ) | Sets the output stream of the StaxWriter object to a TEXT large object, and starts the streaming. |

| Name | Description |
| --- | --- |
| dtd( dtd STRING ) | Writes a DTD to the StaxWriter stream. |
| endDocument() | Closes any open tags and writes corresponding end tags. |
| startDocument( encoding STRING, version STRING, standalone INTEGER ) | Writes a XML declaration to the StaxWriter stream. |

| Name | Description |
| --- | --- |
| declareDefaultNamespace( ns STRING ) | Binds a namespace URI to the default namespace, and forces the output to the StaxWriter stream. |
| declareNamespace( prefix STRING, ns STRING ) | Binds a namespace URI to a prefix, and forces the output of the XML namespace definition to the StaxWriter stream. |
| setDefaultNamespace( ns STRING ) | Binds a namespace URI to the default namespace. |
| setPrefix( prefix STRING, ns STRING ) | Binds a namespace URI to a prefix. |

| Name | Description |
| --- | --- |
| attribute( name STRING, value STRING ) | Writes a XML attribute to the StaxWriter stream. |
| attributeNS( name STRING, ns STRING, value STRING ) | Writes a XML namespace qualified attribute to the StaxWriter stream. |
| cdata( cdata STRING ) | Writes a XML CData to the StaxWriter stream. |
| characters( characters STRING ) | Writes a XML text to the StaxWriter stream. |
| comment( comment STRING ) | Writes a XML comment to the StaxWriter stream. |
| emptyElement( name STRING ) | Writes an empty XML element to the StaxWriter stream. |
| emptyElementNS( name STRING, ns STRING ) | Writes an empty namespace qualified XML element to the StaxWriter stream. |
| endElement() | Writes an end tag to the StaxWriter stream. |
| entityRef( name STRING ) | Writes a XML EntityReference to the StaxWriter stream. |
| processingInstruction( target STRING, data STRING ) | Writes a XML ProcessingInstruction to the StaxWriter stream |
| startElement( name STRING ) | Writes a XML start element to the StaxWriter stream. |
| startElementNS( name STRING, ns STRING ) | Writes a namespace-qualified XML start element to the StaxWriter stream. |
