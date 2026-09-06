---
title: "xml.Serializer.GetOption"
source: "fgl-topics/c_gws_XmlSerializer_getOption.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.GetOption"
type: "concept"
---

# xml.Serializer.GetOption

> Gets a global option value from the serializer engine.

## Syntax

```
xml.Serializer.GetOption(
   str STRING )
  RETURNS INTEGER
```

1. str defines the [option flag](4189-serialization-option-flags.md "Serialization option flags for the xml.Serializer class.").

## Usage

This method gets an option value from the serializer engine for an option defined by
str. It returns the value of the [flag](4189-serialization-option-flags.md "Serialization option flags for the xml.Serializer class.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
