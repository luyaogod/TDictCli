---
title: "Defining keys for WIDGET fields"
source: "fgl-topics/c_fgl_Mig0000_038.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > Defining keys for WIDGET fields"
type: "concept"
---

# Defining keys for WIDGET fields

> The method for binding a keyboard function or control key differ between Four Js BDS and Genero BDL.

With Four Js Business Development Suite (BDS), a form field using the `WIDGET`
attribute as `"BUTTON"` or as `"CHECK"` or `"RADIO"`
with `CLASS="KEY"`, could specify the key to be sent to the program with the
`CONFIG` attribute:

```
DATABASE FORMONLY
SCREEN
{
[f01     ][f02     ][f03     ]
}
END
ATTRIBUTES
f01 = FORMONLY.b_print, WIDGET="BUTTON", CONFIG="F10";
f02 = FORMONLY.b_ok, WIDGET="BUTTON", CONFIG="KEY_accept";
f03 = FORMONLY.b_cancel, WIDGET="BUTTON", CONFIG="KEY_interrupt";
END
```

The key specification could be a function key like `F10`, a control key like
`Control-P`, or a virtual key like `KEY_accept` (by default, ESC) or
`KEY_interrupt` (by default, Control-C).

With Genero BDL, use the `ACCELERATOR` action attribute to bind a keyboard
function or control key to a named action that is controlled in the program with an `ON
ACTION action-name` handler.

## Related links

**Related concepts**  

[Defining keyboard accelerators for actions](../11_user-interface/2265-defining-keyboard-accelerators-for-actions.md "Defining keyboard accelerators for actions")
