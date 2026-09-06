---
title: "om.DomDocument.createEntity"
source: "fgl-topics/c_fgl_ClassDomDocument_createEntity.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.createEntity"
type: "concept"
---

# om.DomDocument.createEntity

> Create a new entity node in the DOM document.

## Syntax

```
createEntity(
   value STRING )
  RETURNS om.DomNode
```

1. value defines the name of the entity node.

## Usage

Use the method `createEntity()` to create a new `om.DomNode`
entity node. The entity name must be passed as parameter.

The text representation of a entity node is `&value;`.

The created node will have the reserved tagName "`@entity`", with a single
attribute named "`@entity`" containing the text of the entity.

To hold the reference to the new node, define a variable with the `om.DomNode`
type.

## Example

```
MAIN
  DEFINE mydoc om.DomDocument
  DEFINE n om.DomNode
  LET mydoc = om.DomDocument.create("Test")
  LET n = mydoc.createEntity("quote")
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
