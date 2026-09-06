---
title: "xml.DomDocument.setFeature"
source: "fgl-topics/c_gws_XmlDomDocument_setFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.setFeature"
type: "concept"
---

# xml.DomDocument.setFeature

> Sets a feature for this xml.DomDocument object.

## Syntax

```
setFeature(
   property STRING,
   value STRING)
```

1. property is the name of the input parameter
   defining the name of a `xml.DomDocument` feature.
2. value is the value of a feature.

## Usage

Use this method to set a feature for the `xml.DomDocument` object, where
property is the name of a [DomDocument feature](4016-domdocument-features.md "A list of features for the xml.DomDocument class."), and
value is the value of a feature.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
