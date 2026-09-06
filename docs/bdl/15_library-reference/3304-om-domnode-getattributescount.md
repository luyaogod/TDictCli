---
title: "om.DomNode.getAttributesCount"
source: "fgl-topics/c_fgl_ClassDomNode_getAttributesCount.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getAttributesCount"
type: "concept"
---

# om.DomNode.getAttributesCount

> Returns the number of attributes in the DOM node.

## Syntax

```
getAttributesCount()
  RETURNS INTEGER
```

## Usage

The `getAttributesCount()` method
returns the number of attributes defined in the current node.

This
method is typically used to scan all the attributes of a node by position,
with the `getAttributeName()` and `getAttributeValue()`
methods.

## Example

```
FUNCTION display_attribute_names(node om.DomNode)
    DEFINE index, count INTEGER
    LET count = node.getAttributesCount()
    FOR index = 1 TO count
        DISPLAY node.getAttributeName(index)
    END FOR
END FUNCTION
```

## Related links

**Related concepts**  

[om.DomNode.getAttributeName](3306-om-domnode-getattributename.md "Returns the name of a DOM node attribute by position.")

[om.DomNode.getAttributeValue](3307-om-domnode-getattributevalue.md "Returns the value of a DOM node attribute by position.")
