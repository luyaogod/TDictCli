---
title: "om.XmlReader.getAttributes"
source: "fgl-topics/c_fgl_ClassXMLReader_getAttributes.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods > om.XmlReader.getAttributes"
type: "concept"
---

# om.XmlReader.getAttributes

> Builds an attribute list for the current processed element.

## Syntax

```
getAttributes()
  RETURNS om.SaxAttributes
```

## Usage

Use the `getAttributes()` method to create a list of attributes of an
`om.SaxAttributes` object, from the current processed element, in the
`StartElement` or `EndElement` event context.

Declare a variable with the `om.SaxAttributes` type to reference
the attribute list.

Note that once created with the `getAttributes()` method, the
`om.SaxAttributes` object is automatically updated based on the
element currently processed by the `om.XmlReader`.

## Example

```
DEFINE r om.XmlReader,
       e STRING, i INT
       a om.SaxAttributes
   ...
   LET e = r.read()
   WHILE e IS NOT NULL
     CASE e 
       ...
       WHEN "StartElement"
         LET a = r.getAttributes()
         FOR i=1 to a.getLength()
             ...
```

For a complete example, see [Example 1: Parsing an XML file](3376-example-1-parsing-an-xml-file.md).

## Related links

**Related concepts**  

[The SaxAttributes class](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.")
