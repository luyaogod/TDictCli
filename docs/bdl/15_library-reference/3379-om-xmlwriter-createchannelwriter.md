---
title: "om.XmlWriter.createChannelWriter"
source: "fgl-topics/c_fgl_ClassXMLWriter_createChannelWriter.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class > om.XmlWriter methods > om.XmlWriter.createChannelWriter"
type: "concept"
---

# om.XmlWriter.createChannelWriter

> Creates an om.SaxDocumentHandler object writing to a channel object.

## Syntax

```
om.XmlWriter.createChannelWriter(
   channel base.Channel )
  RETURNS om.SaxDocumentHandler
```

1. channel is a `base.Channel` object reference.

## Usage

The `om.XmlWriter.createChannelWriter()` class method creates an
`om.SaxDocumentHandler` object that will write to the specified channel object, when
using the `om.SaxDocumentHandler` methods.

The [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") object must exist
and be open to receive data from the SAX document handler.

## Example

This example uses the channel to write to stdout, by passing `NULL` as file name
to the `base.Channel.openFile()` method:

```
DEFINE w om.SaxDocumentHandler
DEFINE ch base.Channel
...
LET ch = base.Channel.create()
CALL ch.openFile(NULL,"w")
LET w = om.XmlWriter.createChannelWriter(ch)
...
```

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")

[The SaxDocumentHandler class](3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
