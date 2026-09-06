---
title: "om.NodeList.item"
source: "fgl-topics/c_fgl_ClassNodeList_item.html"
breadcrumb: "Library reference > Built-in packages > The om package > The NodeList class > om.NodeList methods > om.NodeList.item"
type: "concept"
---

# om.NodeList.item

> Returns a DOM node element by position in the node list.

## Syntax

```
item( index INTEGER )
  RETURNS om.DomNode
```

1. index is the ordinal position of the node in the list.

## Usage

The `item()` method returns the `om.DomNode` object
at the position specified.

First element is at position 1.

If there is no element at the specified index, the method returns `NULL`.

## Example

```
DEFINE list om.NodeList,
       node om.DomNode
...
LET node = list.item(12)
```

For a complete example, see [Example 1: Search nodes by tag name](3335-example-1-search-nodes-by-tag-name.md).

## Related links

**Related concepts**  

[om.DomNode.selectByTagName](3321-om-domnode-selectbytagname.md "Finds descendant DOM nodes based on a tag name.")

[om.DomNode.selectByPath](3320-om-domnode-selectbypath.md "Finds descendant DOM nodes from an XPath-like pattern.")
