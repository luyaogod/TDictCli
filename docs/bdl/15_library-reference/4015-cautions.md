---
title: "Cautions"
source: "fgl-topics/c_gws_XML_DomDocument_class_012.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > Cautions"
type: "concept"
---

# Cautions

> Some things you need to be aware of when working with the xml.DomDocument class.

Whitespaces, line feeds and carriage returns between elements are represented as text nodes in
memory. A XML document written in a single line and a human readable (pretty printed format) do not
have the same representation in the `DomDocument`. Take this into account when
navigating in the document.

If a `xml.DomNode` is not attached to a `DomDocument` and not
referenced by any variable it can be destroyed. If one child of this node is still referenced, this
child is not destroyed but its parent and the other nodes of the subtree are destroyed. To check if
a node is attached to a `DomDocument` use [isAttached](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") method.

The `DomDocument` remains in memory if any of its nodes are still referenced in a
variable.
