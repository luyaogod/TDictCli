---
title: "om.DomNode.getChildByIndex"
source: "fgl-topics/c_fgl_ClassDomNode_getChildByIndex.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getChildByIndex"
type: "concept"
---

# om.DomNode.getChildByIndex

> Returns a child DOM node by position.

## Syntax

```
getChildByIndex(
   index INTEGER )
  RETURNS om.DomNode
```

1. index is the index of the child node, starts at 1.

## Usage

The `getChildByIndex()` method returns the child DOM node by position
in the current node.

If there is no child node at the given position, the method returns `NULL`.

## Related links

**Related concepts**  

[om.DomNode.getChildCount](3309-om-domnode-getchildcount.md "Returns the number of children nodes.")
