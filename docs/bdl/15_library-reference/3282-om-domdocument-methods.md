---
title: "om.DomDocument methods"
source: "fgl-topics/c_fgl_ClassDomDocument_methods.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods"
type: "concept"
---

# om.DomDocument methods

> Methods of the om.DomDocument class.

| Name | Description |
| --- | --- |
| om.DomDocument.create( tagName STRING ) RETURNS om.DomDocument | Create a new empty `om.DomDocument` object. |
| om.DomDocument.createFromString( s STRING ) RETURNS om.DomDocument | Create a new `om.DomDocument` object from an XML string. |
| om.DomDocument.createFromXmlFile( path STRING ) RETURNS om.DomDocument | Create a new `om.DomDocument` object from a XML file. |

| Name | Description |
| --- | --- |
| createChars( value STRING ) RETURNS om.DomNode | Create a new text node in the DOM document. |
| createElement( tagName STRING ) RETURNS om.DomNode | Create a new element node in the DOM document. |
| createEntity( value STRING ) RETURNS om.DomNode | Create a new entity node in the DOM document. |
| copy( old om.DomNode, deep INTEGER ) RETURNS om.DomNode | Create a new element node by copying an existing node. |
| getElementById( id INTEGER ) RETURNS om.DomNode | Returns a node element ID based on the internal AUI tree id. |
| getDocumentElement() RETURNS om.DomNode | Returns the root node element of the DOM document. |
| removeElement( oldChild om.DomNode ) | Remove a DomNode object and all its descendants. |
