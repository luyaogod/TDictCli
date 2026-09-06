---
title: "Ending dynamic dialogs"
source: "fgl-topics/c_fgl_dynamic_dialogs_terminate.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Ending dynamic dialogs"
type: "concept"
---

# Ending dynamic dialogs

> Describes how to terminate dynamic dialogs.

## Implementing the accept and cancel actions

Regular static dialog instructions implement the accept and cancel actions, to respectively
validate or abort the dialog.

These actions are created automatically for static dialogs, but must be created by hand for
dynamic dialogs.

In the case of cancel, you can mimic the behavior of static dialogs by setting the
`int_flag` register to `TRUE` and then directly leave the
`WHILE` loop with an `EXIT WHILE`.

For the accept action, call the `ui.Dialog.accept()` method to validate field
input and leave the dialog, and put the `EXIT WHILE` in the termination event, to
leave the dialog events loop.

For example, to implement the accept and cancel actions for a simple record
input:

```
DEFINE d ui.Dialog,
       t STRING
...
LET d = ui.Dialog.createInputByName(fields)
CALL d.addTrigger("ON ACTION cancel")
CALL d.addTrigger("ON ACTION accept")
...
WHILE (t := d.nextEvent()) IS NOT NULL
    CASE t
    WHEN "ON ACTION cancel"
      LET int_flag = TRUE
      EXIT WHILE
    WHEN "ON ACTION accept"
      CALL d.accept()
    WHEN "AFTER INPUT"
      EXIT WHILE
  END CASE
END WHILE
```

For a multiple dynamic dialog, the termination event is `"AFTER DIALOG"`, and the
`EXIT WHILE` must be done when that event occurs (for the accept
action):

```
...
WHILE (t := d.nextEvent()) IS NOT NULL
    CASE t
    ...
    WHEN "AFTER DIALOG"
      EXIT WHILE
  END CASE
END WHILE
```

In multiple dialogs, the `AFTER INPUT`, `AFTER DISPLAY`,
`AFTER CONSTRUCT` triggers are used to indicate that the focus leaves the
corresponding sub-dialog. Do not perform `EXIT WHILE` in these triggers for a
multiple dialog.

## Terminating the dialog

Some synchronization code needs to be implemented to properly destroy the dynamic dialog.

A dialog needs to be destroyed before closing its corresponding window/form.

In order to terminate a dialog, call the [`close()`](../15_library-reference/3189-ui-dialog-close.md "Closes a dynamic dialog.") dialog method and assign `NULL` to the
`ui.Dialog` variable referencing the dialog object. This will close the dialog and
destroy the corresponding object, if no other variables references it.

When the dialog object is terminated, the corresponding window can also be
closed:

```
...
  WHEN "ON ACTION cancel"
    EXIT WHILE
END WHILE
CALL d.close()
LET d = NULL
CLOSE WINDOW w1
```
