---
title: "Defining control keys"
source: "fgl-topics/c_fgl_programs_013.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Defining control keys"
type: "concept"
---

# Defining control keys

> The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).

## Syntax

```
OPTIONS logical-key KEY physical-key
```

where logical-key is one
of:

```
{ ACCEPT | HELP | INSERT | DELETE | NEXT | PREVIOUS }
```

1. logical-key identifies a logical key corresponding to a dialog action or
   operation.
2. physical-key defines the keyboard accelerator (F1, CTRL-Z), to be associated
   with the logical-key.

## Usage

The `OPTIONS logical-key KEY physical-key`
instruction can be used to associate physical keys (F2) for logical keys correspongding to a dialog
operation (screen record page up/down) or dialog action (such as `insert`), to be
triggered in the current interactive instruction.

This form of the `OPTIONS` instruction is provided for backward compatibility with
the TUI mode: In GUI mode, use the [action
configuration](../11_user-interface/2265-defining-keyboard-accelerators-for-actions.md) to define accelerator keys for actions.

| OPTIONS clause | Corresponding action / dialog function | Default Physical Key |
| --- | --- | --- |
| `OPTIONS ACCEPT KEY` | The [`accept` action](../11_user-interface/2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions."), that validates the current dialog. | `ESCAPE` |
| `OPTIONS HELP KEY` | The `help` action, to display [help messages](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs."). | `CONTROL-W` |
| `OPTIONS INSERT KEY` | The `insert` action, to insert a new row in an [`INPUT ARRAY`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."). | `F1` |
| `OPTIONS DELETE KEY` | The `delete` action, to deletes the current row in an `INPUT ARRAY`. | `F2` |
| `OPTIONS NEXT KEY` | Scrolls to the next screen record page in an `INPUT ARRAY` or [`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions."). | `F3` |
| `OPTIONS PREVIOUS KEY` | Scrolls to the previous screen record page in an `INPUT ARRAY` or `DISPLAY ARRAY`. | `F4` |

In TUI mode, the key to cancel a dialog is the interruption key `CTRL-C`. In this
context, the key to cancel dialogs can only be configured with the TTY settings, to raise an
interruption signal (see stty command).

You can specify the following keywords for the physical key names:

| Key Name | Description |
| --- | --- |
| `ESC` or `ESCAPE` | The ESC key (not recommended, use `ACCEPT` instead). |
| `INTERRUPT` | The interruption key (on UNIX™, interruption signal). |
| `TAB` | The TAB key (not recommended). |
| `CONTROL-char` | A control key where *char* can be any character except A, D, H, I, J, K, L, M, R, or X |
| `F1` through `F255` | A function key. |
| `LEFT` | The left arrow key. |
| `RETURN` or `ENTER` | The return key. |
| `RIGHT` | The right arrow key. |
| `DOWN` | The down arrow key. |
| `UP` | The up arrow key. |
| `PREVIOUS` or `PREVPAGE` | The previous page key. |
| `NEXT` or `NEXTPAGE` | The next page key. |

You might not be able to use other keys that have special meaning to your version of
the operating system. For example, `CONTROL-C`, `CONTROL-Q`,
and `CONTROL-S` specify the Interrupt, XON, and XOFF signals on many
UNIX systems.

## Related links

**Related concepts**  

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
