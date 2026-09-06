---
title: "fgl_getkeylabel()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETKEYLABEL.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getkeylabel()"
type: "concept"
---

# fgl_getkeylabel()

> Returns the default label associated to a key.

## Syntax

```
FUNCTION fgl_getkeylabel(
   keyName STRING )
  RETURNS STRING
```

1. keyName is the logical name of a key such as `F11` or
   `DELETE`, `INSERT`, `CANCEL`.

## Usage

The `fgl_getkeylabel()` function returns the default
label defined for the function or control key passed as parameter.

This function returns the default key labels defined for all dialogs. There are different levels
of key label definitions.

This function is provided for backward compatibility, use action
defaults to define action view texts.

## Related links

**Related concepts**  

[Setting action key labels](../11_user-interface/2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
