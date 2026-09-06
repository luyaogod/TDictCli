---
title: "Accelerator key names"
source: "fgl-topics/c_fgl_prog_dialogs_accelerator_keys.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Keyboard accelerator names > Accelerator key names"
type: "concept"
---

# Accelerator key names

> Accelerators keys are attributes defining the keyboard shortcuts for actions.

Action accelerator keys are typically associated with action default attributes. For a complete
usage description, see [Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md).

Keyboard key names can be combined with [key
modifiers](2295-accelerator-key-modifiers.md "Key modifiers define keyboard control key combinations in accelerator key names."), to define key combinations such as [Ctrl] + [P].

> **Tip:**
>
> To force an action to have no accelerator, specify "`none`" as the
> accelerator name.

| Key Name | Description |
| --- | --- |
| `none` | Special name indicating the runtime system must not set any accelerator for the action. |
| `0-9` | Decimal digit keys from [0] to [9] |
| `A-Z` | Letters keys from [A] to [Z] |
| `F1-F35` | The functions keys like [F10] |
| `BackSpace` | The [Backspace] or [←] key |
| `Del` | The [Del] key (numeric keypad) |
| `Delete` | The [Delete] key (navigation keyboard group) |
| `Down` | The [Down] / [↓] key |
| `End` | The [End] key |
| `Enter` | The [Enter] key (numeric keypad, see [Note](2294-accelerator-key-names.md)) |
| `Escape` | The [Esc] key |
| `Home` | The [Home] key |
| `Ins` | The [Ins] key (numeric keypad) |
| `Insert` | The [Insert] key (navigation keyboard group) |
| `Left` | The [Left] / [←] key |
| `Minus` | The [-] minus sign key (hyphen) |
| `Prior` | The [PgUp] key |
| `Next` | The [PgDn] key |
| `Return` | The [Return] key (alphanumeric keypad, see [Note](2294-accelerator-key-names.md)) |
| `Right` | The [→] key |
| `Space` | The [\_\_\_\_(spacebar)\_\_\_\_] key |
| `Tab` | The [Tab] key |
| `Up` | The [↑] key |

> **Note:**
>
> The "`Enter`" accelerator key name represents the `[ENTER]` key
> available on the numeric keypad of standard keyboards, while "`Return`" represents
> the `[RETURN]` key of the alphanumeric keyboard. By default, the "accept" validation
> action is configured to accept both "`Enter`" and "`Return`" keys.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")
