---
title: "xml.StaxWriter.characters"
source: "fgl-topics/c_gws_XmlStaxWriter_characters.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.characters"
type: "concept"
---

# xml.StaxWriter.characters

> Writes a XML text to the StaxWriter stream.

## Syntax

```
characters(
   characters STRING )
```

1. characters defines the value to write.

## Usage

This method writes the character string passed as parameter as a text element.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
