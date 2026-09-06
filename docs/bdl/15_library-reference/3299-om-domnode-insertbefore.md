---
title: "om.DomNode.insertBefore"
source: "fgl-topics/c_fgl_ClassDomNode_insertBefore.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.insertBefore"
type: "concept"
---

# om.DomNode.insertBefore

> Inserts an existing node before the existing node specified.

## Syntax

```
insertBefore(
   newChild om.DomNode,
   oldChild om.DomNode)
```

1. newChild is a reference to a new created node.
2. oldChild is a reference to a child node existing in the current node.

## Usage

The `insertBefore()` method takes an existing `om.DomNode` element
node and inserts it before the child node passed as second parameter, in the list of nodes of the
object node calling the method.

The child node passed to the `insertBefore()` method must have been created from
the same DOM document object, for example with the `om.DomDocument.createElement()`
method.

> **Important:**
>
> When using the `om` API to update the AUI tree, it is not
> possible to insert nodes, if the parent node has already been synchronized with the front-end. Child
> nodes can be inserted, only if the parent node is not yet synchronized with the front-end. For more
> details, see [Manipulating the abstract user interface tree](../11_user-interface/1512-manipulating-the-abstract-user-interface-tree.md).

## Example

```
DEFINE parent, other, child om.DomNode
...
LET child = mydoc.createElement("Item")
CALL parent.insertBefore(child, other)
```

## Related links

**Related concepts**  

[om.DomDocument.createElement](3289-om-domdocument-createelement.md "Create a new element node in the DOM document.")
