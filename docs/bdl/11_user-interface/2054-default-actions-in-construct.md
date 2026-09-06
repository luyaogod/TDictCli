---
title: "Default actions IN CONSTRUCT"
source: "fgl-topics/c_fgl_Construct_009.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > Default actions IN CONSTRUCT"
type: "concept"
description: "When an CONSTRUCT instruction executes, the runtime system creates a set of default actions . Depending on the invoked default action, field validation occurs and different CONSTRUCT control blocks ..."
---

# Default actions IN CONSTRUCT

When an `CONSTRUCT` instruction executes, the runtime
system creates a set of [default actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.").

Depending on the invoked default action, field validation occurs and different
`CONSTRUCT` control blocks are executed.

This table lists the default actions created for this dialog:

| Default action | Description |
| --- | --- |
| `accept` | Validates the `CONSTRUCT` dialog (validates field criteria)*Creation can be avoided with* `ACCEPT` *attribute.* |
| `cancel` | Cancels the `CONSTRUCT` dialog (no field validation, [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") is set to `TRUE`)*Creation can be avoided with* `CANCEL` *attribute.* |
| `close` | By default, cancels the `CONSTRUCT` dialog (no validation, int\_flag is set)Default action view is hidden. See [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). |
| `help` | Shows the help topic defined by the `HELP` clause.*Only created when a* `HELP` *clause is defined.* |

The accept and cancel default actions can be avoided with the `ACCEPT` and
`CANCEL` dialog control
attributes:

```
CONSTRUCT BY NAME cond ON field1 ATTRIBUTES (CANCEL=FALSE)
    ...
```

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
