---
title: "PAUSE"
source: "fgl-topics/c_fgl_reports_PAUSE.html"
breadcrumb: "Reports > Report instructions > PAUSE"
type: "concept"
---

# PAUSE

> Pauses a report displayed to the screen.

## Syntax

```
PAUSE [ "comment" ]
```

1. comment is an optional comment to be displayed.

## Usage

Output is sent by default to the screen unless the `START REPORT` statement or the
`OUTPUT` section
specifies a destination for report output.

The `PAUSE` statement can be executed only if the report sends its output
to the screen. It has no effect if you include a `TO` clause in either of
these contexts:

- In the `OUTPUT` section of the report definition.
- In the `START REPORT` statement of the report driver.

Include the `PAUSE` statement in the `PAGE HEADER` or `PAGE TRAILER` block of the report.
For example, the following code causes the runtime system to skip a line and pause at the
end of each page of report output displayed on the screen:

```
  PAGE TRAILER
    SKIP 1 LINE
    PAUSE "Press return to continue"
```

## Related links

**Related concepts**  

[SKIP](2495-skip.md "Skips a given number of lines in a report.")
