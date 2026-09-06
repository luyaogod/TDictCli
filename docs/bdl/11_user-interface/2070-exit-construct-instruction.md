---
title: "EXIT CONSTRUCT instruction"
source: "fgl-topics/c_fgl_Construct_EXIT_CONSTRUCT.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control instructions > EXIT CONSTRUCT instruction"
type: "concept"
description: "The EXIT CONSTRUCT instruction terminates the CONSTRUCT instruction and resumes the program execution at the instruction following the CONSTRUCT block. CONSTRUCT BY NAME where_part ON ... ... ON ..."
---

# EXIT CONSTRUCT instruction

The `EXIT CONSTRUCT` instruction terminates the `CONSTRUCT`
instruction and resumes the program execution at the instruction following the
`CONSTRUCT` block.

```
CONSTRUCT BY NAME where_part ON ...
  ...
  ON ACTION exit_now
    EXIT CONSTRUCT
  ...
END CONSTRUCT
```

Performing an `EXIT CONSTRUCT` instruction during a dialog is equivalent to
canceling the dialog: No field validation will occur, and the [`AFTER FIELD`](1944-after-field-block.md) or [`AFTER CONSTRUCT`](2058-after-construct-block.md) blocks will
not be executed.

`EXIT CONSTRUCT` does not set `int_flag` to `TRUE` as when the cancel action is fired.

The `EXIT CONSTRUCT` instruction can only be used in a singular
`CONSTRUCT` dialog, it cannot be used in a `DIALOG / END DIALOG`
multiple dialog block.

## Related links

**Related concepts**  

[ACCEPT CONSTRUCT instruction](2068-accept-construct-instruction.md "ACCEPT CONSTRUCT instruction")

[CONTINUE CONSTRUCT instruction](2069-continue-construct-instruction.md "CONTINUE CONSTRUCT instruction")
