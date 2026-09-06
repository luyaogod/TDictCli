---
title: "xml.StaxReader.getEventType"
source: "fgl-topics/c_gws_XmlStaxReader_getEventType.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getEventType"
type: "concept"
---

# xml.StaxReader.getEventType

> Returns a string that indicates the type of event the cursor of the StaxReader object is pointing to.

## Syntax

```
getEventType()
  RETURNS STRING
```

## Usage

This method returns the name of the event type the StaxReader object cursor is pointing to. See
[StaxReader Event Types](4167-staxreader-event-types.md "Event types of the xml.StaxReader class.") for the full list of StaxReader event types.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
