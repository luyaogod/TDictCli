---
title: "Accelerator key modifiers"
source: "fgl-topics/c_fgl_prog_dialogs_key_modifiers.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Keyboard accelerator names > Accelerator key modifiers"
type: "concept"
---

# Accelerator key modifiers

> Key modifiers define keyboard control key combinations in accelerator key names.

The accelerator key names such as `"F10"` can be combined with key modifiers, by
using a minus sign (hyphen) as
separator:

```
key-modifier - [ key-modifier - [ key-modifier - ] ] key-name
```

| Key Modifier | Description |
| --- | --- |
| `Control` | The left or right `[Ctrl]` key. |
| `Shift` | The left of right `[Shift]` key. |
| `Alt` | The left or right `[Alt]` key. |

For example:

```
Control-P 
Shift-Alt-F12 
Control-Shift-Alt-Z
```

## Related links

**Related concepts**  

[Accelerator key names](2294-accelerator-key-names.md "Accelerators keys are attributes defining the keyboard shortcuts for actions.")
