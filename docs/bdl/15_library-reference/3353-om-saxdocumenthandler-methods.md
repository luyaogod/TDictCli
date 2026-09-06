---
title: "om.SaxDocumentHandler methods"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_methods.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods"
type: "concept"
---

# om.SaxDocumentHandler methods

> Methods of the om.SaxDocumentHandler class.

| Name | Description |
| --- | --- |
| om.SaxDocumentHandler.createForName( moduleName STRING ) RETURNS om.SaxDocumentHandler | Creates a new SAX document handler object for the given .4gl module. |

| Name | Description |
| --- | --- |
| characters( chars STRING ) | Processes a text node. |
| endDocument() | Processes the end of the document. |
| endElement( name STRING ) | Processes the end of an element. |
| processingInstruction( name STRING, data STRING ) | Processes a processing instruction. |
| readXmlFile( path STRING ) | Reads and processes an XML file with the SAX document handler. |
| setIndent( indenting BOOLEAN ) | Controls indentation in XML output. |
| startDocument() | Processes the beginning of the document. |
| startElement( name STRING, atts om.SaxAttributes ) | Processes the beginning of an element. |
| skippedEntity( name STRING ) | Processes an unresolved entity. |
