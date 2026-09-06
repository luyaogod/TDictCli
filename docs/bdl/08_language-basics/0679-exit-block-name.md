---
title: "EXIT block-name"
source: "fgl-topics/c_fgl_FlowControl_EXIT.html"
breadcrumb: "Language basics > Flow control > EXIT block-name"
type: "concept"
---

# EXIT block-name

> The EXIT block instruction transfers control out of the current program block.

## Syntax

```
EXIT
 { CASE
 | FOR
 | FOREACH
 | WHILE
 | MENU
 | CONSTRUCT
 | REPORT
 | DISPLAY
 | INPUT
 | DIALOG }
```

## Usage

The `EXIT block-name` instruction
transfers control out of a control structure (a block, a loop,
a `CASE` statement, or an interface instruction).

The `EXIT block-name` instruction
must be used inside the control structure specified by block-name.
For example, `EXIT FOR` can only appear inside
a `FOR ... END FOR` iteration block.

`EXIT DISPLAY` exits the [`DISPLAY
ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") instruction and `EXIT INPUT` exits an [`INPUT`](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") or an [`INPUT ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") block. `EXIT
CONSTRUCT` exits current [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") block. `EXIT DIALOG` exits current [`DIALOG`](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.") block.

To
exit a function, use the `RETURN` instruction. To terminate
a program, use the `EXIT PROGRAM` instruction.

## Example

```
MAIN
  DEFINE i INTEGER
  LET i = 0
  WHILE TRUE
    DISPLAY "This is an infinite loop. How would you get out of here?"
    LET i = i + 1
    IF i = 100 THEN
       EXIT WHILE
    END IF
  END WHILE
  DISPLAY "Done."
END MAIN
```

## Related links

**Related concepts**  

[RETURN](0676-return.md "The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.")

[EXIT PROGRAM](../09_advanced-features/0831-exit-program.md "The EXIT PROGRAM instruction terminates the execution of the program.")
