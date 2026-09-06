---
title: "om.DomNode.getId"
source: "fgl-topics/c_fgl_ClassDomNode_getId.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getId"
type: "concept"
---

# om.DomNode.getId

> Returns the internal AUI tree id of a DOM node.

## Syntax

```
getId()
  RETURNS INTEGER
```

## Usage

The `getId()` method returns an internal integer identifier generated
automatically for any `om.DomNode` object
created in the [abstract user interface](../11_user-interface/1514-actions-in-the-abstract-user-interface-tree.md) tree.

The internal id is typically used to reference a DOM node in an attribute of another node, to
link nodes logically together.

If the DOM node does not belong to the AUI tree, the method returns zero.

## Related links

**Related concepts**  

[om.DomDocument.getElementById](3287-om-domdocument-getelementbyid.md "Returns a node element ID based on the internal AUI tree id.")
