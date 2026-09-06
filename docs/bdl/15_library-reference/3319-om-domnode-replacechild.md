---
title: "om.DomNode.replaceChild"
source: "fgl-topics/c_fgl_ClassDomNode_replaceChild.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.replaceChild"
type: "concept"
---

# om.DomNode.replaceChild

> Replaces a node by another in the child nodes of the current node.

## Syntax

```
replaceChild(
   newChild om.DomNode,
   oldChild om.DomNode)
```

1. newChild is a reference to the new node.
2. oldChild is the node to be replaced.

## Usage

The `replaceChild()` method puts the `om.DomNode` element passed as
first parameter at the place of the node referenced by the second parameter, in the children list of
the object node calling the method.

The new child node passed to the `replaceChild()` method must have been created
from the same DOM document object, for example with the
`om.DomDocument.createElement()` method.

The old node is not destroyed, if it is still referenced by a variable. The old node
still exists in the DOM document, but it is an orphan node, that can be attached to
another parent node in the document.

## Example

```
MAIN
    DEFINE doc  om.DomDocument,
           r om.DomNode,
           p om.DomNode,
           o om.DomNode,
           n om.DomNode

    LET doc = om.DomDocument.create("Items")

    LET r = doc.createElement("Zoo")

    LET p = doc.createElement("DodoList")
    CALL r.appendChild(p)

    LET o = doc.createElement("Dodo")
    CALL o.setAttribute("name", "momo")
    CALL o.setAttribute("gender", "male")
    CALL p.appendChild(o)

    CALL r.writeXml("file1.xml")

    LET n = doc.createElement("Dodo")
    CALL n.setAttribute("name", "kiki")
    CALL n.setAttribute("gender", "female")

    CALL p.replaceChild(n, o)

    -- o is orphan but still exists
    CALL o.writeXml("file2.xml")
    LET o = NULL -- unref/destroy the node

    CALL r.writeXml("file3.xml")
END MAIN
```

The above program will produce following files:

file1.xml

```
<?xml version='1.0' encoding='ASCII'?>
<Zoo>
  <DodoList>
    <Dodo name="momo" gender="male"/>
  </DodoList>
</Zoo>
```

file2.xml

```
<?xml version='1.0' encoding='ASCII'?>
<Dodo name="momo" gender="male"/>
```

file3.xml

```
<?xml version='1.0' encoding='ASCII'?>
<Zoo>
  <DodoList>
    <Dodo name="kiki" gender="female"/>
  </DodoList>
</Zoo>
```

## Related links

**Related concepts**  

[om.DomNode.appendChild](3297-om-domnode-appendchild.md "Adds an existing node at the end of the list of children in the current node.")

[om.DomNode.removeChild](3318-om-domnode-removechild.md "Deletes the specified child node from the current node.")

[om.DomDocument.createElement](3289-om-domdocument-createelement.md "Create a new element node in the DOM document.")
