---
title: "WHILE"
source: "fgl-topics/c_fgl_FlowControl_WHILE.html"
breadcrumb: "Language basics > Flow control > WHILE"
type: "concept"
---

# WHILE

> The WHILE statement executes a block of statements until the specified condition becomes false.

## Syntax

```
WHILE condition
   { statement
   | EXIT WHILE
   | CONTINUE WHILE }
   [...]
END WHILE
```

1. condition must be a boolean expression.
2. statement is any instruction supported by the
   language.

## Usage

As long as the condition specified
after a `WHILE` keyword is `TRUE`,
all statements inside the `WHILE ... END WHILE`
block are executed. After executing the last statement of the
block, the runtime system again evaluates the condition, and
if it is still `TRUE`, continues with the first
statement in the block.

The loop stops when the condition becomes `FALSE` or when an [`EXIT WHILE`](0679-exit-block-name.md "The EXIT block instruction transfers control out of the current program block.") is reached.

Use the [`CONTINUE WHILE`](0678-continue-block-name.md "The CONTINUE block-name instruction resumes execution of a loop or dialog statement.")
instruction to skip the next statements and continue with the loop.

To avoid unending loops, make sure that the condition will become `FALSE` at
some point, or that an `EXIT WHILE` statement
will be executed.

## Example

```
MAIN
  DEFINE cnt INTEGER
  LET cnt = 1
  WHILE cnt <= 100
    DISPLAY "Iter: " || cnt 
    LET cnt = cnt + 1
    IF int_flag THEN
      EXIT WHILE
    END IF
  END WHILE
END MAIN
```

## Related links

**Related concepts**  

[Boolean expressions](0594-boolean-expressions.md "This section covers boolean expression evaluation rules.")
