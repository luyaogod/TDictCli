---
title: "om.XmlReader.getCharacters"
source: "fgl-topics/c_fgl_ClassXMLReader_getCharacters.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods > om.XmlReader.getCharacters"
type: "concept"
---

# om.XmlReader.getCharacters

> Returns the character data of the current processed element.

## Syntax

```
getCharacters()
  RETURNS STRING
```

## Usage

Use the `getCharacters()` method
to get the character data of the current processed element, in the `Characters` event
context.

## Example

```
DEFINE r om.XmlReader,
       e STRING
   ...
   LET e = r.read()
   WHILE e IS NOT NULL
     CASE e 
       ...
       WHEN "Characters"
         DISPLAY "Characters:'",r.getCharacters(),"'"
     ...
```

For a complete example, see [Example 1: Parsing an XML file](3376-example-1-parsing-an-xml-file.md).
