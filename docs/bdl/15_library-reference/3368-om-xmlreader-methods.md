---
title: "om.XmlReader methods"
source: "fgl-topics/c_fgl_ClassXMLReader_methods.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > om.XmlReader methods"
type: "concept"
---

# om.XmlReader methods

> Methods of the om.XmlReader class.

| Name | Description |
| --- | --- |
| om.XmlReader.createFileReader( path STRING ) RETURNS om.XmlReader | Creates an XML reader object from a file. |

| Name | Description |
| --- | --- |
| getCharacters() RETURNS STRING | Returns the character data of the current processed element. |
| getAttributes() RETURNS om.SaxAttributes | Builds an attribute list for the current processed element. |
| getTagName() RETURNS STRING | Returns the tag name of the current processed element. |
| read() RETURNS STRING | Reads the next SAX event to process. |
| skippedEntity() RETURNS STRING | Returns the name of an unresolved entity. |
