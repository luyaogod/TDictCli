---
title: "Writing an XML report driver and routine"
source: "fgl-topics/c_fgl_reports_007.html"
breadcrumb: "Reports > XML output for reports > Writing an XML report driver and routine"
type: "concept"
description: "Generating XML output To produce an XML report, initiate the report with the START REPORT instruction followed by the TO XML HANDLER clause, to specify the SAX document handler that will process the ..."
---

# Writing an XML report driver and routine

## Generating XML output

To produce an XML report, initiate the report with the [`START REPORT`](2470-start-report.md "The START REPORT instruction initializes a report execution.") instruction followed
by the `TO XML HANDLER` clause, to specify the [SAX document handler](../15_library-reference/3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.") that will process the XML
report output:

```
START REPORT order_report 
      TO XML HANDLER om.XmlWriter.createFileWriter("orders.xml")
```

In the report routine, use either [`PRINT`](2491-print.md "Formats and prints a row of data in a report routine.") or [`PRINTX`](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.") statement to generate XML output. Each `PRINT` or
`PRINTX` instruction will produce a `<Print/>` XML element with
`<Item/>` child elements containing the values and data types. The
`PRINTX` instruction is equivalent to `PRINT`, except that you can
specify the name of the `<Print/>` XML element with the `NAME`
clause:

```
REPORT order_report(r_ord)
  ...
  FORMAT
    ON EVERY ROW
      PRINTX NAME = order r_ord.*
  ...
END REPORT
```

Example of XML output produced by the `PRINTX NAME = order r_ord.*`
instruction:

```
<Print name="order">
   <Item name="r_ord.ord_id" type="INTEGER" value="1001" isoValue="1001"/>
   <Item name="r_ord.ord_cust" type="INTEGER" value="101" isoValue="101"/>
   <Item name="r_ord.ord_date" type="DATE" value="11/21/2017" isoValue="2017-11-21"/>
</Print>
```

## Nested XML reports

If a new report is started with `START REPORT` instruction inside a
`REPORT` routine producing XML, and if there is no destination specified in the
`START REPORT` instruction, the sub-report inherits the XML output target of the
parent, and sub-report nodes will be merged into the parent XML output:

```
REPORT order_report(r_ord)
  ...
  FORMAT
    ON EVERY ROW
      PRINTX NAME = order r_ord.*
      -- Merges sub-report output to parent report XML handler
      START REPORT sub_report
      FOR ...
          OUTPUT TO REPORT sub_report(...)
      END FOR
      FINISH REPORT sub_report
  ...
END REPORT
```

## API for global XML handler

The [`fgl_report_set_document_handler()`](../15_library-reference/2776-fgl-report-set-document-handler.md "Redirects the next report to an XML document handler.") built-in function can be used to
specify a general XML handler, for `START REPORT` instructions which do not
use the `TO XML HANDLER` clause:

```
MAIN
  ...
  CALL fgl_report_set_document_handler( om.XmlWriter.createFileWriter("orders.xml") )
  ...
  START REPORT order_report -- Produces XML output to "orders.xml"
  ...
END MAIN
```

The `fgl_report_set_document_handler()` function is supported for backward
compatibility, it is recommended to use `START REPORT … TO XML HANDLER` instead.

## Related links

**Related concepts**  

[The SaxDocumentHandler class](../15_library-reference/3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
