---
title: "ACCELERATOR attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ACCELERATOR.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ACCELERATOR attribute"
type: "concept"
---

# ACCELERATOR attribute

> The ACCELERATOR is an action attribute defining the primary accelerator key for an action.

## Syntax

```
ACCELERATOR = { key | "key" }
```

1. key defines the accelerator key. This can be a
   combination of `CONTROL-`, `SHIFT-`, `ALT-` and
   identifier or a digit. The string literal is allowed to define accelerators such as
   `"*"` to specify an asterisk character.

## Usage

This attribute is an action attribute that can be specified in form `ACTION
DEFAULTS`, for more details, see [ACCELERATOR action attribute](2267-accelerator-action-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action.").

## Example:

```
ACTION DEFAULTS
  ...
  ACTION print ( TEXT="Print", ACCELERATOR = CONTROL-P )
  ...
END
```

## Related links

**Related concepts**  

[Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")
