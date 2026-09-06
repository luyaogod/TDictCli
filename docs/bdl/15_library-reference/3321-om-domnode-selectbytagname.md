---
title: "om.DomNode.selectByTagName"
source: "fgl-topics/c_fgl_ClassDomNode_selectByTagName.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.selectByTagName"
type: "concept"
---

# om.DomNode.selectByTagName

> Finds descendant DOM nodes based on a tag name.

## Syntax

```
selectByTagName(
   tagName STRING )
  RETURNS om.NodeList
```

1. tagName is a tag name for the search.

## Usage

The `selectByTagName()` method scans the DOM tree for descendant
nodes defined with the tag name specified as parameter.

DOM node tag names are case-sensitive.

The method creates a list of nodes as a `om.NodeList` object.
This list object is then used to process the nodes found.

## Example

```
DEFINE node om.DomNode,
       nodelist om.NodeList
...
LET nodelist = node.selectByTagName("Car")
```

For a complete example, see [Example 1: Search nodes by tag name](3335-example-1-search-nodes-by-tag-name.md).

## Related links

**Related concepts**  

[The NodeList class](3330-the-nodelist-class.md "A om.NodeList object hold a list of DOM nodes.")
