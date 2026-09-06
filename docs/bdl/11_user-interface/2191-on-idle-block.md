---
title: "ON IDLE block"
source: "fgl-topics/c_fgl_dialog_ON_IDLE_8.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON IDLE block"
type: "concept"
description: "Syntax ON IDLE seconds instruction [...] Usage The ON IDLE seconds clause defines a set of instructions that must be executed after a given period of user inactivity. This interaction block can be ..."
---

# ON IDLE block

## Syntax

```
ON IDLE seconds
   instruction [...]
```

## Usage

The `ON IDLE seconds` clause defines a set of instructions that must be
executed after a given period of user inactivity. This interaction block can be used, for example,
to quit the dialog after the user has not interacted with the program for a specified period of
time.

Do not mix [`ON TIMER`](1899-on-timer-block.md) and
`ON IDLE` clauses.

If the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") option is used in an input dialog, [`ON IDLE`](1897-on-idle-block.md) and [`ON TIMER`](1899-on-timer-block.md) will flush and validate the
current field content. If the field contains invalid data, an error message is displayed, and the
corresponding code block is not executed. Therefore, `ON IDLE` and `ON
TIMER` are not recommended in input dialogs using the `UNBUFFERED`
option.

When implementing multiple dialogs with `DIALOG` block, do not mix several
`ON IDLE` clauses in different sub-dialog blocks: specify a unique `ON
IDLE` clause at the `DIALOG` block level.

The parameter of `ON IDLE` must be an integer literal or variable. If the value is
zero, the dialog timeout is disabled.

It is not recommended to use the `ON IDLE` trigger with a short timeout period
such as 1 or 2 seconds; The purpose of this trigger is to give the control back to the program after
a relatively long period of inactivity (10, 30 or 60 seconds). This is typically the case when the
end user leaves the workstation, or gets a phone call. The program can then execute some code before
the user gets the control back.

```
ON IDLE 30
   IF ask_question(
      "Do you want to reload information from the database?") THEN
      -- Fetch data back from the db server
   END IF
```

The timeout value is taken into account when the dialog initializes its internal data structures.
If you use a program variable instead of an integer constant, any change of the variable will have
no effect if the variable is changed after the dialog has initialized. If you want to change the
value of the timeout variable, it must be done before the dialog block.

The [`PROMPT`](1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
dialog is automatically terminated after `ON
IDLE`, `ON TIMER`,
`ON ACTION`, or `ON KEY` block execution.

## Related links

**Related concepts**  

[Get program control if user is inactive](2226-get-program-control-if-user-is-inactive.md "Execute some code after a given number of seconds, when the user does not interact with the program.")

[ON TIMER block](1899-on-timer-block.md "ON TIMER block")
