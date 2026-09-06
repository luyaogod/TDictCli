---
title: "Understanding predefined actions"
source: "fgl-topics/c_fgl_prog_dialogs_predact_basics.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Predefined actions > Understanding predefined actions"
type: "concept"
description: "Basics Predefined actions are different from user-defined action, in the sense that the name of a predefined action is reserved, and the action may have an ON ACTION handler, while user-defined ..."
---

# Understanding predefined actions

## Basics

Predefined actions are different from user-defined action, in the sense that the
name of a predefined action is reserved, and the action may have an `ON ACTION`
handler, while user-defined actions have a specific name, and must be implemented with an `ON
ACTION` handler.

There are three types of predefined actions:

- [Automatic actions](2257-list-of-predefined-actions.md): actions that are automatically created and handled by the
  program dialog, like `accept`, `cancel`, `insert`.
- [Particular actions](2257-list-of-predefined-actions.md): actions that do not need an `ON ACTION`
  handler, like `interrupt`.
- [Special actions](2257-list-of-predefined-actions.md): actions with a special usage, such as
  `dialogtouched`.

Default decoration attributes and keyboard shortcuts are defined with [action defaults](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes."), like for user-defined
actions.

## Context-depend activiation

Some predefined actions (such as insert, append and delete in `INPUT ARRAY`) are
enabled and disabled automatically by the dialog, depending on the context. For example, when a
static array used by the `INPUT ARRAY` is full, the insert and append actions get
disabled.

Even when overwriting such actions with your own action handler, the runtime system will continue
to enable and disable the actions automatically.

Overwriting predefined actions is not recommended.

## Binding action views to predefined actions

As for user-defined actions, if you design forms with action views using predefined action names,
they will automatically [attach themselves](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")
to the actions of the interactive instructions.

It is also possible to define default images, texts, comments and accelerator keys in the action
defaults resource file for the predefined actions.

## Overwriting predefined actions

If you define your own `ON ACTION`
handler with the name of a predefined action, the default action processing is bypassed and the
program code is executed instead.

This code example defines an `ON ACTION` clause with the `accept`
predefined action name:

```
INPUT BY NAME customer.*
  ON ACTION accept
    ...
END INPUT
```

In this case, the default behavior of the automatic accept action is not performed; the user code
is executed instead.
