---
title: "TERMINATE REPORT"
source: "fgl-topics/c_fgl_reports_TERMINATE_REPORT.html"
breadcrumb: "Reports > The report driver > TERMINATE REPORT"
type: "concept"
---

# TERMINATE REPORT

> The TERMINATE REPORT instruction cancels a report execution.

## Syntax

```
TERMINATE REPORT report-name
```

1. report-name is the name of the report to be canceled.

## Usage

`TERMINATE REPORT` cancels the report processing. It is typically
used when the program (or the user) becomes aware that a problem prevents the report
from producing part of its intended output, or when the user interrupted the report
processing.

`TERMINATE REPORT` has the following effects:

- Terminates the processing of the current report.
- Deletes any intermediate files or temporary tables that were created in processing
  the report.

The `EXIT REPORT` instruction has the same effect, except that it
can be used inside the report definition.

## Related links

**Related concepts**  

[FINISH REPORT](2472-finish-report.md "The FINISH REPORT instruction finalizes a report execution.")

[EXIT REPORT](2490-exit-report.md "Cancels the report processing.")
