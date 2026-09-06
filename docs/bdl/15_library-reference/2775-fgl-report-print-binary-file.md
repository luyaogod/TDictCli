---
title: "fgl_report_print_binary_file()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_REPORT_PRINT_BINARY_FILE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_report_print_binary_file()"
type: "concept"
---

# fgl_report_print_binary_file()

> Prints a file containing binary data during a report.

## Syntax

```
FUNCTION fgl_report_print_binary_file(
   path STRING )
```

1. path is the name of the binary file.

## Usage

This function prints a file containing binary data during a
[report](../12_reports/2461-reports.md).

This function is provided for backward compatibility and must only be using
inside a `REPORT` routine.
