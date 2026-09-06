---
title: "om.XmlReader.read"
source: "fgl-topics/c_fgl_ClassXMLReader_read.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods > om.XmlReader.read"
type: "concept"
---

# om.XmlReader.read

> Reads the next SAX event to process.

## Syntax

```
read()
  RETURNS STRING
```

## Usage

The `read()` method reads the next XML fragment and returns the
name of the SAX event to process.

| Event name | Description | Action |
| --- | --- | --- |
| StartDocument | Beginning of the document | Prepare processing (allocate resources) |
| StartElement | Beginning of a node | Get current element's tag name or attributes with `getTagName() getAttributes()` |
| Characters | Value of the current element | Get current text element's value with `getCharacters()` |
| SkippedEntity | Reached skipped entity | Get current skipped entity element's value with `skippedEntity()` |
| EndElement | Ending of a node | Get current element's tagname with `getTagName()` |
| EndDocument | Ending of the document | Finish processing (release resources) |

## Example

```
DEFINE r om.XmlReader,
       e STRING
   ...
   LET e = r.read()
   WHILE e IS NOT NULL
     CASE e 
       ...
     END CASE
     LET e = r.read()
   END WHILE
```

For a complete example, see [Example 1: Parsing an XML file](3376-example-1-parsing-an-xml-file.md).
