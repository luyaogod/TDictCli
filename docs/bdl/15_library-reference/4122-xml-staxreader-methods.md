---
title: "xml.StaxReader methods"
source: "fgl-topics/c_gws_XmlStaxReader_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods"
type: "concept"
---

# xml.StaxReader methods

> Methods for the xml.StaxReader class.

| Name | Description |
| --- | --- |
| xml.StaxReader.Create() RETURNS xml.StaxReader | Constructor of a StaxReader object. |

| Name | Description |
| --- | --- |
| setFeature( property STRING, value STRING ) | Sets a feature of a StaxReader object. |
| getFeature( property STRING ) RETURNS STRING | Gets a feature of a StaxReader object. |

| Name | Description |
| --- | --- |
| readFrom( name STRING ) | Sets the input stream of the StaxReader object to a file or a URL and starts the streaming |
| readFromDocument( _doc xml.DomDocument ) | Sets the input stream of the StaxReader object to a DomDocument object and starts the streaming. |
| readFromText( _txt TEXT ) | Sets the input stream of the StaxReader object to a TEXT large object and starts the streaming. |
| readFromPipe( name STRING ) | Sets the input stream of the StaxReader object to a PIPE and starts the streaming. |
| close() | Closes the StaxReader streaming and releases all associated resources. |

| Name | Description |
| --- | --- |
| getEventType() RETURNS STRING | Returns a string that indicates the type of event the cursor of the StaxReader object is pointing to. |
| hasName() RETURNS INTEGER | Checks whether the StaxReader cursor points to a node with a name. |
| hasText() RETURNS INTEGER | Checks whether the StaxReader cursor points to a node with a text value. |
| isEmptyElement() RETURNS INTEGER | Checks whether the StaxReader cursor points to an empty element node. |
| isStartElement() RETURNS INTEGER | Checks whether the StaxReader cursor points to a start element node. |
| isEndElement() RETURNS INTEGER | Checks whether the StaxReader cursor points to an end element node. |
| isCharacters() RETURNS INTEGER | Checks whether the StaxReader cursor points to a text node. |
| isIgnorableWhitespace() RETURNS INTEGER | Checks whether the StaxReader cursor points to ignorable whitespace. |
| getBytesRead() RETURNS INTEGER | Returns the number of bytes read by the StaxReader from the current XML document. |

| Name | Description |
| --- | --- |
| getEncoding() RETURNS STRING | Returns the document encoding defined in the XML Document declaration, or NULL. |
| getVersion() RETURNS STRING | Returns the document version defined in the XML Document declaration, or NULL. |
| isStandalone() RETURNS INTEGER | Checks whether the document standalone attribute defined in the XML Document declaration is set to yes. |
| standaloneSet() RETURNS INTEGER | Checks whether the document standalone attribute is defined in the XML Document declaration. |

| Name | Description |
| --- | --- |
| getPrefix() RETURNS STRING | Returns the prefix of the current XML node, or NULL. |
| getLocalName() RETURNS STRING | Returns the local name of the current XML node, or NULL. |
| getName() RETURNS STRING | Returns the qualified name of the current XML node, or NULL. |
| getNamespace() RETURNS STRING | Returns the namespace URI of the current XML node, or NULL. |
| getText() RETURNS STRING | Returns as a string the value of the current XML node, or NULL. |

| Name | Description |
| --- | --- |
| getPITarget() RETURNS STRING | Returns the target part of a XML Processing Instruction node, or NULL. |
| getPIData() RETURNS STRING | Returns the data part of a XML Processing Instruction node, or NULL. |

| Name | Description |
| --- | --- |
| getAttributeCount() RETURNS INTEGER | Returns the number of XML attributes defined on the current XML node, or zero. |
| getAttributeLocalName( index INTEGER ) RETURNS STRING | Returns the local name of a XML attribute defined at a given position on the current XML node, or NULL. |
| getAttributeNamespace( index INTEGER ) RETURNS STRING | Returns the namespace URI of a XML attribute defined at a given position on the current XML node, or NULL. |
| getAttributePrefix( index INTEGER ) RETURNS STRING | Returns the prefix of a XML attribute defined at a given position on the current XML node, or NULL. |
| getAttributeValue( index INTEGER ) RETURNS STRING | Returns the value of a XML attribute defined at a given position on the current XML node, or NULL. |
| findAttributeValue( name STRING, ns STRING ) RETURNS STRING | Returns the value of an XML attribute of a given name and/or namespace. |

| Name | Description |
| --- | --- |
| lookupNamespace( prefix STRING ) RETURNS STRING | Looks up the namespace URI associated with a given prefix starting from the current XML node the StaxReader cursor is pointing to. |
| lookupPrefix( ns STRING ) RETURNS STRING | Looks up the prefix associated with a given namespace URI, starting from the current XML node the StaxReader cursor is pointing to. |
| getNamespaceCount() RETURNS INTEGER | Returns the number of namespace declarations defined on the current XML node, or zero. |
| getNamespacePrefix( index INTEGER ) RETURNS STRING | Returns the prefix of a namespace declaration defined at a given position on the current XML node, or NULL. |
| getNamespaceURI( index INTEGER ) RETURNS STRING | Returns the URI of a namespace declaration defined at a given position on the current XML node, or NULL. |

| Name | Description |
| --- | --- |
| hasNext() RETURNS INTEGER | Checks whether the StaxReader cursor can be moved to a XML node next to it. |
| next() | Moves the StaxReader cursor to the next XML node. |
| nextTag() | Moves the StaxReader cursor to the next XML open or end tag |
| nextSibling() | Moves the StaxReader cursor to the immediate next sibling XML Element of the current node, skipping all its child nodes. |
