---
title: "VALIDATE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_VALIDATE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > VALIDATE attribute"
type: "concept"
---

# VALIDATE attribute

> The VALIDATE action attribute defines the data validation level for a given action.

## Syntax

```
VALIDATE = NO
```

## Usage

This attribute is an action attribute that can be specified in form `ACTION
DEFAULTS`.

When the `VALIDATE` action attribute is set to `NO`, it indicates
that no data validation must occur for this action. However, current input buffer contains the text
modified by the user before triggering the action.

For more details, see [VALIDATE action attribute](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.").
