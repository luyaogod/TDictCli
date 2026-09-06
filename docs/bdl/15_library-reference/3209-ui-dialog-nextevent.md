---
title: "ui.Dialog.nextEvent"
source: "fgl-topics/c_fgl_ClassDialog_nextEvent.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.nextEvent"
type: "concept"
---

# ui.Dialog.nextEvent

> Waits for a dialog event.

## Syntax

```
nextEvent()
   RETURNS STRING
```

1. event is the name of the dialog event that raised.

## Usage

The `nextEvent()` waits for a dialog event to occur, and returns a string that
identifies the dialog event that is raised.

This method must be used in a `WHILE` loop, to implement a dynamic dialog.

The method can return `NULL`, if a dialog error occurs, or if the dialog
terminates (with [`ui.Dialog.close()`](3189-ui-dialog-close.md "Closes a dynamic dialog.")).

The `nextEvent()` method returns only basic event description. For detailed
information about the last event that has occurred, use the [`ui.Dialog.getEventDescription()`](3196-ui-dialog-geteventdescription.md "Returns a detailed description of the last event that has occurred in a dynamic dialog.") method.

The recommended programming pattern for the event `WHILE` loop is to test for
nulls:

```
DEFINE d ui.Dialog,
       t STRING
...
WHILE (t := d.nextEvent()) IS NOT NULL
    CASE t
        WHEN "BEFORE FIELD cust_name"
...
```

A dialog event can be a user-defined trigger such as `"ON ACTION print"`, or an
implicit trigger such as `"BEFORE ROW"`, corresponding to the control blocks that can
be defined in static dialog instructions such as [`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.").

User-defined triggers are added to the dynamic dialog with the [`addTrigger()`](3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog") method:

| Trigger name | Description | Dialog block equivalent |
| --- | --- | --- |
| `ON ACTION action-name` | Action handler for the action identified by action-name. | [`ON ACTION` block](../11_user-interface/1918-on-action-block.md) |
| `ON APPEND` | Row addition action handler for a display array dynamic dialog. | [`ON APPEND` block](../11_user-interface/1984-on-append-block.md) |
| `ON DELETE` | Row deletion action handler for a display array dynamic dialog. | [`ON DELETE` block](../11_user-interface/1987-on-delete-block.md) |
| `ON FILL BUFFER` | Trigger to fill the current page of a paged mode display array dynamic dialog. | [`ON FILL BUFFER` block](../11_user-interface/1968-on-fill-buffer-block.md) |
| `ON INSERT` | Row insertion action handler for a display array dynamic dialog. | [`ON INSERT` block](../11_user-interface/1985-on-insert-block.md) |
| `ON IDLE seconds` | Idle timout trigger. | [`ON IDLE` block](../11_user-interface/1897-on-idle-block.md) |
| `ON SORT` | Rowset sort event. | [`ON SORT` block](../11_user-interface/1989-on-sort-block.md) |
| `ON TIMER seconds` | Timer trigger. | [`ON TIMER` block](../11_user-interface/1899-on-timer-block.md) |
| `ON UPDATE` | Row modification action handler for a display array dynamic dialog. | [`ON UPDATE` block](../11_user-interface/1986-on-update-block.md) |

Implicit dialog triggers are predefined and can be detected and handled in the dialog
`WHILE` loop if needed:

| Trigger name | Description | Dialog block equivalent |
| --- | --- | --- |
| `BEFORE DIALOG` | Initialization of the multiple dynamic dialog. | [`BEFORE DIALOG` block](../11_user-interface/2101-before-dialog-block.md) |
| `AFTER DIALOG` | End of the multiple dynamic dialog. | [`AFTER DIALOG` block](../11_user-interface/2102-after-dialog-block.md) |
| `BEFORE DISPLAY` | In singular dynamic dialog, initialization of the display array. In multiple dynamic dialog, when display array gets focus. | [`BEFORE DISPLAY` block](../11_user-interface/1973-before-display-block.md) |
| `AFTER DISPLAY` | In singular dynamic dialog, end of the display array. In multiple dynamic dialog, when display array loses focus. | [`AFTER DISPLAY` block](../11_user-interface/1974-after-display-block.md) |
| `BEFORE INPUT` | In singular dynamic dialog, initialization of the input by name. In multiple dynamic dialog, when input by name gets focus. | [`BEFORE INPUT` block](../11_user-interface/1940-before-input-block.md) |
| `AFTER INPUT` | In singular dynamic dialog, end of the input by name. In multiple dynamic dialog, when input by name loses focus. | [`AFTER INPUT` block](../11_user-interface/1941-after-input-block.md) |
| `BEFORE CONSTRUCT` | In singular dynamic dialog, initialization of the construct. In multiple dynamic dialog, when contruct gets focus. | [`BEFORE INPUT` block](../11_user-interface/2057-before-construct-block.md) |
| `AFTER CONSTRUCT` | In singular dynamic dialog, end of the construct. In multiple dynamic dialog, when construct loses focus. | [`AFTER INPUT` block](../11_user-interface/2058-after-construct-block.md) |
| `BEFORE ROW` | Moving to a new row in a display array or input array dynamic dialog. | [`BEFORE ROW` block](../11_user-interface/1975-before-row-block.md) |
| `AFTER ROW` | Leaving the current row in a display array or input array dynamic dialog. | [`AFTER ROW` block](../11_user-interface/1976-after-row-block.md) |
| `BEFORE INSERT` | Before a new row is created in an input array dynamic dialog. | [`BEFORE INSERT` block](../11_user-interface/2020-before-insert-block.md) |
| `AFTER INSERT` | After a new row is created in an input array dynamic dialog. | [`AFTER INSERT` block](../11_user-interface/2021-after-insert-block.md) |
| `BEFORE DELETE` | Before a new row is deleted in an input array dynamic dialog. | [`BEFORE DELETE` block](../11_user-interface/2022-before-delete-block.md) |
| `AFTER DELETE` | After a new row is deleted in an input array dynamic dialog. | [`AFTER DELETE` block](../11_user-interface/2023-after-delete-block.md) |
| `BEFORE FIELD field-name` | Entering the field field-name in an input dynamic dialog. | [`BEFORE FIELD` block](../11_user-interface/1942-before-field-block.md) |
| `AFTER FIELD field-name` | Leaving the field field-name in an input dynamic dialog. | [`AFTER FIELD` block](../11_user-interface/1944-after-field-block.md) |
| `ON CHANGE field-name` | Value of field field-name changed in an input dynamic dialog. | [`ON CHANGE` block](../11_user-interface/1943-on-change-block.md) |

## Related links

**Related concepts**  

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
