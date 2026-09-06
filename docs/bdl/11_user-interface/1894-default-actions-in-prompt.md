---
title: "Default actions in PROMPT"
source: "fgl-topics/c_fgl_prompt_007.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Using simple prompt inputs > Default actions in PROMPT"
type: "concept"
description: "When a PROMPT instruction executes, the runtime system creates a set of default actions . Depending on the invoked default action, field validation occurs and different PROMPT control blocks are ..."
---

# Default actions in PROMPT

When a `PROMPT` instruction executes, the runtime
system creates a set of [default actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.").

Depending on the invoked default action, field validation occurs and different
`PROMPT` control blocks are executed.

This table lists the default actions created for this dialog:

| Default action | Description |
| --- | --- |
| `accept` | Validates the `PROMPT` dialog (validates field criteria)*Creation can be avoided with the* `ACCEPT` *attribute.* |
| `cancel` | Cancels the `PROMPT` dialog (no validation, `int_flag` is set)*Creation can be avoided with the* `CANCEL` *attribute.* |
| `close` | By default, cancels the `PROMPT` dialog (no validation, `int_flag` is set)Default action view is hidden. See [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). |
| `help` | Shows the help topic defined by the `HELP` clause.*Only created when a* `HELP` *clause is defined.* |

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
