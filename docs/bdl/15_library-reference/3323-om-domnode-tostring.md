---
title: "om.DomNode.toString"
source: "fgl-topics/c_fgl_ClassDomNode_toString.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.toString"
type: "concept"
---

# om.DomNode.toString

> Serializes the current node into an XML formatted string.

## Syntax

```
toString()
  RETURNS STRING
```

## Usage

The `toString()` method builds an XML formatted string with
the DOM structure of the current node and returns the string.

## Example

```
DEFINE node om.DomNode, s STRING
...
LET s = node.toString()
```

## Related links

**Related concepts**  

[om.DomNode.parse](3301-om-domnode-parse.md "Parses an XML formatted string and creates the DOM structure in the current node.")
