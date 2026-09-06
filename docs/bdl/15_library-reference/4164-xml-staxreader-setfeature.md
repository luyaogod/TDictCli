---
title: "xml.StaxReader.setFeature"
source: "fgl-topics/c_gws_XmlStaxReader_setFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.setFeature"
type: "concept"
---

# xml.StaxReader.setFeature

> Sets a feature of a StaxReader object.

## Syntax

```
setFeature(
   property STRING,
   value STRING )
```

1. [property](4166-staxreader-features.md "Features of the xml.StaxReader class.") defines the name of a feature.
2. value defines the value of the feature.

## Usage

Use this method to set a feature for the StaxReader object, where property is
the name of a [XmlStaxReader feature](4166-staxreader-features.md "Features of the xml.StaxReader class."), and value is the
value of a feature.

The features can be changed at any time, but they will only be taken into account at the
beginning of a new stream (see [`readFrom`](4160-xml-staxreader-readfrom.md "Sets the input stream of the StaxReader object to a file or a URL and starts the streaming") or [`readFromDocument`](4161-xml-staxreader-readfromdocument.md "Sets the input stream of the StaxReader object to a DomDocument object and starts the streaming.")).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
