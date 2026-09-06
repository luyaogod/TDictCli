---
title: "fgl_dialog_getkeylabel()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETKEYLABEL.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getkeylabel()"
type: "concept"
---

# fgl_dialog_getkeylabel()

> Returns the label associated to a key for the current interactive instruction.

## Syntax

```
FUNCTION fgl_dialog_getkeylabel(
   keyName STRING )
  RETURNS STRING
```

1. keyName is the logical name of a key such as `F11`
   or `DELETE`, `INSERT`, `CANCEL`.

## Usage

The `fgl_dialog_getkeylabel()` function returns the label
defined for the function or control key passed as parameter, for the
current interactive instruction.

This function returns the key labels defined for the current dialog.
There are different levels of key label definitions.

This function is provided for backward compatibility, use action defaults
to define action view texts.

## Related links

**Related concepts**  

[Setting action key labels](../11_user-interface/2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
