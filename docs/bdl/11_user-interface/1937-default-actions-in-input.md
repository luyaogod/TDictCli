---
title: "Default actions in INPUT"
source: "fgl-topics/c_fgl_record_input_008.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > Default actions in INPUT"
type: "concept"
description: "When an INPUT instruction executes, the runtime system creates a set of default actions . Depending on the invoked default action, field validation occurs and different INPUT control blocks are ..."
---

# Default actions in INPUT

When an `INPUT` instruction executes, the runtime
system creates a set of
[default actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.").

Depending on the invoked default action, field validation occurs and different
`INPUT` control blocks are executed.

This table lists the default actions created for this dialog:

| Default action | Description |
| --- | --- |
| `accept` | Validates the `INPUT` dialog (validates fields and leaves the dialog)*Creation can be avoided with* `ACCEPT` *attribute.* |
| `cancel` | Cancels the `INPUT` dialog (no validation, `int_flag` is set to `TRUE`)*Creation can be avoided with* `CANCEL` *attribute.* |
| `close` | By default, cancels the `INPUT` dialog (no validation, `int_flag` is set to `TRUE`)Default action view is hidden. See [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). |
| `help` | Shows the help topic defined by the `HELP` clause.*Only created when a* `HELP` *clause is defined.* |

The accept and cancel default actions can be avoided with the `ACCEPT` and
`CANCEL` dialog control attributes:

```
INPUT BY NAME field1 ATTRIBUTES ( CANCEL=FALSE )
  ...
```

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
