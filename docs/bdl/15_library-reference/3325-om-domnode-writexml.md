---
title: "om.DomNode.writeXml"
source: "fgl-topics/c_fgl_ClassDomNode_writeXml.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.writeXml"
type: "concept"
---

# om.DomNode.writeXml

> Creates an XML file from the current DOM node.

## Syntax

```
writeXml(
   path STRING )
```

1. path is the path to the XML file.

## Usage

The `writeXml()` method writes the content of the current DOM
node to the file passed as parameter.

## Example

```
DEFINE node om.DomNode
...
CALL node.writeXml("output.xml")
```

## Related links

**Related concepts**  

[om.DomNode.loadXml](3300-om-domnode-loadxml.md "Load an XML file into the current node.")
