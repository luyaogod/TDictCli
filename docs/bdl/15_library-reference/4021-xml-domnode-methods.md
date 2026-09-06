---
title: "xml.DomNode methods"
source: "fgl-topics/c_gws_XmlDomNode_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods"
type: "concept"
---

# xml.DomNode methods

> Methods for the xml.DomNode class.

| Name | Description |
| --- | --- |
| getChildrenCount() RETURNS INTEGER | Returns the number of child DomNode objects for a DomNode object. |
| getChildNodeItem( index INTEGER ) RETURNS xml.DomNode | Returns the child DomNode object at a given position for a DomNode object. |
| getFirstChild() RETURNS xml.DomNode | Returns the first child DomNode object for this XML Element DomNode object. |
| getFirstChildElement() RETURNS xml.DomNode | Returns the first XML Element child DomNode object for this DomNode object. |
| getLastChild() RETURNS xml.DomNode | Returns the last child DomNode object for a XML Element DomNode object. |
| getLastChildElement() RETURNS xml.DomNode | Returns the last child XML element DomNode object for this DomNode object. |
| getNextSibling() RETURNS xml.DomNode | Returns the DomNode object immediately following a DomNode object. |
| getNextSiblingElement() RETURNS xml.DomNode | Returns the XML Element DomNode object immediately following a DomNode object. |
| getParentNode() RETURNS xml.DomNode | Returns the parent DomNode object for this DomNode object. |
| getOwnerDocument() RETURNS xml.DomDocument | Returns the DomDocument object containing this DomNode object. |
| getPreviousSibling() RETURNS xml.DomNode | Returns the DomNode object immediately preceding a DomNode object. |
| getPreviousSiblingElement() RETURNS xml.DomNode | Returns the XML Element DomNode object immediately preceding a DomNode object. |
| hasChildNodes() RETURNS INTEGER | Returns TRUE if a node has child nodes. |

| Name | Description |
| --- | --- |
| clone( deep INTEGER ) RETURNS xml.DomNode | Returns a duplicate DomNode object of a node. |
| appendChild( newChild xml.DomNode ) | Adds a child DomNode object to the end of the child list for a DomNode object |
| appendChildElement( name STRING ) RETURNS xml.DomNode | Creates and adds a child XML Element node to the end of the list of child nodes for an XML Element DomNode object. |
| appendChildElementNS( prefix STRING, name STRING, ns STRING ) RETURNS xml.DomNode | Creates and adds a child namespace qualified XML Element node to the end of the list of child nodes for an XML Element DomNode object. |
| addNextSibling( newNode xml.DomNode ) | Adds a DomNode object as the next sibling of a DomNode object. |
| addPreviousSibling( newNode xml.DomNode ) | Adds a DomNode object as the previous sibling of a DomNode object. |
| insertBeforeChild( newChild xml.DomNode, refChild xml.DomNode ) | Inserts a DomNode object before an existing child DomNode object. |
| insertAfterChild( newChild xml.DomNode, refChild xml.DomNode ) | Inserts a DomNode object after an existing child DomNode object. |
| prependChild( newChild xml.DomNode ) | Adds a child DomNode object to the beginning of the child list for a DomNode object. |
| prependChildElement( name STRING ) RETURNS xml.DomNode | Creates and adds a child XML Element node to the beginning of the list of child nodes for this XML Element DomNode object. |
| prependChildElementNS( prefix STRING, name STRING, ns STRING ) RETURNS xml.DomNode | Creates and adds a child namespace-qualified XML Element node to the beginning of the list of child nodes for an XML Element DomNode object. |
| removeAllChildren() | Removes all child DomNode objects from a DomNode object. |
| removeChild( oldChild xml.DomNode ) | Removes a child DomNode object from the list of child DomNode objects. |
| replaceChild( newChild xml.DomNode, oldChild xml.DomNode ) | Replaces an existing child DomNode with another child DomNode object. |

| Name | Description |
| --- | --- |
| getLocalName() RETURNS STRING | Gets the local name for a DomNode object. |
| getNamespaceURI() RETURNS STRING | Returns the namespace URI for a DomNode object. |
| getNodeName() RETURNS STRING | Gets the name for a DomNode object. |
| getNodeType() RETURNS STRING | Gets the XML type for this DomNode object. |
| getNodeValue() RETURNS STRING | Returns the value for a DomNode object. |
| getPrefix() RETURNS STRING | Returns the prefix for a DomNode object. |
| isAttached() RETURNS INTEGER | Returns whether the node is attached to the XML document. |

| Name | Description |
| --- | --- |
| setNodeValue( value STRING ) | Sets the node value for a DomNode object. |
| setPrefix( prefix STRING ) | Sets the prefix for a DomNode object. |
| toString() RETURNS STRING | Returns a string representation of a DomNode object. |

| Name | Description |
| --- | --- |
| hasAttribute( name STRING ) RETURNS INTEGER | Checks whether an XML Element DomNode object has the XML Attribute specified by a specified name. |
| hasAttributeNS( name STRING, ns STRING ) RETURNS INTEGER | Checks whether a namespace qualified XML Attribute of a given name is carried by an XML Element DomNode object. |
| getAttributesCount() RETURNS INTEGER | Returns the number of XML Attribute DomNode objects on this XML Element DomNode object. |
| getAttributeNode( name STRING ) RETURNS xml.DomNode | Returns an XML Attribute DomNode object for an XML Element DomNode object |
| getAttributeNodeItem( index INTEGER ) RETURNS xml.DomNode | Returns the XML Attribute DomNode object at a given position on this XML Element DomNode object. |
| getAttributeNodeNS( name STRING, ns STRING ) RETURNS xml.DomNode | Returns a namespace-qualified XML Attribute DomNode object for an XML Element DomNode object |
| getAttribute( name STRING ) RETURNS STRING | Returns the value of a XML Attribute for an `xml.DomNode` object |
| getAttributeNS( name STRING, ns STRING ) RETURNS STRING | Returns the value of a namespace qualified XML Attribute for an `xml.DomNode` object |
| hasAttributes() RETURNS INTEGER | Identifies whether a node has XML Attribute nodes. |
| setAttribute( name STRING, value STRING ) | Sets (or resets) an XML Attribute for an XML Element DomNode object. |
| setAttributeNode( attr xml.DomNode ) | Sets (or resets) an XML Attribute DomNode object to an XML Element DomNode object. |
| setAttributeNodeNS( attr xml.DomNode ) | Sets (or resets) a namespace-qualified XML Attribute DomNode object to an XML Element DomNode object. |
| setAttributeNS( prefix STRING, name STRING, ns STRING, value STRING ) | Sets (or resets) a namespace-qualified XML Attribute for an XML Element DomNode object. |
| setIdAttribute( name STRING, isId INTEGER ) | Set the XML Attribute of given name to be of type ID. Declare (or undeclare) the ID as user-determined. |
| setIdAttributeNS( name STRING, ns STRING, isId INTEGER ) | Set the namespace-qualified XML Attribute of given name and namespace to be of type ID. Declare (or undeclare) the ID as user-determined. |
| removeAttribute( name STRING ) | Removes an XML Attribute for an XML Element DomNode object. |
| removeAttributeNS( name STRING, ns STRING ) | Removes a namespace qualified XML Attribute for an XML Element DomNode object |

| Name | Description |
| --- | --- |
| getElementsByTagName( tag STRING ) RETURNS xml.DomNodeList | Returns a DomNodeList object containing all XML Element DomNode objects with the same tag name. |
| getElementsByTagNameNS( tag STRING, ns STRING ) RETURNS xml.DomNodeList | Returns a DomNodeList object containing all namespace-qualified XML Element DomNode objects with the same tag name and namespace. |
| isDefaultNamespace( ns STRING ) RETURNS INTEGER | Checks whether the specified namespace URI is the default namespace. |
| lookupNamespaceURI( prefix STRING ) RETURNS STRING | Looks up the namespace URI associated to a prefix, starting from a specified node. |
| lookupPrefix( ns STRING ) RETURNS STRING | Looks up the prefix associated to a namespace URI, starting from the specified node. |
| selectByXPath( expr STRING [, args ] ) RETURNS xml.DomNodeList | Returns an `xml.DomNodeList` object containing all `xml.DomNode` objects matching an XPath 1.0 expression. |
