---
title: "om.SaxDocumentHandler.createForName"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_createForName.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods > om.SaxDocumentHandler.createForName"
type: "concept"
---

# om.SaxDocumentHandler.createForName

> Creates a new SAX document handler object for the given .4gl module.

## Syntax

```
om.SaxDocumentHandler.createForName(
   moduleName STRING )
  RETURNS om.SaxDocumentHandler
```

1. moduleName is the name of the .4gl module defining the
   document handler events.

## Usage

The `om.SaxDocumentHandler.createForName()` method creates an
`om.SaxDocumentHandler` instance and binds the .42m module
passed as argument to the object.

To hold the reference to a SAX document handler object, define a variable with the
`om.SaxDocumentHandler` type.

The module must be available as a compiled .42m file, which is loadable
based on the environment settings (FGLLDPATH).

The module must implement the following functions to process the SAX filter events:

| Function | Description |
| --- | --- |
| startDocument() | Called once at the beginning of the document processing. |
| endDocument() | Called once at the end of the document processing. |
| startElement( tagname STRING, attrs om.SaxAttributes )tagname is the tag name of element.attrs is list of attributes. | Called when an XML element is reached. Use the [`om.SaxAttributes`](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.") methods to handle the attributes of the processed element. |
| endElement( tagname STRING )tagname is the tag name of element. | Called when the end of an XML element is reached. |
| processingInstruction( piname STRING, data STRING )piname is the name of the processing instruction.data is the content of the processing instruction. | Called when a processing instruction is reached. |
| characters( data STRING )data is the text data. | Called when a text node is reached. |
| skippedEntity( name STRING )name is the name of the unknown entity. | Called when an unknown entity node is reached (like &xxx; for example). |

## Example

```
DEFINE f om.SaxDocumentHandler
LET f = om.SaxDocumentHandler.createForName("mysaxmod")
...
```

For a complete example, see [Example 1: Extracting phone numbers from a directory](3365-example-1-extracting-phone-numbers-from-a-directory.md).

## Related links

**Related concepts**  

[The SaxAttributes class](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.")

[FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.")
