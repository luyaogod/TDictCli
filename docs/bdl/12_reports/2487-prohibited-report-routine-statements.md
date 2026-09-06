---
title: "Prohibited report routine statements"
source: "fgl-topics/c_fgl_reports_Prohibited_Statements.html"
breadcrumb: "Reports > The report routine > Prohibited report routine statements"
type: "concept"
description: "Language statements that have no meaning inside a report definition routine are prohibited. These statements are some of the statements that are not valid within any control block of the FORMAT ..."
---

# Prohibited report routine statements

Language statements that have no meaning inside a report definition routine are
prohibited. These statements are some of the statements that are not valid within
any control block of the `FORMAT` section of a `REPORT`
program block, such as interactive statements (`CONSTRUCT`,
`INPUT`, `DIALOG`, `MENU`), program
block definitions (`FUNCTION`, `REPORT`), and some
flow control instructions like `RETURN`.

A compile-time error is issued if you attempt to include any of these statements
in a control block of a report. You can call a function that includes some of these
statements, but this is not recommended.

## Related links

**Related concepts**  

[Report instructions](2489-report-instructions.md "The report instruction listed in this section can appear only in control blocks of the FORMAT section of a report routine.")
