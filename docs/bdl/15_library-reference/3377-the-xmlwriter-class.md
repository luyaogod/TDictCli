---
title: "The XmlWriter class"
source: "fgl-topics/c_fgl_ClassXMLWriter.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class"
type: "concept"
---

# The XmlWriter class

> The om.XmlWriter class implements methods to write XML to a stream.

The `om.XmlWriter` class implements methods to create a `om.SaxDocumentHandler` object.

Steps to use a XML writer:

1. Declare a variable with the `om.SaxDocumentHandler` type.
2. Create the writer object with one of the class methods of `om.XmlWriter`
   and assign the reference to the variable.
   - [`om.XmlWriter.createFileWriter(filename)`](3380-om-xmlwriter-createfilewriter.md "Creates an om.SaxDocumentHandler object writing to a file.")
     creates an object writing to a file.
   - [`om.XmlWriter.createPipeWriter(command)`](3381-om-xmlwriter-createpipewriter.md "Creates an om.SaxDocumentHandler object writing to a pipe created for a process.")
     creates an object writing to a pipe opened by a sub-process.
   - [`om.XmlWriter.createSocketWriter(hostname,portnum)`](3382-om-xmlwriter-createsocketwriter.md "Creates an om.SaxDocumentHandler object writing to a socket.")
     creates an object writing to the TCP socket.
3. Output XML data with the methods of the `om.SaxDocumentHandler` object:
   1. Use the method [`startDocument()`](3361-om-saxdocumenthandler-startdocument.md "Processes the beginning of the document.") to start writing to the output.
   2. From this point, the order of method calls defines the structure of the XML document. To write an element, fill an
      [`om.SaxAttributes`](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.") object with attributes.
   3. Then, initiate the element output with the method [`startElement()`](3362-om-saxdocumenthandler-startelement.md "Processes the beginning of an element.").
   4. Write element data with the [`characters()`](3355-om-saxdocumenthandler-characters.md "Processes a text node.") method.
   5. Entity nodes are created with the [`skippedEntity()`](3363-om-saxdocumenthandler-skippedentity.md "Processes an unresolved entity.") method.
   6. Finish element output with a call to the [`endElement()`](3357-om-saxdocumenthandler-endelement.md "Processes the end of an element.") method.
   7. Repeat these steps as many times as you have elements to write.
   8. Instead of using the `startElement()` method, you can generate processing instruction elements with
      [`processingInstruction()`](3358-om-saxdocumenthandler-processinginstruction.md "Processes a processing instruction.").
   9. Finally, you must finish the document output with a [`endDocument()`](3356-om-saxdocumenthandler-enddocument.md "Processes the end of the document.") call.

> **Important:**
>
> Build-in XML support classes `om.*` have known limitations. Refere to the [Limitations of XML built-in classes](../09_advanced-features/0951-limitations-of-xml-built-in-classes.md "Built-in XML classes have some limitations you must be aware of.") page for more details.

## Child topics

- [om.XmlWriter methods](3378-om-xmlwriter-methods.md): Methods of the om.XmlWriter class.
- [Examples](3383-examples.md): om.XmlWriter usage examples.
