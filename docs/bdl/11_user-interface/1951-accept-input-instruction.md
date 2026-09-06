---
title: "ACCEPT INPUT instruction"
source: "fgl-topics/c_fgl_record_input_ACCEPT_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT control instructions > ACCEPT INPUT instruction"
type: "concept"
description: "Syntax ACCEPT INPUT Usage The ACCEPT INPUT instruction validates and exits the INPUT instruction, if no error is raised. The AFTER FIELD , ON CHANGE , etc. control blocks will be executed. The ..."
---

# ACCEPT INPUT instruction

## Syntax

```
ACCEPT INPUT
```

## Usage

The `ACCEPT INPUT` instruction validates and exits the `INPUT`
instruction, if no error is raised.

The [`AFTER FIELD`](1944-after-field-block.md), [`ON CHANGE`](1943-on-change-block.md), etc. control blocks will be
executed.

The statements after the `ACCEPT INPUT` instruction will be skipped.

```
INPUT BY NAME cust_rec.*
  ...
  ON ACTION process_order
    CALL set_missing_defaults()
    ACCEPT INPUT
  ...
END INPUT
```

The `INPUT` instruction creates the default accept action to let the user validate
the dialog. Use of the `ACCEPT INPUT` instruction is recommended only in specific
cases when the default accept action is not appropriated.

The `ACCEPT INPUT` instruction can only be used in a singular
`INPUT` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple
dialog block.

## Related links

**Related concepts**  

[EXIT INPUT instruction](1953-exit-input-instruction.md "EXIT INPUT instruction")

[CONTINUE INPUT instruction](1952-continue-input-instruction.md "CONTINUE INPUT instruction")

**Related reference**  

[INPUT control blocks execution order](1939-input-control-blocks-execution-order.md "INPUT control blocks execution order")
