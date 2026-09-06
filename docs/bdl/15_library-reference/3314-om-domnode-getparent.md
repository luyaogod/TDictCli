---
title: "om.DomNode.getParent"
source: "fgl-topics/c_fgl_ClassDomNode_getParent.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getParent"
type: "concept"
---

# om.DomNode.getParent

> Returns the parent DOM node.

## Syntax

```
getParent()
  RETURNS om.DomNode
```

## Usage

The `getParent()` method returns the parent DOM node of
the current node.

If the current node is the root node, the method returns
`NULL`.

## Example

```
DEFINE parent, current om.DomNode
...
LET parent = current.getParent()
```

## Related links

**Related concepts**  

[om.DomNode.getFirstChild](3310-om-domnode-getfirstchild.md "Returns the first child DOM node.")

[om.DomNode.getLastChild](3312-om-domnode-getlastchild.md "Returns the last child DOM node.")
