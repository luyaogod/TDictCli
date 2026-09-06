---
title: "The DomNode class"
source: "fgl-topics/c_fgl_ClassDomNode.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class"
type: "concept"
---

# The DomNode class

> The om.DomNode class provides methods to manipulate a DOM node of a data tree.

This class follows the [DOM](../09_advanced-features/0948-xml-support.md "Introduces to DOM and SAX standards and describes the XML utility classes built-in the language.") standards.

A DomNode object is a node (or element) of a [DomDocument](3281-the-domdocument-class.md "The om.DomDocument class provides methods to manipulate a data tree, following the DOM standards.").

Tag and attribute names of DOM nodes are case sensitive; "Wheel" is not the same as "wheel".

Text nodes cannot have attributes, but they have plain text. In text nodes, the characters can
be accessed with the `@chars` attribute name. In XML representation, a text node
is the text itself. Do not confuse it with the parent node.
For example, `<Item id="32">Red shoes</Item>` represents 2 nodes:
The parent 'Item' node and a text node with string 'Red shoes'.

If you need to identify an element, use a common attribute like "name". If you need to label an
element, use a common attribute like "text".

> **Important:**
>
> Build-in XML support classes `om.*` have known limitations. Refere to the [Limitations of XML built-in classes](../09_advanced-features/0951-limitations-of-xml-built-in-classes.md "Built-in XML classes have some limitations you must be aware of.") page for more details.

## Child topics

- [om.DomNode methods](3296-om-domnode-methods.md): Methods of the om.DomNode class.
- [Examples](3326-examples.md): om.DomNode usage examples.
