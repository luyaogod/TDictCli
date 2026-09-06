---
title: "xml.StaxWriter.setFeature"
source: "fgl-topics/c_gws_XmlStaxWriter_setFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.setFeature"
type: "concept"
---

# xml.StaxWriter.setFeature

> Sets a feature of a StaxWriter object.

## Syntax

```
setFeature(
   property STRING,
   value STRING )
```

1. property defines the name of a feature.
2. value is the value of the feature.

## Usage

Use this method to set a feature of a `StaxWriter` object, where
property is the name of a [feature](4119-staxwriter-features.md "Features of the xml.StaxWriter class."), and value is the value of the feature. The features can be
changed at any time, but will only be taken into account at the beginning of a new stream (see [`writeTo`](4115-xml-staxwriter-writeto.md "Sets the output stream of the StaxWriter object to a file or an URL, and starts the streaming.") or [`writeToDocument`](4116-xml-staxwriter-writetodocument.md "Sets the output stream of the StaxWriter object to an xml.DomDocument object, and starts the streaming.")).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
