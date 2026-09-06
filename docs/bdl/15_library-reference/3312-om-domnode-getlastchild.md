---
title: "om.DomNode.getLastChild"
source: "fgl-topics/c_fgl_ClassDomNode_getLastChild.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getLastChild"
type: "concept"
---

# om.DomNode.getLastChild

> Returns the last child DOM node.

## Syntax

```
getLastChild()
  RETURNS om.DomNode
```

## Usage

The `getLastChild()` method returns the last child DOM node
in the current node.

This method is typically used to scan child nodes with the `getPrevious()` method,
until `getPrevious()` returns `NULL`.

## Example

```
FUNCTION display_children(node om.DomNode)
    DEFINE child om.DomNode
    LET child = node.getLastChild()
    WHILE child IS NOT NULL
        DISPLAY child.toString()
        LET child = child.getPrevious()
    END WHILE
END FUNCTION
```

## Related links

**Related concepts**  

[om.DomNode.getPrevious](3315-om-domnode-getprevious.md "Returns previous sibling DOM node of this node.")

[om.DomNode.getNext](3313-om-domnode-getnext.md "Returns the next sibling DOM node of this node.")
