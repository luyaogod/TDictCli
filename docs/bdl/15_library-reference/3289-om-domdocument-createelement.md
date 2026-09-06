---
title: "om.DomDocument.createElement"
source: "fgl-topics/c_fgl_ClassDomDocument_createElement.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.createElement"
type: "concept"
---

# om.DomDocument.createElement

> Create a new element node in the DOM document.

## Syntax

```
createElement(
   tagName STRING )
  RETURNS om.DomNode
```

1. tagName defines the tag name of the node.

## Usage

Use the method `createElement()` to create a new `om.DomNode`
element node. The tag name of the element must be passed as parameter.

To hold the reference to the new node, define a variable with the `om.DomNode`
type.

## Example

```
MAIN
  DEFINE mydoc om.DomDocument
  DEFINE n om.DomNode
  LET mydoc = om.DomDocument.create("Test")
  LET n = mydoc.createElement("Car")
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
