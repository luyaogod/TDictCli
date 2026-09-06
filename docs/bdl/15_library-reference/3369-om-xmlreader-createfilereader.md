---
title: "om.XmlReader.createFileReader"
source: "fgl-topics/c_fgl_ClassXMLReader_createFileReader.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods > om.XmlReader.createFileReader"
type: "concept"
---

# om.XmlReader.createFileReader

> Creates an XML reader object from a file.

## Syntax

```
om.XmlReader.createFileReader(
   path STRING )
  RETURNS om.XmlReader
```

1. path is the path to an XML formatted file.

## Usage

Use the `om.XmlReader.createFileReader()` method to create a new
`om.XmlReader` object, to process the XML data from a file input stream.

To hold the reference to an XmlReader object, define a variable with the
`om.XmlReader` type.

> **Important:**
>
> Files encoded in UTF-8 can start with the UTF-8 Byte Order Mark (BOM), a sequence of `0xEF
> 0xBB 0xBF` bytes, also known as UNICODE `U+FEFF`. When reading files, Genero
> BDL will ignore the UTF-8 BOM, if it is present at the beginning of the file. This applies to
> instructions such as `LOAD`, as well as I/O APIs such as
> `base.Channel.read()` and `readLine()`.

## Example

```
DEFINE r om.XmlReader
LET r = om.XmlReader.createFileReader("cars.xml")
...
```

For a complete example, see [Example 1: Parsing an XML file](3376-example-1-parsing-an-xml-file.md).

## Related links

**Related concepts**  

[om.XmlReader.read](3373-om-xmlreader-read.md "Reads the next SAX event to process.")
