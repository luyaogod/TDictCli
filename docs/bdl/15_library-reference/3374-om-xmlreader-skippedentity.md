---
title: "om.XmlReader.skippedEntity"
source: "fgl-topics/c_fgl_ClassXMLReader_skippedEntity.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods > om.XmlReader.skippedEntity"
type: "concept"
---

# om.XmlReader.skippedEntity

> Returns the name of an unresolved entity.

## Syntax

```
skippedEntity()
  RETURNS STRING
```

## Usage

The `skippedEntity()` method
returns the name of the unresolved entity, in the `SkippedEntity` event
context.

The parser identifies well know character
entities such as `&amp;` / `&apos;` / `&lt;` / `&gt;` /
`&quot;`, other character entities
are treated as skipped entities and can be processed in
the `SkippedEntity` event.

## Example

```
DEFINE r om.XmlReader,
       e STRING
   ...
   LET e = r.read()
   WHILE e IS NOT NULL
     CASE e 
       ...
       WHEN "SkippedEntity"
         DISPLAY "Entity:'",r.skippedEntity(),"'"
     ...
```

For a complete example, see [Example 1: Parsing an XML file](3376-example-1-parsing-an-xml-file.md).

## Related links

**Related concepts**  

[The SaxAttributes class](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.")

[FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.")
