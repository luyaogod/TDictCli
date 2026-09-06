---
title: "om.NodeList.getLength"
source: "fgl-topics/c_fgl_ClassNodeList_getLength.html"
breadcrumb: "Library reference > Built-in packages > The om package > The NodeList class > om.NodeList methods > om.NodeList.getLength"
type: "concept"
---

# om.NodeList.getLength

> Returns the number of elements in the node list.

## Syntax

```
getLength()
  RETURNS INTEGER
```

1. node is a reference to a node.

## Usage

The `getLength()` method
returns the size of the node list.

Query the node list for elements with the `item()` method, in the range 1 to
`getLength()`.

## Example

```
DEFINE list om.NodeList
...
DISPLAY list.getLength()
```

For a complete example, see [Example 1: Search nodes by tag name](3335-example-1-search-nodes-by-tag-name.md).

## Related links

**Related concepts**  

[om.DomNode.selectByTagName](3321-om-domnode-selectbytagname.md "Finds descendant DOM nodes based on a tag name.")

[om.DomNode.selectByPath](3320-om-domnode-selectbypath.md "Finds descendant DOM nodes from an XPath-like pattern.")
