---
title: "The DEFINE clause"
source: "fgl-topics/c_fgl_DIALOG_block_DEFINE.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Structure of a declarative DIALOG block > The DEFINE clause"
type: "concept"
---

# The DEFINE clause

> The DEFINE clause can be used to define program variables with a scope that is local to the declarative dialog block.

This clause must be placed before any other sub-dialog
block:

```
DIALOG customer_input(p_cust t_cust INOUT)
    DEFINE checked BOOLEAN
    DEFINE tmp STRING

    INPUT BY NAME p_cust.*
        ...
    END INPUT

END DIALOG
```

The `DEFINE` clause is only allowed in declarative dialog blocks. Variables used
locally in a procedural dialog block must be defined in the scope of the function containing the
procedural dialog block.

## Related links

**Related concepts**  

[Variables](../08_language-basics/0686-variables.md "Explains how to define program variables.")
