---
title: "Get program control on a regular (timed) basis"
source: "fgl-topics/c_fgl_prog_dialogs_on_timer.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Get program control on a regular (timed) basis"
type: "concept"
---

# Get program control on a regular (timed) basis

> Execute some code after a given number of seconds, with or without user interaction with the program.

## When to use the ON TIMER trigger?

In some cases, an application needs to execute code at a regular interval, for example to process
a message arrived in a queue, refresh data on a dashboard, or display resources in time-based
graphs.

> **Important:**
>
> Unlike the `ON IDLE` trigger which executes when there
> is no user activity, the `ON TIMER`
> trigger executes even when the user interacts with the application. Therefore, the code executed in
> an `ON TIMER` trigger must perform quickly, otherwise the end user will experience
> poor performance. As a general rule, make sure the time spent in the `ON TIMER` code
> is less than the timer interval. For example, if the processing time takes about 2 seconds, it doesn't
> make sense to have an `ON TIMER` that triggers every second.

## Implementing the ON TIMER trigger

To return control to the program on regular intervals, use the `ON TIMER
seconds` trigger in dialogs. This trigger is dialog specific. It is
typically defined in the main dialog of the program, but it can be defined in every dialog.

> **Important:**
>
> - Consider using the `ON TIMER` only in dialogs that do not handle field input,
>   such as `DISPLAY ARRAY` and `MENU`. If used in input dialogs, this
>   trigger may execute in the middle of a field input, which can produce field value validation and
>   raise an input error. However, `ON TIMER` can be used in input dialogs where the user
>   cannot enter invalid values (for example when using `CHECKBOX`,
>   `RADIOGROUP`, `COMBOBOX`, and character-type fields like
>   `TEXTEDIT`)
> - When implementing multiple dialogs with `DIALOG` block, do not mix several
>   `ON TIMER` clauses in different sub-dialog blocks: Specify a unique `ON
>   TIMER` clause at the `DIALOG` block level. Obviously, it does not make much
>   sense to mix `ON TIMER` and `ON IDLE` clauses.

For example:

```
DEFINE seconds SMALLINT
LET seconds = 120
DISPLAY ARRAY ...
   ...
   ON TIMER seconds
      MESSAGE "Check for messages in queue..."
      -- Query the message server for new messages.
...
```

Note that the parameter of the `ON TIMER` trigger can be an integer variable, but
it will only be read when the dialog is started. Changing the variable during dialog execution will
have no effect.

A value of zero or less than zero disables the timeout trigger.

## Related links

**Related concepts**  

[Get program control if user is inactive](2226-get-program-control-if-user-is-inactive.md "Execute some code after a given number of seconds, when the user does not interact with the program.")
