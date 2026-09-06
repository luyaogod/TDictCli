---
title: "om.XmlReader.getTagName"
source: "fgl-topics/c_fgl_ClassXMLReader_getTagName.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods > om.XmlReader.getTagName"
type: "concept"
---

# om.XmlReader.getTagName

> Returns the tag name of the current processed element.

## Syntax

```
getTagName()
  RETURNS STRING
```

## Usage

Use the `readXmlFile()` method
to get the tag name of the current processed element, in the `StartElement` or `EndElement`
event context.

## Example

```
DEFINE r om.XmlReader,
       e STRING
   ...
   LET e = r.read()
   WHILE e IS NOT NULL
     CASE e 
       ...
       WHEN "StartElement"
         DISPLAY "TagName = ", r.getTagName()
             ...
```

For a complete example, see [Example 1: Parsing an XML file](3376-example-1-parsing-an-xml-file.md).
