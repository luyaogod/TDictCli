---
title: "KEYS section"
source: "fgl-topics/c_fgl_FormSpecFiles_KEYS_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > KEYS section"
type: "concept"
---

# KEYS section

> The KEYS section can be used to define default key labels for the current form.

## Syntax

> **Important:**
>
> This feature is supported for backward compatibility. Consider using [action attributes](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") to define accelerator keys
> and decorate actions.

```
KEYS
key-name = [%]"label"
[...]
[END]
```

1. key-name is the name of a key (like `F10`,
   `Control-z`).
2. label is the text to be displayed in the button corresponding to the
   key.

## Usage

The `KEYS` section can be used to define default key labels at the form level.

The `KEYS` section must appear in the sequence described
in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `KEYS` section is optional in a form definition.

The `END` keyword is optional.

## Example

```
KEYS
  F10 = "City list"
  F11 = "State list"
  F15 = "Validate"
END
```

## Related links

**Related concepts**  

[Setting action key labels](2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.")
