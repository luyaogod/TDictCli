---
title: "standard.setReportPrinter"
source: "fgl-topics/c_fgl_frontcall_standard_setreportprinter.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.setReportPrinter"
type: "concept"
---

# standard.setReportPrinter

> Override the GDC printer configuration used for report generation for the current application (DBPRINT=FGLSERVER).

> **Important:**
>
> This feature is provided for development and
> debug purpose, and must not be used in a production environment.

## Syntax

```
ui.Interface.frontCall("standard", "setReportPrinter",
  [printer], [result])
```

1. printer - A string describing the printer to use for report generation.
2. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`setReportPrinter`" front call overrides the printer configuration used for
report generation for the current application, when using [DBPRINT=FGLSERVER](../07_configuration/0514-dbprint.md "Defines the print device to be used by reports.").

An empty or null string resets it to the default behavior.

The printer parameter is a string that describes the printer to use for report
generation. For example: `"moliere, Portrait, A4, 96 dpi, 1 copy, Ascendent, Color,
Auto"`.

The printer string can be can in the "Report To Printer" printer panel of the GDC Monitor.

Alternatively, you can specify `"<ASK_ONCE>"` ,
`"<ASK_ALWAYS>"` , `"<USER_DEFINED>"` or
`"<USE_DEFAULT>"` " which will perform the corresponding actions.

## Example

```
DEFINE result BOOLEAN
CALL ui.interface.frontCall("standard","setReportPrinter",
        ["moliere, Portrait, A4, 96 dpi, 1 copy, Ascendent, Color, Auto"],
        [result])
```
