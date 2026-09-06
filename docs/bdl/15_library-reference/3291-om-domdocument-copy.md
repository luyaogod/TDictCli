---
title: "om.DomDocument.copy"
source: "fgl-topics/c_fgl_ClassDomDocument_copy.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.copy"
type: "concept"
---

# om.DomDocument.copy

> Create a new element node by copying an existing node.

## Syntax

```
copy(
   old om.DomNode,
   deep INTEGER )
  RETURNS om.DomNode
```

1. old references the source node to copy.
2. deep is a boolean to control the recursive node copy.

## Usage

Use the method `copy()` to create a new `om.DomNode`
element node from an existing node.

Pass `TRUE` as second parameter to clone a complete tree of nodes.

To hold the reference to the new node, define a variable with the `om.DomNode`
type.

## Example

```
MAIN
  DEFINE mydoc om.DomDocument
  DEFINE n, s om.DomNode
  LET s = mydoc.createElement("Car")
  LET n = mydoc.copy(s, TRUE)
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
