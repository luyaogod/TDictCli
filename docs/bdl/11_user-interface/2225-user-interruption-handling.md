---
title: "User interruption handling"
source: "fgl-topics/c_fgl_prog_dialogs_interruption.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > User interruption handling"
type: "concept"
---

# User interruption handling

> Allow the end user to cancel a dialog or a long running procedure.

## Basics

To properly understand how Genero handles user interruption, we first need to identify the
different events that can occur for a Genero process, and how they are triggered:

1. The Unix "interrupt signal" (SIGINT): Can be triggered by a kill -s INT
   command, or by a `CTRL-C` key in a terminal (see stty -a)). To
   detect the `CTRL-C` key on Microsoft Windows platforms, the concept is a bit
   different and managed by Console Control Handlers. Note that a Unix SIGINT signal can occur in TUI
   or GUI mode.
2. The Unix "quit signal" (SIGQUIT): Can be triggered by a kill -s QUIT command,
   or by a `CTRL-\` key in a terminal (see stty -a)). Like SIGINT, a
   SIGQUIT signal can occur in TUI or GUI mode on Unix.
3. The 'cancel' (and 'close') predefined actions of singular dialogs, when the runtime system is
   waiting for a user action.
4. The asynchronous user interruption event, fired by a action view (button) with the name
   'interrupt', when the runtime system is in processing mode (FOR loop, SQL query)

If the program executes an interactive instruction, the GUI front-end can send action events
triggered by the user such as 'accept' and 'cancel'. But when the program performs a long process
like a loop, a report, or a database query, the front-end has no control. In the second case, user
interruption can be detected by using the [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") register and the 'interrupt' action name.

## Tips for SIGINT and SIGQUIT signals

Typical Genero programs should use [`DEFER
INTERRUPT` and `DEFER QUIT`](../09_advanced-features/0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received."). Otherwise, any SIGINT or SIGQUIT signal
will just stop the process.

Sending a SIGINT when a program is in a [SLEEP](../08_language-basics/0684-sleep.md "The SLEEP instruction causes the program to pause for the specified number of seconds.") will cancel the sleep, because this is a system call that is
interruptible.

In GUI mode, one can still send SIGINT signals to the fglrun process, but
that's not good practice and will not stop the current dialog as in TUI mode.

## Detecting dialog cancellation

When executing a [singular dialogs](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.") such as
`INPUT BY NAME`, the runtime system creates automatically the 'accept' and 'cancel'
[predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions."), to let the user
validate or reject the dialog. When executing a multiple dialog, there is no predefined 'cancel'
action created, and `int_flag` is not set, even for user-defined `ON ACTION
cancel` handlers.

With singular dialogs, the [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") register is typically tested after the dialog block (or in the
`AFTER INPUT` block for example), to detect if the user wants the dialog input to be
validated or rejected. Note that `int_flag` is not reset to `FALSE`
when the automatic 'accept' action is fired. Therefore, it is good practice to initialize
`int_flag` to `FALSE`, before starting a singular dialog.

In TUI mode, to cancel a singular dialog, user presses the `CTRL-C` key in the
terminal, which produces a SIGINT signal and leaves the dialog instruction, after setting
`int_flag` to `1/TRUE`. Since the trigger to cancel a singular dialog
in TUI mode is the SIGINT signal, the dialog can also be canceled by sending the SIGINT signal to
the `fglrun` process. This is however not good practice. Dialog cancellation with
SIGINT is supported for legacy TUI applications.

In GUI mode, to cancel a singular dialog, user fires the 'cancel' (or 'close') predefined action,
to leave the dialog instruction, and `int_flag` is set to `1/TRUE`.
There is no interrupt signal or interrupt event used in this case: This is triggered by a regular
GUI action event.

## Detecting asynchronous user interruptions

To detect user interruptions coming from a GUI front-end while a program is in processing mode,
define an action view with the name
'interrupt':

```
BUTTON sb: interrupt, TEXT="Stop";
```

When the runtime system takes control to process program code or execute a long running SQL
query, the front-end automatically enables the local 'interrupt' action to let the user send an
asynchronous interruption request to the program.

A program (the runtime system) can also receive a SIGINT interruption signal from the
operating system. The interruption request that comes from the front-end is a different source.
However, the runtime system handles both type of interruption events the same way.

When receiving an interrupt event from the front-end (from the particular 'interrupt' action), or
from the system (SIGINT) the runtime system sets the [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") register to `TRUE`.

Consider using `DEFER INTERRUPT` and
test the `int_flag` register to properly handle user interruptions, and avoid
immediate program termination. If the `DEFER INTERRUPT` instruction is not used, the
program will stop immediately when an interruption event is caught. With `DEFER
INTERRUPT`, the program continues, and can test `int_flag` to check if an
interruption event occurred. It is good practice to reset `int_flag` to
`FALSE` after detecting interruption:

```
WHILE ...
   IF int_flag THEN
      LET int_flag=FALSE
      ERROR "Procedure was interrupted by the user"
      EXIT WHILE
   END IF
   ...
END WHILE
```

SQL queries can be interrupted too, if the target database supports this feature and the
`OPTIONS SQL INTERRUPT ON` instruction is used. However, since the control is on the
database server side while the SQL statement is running, it is not possible to execute program code
to check `int_flag`. In order to detect an SQL interruption, check the sqlca.sqlcode
register after the query for SQL error -213, indicating that the last SQL statement was
interrupted.

> **Tip:**
>
> It is good practice to surround the SQL statement with `OPTIONS
> SQL INTERRUPT ON/OFF` and `WHENEVER ERROR CONTINUE/STOP` (or a
> `TRY/CATCH` block): You don't want other SQL statements to be interruptible, and if
> they fail but should not (for example if the table does not exist), the program must
> stop.

```
OPTIONS SQL INTERRUPT ON
WHENEVER ERROR CONTINUE
SELECT COUNT(*) INTO cnt FROM ...  -- Long running SQL statement
WHENEVER ERROR STOP
OPTIONS SQL INTERRUPT OFF
IF sqlca.sqlcode == -213 THEN
   ERROR "Database query interrupted by user"
   ...
END IF
```

When not using `DEFER INTERRUPT`, if the program enters in a long running
procedure, a button with the action name 'interrupt' will become active. The user can then
press that button, and the runtime system will stop the program, since `DEFER
INTERRUPT` is not used. However, this will not happen when a dialog is active,
because the 'interrupt' button will be automatically disabled in that context. Such
situation can confuse the end user, expecting that the 'interrupt' button can stop the
program in any context.

Note that the front-end can not handle interruption requests properly, if the application code
generates a lot of AUI traffic, by doing many [`ui.Interface.refresh()`](../15_library-reference/3117-ui-interface-refresh.md "Synchronize the user interface with the front-end.") calls. This refresh method should not be called more
often than once in one second.

## Implementing interruption of a long running SQL query

```
-- db_busy.per
LAYOUT
 GRID
 {
  Database query in progress...
        [sb            ]
 }
 END
END
ATTRIBUTES
 BUTTON sb: interrupt, TEXT="Stop";
END
```

```
MAIN
  DEFINE oc INT
  DEFER INTERRUPT
  DATABASE stores
  OPEN FORM f FROM "db_busy"
  DISPLAY FORM f
  CALL ui.Interface.refresh()
  OPTIONS SQL INTERRUPT ON
  WHENEVER ERROR CONTINUE
  SELECT COUNT(*) INTO oc FROM orders
  WHENEVER ERROR STOP
  OPTIONS SQL INTERRUPT OFF
  IF sqlca.sqlcode == -213 THEN
    ERROR "Database query has been interrupted..."
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Using SQL interruption](../10_sql-support/0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
