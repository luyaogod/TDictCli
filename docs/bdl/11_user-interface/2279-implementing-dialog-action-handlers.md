---
title: "Implementing dialog action handlers"
source: "fgl-topics/c_fgl_prog_dialogs_action_handler.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Implementing dialog action handlers"
type: "concept"
---

# Implementing dialog action handlers

> How to execute user code in ON ACTION blocks when an action is fired.

Actions handlers are typically defined in dialog instructions with the `ON ACTION` interaction block. You must
specify the name of the action after the `ON ACTION` keywords:

```
INPUT BY NAME ...
   ...
   ON ACTION print
      -- user code
   ...
```

Action handlers can also be defined with the `COMMAND` syntax in
`MENU` and `DIALOG` instructions:

```
MENU ...
   ...
   COMMAND "Print" "Print the current record"
      -- user code
   ...
```

`ON ACTION` blocks provide better abstraction than `COMMAND`
blocks by using simple action identifiers and leaving the decoration in the form files or
action defaults files.

The `ON ACTION` block defines an action handler with a simple action
name.

Built-in class methods such as `ui.Dialog.setActionActive()` take an action name
as parameter. For more details, see [Identifying actions in ui.Dialog methods](../15_library-reference/3238-identifying-actions-in-ui-dialog-methods.md).

The `COMMAND` block defines an action handler with an action name, but it also
defines decoration attributes, such as the label and comment. Keyboard accelerators and help topic
numbers can also be defined.

Action views controlled by `ON ACTION` handlers cannot get the focus. When using
the `COMMAND` action handler, action views such as a `BUTTON` defined
in the form layout can get the focus and are part of the tabbing item list.

Action handlers are bound to [action views](2278-defining-action-views-in-forms.md "How to define action views that will fire action events.")
by name.

## Related links

**Related concepts**  

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")

[COMMAND [KEY] block](2124-command-key-block.md "COMMAND [KEY] block")
