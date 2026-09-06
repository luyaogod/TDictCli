---
title: "xml.StaxWriter.writeTo"
source: "fgl-topics/c_gws_XmlStaxWriter_writeTo.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.writeTo"
type: "concept"
---

# xml.StaxWriter.writeTo

> Sets the output stream of the StaxWriter object to a file or an URL, and starts the streaming.

## Syntax

```
writeTo(
   name STRING )
```

1. name defines a valid URL or the name
   of the file that will contain the resulting XML document.

## Usage

This method sets the output stream of the `StaxWriter` object to a file or a URL
specified in name, and starts the streaming.

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

## Examples

```
writeTo("printerList.xml")
```

```
writeTo("http://myserver:1100/documents/printerList.xml")
```

```
writeTo("https://myserver:1100/documents/printerList.xml")
```

```
writeTo("alias://printerlist")
```

In the example `printerlist`
alias is defined in fglprofile as `ws.printerlist.url =
"http://myserver:1100/documents/ptinterList.xml"`.
