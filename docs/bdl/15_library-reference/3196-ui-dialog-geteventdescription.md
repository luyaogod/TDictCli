---
title: "ui.Dialog.getEventDescription"
source: "fgl-topics/c_fgl_ClassDialog_getEventDescription.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getEventDescription"
type: "concept"
---

# ui.Dialog.getEventDescription

> Returns a detailed description of the last event that has occurred in a dynamic dialog.

## Syntax

```
getEventDescription( )
   RETURNS STRING
```

## Usage

In the loop implementing a dynamic dialog, the `getEventDescription()` method can
be used to get detailed information about the last event that has occurred, typically after calling
the `nextEvent()` method.

The `nextEvent()` method returns only basic information about the event that has
occurred. For example, a `"BEFORE FIELD ..."` events contains only the field name
(without any table name or screen-record prefix). If several fields use the same name (especially
in a multiple-dialog), use the `getEventDescription()` method, to identify what
sub-dialog the field belongs to.

## Example

```
DEFINE d ui.Dialog,
       t, desc STRING
...
WHILE (t := d.nextEvent()) IS NOT NULL
    IF t MATCHES "BEFORE FIELD*" THEN
        LET desc = d.getEventDescription()
        IF desc == "BEFORE FIELD customer.cust_id" THEN
            ...
...
```

## Related links

**Related concepts**  

[ui.Dialog.nextEvent](3209-ui-dialog-nextevent.md "Waits for a dialog event.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
