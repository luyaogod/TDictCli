---
title: "ON KEY block"
source: "fgl-topics/c_fgl_dialog_ON_KEY_6.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG interaction blocks > ON KEY block"
type: "concept"
description: "Syntax ON KEY ( key-name ) instruction [...] ON KEY basics An ON KEY ( key-name ) block defines an action with a hidden action view (no default button is visible), that executes a sequence of ..."
---

# ON KEY block

## Syntax

```
ON KEY (key-name)
   instruction [...]
```

## ON KEY basics

An `ON KEY (key-name)` block defines an action with a hidden
action view (no default button is visible), that executes a sequence of instructions when the user
presses the specified key.

The `ON KEY` block is supported for backward compatibility with TUI mode
applications. In new developments, consider using [`ON
ACTION`](1896-on-action-block.md) with accelerators defined in action defaults.

An `ON KEY` block can specify up to four accelerator keys. Each key creates a
specific action object that will be identified by the key name in lowercase.

For example, `ON KEY(F5,F6)` creates two actions with the names
`f5` and `f6`. Each action object will get an
`acceleratorName` attribute assigned, with the corresponding accelerator name. The
specified keys must be one of [the virtual
keys](2293-on-key-virtual-keys.md "Virtual keys are the key names that can be used in program instructions such as ON KEY and COMMAND KEY.").

> **Note:**
>
> The `KEY()` clause allows a comma-separated list of keys. Up to four keys can be
> specified. For new developments, consider using a single key, or prefer [`ON ACTION`](1896-on-action-block.md) handlers with a single accelerator
> definition in action defaults.

In GUI mode, action defaults are applied for `ON KEY` actions by using the name of
the action (the key name). You can define secondary accelerator keys, as well as default decoration
attributes like button text and image, by using the key name as action identifier. The action name
is always in lowercase letters.

Check carefully `ON KEY CONTROL-?` statements to avoid having duplicate
accelerators for multiple actions due to the accelerators defined by action defaults. Additionally,
`ON KEY` statements used with `ESC`, `TAB`,
`UP`, `DOWN`, `LEFT`, `RIGHT`,
`HELP`, `NEXT`, `PREVIOUS`, `INSERT`,
`CONTROL-M`, `CONTROL-X`, `CONTROL-V`,
`CONTROL-C` and `CONTROL-A` should be avoided for use in GUI programs,
because it's very likely to clash with default accelerators defined in the factory action defaults
file provided by default.

By default, `ON KEY` actions are not decorated with a default button in the action
frame (the default action view). You can show the default button by configuring a
`text` attribute with the action
defaults.

```
ON KEY (CONTROL-Z)
   CALL open_zoom()
```

## ON KEY(INTERRUPT)

In a singular dialog like `INPUT`, `ON KEY (INTERRUPT)` is mapped
to the [automatic "cancel" action](2257-list-of-predefined-actions.md). As result, the front-end sees only a "cancel" action and the code
of the `ON KEY(INTERRUPT)` is executed when the user clicks on the Cancel button.

In a multiple dialog, `ON KEY (INTERRUPT)` creates a regular user-defined action,
without any accelerator key.

Avoid to combine `ON KEY (INTERRUPT)` with `ON ACTION cancel`. In a
singular dialog, if both are used, the code of the last declared handler will be executed when the
user cancels the dialog.

## ON KEY in PROMPT

The [`PROMPT`](1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
dialog is automatically terminated after `ON
IDLE`, `ON TIMER`,
`ON ACTION`, or `ON KEY` block execution.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.")
