---
title: "ACCELERATOR3 attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ACCELERATOR3.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ACCELERATOR3 attribute"
type: "concept"
---

# ACCELERATOR3 attribute

> The ACCELERATOR3 is an action attribute defining the third accelerator key for an action.

## Syntax

```
ACCELERATOR3 = { key | "key" }
```

1. key defines the accelerator key. This can be a
   combination of `CONTROL-`, `SHIFT-`, `ALT-` and
   identifier or a digit. The string literal is allowed to define accelerators such as
   `"*"` to specify an asterisk character.

## Usage

This attribute is an action attribute that can be specified in form `ACTION
DEFAULTS`, for more details, see [ACCELERATOR3 action attribute](2269-accelerator3-action-attribute.md "The ACCELERATOR3 is an action attribute defining the third accelerator key for an action.").

## Related links

**Related concepts**  

[Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")
