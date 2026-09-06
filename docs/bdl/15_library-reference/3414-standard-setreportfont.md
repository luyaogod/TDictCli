---
title: "standard.setReportFont"
source: "fgl-topics/c_fgl_frontcall_standard_setreportfont.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.setReportFont"
type: "concept"
---

# standard.setReportFont

> Override the font used for GDC report generation for the current application (DBPRINT=FGLSERVER).

> **Important:**
>
> This feature is provided for development and
> debug purpose, and must not be used in a production environment.

## Syntax

```
ui.Interface.frontCall("standard", "setReportFont",
  [font], [result])
```

1. font - A string describing the font to use for report generation.
2. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`setReportFont`" front call overrides the default font used by GDC for report
generation for the current application, when using [DBPRINT=FGLSERVER](../07_configuration/0514-dbprint.md "Defines the print device to be used by reports.").

An empty or null string resets it to the default behavior.

The font parameter is a string that describe the font to use for report
generation. For example: `"Helvetica, Bold, Italic, 13"`.

> **Tip:**
>
> Copy/paste the font string from the "Report To Printer" font panel from GDC
> Monitor.

Alternatively, you can specify `"<ASK_ONCE>"` ,
`"<ASK_ALWAYS>"` , `"<USER_DEFINED>"` or
`"<USE_DEFAULT>"` " which will perform the corresponding actions.

## Example

```
DEFINE result BOOLEAN
CALL ui.interface.frontCall("standard","setReportFont",
        ["Helvetica, Bold, Italic, 13"],
        [result])
```
