---
title: "om.DomNode.getChildCount"
source: "fgl-topics/c_fgl_ClassDomNode_getChildCount.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getChildCount"
type: "concept"
---

# om.DomNode.getChildCount

> Returns the number of children nodes.

## Syntax

```
getChildCount()
  RETURNS INTEGER
```

## Usage

The `getChildCount()` method returns the number of child nodes in the current
node.

This method is typically used to scan the child nodes of a DOM node, with the
`getChildByIndex()` method.

## Example

```
FUNCTION display_children(node om.DomNode)
    DEFINE index, count INTEGER
    DEFINE child om.DomNode
    LET count = node.getChildCount()
    FOR index = 1 TO count
        LET child = node.getChildByIndex(index)
        DISPLAY child.toString()
    END FOR
END FUNCTION
```

## Related links

**Related concepts**  

[om.DomNode.getChildByIndex](3308-om-domnode-getchildbyindex.md "Returns a child DOM node by position.")
