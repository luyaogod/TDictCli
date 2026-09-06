---
title: "om.DomNode.removeAttribute"
source: "fgl-topics/c_fgl_ClassDomNode_removeAttribute.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.removeAttribute"
type: "concept"
---

# om.DomNode.removeAttribute

> Delete the specified attribute from the DOM node.

## Syntax

```
removeAttribute(
   name STRING )
```

1. name is the name of the attribute.

## Usage

The `removeAttribute()` method deletes the attribute identified by
the name passed as parameter.

DOM node attribute names are case-sensitive.

If the attribute does not exist for this node the method returns silently.

## Example

```
DEFINE node om.DomNode
...
CALL node.removeAttribute("comments")
```

## Related links

**Related concepts**  

[om.DomNode.setAttribute](3322-om-domnode-setattribute.md "Sets the value of a DOM node attribute.")
