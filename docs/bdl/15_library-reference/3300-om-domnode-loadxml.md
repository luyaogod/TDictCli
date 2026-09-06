---
title: "om.DomNode.loadXml"
source: "fgl-topics/c_fgl_ClassDomNode_loadXml.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.loadXml"
type: "concept"
---

# om.DomNode.loadXml

> Load an XML file into the current node.

## Syntax

```
loadXml(
   path STRING )
  RETURNS om.DomNode
```

1. path is the path to the XML file.

## Usage

The `loadXml()` method takes a file path as parameter and loads the XML content
into the current node, by creating a new DOM structure in memory. The method then returns the
created child DOM node.

To hold the reference to the new node, define a variable with the `om.DomNode` type.

In case of XML parsing error, the method does not raise a runtime error
and therefore cannot be trapped with `WHENEVER ERROR` or in a
`TRY/CATCH` block. Instead, `NULL` is returned and the
`status` variable is set with an error number such as -8022 or -8050.

> **Important:**
>
> Files encoded in UTF-8 can start with the UTF-8 Byte Order Mark (BOM), a sequence of `0xEF
> 0xBB 0xBF` bytes, also known as UNICODE `U+FEFF`. When reading files, Genero
> BDL will ignore the UTF-8 BOM, if it is present at the beginning of the file. This applies to
> instructions such as `LOAD`, as well as I/O APIs such as
> `base.Channel.read()` and `readLine()`.

## Example

```
DEFINE parent, new om.DomNode
...
LET new = parent.loadXml("myfile.xml")
IF status != 0 THEN
   DISPLAY "ERROR: ", status
   EXIT PROGRAM 1
END IF
...
```

## Related links

**Related concepts**  

[om.DomNode.writeXml](3325-om-domnode-writexml.md "Creates an XML file from the current DOM node.")
