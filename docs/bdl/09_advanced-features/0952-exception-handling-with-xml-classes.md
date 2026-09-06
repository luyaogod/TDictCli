---
title: "Exception handling with XML classes"
source: "fgl-topics/c_fgl_xml_utils_006.html"
breadcrumb: "Advanced features > XML support > Exception handling with XML classes"
type: "concept"
---

# Exception handling with XML classes

> Errors can occur while using XML built-in classes.

For example, calling methods of a SAX handler in an invalid order raises the runtime
error [-8004](../15_library-reference/4483-genero-bdl-errors.md).

By default, the program stops in case of exception. XML errors can be trapped with
the `WHENEVER ERROR` or `TRY/CATCH` exception handlers of
Genero. If an error occurs during a method call of an XML class, the runtime system sets
the [status](0939-status.md "status is a predefined variable that contains the execution status of the last instruction.") variable.

This code example shows the trapping of XML classes
errors.

```
MAIN
  DEFINE w om.SaxDocumentHandler
  LET w = om.SaxDocumentHandler.createFileWriter("sample.xml")
  TRY
    CALL w.endDocument()
  CATCH
    DISPLAY "ERROR: ", status
  END TRY
END
```

## Related links

**Related concepts**  

[The SaxDocumentHandler class](../15_library-reference/3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")

[Exceptions](0848-exceptions.md "Describes exception (error) handling in the programs.")
