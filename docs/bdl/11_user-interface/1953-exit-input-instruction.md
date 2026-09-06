---
title: "EXIT INPUT instruction"
source: "fgl-topics/c_fgl_record_input_EXIT_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT control instructions > EXIT INPUT instruction"
type: "concept"
description: "Syntax EXIT INPUT Usage The EXIT INPUT instruction terminates the INPUT instruction and resumes the program execution at the instruction following the INPUT block. The dialog is exited immediately. ..."
---

# EXIT INPUT instruction

## Syntax

```
EXIT INPUT
```

## Usage

The `EXIT INPUT` instruction terminates the `INPUT` instruction and
resumes the program execution at the instruction following the `INPUT` block. The
dialog is exited immediately.

```
INPUT BY NAME cust_rec.*
  ...
  ON ACTION exit_now
    EXIT INPUT
  ...
END INPUT
```

Performing an `EXIT INPUT` instruction during a dialog is equivalent to canceling
the dialog: No field validation will occur, and the [`AFTER FIELD`](1944-after-field-block.md) or [`AFTER
INPUT`](1941-after-input-block.md) blocks will not be executed.

`EXIT INPUT` does not set `int_flag` to `TRUE` as when the cancel action is fired.

The `EXIT INPUT` instruction can only be used in a singular `INPUT`
dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog block.

## Related links

**Related concepts**  

[ACCEPT INPUT instruction](1951-accept-input-instruction.md "ACCEPT INPUT instruction")

[CONTINUE INPUT instruction](1952-continue-input-instruction.md "CONTINUE INPUT instruction")

**Related reference**  

[INPUT control blocks execution order](1939-input-control-blocks-execution-order.md "INPUT control blocks execution order")
