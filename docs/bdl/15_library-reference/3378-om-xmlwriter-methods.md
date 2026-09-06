---
title: "om.XmlWriter methods"
source: "fgl-topics/c_fgl_ClassXMLWriter_methods.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class > om.XmlWriter methods"
type: "concept"
---

# om.XmlWriter methods

> Methods of the om.XmlWriter class.

| Name | Description |
| --- | --- |
| om.XmlWriter.createChannelWriter( channel base.Channel ) RETURNS om.SaxDocumentHandler | Creates an `om.SaxDocumentHandler` object writing to a channel object. |
| om.XmlWriter.createFileWriter( path STRING ) RETURNS om.SaxDocumentHandler | Creates an `om.SaxDocumentHandler` object writing to a file. |
| om.XmlWriter.createPipeWriter( command STRING ) RETURNS om.SaxDocumentHandler | Creates an `om.SaxDocumentHandler` object writing to a pipe created for a process. |
| om.XmlWriter.createSocketWriter( host STRING, port INTEGER ) RETURNS om.SaxDocumentHandler | Creates an `om.SaxDocumentHandler` object writing to a socket. |
