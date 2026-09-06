---
title: "xml.StaxReader.readFrom"
source: "fgl-topics/c_gws_XmlStaxReader_readFrom.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.readFrom"
type: "concept"
---

# xml.StaxReader.readFrom

> Sets the input stream of the StaxReader object to a file or a URL and starts the streaming

## Syntax

```
readFrom(
   name STRING )
```

1. name defines a valid URL or the name of the file to read.

## Usage

This method sets the input stream of the StaxReader object to a file or a URL, where
name is a valid URL or the name of the file.

Only the following kinds of URLs are supported:

- http://
- https://
- tcp://
- tcps://
- file:///
- alias://

See [fglprofile
Configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") for more details about URL mapping with aliases, and for proxy and security
configuration.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
