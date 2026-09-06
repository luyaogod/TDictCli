---
title: "ON TIMER block"
source: "fgl-topics/c_fgl_dialog_ON_TIMER_4.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY interaction blocks > ON TIMER block"
type: "concept"
description: "Syntax ON TIMER seconds instruction [...] Usage The ON TIMER seconds clause defines a set of instructions that must be executed at regular intervals. This interaction block can be used, for example, ..."
---

# ON TIMER block

## Syntax

```
ON TIMER seconds
   instruction [...]
```

## Usage

The `ON TIMER seconds` clause defines a set of instructions that must be
executed at regular intervals. This interaction block can be used, for example, to check if a
message has arrived in a queue, and needs to be processed.

Do not mix `ON TIMER` and [`ON
IDLE`](1897-on-idle-block.md) clauses.

If the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") option is used in an input dialog, [`ON IDLE`](1897-on-idle-block.md) and [`ON TIMER`](1899-on-timer-block.md) will flush and validate the
current field content. If the field contains invalid data, an error message is displayed, and the
corresponding code block is not executed. Therefore, `ON IDLE` and `ON
TIMER` are not recommended in input dialogs using the `UNBUFFERED`
option.

As `ON TIMER` can fire field input validation, it is therefore not recommended in
dialogs allowing input.

When implementing multiple dialogs with `DIALOG` block, do not mix several
`ON IDLE` clauses in different sub-dialog blocks: specify a unique `ON
TIMER` clause at the `DIALOG` block level.

The parameter of `ON TIMER` must be an integer literal or variable. If the value
is zero, the dialog timeout is disabled.

It is not recommended to use the `ON TIMER` trigger with a short timeout period,
such as 1 or 2 seconds. The purpose of this trigger is to give the control back to the program after
a reasonable period of time, such as 10, 20 or 60 seconds.

```
ON TIMER 30
   CALL check_for_messages()
```

The timer value is taken into account when the dialog initializes its internal data structures.
If you use a program variable instead of an integer constant, a change of the variable has no effect
if the change takes place after the dialog has initialized. If you want to change the value of the
timeout variable, it must be done before the dialog block.

The [`PROMPT`](1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
dialog is automatically terminated after `ON
IDLE`, `ON TIMER`,
`ON ACTION`, or `ON KEY` block execution.

## Related links

**Related concepts**  

[Get program control on a regular (timed) basis](2227-get-program-control-on-a-regular-timed-basis.md "Execute some code after a given number of seconds, with or without user interaction with the program.")

[ON IDLE block](1897-on-idle-block.md "ON IDLE block")
