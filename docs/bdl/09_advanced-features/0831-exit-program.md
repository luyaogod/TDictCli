---
title: "EXIT PROGRAM"
source: "fgl-topics/c_fgl_programs_EXIT_PROGRAM.html"
breadcrumb: "Advanced features > Program execution > EXIT PROGRAM"
type: "concept"
---

# EXIT PROGRAM

> The EXIT PROGRAM instruction terminates the execution of the program.

## Syntax

```
EXIT PROGRAM [ exit-code ]
```

1. exit-code is a valid integer expression that
   can be read by the process which invoked the program.

## Usage

Use the `EXIT PROGRAM` instruction
to stop the execution of the current program instance.

exit-code must
be zero by default for normal, successful program termination.

exit-code is
converted into a positive integer between 0 and 255 (8 bits).

## Example

```
MAIN
  DISPLAY "Emergency exit."
  EXIT PROGRAM -1
  DISPLAY "This will never be displayed."
END MAIN
```

## Related links

**Related concepts**  

[RUN](0830-run.md "The RUN instruction executes the command passed as argument.")
