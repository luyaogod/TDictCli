---
title: "om.DomNode.setAttribute"
source: "fgl-topics/c_fgl_ClassDomNode_setAttribute.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.setAttribute"
type: "concept"
---

# om.DomNode.setAttribute

> Sets the value of a DOM node attribute.

## Syntax

```
setAttribute(
   name STRING,
   value STRING )
```

1. name is the name of the attribute.
2. value is the attribute value.

## Usage

The `setAttribute()` method sets the value of an attribute in
the current node.

DOM node attribute names are case-sensitive.

Make sure that the strings passed to the method do not contain illegal XML
characters. Illegal XML characters will be silently ignored. Illegal XML characters are any
character below space (ASCII 32), except \r (ASCII 13), \n (ASCII 10) and \t (ASCII 9).

For character nodes (created for example with the [`createChars()`](3288-om-domdocument-createchars.md "Create a new text node in the DOM document.") of a
`DomDocument` object), you can set the text value by passing the
`@chars` attribute name to the method.

## Example

```
DEFINE doc om.DomDocument
DEFINE node om.DomNode
DEFINE text_node om.DomNode
...
CALL node.setAttribute("name", "tiger")
...
LET text_node = doc.createChars("Initial text")
CALL text_node.setAttribute("@chars", "New text...")
CALL node.appendChild(text_node)
```

For a complete example, see [Example 1: Creating a DOM tree](3327-example-1-creating-a-dom-tree.md).

## Related links

**Related concepts**  

[om.DomNode.getAttribute](3302-om-domnode-getattribute.md "Returns the value of a DOM node attribute.")

[om.DomNode.getAttributeString](3305-om-domnode-getattributestring.md "Returns the value of a DOM node attribute, with default string value.")

[om.DomNode.getAttributeInteger](3303-om-domnode-getattributeinteger.md "Returns the value of a DOM node attribute, with default integer value.")
