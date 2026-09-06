---
title: "om.DomNode.appendChild"
source: "fgl-topics/c_fgl_ClassDomNode_appendChild.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.appendChild"
type: "concept"
---

# om.DomNode.appendChild

> Adds an existing node at the end of the list of children in the current node.

## Syntax

```
appendChild(
   newChild om.DomNode )
```

1. newChild is a reference to a node.

## Usage

The `appendChild()` method takes an existing `om.DomNode` element
node and adds it at the end of the children of the object node calling the method.

The child node passed to the `appendChild()` method must have been created from
the same DOM document object, for example with the `om.DomDocument.createElement()`
method.

If the node passed to the `appendChild()` method is already attached to another
parent node, it will be detached from that parent node before being attached to the new parent
node.

## Example

```
DEFINE parent, child om.DomNode
...
CALL parent.appendChild(child)
```

For a complete example, see [Example 1: Creating a DOM tree](3327-example-1-creating-a-dom-tree.md).

## Related links

**Related concepts**  

[om.DomNode.replaceChild](3319-om-domnode-replacechild.md "Replaces a node by another in the child nodes of the current node.")

[om.DomNode.removeChild](3318-om-domnode-removechild.md "Deletes the specified child node from the current node.")

[om.DomDocument.createElement](3289-om-domdocument-createelement.md "Create a new element node in the DOM document.")
