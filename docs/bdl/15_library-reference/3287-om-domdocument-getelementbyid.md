---
title: "om.DomDocument.getElementById"
source: "fgl-topics/c_fgl_ClassDomDocument_getElementById.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.getElementById"
type: "concept"
---

# om.DomDocument.getElementById

> Returns a node element ID based on the internal AUI tree id.

## Syntax

```
getElementById(
   id INTEGER )
  RETURNS om.DomNode
```

## Usage

The method `getElementById()` returns the `om.DomNode` element of
the DOM document based on the internal id number passed as parameter.

Each DOM node gets an internal integer id when it is created in the [abstract user interface](../11_user-interface/1514-actions-in-the-abstract-user-interface-tree.md) tree. It can be referenced by this
unique id. The node id is typically used in other nodes, to reference a node in the DOM
document.

To hold the reference to the root node, define a variable with the `om.DomNode`
type.

## Example

```
MAIN
    DEFINE uid om.DomDocument
    DEFINE n om.DomNode
    LET uid = ui.Interface.getDocument()
    MENU "test"
        COMMAND "Get UI node by ID"
            LET n = uid.getElementById(1)
            IF n IS NOT NULL THEN
               DISPLAY n.toString()
            END IF
        COMMAND "Exit"
            EXIT MENU
    END MENU
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[om.DomNode.getId](3311-om-domnode-getid.md "Returns the internal AUI tree id of a DOM node.")
