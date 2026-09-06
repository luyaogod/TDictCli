---
title: "ui.Dialog.addTrigger"
source: "fgl-topics/c_fgl_ClassDialog_addTrigger.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.addTrigger"
type: "concept"
---

# ui.Dialog.addTrigger

> Adds an event trigger to the dynamic dialog

## Syntax

```
addTrigger(
   trigger STRING )
```

1. trigger is the name of the dialog.

## Usage

When implementing a dynamic dialog, the `addTrigger()` method must be used to
register user-defined triggers:

```
CALL d.addTrigger("ON ACTION print")
```

Registered dialog triggers are then typically managed in a `WHILE` loop using the
[nextEvent()](3209-ui-dialog-nextevent.md "Waits for a dialog event.") method, to wait for dialog
events.

Predefined triggers such as "`BEFORE ROW`", "`BEFORE FIELD`",
"`ON CHANGE`" are already registered in dynamic dialogs. Such triggers do not have to
be added with the `addTrigger()` method. User code can be implemented for predefined
triggers when returned from the `nextEvent()` method.

Dynamic multiple dialogs are implemented with [ui.Dialog.createMultipleDialog](3175-ui-dialog-createmultipledialog.md "Creates an ui.Dialog object to implement a dynamic DIALOG multiple-dialog."). When calling the `addTrigger()`
method just after `createMultipleDialog()`, it will add a global dialog trigger. When
called after methods such as [ui.Dialog.addDisplayArrayTo](3180-ui-dialog-adddisplayarrayto.md "Adds a sub-dialog of type DISPLAY ARRAY TO to an existing ui.Dialog dynamic dialog."), it will add a
local trigger to the last added sub-dialog.

Sub-dialog actions that are created in the context of a sub-dialog with `addTrigger("ON
ACTION action-name")` will be returned as `"ON ACTION
sub-dialog-name.action-name"` from the
`nextEvent()` method.

The following triggers are accepted by the `addTrigger()` method:

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

## Related links

**Related concepts**  

[ui.Dialog.nextEvent](3209-ui-dialog-nextevent.md "Waits for a dialog event.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
