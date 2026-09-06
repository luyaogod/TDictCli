---
title: "Runtime error raised when report dimensions are invalid"
source: "fgl-topics/c_fgl_Migrate_to_240_report_dimensions.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Runtime error raised when report dimensions are invalid"
type: "concept"
---

# Runtime error raised when report dimensions are invalid

> Report page length checking error -4375 might occur at compile time or runtime.

Several fixes related to PAGE LENGTH error [-4375](../15_library-reference/4483-genero-bdl-errors.md) have been done in
subsequent releases.

Below is the history of changes for this report definition error:

1. Version 2.40 fixes bug FGL-3035: A START REPORT instruction raises error -4375 at
   runtime, when the top/bottom margin sizes do not fit the page length. In this
   version, the error is not returned at compile time, because report dimensions can be
   specified with variables in START REPORT.
2. Version 2.50.00 fixes bug FGL-3711: The compiler also raises error -4375, if PAGE
   LENGTH is too short to cover the specified page header and trailer lengths.
3. Version 2.51.07 fixes bug FGL-651: FIRST PAGE HEADER blocks can have the same number
   of rows as the PAGE LENGTH.
4. Version 2.50.26 fixes bug FGL-4223: Reports with PAGE LENGTH = 1 TOP MARGING 0
   BOTTOM MARGIN 0 are often used to produce reports without any page formatting. A
   page length of 1 denies printing in PAGE HEADER and PAGE TRAILER. However, it is
   possible to print any number of lines in FIRST PAGE HEADER and ON LAST ROW is
   possible.

## Related links

**Related concepts**  

[Reports](../12_reports/2461-reports.md "Reports")
