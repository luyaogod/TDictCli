---
title: "om.XmlWriter.createFileWriter"
source: "fgl-topics/c_fgl_ClassXMLWriter_createFileWriter.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class > om.XmlWriter methods > om.XmlWriter.createFileWriter"
type: "concept"
---

# om.XmlWriter.createFileWriter

> Creates an om.SaxDocumentHandler object writing to a file.

## Syntax

```
om.XmlWriter.createFileWriter(
   path STRING )
  RETURNS om.SaxDocumentHandler
```

1. path is the path to the file.

## Usage

The `om.XmlWriter.createFileWriter()` class method creates an
`om.SaxDocumentHandler` object that will write to the specified
file when using the `om.SaxDocumentHandler` methods.

The file is created if it does not exist. If the file cannot be created, the
method returns `NULL`.

When passing `NULL` as file name, the XmlWriter can be used
to write to stdout.

## Example

```
DEFINE w om.SaxDocumentHandler
...
LET w = om.XmlWriter.createFileWriter("mydata.xml")
IF w IS NULL THEN
   ERROR "Could not create file."
   EXIT PROGRAM 1
END IF
...

-- Create an XmlWriter object to write to stdout:
LET w = om.XmlWriter.createFileWriter(NULL)
...
```

## Related links

**Related concepts**  

[The SaxDocumentHandler class](3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
