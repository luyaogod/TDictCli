---
title: "om.DomNode.getFirstChild"
source: "fgl-topics/c_fgl_ClassDomNode_getFirstChild.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getFirstChild"
type: "concept"
---

# om.DomNode.getFirstChild

> Returns the first child DOM node.

## Syntax

```
getFirstChild()
  RETURNS om.DomNode
```

## Usage

The `getFirstChild()` method returns the first child DOM node in
the current node.

This method is typically used to scan child nodes with the `getNext()` method,
until `getNext()` returns `NULL`.

## Example

```
FUNCTION display_children(node om.DomNode)
    DEFINE child om.DomNode
    LET child = node.getFirstChild()
    WHILE child IS NOT NULL
        DISPLAY child.toString()
        LET child = child.getNext()
    END WHILE
END FUNCTION
```

## Related links

**Related concepts**  

[om.DomNode.getNext](3313-om-domnode-getnext.md "Returns the next sibling DOM node of this node.")

[om.DomNode.getPrevious](3315-om-domnode-getprevious.md "Returns previous sibling DOM node of this node.")
