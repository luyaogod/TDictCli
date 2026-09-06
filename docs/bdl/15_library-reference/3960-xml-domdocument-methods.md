---
title: "xml.DomDocument methods"
source: "fgl-topics/c_gws_XmlDomDocument_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods"
type: "concept"
---

# xml.DomDocument methods

> Methods for the xml.DomDocument class.

| Name | Description |
| --- | --- |
| xml.DomDocument.Create() RETURNS xml.DomDocument | Constructor of an empty `xml.DomDocument` object. |
| xml.DomDocument.CreateDocument( name STRING ) RETURNS xml.DomDocument | Constructor of an `xml.DomDocument` with an XML root element. |
| xml.DomDocument.CreateDocumentNS( prefix STRING, name STRING, ns STRING ) RETURNS xml.DomDocument | Constructor of a `xml.DomDocument` with a root namespace-qualified XML root element |

| Name | Description |
| --- | --- |
| getDocumentElement() RETURNS xml.DomNode | Returns the root XML Element `xml.DomNode` object for this `xml.DomDocument` object. |
| getDocumentNodesCount() RETURNS INTEGER | Returns the number of child `xml.DomNode` objects for this `xml.DomDocument` object. |
| getDocumentNodeItem( index INTEGER ) RETURNS xml.DomNode | Returns the child `xml.DomNode` object at a given position for this `xml.DomDocument` object. |
| getElementById( id STRING ) RETURNS xml.DomNode | Returns the `xml.DomNode` element that has an attribute of type ID with the given value. |
| getElementsByTagName( tag STRING ) RETURNS xml.DomNodeList | Returns a `xml.DomNodeList` object containing all XML Element `xml.DomNode` objects with the same tag name in the document. |
| getElementsByTagNameNS( tag STRING, ns STRING ) RETURNS xml.DomNodeList | Returns an `xml.DomNodeList` object containing all namespace qualified XML Element `xml.DomNode` objects with the same tag name and namespace in the entire document |
| getFirstDocumentNode() RETURNS xml.DomNode | Returns the first child `xml.DomNode` object for this `xml.DomDocument` object. |
| getLastDocumentNode() RETURNS xml.DomNode | Returns the last child `xml.DomNode` object in this `xml.DomDocument` object. |
| selectByXPath( expr STRING [, args ] ) RETURNS xml.DomNodeList | Returns a `xml.DomNodeList` object containing all `xml.DomNode` objects matching an XPath 1.0 expression. |

| Name | Description |
| --- | --- |
| appendDocumentNode( n xml.DomNode ) | Adds a child DomNode object to the end of the DomNode children for this DomDocument object. |
| clone() RETURNS xml.DomDocument | Returns a copy of this `xml.DomDocument` object. |
| declareNamespace( node xml.DomNode, alias STRING, ns STRING ) | Forces namespace declaration to a XML Element `xml.DomNode` for this `xml.DomDocument` object. |
| insertAfterDocumentNode( newNode xml.DomNode, ref xml.DomNode ) | Inserts a child `xml.DomNode` object after another child DomNode in this `xml.DomDocument` object. |
| insertBeforeDocumentNode( newNode xml.DomNode, ref xml.DomNode ) | Inserts a child DomNode object before another child `xml.DomNode` for this `xml.DomDocument` object. |
| importNode( n xml.DomNode deep INTEGER ) RETURNS xml.DomNode | Imports a `xml.DomNode` from a `xml.DomDocument` object into its new context (attached to a `xml.DomDocument` object). |
| prependDocumentNode( n xml.DomNode ) | Adds a child `xml.DomNode` object to the beginning of the `xml.DomNode` children of this `xml.DomDocument` object. |
| removeDocumentNode( n xml.DomNode ) | Removes a child `xml.DomNode` object from the `xml.DomNode` children in this `xml.DomDocument` object. |

| Name | Description |
| --- | --- |
| createAttribute( elt STRING ) RETURNS xml.DomNode | Creates a XML Attribute `xml.DomNode` object for an `xml.DomDocument` object. |
| createAttributeNS( prefix STRING, elt STRING, ns STRING ) RETURNS xml.DomNode | Creates a XML namespace-qualified Attribute `xml.DomNode` object for this `xml.DomDocument` object. |
| createCDATASection( data STRING ) RETURNS xml.DomNode | Creates an XML CData `xml.DomNode` object for an `xml.DomDocument` object. |
| createComment( cmt STRING ) RETURNS xml.DomNode | Creates a XML Comment `xml.DomNode` object for an `xml.DomDocument` object. |
| createDocumentFragment() RETURNS xml.DomNode | Creates a XML Document Fragment `xml.DomNode` object for this `xml.DomDocument` object. |
| createDocumentType( name STRING, publicID STRING, systemID STRING, internalSubset STRING ) RETURNS xml.DomNode | Creates a XML Document Type (DTD) `xml.DomNode` object for this `xml.DomDocument` object. |
| createElement( elt STRING ) RETURNS xml.DomNode | Creates a XML Element `xml.DomNode` object for an `xml.DomDocument` object |
| createElementNS( prefix STRING, elt STRING, ns STRING ) RETURNS xml.DomNode | Creates a XML namespace-qualified Element `xml.DomNode` object for an `xml.DomDocument` object. |
| createEntityReference( entity STRING ) RETURNS xml.DomNode | Creates a XML EntityReference `xml.DomNode` object for this `xml.DomDocument` object |
| createNode( str STRING ) RETURNS xml.DomNode | Creates an `xml.DomNode` object from a string for this `xml.DomDocument` object. |
| createProcessingInstruction( target STRING, data STRING ) RETURNS xml.DomNode | Creates a XML Processing Instruction `xml.DomNode` object for this `xml.DomDocument` object. |
| createTextNode( txt STRING ) RETURNS xml.DomNode | Creates a XML Text `xml.DomNode` object for this `xml.DomDocument` object. |

| Name | Description |
| --- | --- |
| load( filename STRING ) | Loads a XML Document into a DomDocument object from a file or an URL. |
| loadFromPipe( name STRING ) | Loads a XML Document into this `xml.DomDocument` object from a PIPE. |
| loadFromString( xmlstr STRING ) | Loads a XML Document into a `xml.DomDocument` object from a string. |
| normalize() | Normalizes the entire Document. |
| save( filename STRING ) | Saves this `xml.DomDocument` object as a XML Document to a file or URL. |
| saveToPipe( name STRING ) | Saves this `xml.DomDocument` object as an XML Document via a PIPE. |
| saveToString() RETURNS STRING | Saves this `xml.DomDocument` object as a XML Document to a string. |

| Name | Description |
| --- | --- |
| getFeature( property STRING) RETURNS STRING | Gets a feature for an `xml.DomDocument` object. |
| getXmlEncoding() RETURNS STRING | Returns the document encoding as defined in the XML document declaration. |
| getXmlVersion() RETURNS STRING | Returns the document version as defined in the XML document declaration. |
| isXmlStandalone() RETURNS INTEGER | Checks whether the XML standalone attribute is set in the XML declaration. |
| setFeature( property STRING, value STRING) | Sets a feature for this `xml.DomDocument` object. |
| setXmlEncoding( encoding STRING ) | Sets the XML document encoding in the XML declaration of this `xml.DomDocument`. |
| setXmlStandalone( alone INTEGER ) | Sets the XML standalone attribute in the XML declaration to "yes" or "no" in the XML declaration, or removes the standalone attribute. |

| Name | Description |
| --- | --- |
| validate( ) RETURNS INTEGER | Performs a DTD or XML Schema validation for this `xml.DomDocument` object. |
| validateOneElement( elt xml.DomNode ) RETURNS INTEGER | Performs a DTD or XML Schema validation of a XML Element `xml.DomNode` object. |

| Name | Description |
| --- | --- |
| getErrorsCount() RETURNS INTEGER | Returns the number of errors encountered during the loading, saving or validation of a XML document. |
| getErrorDescription( index INTEGER ) RETURNS STRING | Returns the error description at the given position. |

## Child topics

- [xml.DomDocument.setXmlStandalone](4008-xml-domdocument-setxmlstandalone.md): Sets the XML standalone attribute in the XML declaration to "yes" or "no" in the XML declaration, or removes the standalone attribute.
