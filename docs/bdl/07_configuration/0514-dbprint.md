---
title: "DBPRINT"
source: "fgl-topics/c_fgl_EnvVariables_DBPRINT.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBPRINT"
type: "concept"
---

# DBPRINT

> Defines the print device to be used by reports.

The DBPRINT environment variable specifies the print device to be used by reports defined
`TO PRINTER`.

On UNIX™ systems, the DBPRINT environment variable
typically contains the printer queue command (such as lp).

On any platform, if the DBPRINT environment variable is not set, it defaults to the
lp command.

When defining `DBPRINT=FGLSERVER`, the report is sent to the printer
configured in the Genero Desktop Client (GDC). See also [standard.setReportFont](../15_library-reference/3414-standard-setreportfont.md "Override the font used for GDC report generation for the current application (DBPRINT=FGLSERVER)."), [standard.setReportPrinter](../15_library-reference/3415-standard-setreportprinter.md "Override the GDC printer configuration used for report generation for the current application (DBPRINT=FGLSERVER).") front calls.

## Related links

**Related concepts**  

[Reports](../12_reports/2461-reports.md "Reports")
