---
title: "ACCEPT CONSTRUCT instruction"
source: "fgl-topics/c_fgl_Construct_ACCEPT_CONSTRUCT.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control instructions > ACCEPT CONSTRUCT instruction"
type: "concept"
description: "The ACCEPT CONSTRUCT instruction validates and exits the CONSTRUCT instruction, if no error is raised. The AFTER FIELD and AFTER CONSTRUCT control blocks will be executed. The statements after the ..."
---

# ACCEPT CONSTRUCT instruction

The `ACCEPT CONSTRUCT` instruction validates and exits the
`CONSTRUCT` instruction, if no error is raised.

The [`AFTER FIELD`](1944-after-field-block.md) and [`AFTER CONSTRUCT`](2058-after-construct-block.md) control blocks will
be executed.

The statements after the `ACCEPT CONSTRUCT` will be skipped.

```
CONSTRUCT BY NAME where_part ON ...
  ...
  ON ACTION default_query
    CALL set_default_filter()
    ACCEPT CONSTRUCT
  ...
END CONSTRUCT
```

The `CONSTRUCT` instruction creates the default accept action to let the user
validate the dialog.

Use the `ACCEPT CONSTRUCT` instruction only in specific cases when the default
accept action is not appropriate.

The `ACCEPT CONSTRUCT` instruction can only be used in a singular
`CONSTRUCT` dialog, it cannot be used in a `DIALOG / END DIALOG`
multiple dialog block.

## Related links

**Related concepts**  

[NEXT FIELD instruction](1955-next-field-instruction.md "NEXT FIELD instruction")

[EXIT CONSTRUCT instruction](2070-exit-construct-instruction.md "EXIT CONSTRUCT instruction")

[CONTINUE CONSTRUCT instruction](2069-continue-construct-instruction.md "CONTINUE CONSTRUCT instruction")
