---
title: "om.DomNode.removeChild"
source: "fgl-topics/c_fgl_ClassDomNode_removeChild.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.removeChild"
type: "concept"
---

# om.DomNode.removeChild

> Deletes the specified child node from the current node.

## Syntax

```
removeChild(
   oldChild om.DomNode )
```

1. oldChild is a reference to a node.

## Usage

The `removeChild()` method detaches an `om.DomNode` element node
from the current node.

The removed node is not destroyed, if it is still referenced by a variable. The
`removeChild()` method will only break the link between the parent node and the child
node. The child node still exists in the DOM document, but it is an orphan node, that can be
attached to another parent node in the document.

## Example

```
MAIN
    DEFINE doc  om.DomDocument,
           r om.DomNode,
           p om.DomNode,
           c om.DomNode

    LET doc = om.DomDocument.create("Items")

    LET r = doc.createElement("Zoo")

    LET p = doc.createElement("DodoList")
    CALL r.appendChild(p)

    LET c = doc.createElement("Dodo")
    CALL c.setAttribute("name", "momo")
    CALL c.setAttribute("gender", "male")
    CALL p.appendChild(c)

    CALL r.writeXml("file1.xml")

    CALL p.removeChild(c)

    -- c is orphan but still exists
    CALL c.writeXml("file2.xml")
    LET c = NULL -- unref/destroy the node

    CALL r.writeXml("file3.xml")
END MAIN
```

The above program will produce the following files:

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
  <DodoList/>
</Zoo>
```

## Related links

**Related concepts**  

[om.DomNode.appendChild](3297-om-domnode-appendchild.md "Adds an existing node at the end of the list of children in the current node.")

[om.DomNode.replaceChild](3319-om-domnode-replacechild.md "Replaces a node by another in the child nodes of the current node.")
