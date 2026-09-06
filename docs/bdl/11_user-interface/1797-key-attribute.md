---
title: "KEY attribute"
source: "fgl-topics/c_fgl_FSFAttributes_KEY.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > KEY attribute"
type: "concept"
---

# KEY attribute

> The KEY attribute is used to define the labels of keys when the field is made current.

## Syntax

```
KEY keyname = [%]"label"
```

1. keyname is the name of a key ( like `F10`,
   `"Control-z"` ).
2. label is the text to be displayed in the button corresponding to the
   key.

## Usage

Use the `KEY` attribute to define a label for the accelerator key corresponding
to an action when the focus is in the field.

The keyname must be specified in quotes if you want to use Control / Shift
/ Alt key modifiers.

See the `KEYS` section to define key labels for the whole form.

> **Note:**
>
> This feature is supported for backward compatibility. Consider using [action attributes](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") to define accelerator keys
> and decorate actions.

## Example

```
EDIT f001 = customer.city, KEY F10 = "City list";
EDIT f002 = customer.state, KEY "Control-z" = "Open Zoom";
```

## Related links

**Related concepts**  

[Setting action key labels](2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.")

[KEYS section](1752-keys-section.md "The KEYS section can be used to define default key labels for the current form.")
