---
title: "ACCELERATOR4 attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ACCELERATOR4.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ACCELERATOR4 attribute"
type: "concept"
---

# ACCELERATOR4 attribute

> The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action.

## Syntax

```
ACCELERATOR4 = { key | "key" }
```

1. key defines the accelerator key. This can be a
   combination of `CONTROL-`, `SHIFT-`, `ALT-` and
   identifier or a digit. The string literal is allowed to define accelerators such as
   `"*"` to specify an asterisk character.

## Usage

This attribute is an action attribute that can be specified in form `ACTION
DEFAULTS`, for more details, see [ACCELERATOR4 action attribute](2270-accelerator4-action-attribute.md "The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action.").

## Related links

**Related concepts**  

[Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")
