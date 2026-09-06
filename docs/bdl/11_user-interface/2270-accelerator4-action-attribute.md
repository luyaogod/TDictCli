---
title: "ACCELERATOR4 action attribute"
source: "fgl-topics/c_fgl_action_attribute_ACCELERATOR4.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > ACCELERATOR4 action attribute"
type: "concept"
---

# ACCELERATOR4 action attribute

> The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action.

## Syntax 1 (Dialog action handlers):

There is no syntax to define accelerator #4 in dialog action handlers.

## Syntax 2 (`ACTION DEFAULTS` section in form files):

```
ACCELERATOR4 = { key | "key" }
```

## Syntax 3 (Global .4ad action defaults file):

```
acceleratorName4 = "key"
```

1. key defines the accelerator key. This can be a
   combination of `CONTROL-`, `SHIFT-`, `ALT-` and
   identifier or a digit. The string literal is allowed to define accelerators such as
   `"*"` to specify an asterisk character.

## Usage

The `ACCELERATOR4` attribute defines the keyboard combination
that can be pressed by the user to send an action to the program.

This attribute is provided for specific cases, consider using only [one accelerator per action](1754-accelerator-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action.").

## Related links

**Related concepts**  

[ACCELERATOR4 attribute](1757-accelerator4-attribute.md "The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action.")

[Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.")

[Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
