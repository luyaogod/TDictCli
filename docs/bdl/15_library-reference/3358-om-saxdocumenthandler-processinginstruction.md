---
title: "om.SaxDocumentHandler.processingInstruction"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_processingInstruction.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods > om.SaxDocumentHandler.processingInstruction"
type: "concept"
---

# om.SaxDocumentHandler.processingInstruction

> Processes a processing instruction.

## Syntax

```
processingInstruction(
   name STRING,
   data STRING )
```

1. name is the name of the processing instruction (token after
   `<?`).
2. data is the string in the processing instruction tag.

## Usage

The `processingInstruction()` method processes a processing
instruction with the SAX interface.

A processing instruction appears in an XML formatted text as:

```
<?name data ?>
```

## Related links

**Related concepts**  

[The SaxAttributes class](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.")
