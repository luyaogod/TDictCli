---
title: "om.SaxDocumentHandler.setIndent"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_setIndent.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods > om.SaxDocumentHandler.setIndent"
type: "concept"
---

# om.SaxDocumentHandler.setIndent

> Controls indentation in XML output.

## Syntax

```
setIndent(
   indenting BOOLEAN )
```

1. indenting: `TRUE` enables indentation; `FALSE`
   disables indentation.

## Usage

By default, the `om.SaxDocumentHandler`
object outputs XML with indentation.

In order to disable indentation, use the `setIndent(FALSE)` method.
