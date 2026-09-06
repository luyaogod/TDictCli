---
title: "Dynamic multiple dialog"
source: "fgl-topics/c_fgl_dynamic_dialogs_multiple.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Dynamic multiple dialog"
type: "concept"
---

# Dynamic multiple dialog

> Dynamic dialogs allows you to create the equivalent of a DIALOG procedural dialog.

## Steps to implement a dynamic multiple dialog

The order in which dialog methods are called defines the structure of the dynamic multiple
dialog.

To build a dynamic multiple dialog:

1. Create the multiple dialog with [`createMultipleDialog()`](../15_library-reference/3175-ui-dialog-createmultipledialog.md "Creates an ui.Dialog object to implement a dynamic DIALOG multiple-dialog.").
2. Add global dialog triggers with [`addTrigger()`](../15_library-reference/3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog").
3. Define the fields for a sub-dialog with a `DYNAMIC ARRAY OF RECORD` with
   `name` and `type` members.
4. Add the sub-dialog with one of the following methods:
   - [`addConstructByName()`](../15_library-reference/3179-ui-dialog-addconstructbyname.md "Adds a sub-dialog of type CONSTRUCT BY NAME to an existing ui.Dialog dynamic dialog.")
   - [`addDisplayArrayTo()`](../15_library-reference/3180-ui-dialog-adddisplayarrayto.md "Adds a sub-dialog of type DISPLAY ARRAY TO to an existing ui.Dialog dynamic dialog.")
   - [`addInputArrayFrom()`](../15_library-reference/3182-ui-dialog-addinputarrayfrom.md "Adds a sub-dialog of type INPUT ARRAY FROM to an existing ui.Dialog dynamic dialog.")
   - [`addInputByName()`](../15_library-reference/3183-ui-dialog-addinputbyname.md "Adds a sub-dialog of type INPUT BY NAME to an existing ui.Dialog dynamic dialog.")
5. Add sub-dialog triggers with [`addTrigger()`](../15_library-reference/3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog"). When calling this method after a sub-dialog item is added, it
   will create the trigger in the sub-dialog context.
6. Implement the `WHILE` loop using the [`nextEvent()`](../15_library-reference/3209-ui-dialog-nextevent.md "Waits for a dialog event.") method, as for a [simple dynamic dialog](2430-the-dynamic-dialog-event-loop.md "Dynamic dialog body is implemented with a WHILE loop using ui.Dialog.nextEvent()."). However, sub-dialog
   triggers must be identified with the sub-dialog name.

## Identifying global and sub-dialog actions

Sub-dialog actions are added with `addTrigger("ON ACTION
action-name")` in the context of a sub-dialog definition, after
calling a method such as `addDisplayArrayTo()`.

In the event loop, the `nextEvent()` method returns sub-dialog actions events with
the sub-dialog name prefix, using the `"ON ACTION
sub-dialog-name.action-name"` text form:

```
...
CALL d.addDisplayArrayTo(fields, "sr_cust")
CALL d.addTrigger("ON ACTION refresh")
...
WHILE (t := d.nextEvent()) IS NOT NULL
  CASE t
     ...
     WHEN "ON ACTION sr_cust.refresh"
       ...
```

## Multiple dialog initialization and termination triggers

As in a static multiple dialog, the `BEFORE DIALOG` and `AFTER
DIALOG` events indicate respectively the initialization and termination of the dynamic
multiple
dialog:

```
WHILE (t := d.nextEvent()) IS NOT NULL
  CASE t
     ...
     WHEN "BEFORE DIALOG"
       ...
     WHEN "AFTER DIALOG"
       ...
```

## Identifying sub-dialog triggers

Sub-dialog triggers such as `BEFORE DISPLAY`, `BEFORE INPUT` can be
identified in the event loop with the sub-dialog identifier. The `nextEvent()` method
returns the event text in the form `"trigger-name
sub-dialog-name"`.

For example, when you define the `DISPLAY ARRAY`-style sub-dialog with the screen
record name "sr\_cust", the trigger `"BEFORE DISPLAY sr_cust"` will be generated when
the `sr_cust` list gets the focus (as in a static `DIALOG / END
DIALOG` multiple dialog block):

```
...
CALL d.addDisplayArrayTo(fields, "sc_cust")
...
CALL d.addDisplayArrayTo(fields, "sc_prod")
...
WHILE (t := d.nextEvent()) IS NOT NULL
  CASE t
     ...
     WHEN "BEFORE DISPLAY sc_cust"
       ...
     WHEN "AFTER DISPLAY sc_cust"
       ...
     WHEN "BEFORE DISPLAY sc_prod"
       ...
     WHEN "AFTER DISPLAY sc_prod"
       ...
```

Note that the `nextEvent()` method returns only basic information about the event
that occurred. For example, `"BEFORE FIELD"` events contain only the field name
without a prefix that allows to identify the sub-dialog. To get detailed information about the last
event, use the [`ui.Dialog.getEventDescription()`](../15_library-reference/3196-ui-dialog-geteventdescription.md "Returns a detailed description of the last event that has occurred in a dynamic dialog.") method.

## Example

This example defines a dynamic multiple dialog with a `DISPLAY ARRAY` and
`CONSTRUCT BY NAME` sub-dialogs, and adds `ON ACTION` triggers at
global and sub-dialog levels:

```
TYPE t_fields DYNAMIC ARRAY OF RECORD
                name STRING,
                type STRING
              END RECORD
DEFINE d ui.Dialog,
       list_fields t_fields,
       qbe_fields t_fields
...
OPEN WINDOW w1 WITH FORM "myform"
...
LET d = ui.Dialog.createMultipleDialog()
-- add global triggers
CALL d.addTrigger("ON ACTION accept")
CALL d.addTrigger("ON ACTION cancel")
...
-- Add a DISPLAY ARRAY sub-dialog with name "sr_cust" (screen array)
CALL d.addDisplayArrayTo(list_fields, "sr_cust")
CALL d.addTrigger("ON ACTION refresh")
...
-- Add CONSTRUCT BY NAME sub-dialog with name "qr_cust"
CALL d.addConstructByName(qbe_fields, "qr_cust")
CALL d.addTrigger("ON ACTION query")
...
WHILE (t := d.nextEvent()) IS NOT NULL
  CASE t
     ...
     WHEN "ON ACTION sr_cust.refresh"
       CALL refresh_list(d)
     ...
     WHEN "ON ACTION qr_cust.query"
       CALL exec_query(d)
     ...
     WHEN "ON ACTION accept"
       CALL d.accept()
     WHEN "ON ACTION close"
       EXIT WHILE
     WHEN "AFTER DIALOG" -- after d.accept()
       EXIT WHILE
  END CASE
END WHILE
CALL d.close()
LET d = NULL
CLOSE WINDOW w1
```

## Related links

**Related concepts**  

[Sub-dialog actions in procedural DIALOG blocks](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.")
