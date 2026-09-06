---
title: "LABEL"
source: "fgl-topics/c_fgl_FlowControl_LABEL.html"
breadcrumb: "Language basics > Flow control > LABEL"
type: "concept"
---

# LABEL

> The LABEL instruction declares a jump point that can be reached by a GOTO or WHENEVER … GOTO.

## Syntax

```
LABEL label-id:
```

1. label-id is a unique identifier in a `MAIN`,
   `REPORT`, or `FUNCTION` program
   block.
2. The label-id must be followed by a colon (`:`).

## Usage

The `LABEL` instruction declares a statement label, making the next statement
one to which a [`GOTO`](0681-goto.md "The GOTO instruction transfers program control to a labeled line within the same program block.") or [`WHENEVER … GOTO`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") statement can transfer
program control.

## Example

```
MAIN

  DISPLAY "Here 1"
  GOTO label1
  DISPLAY "Here 2"
LABEL label1:
  DISPLAY "Here 3"

  WHENEVER ANY ERROR GOTO label2
  DISPLAY 1/0
  DISPLAY "Here 4"
LABEL label2:
  DISPLAY "Here 5"

END MAIN
```

Output:

```
Here 1
Here 3
Here 5
```
