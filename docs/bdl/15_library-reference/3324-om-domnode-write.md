---
title: "om.DomNode.write"
source: "fgl-topics/c_fgl_ClassDomNode_write.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.write"
type: "concept"
---

# om.DomNode.write

> Processes a DOM document with a SAX document handler.

## Syntax

```
write(
   sdh om.SaxDocumentHandler )
```

1. sdh references a SAX document handler.

## Usage

The `write()` method processes the current DOM node content with the SAX
document handler passed as parameter.

See the SAX document handler class for more details.

## Related links

**Related concepts**  

[The SaxDocumentHandler class](3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
