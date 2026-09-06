---
title: "om.DomDocument.getDocumentElement"
source: "fgl-topics/c_fgl_ClassDomDocument_getDocumentElement.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.getDocumentElement"
type: "concept"
---

# om.DomDocument.getDocumentElement

> Returns the root node element of the DOM document.

## Syntax

```
getDocumentElement()
  RETURNS om.DomNode
```

## Usage

The method `getDocumentElement()` returns the root `om.DomNode`
element node of the DOM document.

To hold the reference to the root node, define a variable with the `om.DomNode`
type.

## Example

```
MAIN
  DEFINE mydoc om.DomDocument
  DEFINE n om.DomNode
  LET mydoc = om.DomDocument.create("Test")
  LET n = mydoc.getDocumentElement()
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
