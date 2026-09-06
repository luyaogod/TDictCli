---
title: "om.DomDocument.createChars"
source: "fgl-topics/c_fgl_ClassDomDocument_createChars.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.createChars"
type: "concept"
---

# om.DomDocument.createChars

> Create a new text node in the DOM document.

## Syntax

```
createChars(
   value STRING )
  RETURNS om.DomNode
```

1. value defines the content of the text node.

## Usage

Use the method `createChars()` to create a new `om.DomNode`
text node. The content of the text node must be passed as parameter.

The created node will have the reserved tagName "`@chars`", and a single attribute
named "`@chars`" storing the character data.

To hold the reference to the new node, define a variable with the `om.DomNode`
type.

## Example

```
MAIN
  DEFINE mydoc  om.DomDocument
  DEFINE root, text_node om.DomNode
  LET mydoc = om.DomDocument.create("Test")
  LET root = mydoc.getDocumentElement()
  LET text_node = mydoc.createChars("Hello, world!")
  DISPLAY text_node.getTagName() -- shows "@chars"
  DISPLAY text_node.getAttribute("@chars") -- the text
  CALL root.appendChild(text_node)
  CALL root.writeXML("output.xml")
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
