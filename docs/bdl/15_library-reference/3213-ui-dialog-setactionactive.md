---
title: "ui.Dialog.setActionActive"
source: "fgl-topics/c_fgl_ClassDialog_setActionActive.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setActionActive"
type: "concept"
---

# ui.Dialog.setActionActive

> Enabling and disabling dialog actions.

## Syntax

```
setActionActive(
  name STRING,
  active BOOLEAN )
```

1. name is the name of the action, see [Identifying actions in dialog methods](3238-identifying-actions-in-ui-dialog-methods.md).
2. active is a boolean value.

## Usage

Use the `setActionActive()` method to enable or disable an
action.

```
CALL DIALOG.setActionActive("zoom", FALSE)
```

The second parameter of the method must be a boolean expression that
evaluates to 0 (`FALSE`) or 1 (`TRUE`).

## Related links

**Related concepts**  

[Enabling and disabling actions](../11_user-interface/2282-enabling-and-disabling-actions.md "By default, dialog actions are enabled. However, it is recommended that an action be disabled when not allowed in the current context.")
