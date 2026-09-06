---
title: "EXIT REPORT"
source: "fgl-topics/c_fgl_reports_EXIT_REPORT.html"
breadcrumb: "Reports > Report instructions > EXIT REPORT"
type: "concept"
---

# EXIT REPORT

> Cancels the report processing.

## Syntax

```
EXIT REPORT
```

## Usage

`EXIT REPORT` cancels the report processing. It must appear in the
`FORMAT` section of the report definition. It is useful after the program
(or the user) becomes aware that a problem prevents the report from producing part of
its intended output.

`EXIT REPORT` has the following effects:

- Terminates the processing of the current report.
- Deletes any intermediate files or temporary tables that were created in processing
  the report.

> **Note:**
>
> Do not use the [`RETURN`](../08_language-basics/0768-returning-values.md "A function can return values with the RETURN instruction.")
> statement as a substitute for `EXIT REPORT`. An error is issued if
> `RETURN` is encountered within the definition of a report.

## Related links

**Related concepts**  

[FORMAT section in REPORT](2479-format-section-in-report.md "Defines the formatting directives inside a report routine.")
