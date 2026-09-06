---
title: "xml.Serializer.SetOption"
source: "fgl-topics/c_gws_XmlSerializer_setOption.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.SetOption"
type: "concept"
---

# xml.Serializer.SetOption

> Sets a global option value for the serializer engine

## Syntax

```
xml.Serializer.SetOption(
   optionName STRING,
   optionValue INTEGER )
```

1. optionName specifies the name of the [option flag](4189-serialization-option-flags.md "Serialization option flags for the xml.Serializer class.").
2. optionValue defines the value of the
   flag.

## Usage

This method sets the specified option value for the serializer engine for an option defined by
optionName.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
