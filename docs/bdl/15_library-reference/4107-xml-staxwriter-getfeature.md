---
title: "xml.StaxWriter.getFeature"
source: "fgl-topics/c_gws_XmlStaxWriter_getFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.getFeature"
type: "concept"
---

# xml.StaxWriter.getFeature

> Gets a feature of a StaxWriter object.

## Syntax

```
getFeature(
   property STRING )
  RETURNS STRING
```

1. property defines the name of a [feature](4119-staxwriter-features.md "Features of the xml.StaxWriter class.").

## Usage

This method returns the name of a [feature](4119-staxwriter-features.md "Features of the xml.StaxWriter class."),
specified in property, of the `StaxWriter` stream.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
