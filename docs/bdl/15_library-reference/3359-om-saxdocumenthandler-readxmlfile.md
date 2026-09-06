---
title: "om.SaxDocumentHandler.readXmlFile"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_readXmlFile.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods > om.SaxDocumentHandler.readXmlFile"
type: "concept"
---

# om.SaxDocumentHandler.readXmlFile

> Reads and processes an XML file with the SAX document handler.

## Syntax

```
readXmlFile(
   path STRING )
```

1. path is the path to an XML formatted file.

## Usage

Use the `readXmlFile()` method after creating the
`om.SaxDocumentHandler` object, to process the XML data from a file input
stream.

## Example

```
DEFINE f om.SaxDocumentHandler
LET f = om.SaxDocumentHandler.createForName("mysaxmod")
CALL f.readXmlFile("cars.xml")
...
```

For a complete example, see [Example 1: Extracting phone numbers from a directory](3365-example-1-extracting-phone-numbers-from-a-directory.md).

## Related links

**Related concepts**  

[The SaxAttributes class](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.")

[FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.")
