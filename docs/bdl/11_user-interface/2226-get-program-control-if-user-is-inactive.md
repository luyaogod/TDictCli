---
title: "Get program control if user is inactive"
source: "fgl-topics/c_fgl_prog_dialogs_on_idle.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Get program control if user is inactive"
type: "concept"
---

# Get program control if user is inactive

> Execute some code after a given number of seconds, when the user does not interact with the program.

## When to use the ON IDLE trigger?

If an interactive instruction has the control, the program waits for a user interaction like an
action or field input. If the end user leaves the workstation, or switches to another application,
the program cannot get the control and is frozen until the user comes back. You might want to
execute some code, after a period of inactivity, for example to refresh the displayed data by doing
a new database query, or even after a long period, to terminate the program automatically.

## Implementing the ON IDLE trigger

To detect user inactivity during a dialog, define an `ON IDLE` trigger in the
dialog. This trigger is dialog specific, it is typically defined in the main dialog of the program,
but it can also be defined in every dialog.
> **Important:**
>
> - Consider using the `ON IDLE` only in dialogs that do not handle field input, such
>   as `DISPLAY ARRAY` and `MENU`. If used in input dialogs, this trigger may execute
>   while the current field contains an incomplete value. The trigger will produce field value
>   validation and raise an input error. However, `ON IDLE` can be used in input dialogs
>   where the user cannot enter invalid values (for example when using `CHECKBOX`,
>   `RADIOGROUP`, `COMBOBOX`, and character-type fields like
>   `TEXTEDIT`).
> - When implementing [multiple dialogs](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.") with
>   `DIALOG` block, do not mix several `ON IDLE` clauses in different
>   sub-dialog blocks: specify a unique `ON IDLE` clause at the `DIALOG`
>   block level.
> - Obviously, it does not make much sense to mix `ON
>   TIMER` and `ON IDLE` clauses.

For example:

```
DEFINE seconds SMALLINT
LET seconds = 120
DISPLAY ARRAY ...
   ...
   ON IDLE seconds
      MESSAGE "Automatic data refresh..."
      -- Reload the array with a new database result set
...
```

Note that the parameter of the `ON IDLE` trigger can be an integer variable, but
it will only be read when the dialog is started. Changing the variable during dialog execution will
have no effect.

A value of zero or less than zero disables the timeout trigger.

## Related links

**Related concepts**  

[Get program control on a regular (timed) basis](2227-get-program-control-on-a-regular-timed-basis.md "Execute some code after a given number of seconds, with or without user interaction with the program.")
