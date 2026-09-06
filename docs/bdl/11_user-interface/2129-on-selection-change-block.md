---
title: "ON SELECTION CHANGE block"
source: "fgl-topics/c_fgl_dialog_ON_SELECTION_CHANGE_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG interaction blocks > ON SELECTION CHANGE block"
type: "concept"
description: "Syntax ON SELECTION CHANGE instruction [...] Usage The ON SELECTION CHANGE trigger can be used to enable multi-row selection and detect when rows are selected or de-selected by the end user during a ..."
---

# ON SELECTION CHANGE block

## Syntax

```
ON SELECTION CHANGE
   instruction [...]
```

## Usage

The `ON SELECTION CHANGE` trigger can be used to enable multi-row selection and
detect when rows are selected or de-selected by the end user during a `DISPLAY ARRAY`
dialog.

If this block is defined, multi-row selection is automatically enabled. However, the feature can
be enabled/disabled with the `setSelectionMode()` dialog method.

## Related links

**Related concepts**  

[Multiple row selection](2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.")
