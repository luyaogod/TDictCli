---
title: "om.XmlWriter.createPipeWriter"
source: "fgl-topics/c_fgl_ClassXMLWriter_createPipeWriter.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class > om.XmlWriter methods > om.XmlWriter.createPipeWriter"
type: "concept"
---

# om.XmlWriter.createPipeWriter

> Creates an om.SaxDocumentHandler object writing to a pipe created for a process.

## Syntax

```
om.XmlWriter.createPipeWriter(
   command STRING )
  RETURNS om.SaxDocumentHandler
```

1. command is the command to be executed.

## Usage

The `om.XmlWriter.createPipeWriter()` class method creates an
`om.SaxDocumentHandler` object that will write to a pipe created
for the specified command. XML data will be send through the pipe when using
the `om.SaxDocumentHandler` methods.

If the process or pipe cannot be created, the method returns `NULL`.

## Example

```
DEFINE w om.SaxDocumentHandler
...
LET w = om.XmlWriter.createPipeWriter("sort -u")
IF w IS NULL THEN
   ERROR "Could not create process."
   EXIT PROGRAM 1
END IF
...
```

## Related links

**Related concepts**  

[The SaxDocumentHandler class](3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
