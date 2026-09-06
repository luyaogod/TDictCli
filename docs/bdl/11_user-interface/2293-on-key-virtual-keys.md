---
title: "ON KEY Virtual keys"
source: "fgl-topics/c_fgl_prog_dialogs_virtual_keys.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Keyboard accelerator names > ON KEY Virtual keys"
type: "concept"
---

# ON KEY Virtual keys

> Virtual keys are the key names that can be used in program instructions such as ON KEY and COMMAND KEY.

An `ON KEY` block defines one to four
different action objects that will be identified by the key name in lowercase (`ON KEY(F5,F6)
= creates Action f5 + Action f6`). Each action object will get an
`acceleratorName` attribute assigned. In GUI mode, [action defaults](1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.") are applied for
`ON KEY` actions by using the name of the key. You can define secondary accelerator
keys, as well as default decoration attributes like button text and image, by using the key name as
action identifier. The action name is always in lowercase letters.

Check carefully the `ON KEY CONTROL-?` statements because they may result
in having duplicate accelerators for multiple actions due to the accelerators defined by
action defaults. Additionally, it is recommended to avoid using `ON KEY`
statements with `ESC`, `TAB`, `UP`,
`DOWN`, `LEFT`, `RIGHT`,
`HELP`, `NEXT`, `PREVIOUS`,
`INSERT`, `CONTROL-M`, `CONTROL-X`,
`CONTROL-V`, `CONTROL-C`, and
`CONTROL-A` in GUI programs, because they are very likely to clash
with default accelerators defined in the action defaults.

By default, `ON KEY` actions are not decorated with [default action view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."): To get a visible
default action view for `ON KEY`, specify a `text` attribute with the action defaults.

| Key Name | Description |
| --- | --- |
| `ACCEPT` | The validation key. |
| `INTERRUPT` | The interruption key. |
| `ESC` or `ESCAPE` | The ESC key (not recommended, use `ACCEPT` instead). |
| `TAB` | The TAB key (not recommended). |
| `Control-char` | A control key where *char* can be any character except A, D, H, I, J, K, L, M, R, or X. |
| `F1` through `F255` | A function key. |
| `DELETE` | The key used to delete a new row in an array. |
| `INSERT` | The key used to insert a new row in an array. |
| `HELP` | The help key. |
| `LEFT` | The left arrow key. |
| `RIGHT` | The right arrow key. |
| `DOWN` | The down arrow key. |
| `UP` | The up arrow key. |
| `PREVIOUS` or `PREVPAGE` | The previous page key. |
| `NEXT` or `NEXTPAGE` | The next page key. |

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
