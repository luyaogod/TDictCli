---
title: "om.DomNode.parse"
source: "fgl-topics/c_fgl_ClassDomNode_parse.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.parse"
type: "concept"
---

# om.DomNode.parse

> Parses an XML formatted string and creates the DOM structure in the current node.

## Syntax

```
parse(
   s STRING )
  RETURNS om.DomNode
```

1. s is an XML formatted string.

## Usage

The `parse()` method scans the XML formatted string passed as parameter and
creates the corresponding DOM nodes in the current node. The method then returns the created child
DOM node.

The node must be created before it is passed as parameter to this method,
typically, with `om.DomDocument.createElement()`.

## Example

```
DEFINE parent, child om.DomNode
...
LET child = parent.parse("<Item/>")
```

## Related links

**Related concepts**  

[om.DomNode.toString](3323-om-domnode-tostring.md "Serializes the current node into an XML formatted string.")
