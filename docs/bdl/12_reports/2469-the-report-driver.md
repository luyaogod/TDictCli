---
title: "The report driver"
source: "fgl-topics/c_fgl_reports_006.html"
breadcrumb: "Reports > The report driver"
type: "concept"
---

# The report driver

> The report driver retrieves data, starts the report engine and sends the data (as input records) to be formatted by the REPORT routine.

## Usage

A report driver can be part of the `MAIN` program
block, or it can be in one or more functions.

The report driver typically consists of a loop (such as `WHILE`,
`FOR`, or `FOREACH`) with the following statements to
process the report:

| Instruction | Description |
| --- | --- |
| [`START REPORT`](2470-start-report.md "The START REPORT instruction initializes a report execution.") | This statement is required to instantiate the report driver. |
| [`OUTPUT TO REPORT`](2471-output-to-report.md "The OUTPUT TO REPORT instruction provides a data row to the report execution.") | Provide data for one row to the report driver. |
| [`FINISH REPORT`](2472-finish-report.md "The FINISH REPORT instruction finalizes a report execution.") | Normal termination of the report. |
| [`TERMINATE REPORT`](2473-terminate-report.md "The TERMINATE REPORT instruction cancels a report execution.") | Cancels the processing of the report. |

A report driver is started by the `START REPORT` instruction.
Once started, data can be provided to the report driver through the `OUTPUT TO REPORT`
statement. To instruct the report engine to terminate output processing, use the `FINISH
REPORT` instruction. To cancel a report from outside the report routine, use
`TERMINATE REPORT` (from inside the report routine, you cancel the report
with `EXIT REPORT`).

In order to handle report interruption, the report driver
can check if the `int_flag` variable is `TRUE` in order to stop the
loop if the user interrupts the report execution.

It is possible to execute several report
drivers at the same time. It is even possible to invoke a report driver inside a
`REPORT` routine, which is different from the current driver.

The programmer must make sure that the runtime system will always execute these instructions
in the following order:

1. `START REPORT`
2. `OUTPUT TO REPORT`
3. `FINISH REPORT`

## Example

```
SCHEMA stores
MAIN
  DEFINE rcust RECORD LIKE customer.*
  DATABASE stores
  DECLARE cu1 CURSOR FOR SELECT * FROM customer 
  LET int_flag = FALSE
  START REPORT myrep 
  FOREACH cu1 INTO rcust.*
     IF int_flag THEN EXIT FOREACH END IF
     OUTPUT TO REPORT myrep(rcust.*)
  END FOREACH
  IF int_flag THEN
    TERMINATE REPORT myrep 
  ELSE
    FINISH REPORT myrep 
  END IF
END MAIN
```

## Related links

**Related concepts**  

[int\_flag](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.")

## Child topics

- [START REPORT](2470-start-report.md): The START REPORT instruction initializes a report execution.
- [OUTPUT TO REPORT](2471-output-to-report.md): The OUTPUT TO REPORT instruction provides a data row to the report execution.
- [FINISH REPORT](2472-finish-report.md): The FINISH REPORT instruction finalizes a report execution.
- [TERMINATE REPORT](2473-terminate-report.md): The TERMINATE REPORT instruction cancels a report execution.
