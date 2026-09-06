---
title: "om.DomNode.getAttributeValue"
source: "fgl-topics/c_fgl_ClassDomNode_getAttributeValue.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getAttributeValue"
type: "concept"
---

# om.DomNode.getAttributeValue

> Returns the value of a DOM node attribute by position.

## Syntax

```
getAttributeValue(
   index INTEGER )
  RETURNS STRING
```

1. index is the index of the attribute, starts
   at 1.

## Usage

The `getAttributeValue()` method
returns the value of an attribute by position in the current node.

DOM
node attribute names are case-sensitive.

If the attribute does
not exist at the given position, the method returns `NULL`.

## Example

```
DEFINE node om.DomNode
...
DISPLAY node.getAttributeValue(12)
```

For a complete example, see [Example 2: Displaying a DOM tree recursively](3328-example-2-displaying-a-dom-tree-recursively.md).

## Related links

**Related concepts**  

[om.DomNode.getAttributeName](3306-om-domnode-getattributename.md "Returns the name of a DOM node attribute by position.")

[om.DomNode.getAttributesCount](3304-om-domnode-getattributescount.md "Returns the number of attributes in the DOM node.")
