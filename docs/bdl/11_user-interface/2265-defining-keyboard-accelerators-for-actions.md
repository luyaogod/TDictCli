---
title: "Defining keyboard accelerators for actions"
source: "fgl-topics/c_fgl_prog_dialogs_action_accelerators.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Defining keyboard accelerators for actions"
type: "concept"
description: "Keyboard accelerators are attributes defining the keyboard shortcuts for actions. Keyboard accelerators can be defined at several levels (global action defaults, form file action defaults, dialog ..."
---

# Defining keyboard accelerators for actions

Keyboard accelerators are attributes defining the keyboard shortcuts for actions.

Keyboard accelerators can be defined at several levels (global action defaults, form file action
defaults, dialog instructions).

For example, in a .per form file `ACTION DEFAULTS`
section:

```
ACTION DEFAULTS
   ACTION print (TEXT="Print", ACCELERATOR=Control-P)
END
```

Up to four accelerator keys can be defined for the same action in action defaults. However, as a
general pattern, define only one accelerator per action.

Do not use the same accelerator for different actions that are active in the same dialog
context.

The runtime system sets default accelerators for [predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions."), when no accelerators are
explicitly defined for the action. The default accelerator of predefined actions depends on the
TUI/GUI user interface mode.

For example:

- In TUI mode, the `accept` action gets the `Escape` key, that can
  be changed with the [`OPTIONS ACCEPT KEY`](../09_advanced-features/0932-defining-control-keys.md "The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).")
  instruction.
- In GUI mode, the `accept` action gets the `Return` and
  `Enter` keys from [FGLDIR/lib/default.4ad file](1603-action-default-attributes-reference-4ad.md "This topic contains all attributes you can define in a .4ad action defaults file.").
- In TUI mode, the `INPUT ARRAY` `insert`/`delete`
  predefined actions will respectively get the F1 and F2 accelerators. These can be changed with [`OPTIONS INSERT/DELETE KEY`](../09_advanced-features/0932-defining-control-keys.md "The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).") instructions.
  `INPUT ARRAY` and `DISPLAY ARRAY` navigation (screen record page down
  / page up) is done with F3 / F4. These can be changed with [`OPTIONS NEXT/PREVIOUS KEY`](../09_advanced-features/0932-defining-control-keys.md "The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).") instructions.
- In GUI mode, `INPUT ARRAY` predefined actions and [`DISPLAY ARRAY` modification
  actions](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.") get the default accelerators from the FGLDIR/lib/default.4ad file:
  - `insert` accelerator is F3
  - `append` accelerator is Control-F3
  - `update` accelerator is F2
  - `delete` accelerator is F4

When a user-defined action is configured with an accelerator that would normally be used for a
predefined action, the runtime system does not set that accelerator for the predefined action. For
example (in GUI mode), if you define an `ON ACTION quit` with an action default using
the accelerator "`Escape`", the "cancel" predefined action will not get the
"`Escape`" default accelerator. In this case, user settings take precedence over
defaults.

Text edition and navigation keys such as `[Home]` and `[End]` are
usually local to the widget. Depending on the context, such common keys might be eaten by the
graphical widget and will not invoke the action configured with the corresponding accelerator. For
example, even if the `"firstrow"` action defines the `Home`
accelerator, when using an `INPUT ARRAY`, the `[Home]` key will jump
to the beginning of the edit field, not the first row of the list.

When using the `ON ACTION` clause in a dialog instruction, action accelerators are
used in both GUI and TUI mode. However, for backward compatibility, this is not done in TUI mode
when using the `ON KEY` clause.

The traditional `ON KEY` clause in a dialog like `INPUT` implicitly
defines the `acceleratorName` attribute for the action, and the corresponding action
default accelerator will be ignored. For example, when you define an `ON KEY(F10)`
block, the first accelerator will be "F10", even if an action default defines an accelerator "F5"
for the action "F10". However, you can set other accelerators with the
`acceleratorName2`, `acceleratorName3` and
`acceleratorName4` attributes in action defaults.

> **Important:**
>
> **In TUI mode, actions created with `ON KEY` do not get
> accelerators of action defaults; only actions defined with `ON ACTION` will get
> accelerators of action defaults.**

In a `MENU` dialog, the behavior is a bit different, see the [`COMMAND "option"`](1916-command-key-option-block.md) and [`COMMAND KEY(keyname)`](1917-command-key-block.md) clauses
of `MENU`.

If you want to force an action to have no accelerator, you can use `none` as the
accelerator name.

## Related links

**Related concepts**  

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")

[Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.")
