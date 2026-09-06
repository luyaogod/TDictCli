---
title: "xml.StaxWriter.startDocument"
source: "fgl-topics/c_gws_XmlStaxWriter_startDocument.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.startDocument"
type: "concept"
---

# xml.StaxWriter.startDocument

> Writes a XML declaration to the StaxWriter stream.

## Syntax

```
startDocument(
   encoding STRING,
   version STRING,
   standalone INTEGER )
```

1. encoding defines the encoding declaration in the XML declaration. Passing a
   `NULL` value will use the default UTF-8 encoding.
2. version defines the XML version in the XML declaration. Passing a
   `NULL` value will use the default 1.0 version.
3. standalone defines the XML standalone declaration. Possible values are:
   - `1` : Set `standalone="yes"`.
   - `0` : Set `standalone="no"`.
   - `-1` : Do not set the `standalone` attribute.

## Usage

This method writes a XML declaration to the StaxWriter stream to specify the encoding, the
version, and whether the document is standalone.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

This call:

```
startDocument("utf-8","1.0",1)
```

Produces:

```
<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>
dtd("note [<!ENTITY writer \"Donald Duck.\">]")
```
