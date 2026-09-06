---
title: "om.DomNode methods"
source: "fgl-topics/c_fgl_ClassDomNode_methods.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods"
type: "concept"
---

# om.DomNode methods

> Methods of the om.DomNode class.

| Name | Description |
| --- | --- |
| appendChild( newChild om.DomNode ) | Adds an existing node at the end of the list of children in the current node. |
| createChild( tagName STRING ) RETURNS om.DomNode | Creates and adds a node at the end of the list of children in the current node. |
| insertBefore( newChild om.DomNode, oldChild om.DomNode) | Inserts an existing node before the existing node specified. |
| removeChild( oldChild om.DomNode ) | Deletes the specified child node from the current node. |
| replaceChild( newChild om.DomNode, oldChild om.DomNode) | Replaces a node by another in the child nodes of the current node. |

| Name | Description |
| --- | --- |
| loadXml( path STRING ) RETURNS om.DomNode | Load an XML file into the current node. |
| parse( s STRING ) RETURNS om.DomNode | Parses an XML formatted string and creates the DOM structure in the current node. |
| toString() RETURNS STRING | Serializes the current node into an XML formatted string. |
| write( sdh om.SaxDocumentHandler ) | Processes a DOM document with a SAX document handler. |
| writeXml( path STRING ) | Creates an XML file from the current DOM node. |

| Name | Description |
| --- | --- |
| getId() RETURNS INTEGER | Returns the internal AUI tree id of a DOM node. |
| getTagName() RETURNS STRING | Returns the XML tag name of a DOM node. |

| Name | Description |
| --- | --- |
| getAttribute( name STRING ) RETURNS STRING | Returns the value of a DOM node attribute. |
| getAttributesCount() RETURNS INTEGER | Returns the number of attributes in the DOM node. |
| getAttributeInteger( name STRING, defaultValue STRING ) RETURNS INTEGER | Returns the value of a DOM node attribute, with default integer value. |
| getAttributeString( name STRING, defaultValue STRING ) RETURNS STRING | Returns the value of a DOM node attribute, with default string value. |
| getAttributeName( index INTEGER ) RETURNS STRING | Returns the name of a DOM node attribute by position. |
| getAttributeValue( index INTEGER ) RETURNS STRING | Returns the value of a DOM node attribute by position. |
| setAttribute( name STRING, value STRING ) | Sets the value of a DOM node attribute. |
| removeAttribute( name STRING ) | Delete the specified attribute from the DOM node. |

| Name | Description |
| --- | --- |
| getChildByIndex( index INTEGER ) RETURNS om.DomNode | Returns a child DOM node by position. |
| getChildCount() RETURNS INTEGER | Returns the number of children nodes. |
| getFirstChild() RETURNS om.DomNode | Returns the first child DOM node. |
| getLastChild() RETURNS om.DomNode | Returns the last child DOM node. |
| getNext() RETURNS om.DomNode | Returns the next sibling DOM node of this node. |
| getParent() RETURNS om.DomNode | Returns the parent DOM node. |
| getPrevious() RETURNS om.DomNode | Returns previous sibling DOM node of this node. |
| selectByPath( path STRING ) RETURNS om.NodeList | Finds descendant DOM nodes from an XPath-like pattern. |
| selectByTagName( tagName STRING ) RETURNS om.NodeList | Finds descendant DOM nodes based on a tag name. |
