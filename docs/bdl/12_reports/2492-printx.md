---
title: "PRINTX"
source: "fgl-topics/c_fgl_reports_PRINTX.html"
breadcrumb: "Reports > Report instructions > PRINTX"
type: "concept"
---

# PRINTX

> Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.

## Syntax

```
PRINTX [NAME = identifier] expression
```

1. identifier is the name to be used in the XML node.
2. expression is any legal language expression.

## Usage

The `PRINTX` statement is similar to `PRINT`, except that when XML is produced by
the report, the XML print element will be named as specified. If the `NAME`
clause is omitted or the report is run in non-XML mode, then `PRINTX` does
exactly the same as `PRINT`.

To generate XML output, you must redirect the report output into a SAX document handler
with the `TO XML HANDLER` clause of `START REPORT`:

```
START REPORT orders_report
   TO XML HANDLER om.XmlWriter.createFileWriter("orders.xml")
```

When using XML output, `BYTE` values are converted to Base64 before they are
printed with the `PRINTX` instruction.

## Example

```
REPORT (fname, lname, ...)
  DEFINE fname VARCHAR(20),
         lname VARCHAR(20)
  ...
  FORMAT
    ...
    ON EVERY ROW
       PRINTX NAME=customer fname, lname
...
```

With
the above code, the variable names will appear in the graphical report designer as
"`customer.fname`" and "`customer.lname`".

## Related links

**Related concepts**  

[The SaxDocumentHandler class](../15_library-reference/3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")

[Expressions](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language.")
