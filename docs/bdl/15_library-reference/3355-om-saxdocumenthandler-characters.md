---
title: "om.SaxDocumentHandler.characters"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_characters.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods > om.SaxDocumentHandler.characters"
type: "concept"
---

# om.SaxDocumentHandler.characters

> Processes a text node.

## Syntax

```
characters(
   chars STRING )
```

1. chars is the content of the text node.

## Usage

The `characters()` method processes a text node with the SAX interface.

Make sure that the strings passed to the method do not contain illegal XML
characters. Illegal XML characters will be silently ignored. Illegal XML characters are any
character below space (ASCII 32), except \r (ASCII 13), \n (ASCII 10) and \t (ASCII 9).

## Related links

**Related concepts**  

[om.SaxDocumentHandler.skippedEntity](3363-om-saxdocumenthandler-skippedentity.md "Processes an unresolved entity.")
