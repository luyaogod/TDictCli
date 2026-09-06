---
title: "xml.DomDocument.getFeature"
source: "fgl-topics/c_gws_XmlDomDocument_getFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getFeature"
type: "concept"
---

# xml.DomDocument.getFeature

> Gets a feature for an xml.DomDocument object.

## Syntax

```
getFeature(
   property STRING)
  RETURNS STRING
```

1. property is the name of the input parameter
   defining the name of a `xml.DomDocument` feature.

## Usage

This method returns the name of a feature for the `xml.DomDocument` object, where
property is the name of the [DomDocument feature](4016-domdocument-features.md "A list of features for the xml.DomDocument class.").

Returns the value of the feature.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
