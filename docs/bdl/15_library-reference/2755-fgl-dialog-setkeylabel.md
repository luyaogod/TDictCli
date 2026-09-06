---
title: "fgl_dialog_setkeylabel()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_SETKEYLABEL.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_setkeylabel()"
type: "concept"
---

# fgl_dialog_setkeylabel()

> Sets the label associated to a key for the current interactive instruction.

## Syntax

```
FUNCTION fgl_dialog_setkeylabel(
   keyName STRING,
   text STRING )
```

1. keyName is the logical name of a key such as `F11`
   or `DELETE`,`INSERT`, `CANCEL`.
2. text is the text associated to the key.

## Usage

The `fgl_dialog_setkeylabel()` associates a text description to a function or
control key for the current dialog. Default action views (buttons that appears in the action panel
of a window) will get the label displayed instead of the function or control key name.

This function defines the key labels for the current dialog. There
are different levels of key label definitions.

> **Note:**
>
> This feature is supported for backward compatibility. Consider using [action attributes](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") to define accelerator keys
> and decorate actions.

## Related links

**Related concepts**  

[Setting action key labels](../11_user-interface/2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
