---
title: "om.XmlWriter.createSocketWriter"
source: "fgl-topics/c_fgl_ClassXMLWriter_createSocketWriter.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class > om.XmlWriter methods > om.XmlWriter.createSocketWriter"
type: "concept"
---

# om.XmlWriter.createSocketWriter

> Creates an om.SaxDocumentHandler object writing to a socket.

## Syntax

```
om.XmlWriter.createSocketWriter(
   host STRING,
   port INTEGER )
  RETURNS om.SaxDocumentHandler
```

1. host is the name of the host computer listening to the TCP port.
2. port is the port number to connect to.

## Usage

The `om.XmlWriter.createSocketWriter()` class method creates an
`om.SaxDocumentHandler` object that will write to a socket identified by the host and
port number passed as parameters. XML data will be sent through the socket when using the
`om.SaxDocumentHandler` methods.

If the socket cannot be opened, the method returns `NULL`.
No timeout is used.

## Example

```
DEFINE w om.SaxDocumentHandler
...
LET w = om.XmlWriter.createSocketWriter("myhost",8012)
IF w IS NULL THEN
   ERROR "Could not open socket."
   EXIT PROGRAM 1
END IF
...
```

## Related links

**Related concepts**  

[The SaxDocumentHandler class](3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
