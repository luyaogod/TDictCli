---
title: "fgl_report_set_document_handler()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_REPORT_SET_DOCUMENT_HANDLER.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_report_set_document_handler()"
type: "concept"
---

# fgl_report_set_document_handler()

> Redirects the next report to an XML document handler.

## Syntax

```
FUNCTION fgl_report_set_document_handler(
   handler om.SaxDocumentHandler )
```

1. handler is the document handler variable.

## Usage

This function attaches the specified XML document handler to the next executed report.

The function must be called before the execution of a `START REPORT` instruction
not using the `TO XML HANDLER` clause.

> **Note:**
>
> The `fgl_report_set_document_handler()` function is provided for backward
> compatibility. Use the `TO XML HANDLER` of `START REPORT` instead. See
> [XML output for reports](../12_reports/2463-xml-output-for-reports.md "For better integration with external tools based on XML standards, reports can produce XML output.") for more details.
