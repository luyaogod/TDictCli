---
title: "ACCELERATOR action attribute"
source: "fgl-topics/c_fgl_action_attribute_ACCELERATOR.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > ACCELERATOR action attribute"
type: "concept"
---

# ACCELERATOR action attribute

> The ACCELERATOR is an action attribute defining the primary accelerator key for an action.

## Syntax 1 (Dialog action handlers):

```
ACCELERATOR = "key"
```

## Syntax 2 (`ACTION DEFAULTS` section in form files):

```
ACCELERATOR = { key | "key" }
```

## Syntax 3 (Global .4ad action defaults file)

```
acceleratorName = "key"
```

1. key defines the accelerator key. This can be a
   combination of `CONTROL-`, `SHIFT-`, `ALT-` and
   identifier or a digit. The string literal is allowed to define accelerators such as
   `"*"` to specify an asterisk character.

## Usage

The `ACCELERATOR` attribute defines the keyboard combination
that can be pressed by the user to send an action to the program.

Note that in dialog-specific action attributes, the `ACCELERATOR` must be
specified as a string expression.

This attribute applies to the actions defined by the current dialog
in the current window. It can be specified as action default attribute in a global
.4ad file, in the `ACTION DEFAULTS` section of form files, or as
dialog action attribute (using `ON ACTION name
ATTRIBUTES(...)`).

## Example

```
-- As action handler attribute
ON ACTION print ATTRIBUTES(ACCELERATOR="control-p")

-- As action default
ACTION DEFAULTS
  ACTION print (ACCELERATOR=control-p)
END

-- In a global action defaults file
<ActionDefault name="print" acceleratorName="control-p" ... />
```

## Related links

**Related concepts**  

[ACCELERATOR attribute](1754-accelerator-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action.")

[Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.")

[Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
