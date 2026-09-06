---
title: "ACCELERATOR2 attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ACCELERATOR2.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ACCELERATOR2 attribute"
type: "concept"
---

# ACCELERATOR2 attribute

> The ACCELERATOR2 is an action attribute defining the secondary accelerator key for an action.

## Syntax

```
ACCELERATOR2 = { key | "key" }
```

1. key defines the accelerator key. This can be a
   combination of `CONTROL-`, `SHIFT-`, `ALT-` and
   identifier or a digit. The string literal is allowed to define accelerators such as
   `"*"` to specify an asterisk character.

## Usage

This attribute is an action attribute that can be specified in form `ACTION
DEFAULTS`, for more details, see [ACCELERATOR2 action attribute](2268-accelerator2-action-attribute.md "The ACCELERATOR2 is an action attribute defining the secondary accelerator key for an action.").

## Related links

**Related concepts**  

[Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")
