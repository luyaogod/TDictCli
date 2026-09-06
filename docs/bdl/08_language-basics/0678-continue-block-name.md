---
title: "CONTINUE block-name"
source: "fgl-topics/c_fgl_FlowControl_CONTINUE.html"
breadcrumb: "Language basics > Flow control > CONTINUE block-name"
type: "concept"
---

# CONTINUE block-name

> The CONTINUE block-name instruction resumes execution of a loop or dialog statement.

## Syntax

```
CONTINUE
 { FOR
 | FOREACH
 | WHILE
 | MENU
 | CONSTRUCT
 | INPUT
 | DIALOG
 }
```

## Usage

The `CONTINUE block-name` instruction
transfers the program execution from a statement block to
another location in the compound statement that is currently
being executed.

`CONTINUE block-name` can
only be used within the statement block specified by block-name.
For example, `CONTINUE FOR` can only be used
within a `FOR ... END FOR` statement block.

The `CONTINUE FOR`, `CONTINUE FOREACH`, or `CONTINUE
WHILE` keywords cause the current [`FOR`](0680-for.md "The FOR instruction executes a statement block a specified number of times."), [`FOREACH`](../10_sql-support/1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor."), or [`WHILE`](0685-while.md "The WHILE statement executes a block of statements until the specified condition becomes false.") loop (respectively) to begin a new cycle immediately. If conditions
do not permit a new cycle, however, the looping statement terminates.

The `CONTINUE MENU`, `CONTINUE CONSTRUCT`, `CONTINUE
INPUT` and `CONTINUE DIALOG` statements cause the program to skip all
subsequent statements in the current control block of a [dialog](../11_user-interface/1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation."). The screen cursor returns to the most recently occupied field in the current form,
giving the user another chance to enter data in that field.

`CONTINUE INPUT` is valid in `INPUT` and `INPUT
ARRAY` statements.

## Example

```
MAIN
  DEFINE i INTEGER
  LET i = 0
  WHILE i < 5
    LET i = i + 1
    DISPLAY "i=" || i 
    CONTINUE WHILE
    DISPLAY "This will never be displayed!"
  END WHILE
END MAIN
```

## Related links

**Related concepts**  

[Boolean expressions](0594-boolean-expressions.md "This section covers boolean expression evaluation rules.")
