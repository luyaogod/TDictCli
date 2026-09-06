---
title: "om.DomNode.createChild"
source: "fgl-topics/c_fgl_ClassDomNode_createChild.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.createChild"
type: "concept"
---

# om.DomNode.createChild

> Creates and adds a node at the end of the list of children in the current node.

## Syntax

```
createChild(
   tagName STRING )
  RETURNS om.DomNode
```

1. tagName is the tag name of the new node.

## Usage

The `createChild()` method creates a new `om.DomNode`
element with the tag name passed as parameter, and adds it at the end of the children
of the object node calling the method.

The method returns the reference to the created object.

When using the string `"@chars"` as tag name, the `createChild()`
method creates an XML text node. To set/get the value of a text node, use the
`"@chars"` as attribute name with [`getAttribute()`](3302-om-domnode-getattribute.md "Returns the value of a DOM node attribute.")/[`setAttribute()`](3322-om-domnode-setattribute.md "Sets the value of a DOM node attribute."). Alternatively, the [`createChars()`](3288-om-domdocument-createchars.md "Create a new text node in the DOM document.") method of the
`om.DomDocument` class can be used, to create a text node and set the value in a
single step.

## Example

```
DEFINE parent, child, text_node om.DomNode
...
LET child = parent.createChild("Item")
...
LET text_node = parent.createChild("@chars")
CALL text_node.setAttribute("@chars", "my text goes here")
```

For a complete example, see [Example 1: Creating a DOM tree](3327-example-1-creating-a-dom-tree.md).

## Related links

**Related concepts**  

[om.DomNode.appendChild](3297-om-domnode-appendchild.md "Adds an existing node at the end of the list of children in the current node.")
