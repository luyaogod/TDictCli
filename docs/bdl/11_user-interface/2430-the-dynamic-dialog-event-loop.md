---
title: "The dynamic dialog event loop"
source: "fgl-topics/c_fgl_dynamic_dialogs_triggers.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > The dynamic dialog event loop"
type: "concept"
---

# The dynamic dialog event loop

> Dynamic dialog body is implemented with a WHILE loop using ui.Dialog.nextEvent().

## Dialog trigger identification

The code implementing dynamic dialogs must be generic. Triggers like `ON ACTION
myaction` are identified as a simple string.

## Implicit standard triggers

Standard dialog triggers such as `BEFORE INPUT`, `BEFORE FIELD
field-name`, `AFTER ROW` do not need to be added when
creating a dynamic dialog: These triggers are implicitly declared. To see a reaction to standard
triggers, add a test for the trigger in the event loop.

## Adding user-defined triggers

Dynamic dialogs can be configured with user-defined triggers, for example to execute code when a
specific action is fired.

After creating the dialog object, add user-defined triggers with the [`ui.Dialog.addTrigger()`](../15_library-reference/3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog")
method:

```
DEFINE d ui.Dialog
...
CALL d.addTrigger("ON ACTION print")
CALL d.addTrigger("ON DELETE")
...
```

Note that some triggers must be identified with the user-defined action name, as in `"ON
ACTION print"`.

User-defined triggers will then be handled in the dynamic dialog loop, when the event occurs.

## Handling dialog events

To implement the "body" of a dynamic dialog, mix a `WHILE` loop with the
`ui.Dialog.nextEvent()` method, to handle dialog events.

The `WHILE` loop will act as the main event handler of your dynamic dialog, and
will loop, waiting for dialog events until you explicitly exit the loop with an `EXIT
WHILE` instruction.

The `nextEvent()` method can return `NULL` in case of dialog error
or when the dialog is terminated.

```
DEFINE d ui.Dialog,
       t STRING
...
WHILE (t := d.nextEvent()) IS NOT NULL
    CASE t
       WHEN "BEFORE DISPLAY"
          ...
       WHEN "ON ACTION print"
          ...
       WHEN "ON DELETE"
          ...
       WHEN "AFTER DISPLAY"
          ...
END WHILE
```

The standard triggers such as `"BEFORE ROW"` and `"AFTER FIELD
field-name"` are equivalent to the trigger found in static dialog control
blocks. These allow you to control the behavior of the dynamic dialog.

The `ui.Dialog.nextEvent()` only provides basic information about the event that
occurred. For example, with `BEFORE FIELD` events only the field name is part of the
string returned by `nextEvent()`. Use the
`ui.Dialog.getEventDescription()` method to get detailed information about the last
event.

The user-defined triggers that have been added with the `addTrigger()` method must
also be handled in the dynamic dialog loop.

Inside the `WHILE` loop, control the behavior of the dialog with the methods
provided in the `ui.Dialog` class. For example, to jump to a different field when the
"jump" action is
fired:

```
 ...
 WHEN "ON ACTION jump"
   CALL d.nextField("customer.cust_name")
 ...
```

BEFORE/AFTER FIELD handlers must be identified with the field name (without the table/formonly
prefix):

```
 ...
 WHEN "AFTER FIELD cust_name"
   IF LENGTH(d.getFieldValue("customer.cust_name")) < 3 THEN
      ERROR "Customer name is too short"
      CALL d.nextField("customer.cust_name")
 END IF
 ...
```

For more details, see the [ui.Dialog.nextEvent()](../15_library-reference/3209-ui-dialog-nextevent.md "Waits for a dialog event.") method reference.
