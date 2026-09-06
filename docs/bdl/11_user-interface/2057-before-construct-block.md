---
title: "BEFORE CONSTRUCT block"
source: "fgl-topics/c_fgl_dialog_BEFORE_CONSTRUCT.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control blocks > BEFORE CONSTRUCT block"
type: "concept"
description: "Syntax BEFORE CONSTRUCT instruction [...] BEFORE CONSTRUCT block in singular CONSTRUCT dialogs In a singular CONSTRUCT instruction, the BEFORE CONSTRUCT is only executed once when dialog is started. ..."
---

# BEFORE CONSTRUCT block

## Syntax

```
BEFORE CONSTRUCT
   instruction [...]
```

## BEFORE CONSTRUCT block in singular CONSTRUCT dialogs

In a singular `CONSTRUCT` instruction, the `BEFORE CONSTRUCT` is
only executed once when dialog is started.

The `BEFORE CONSTRUCT` block is executed once at dialog start-up, before the
runtime system gives control to the user for criteria input. This block can be used to display
messages to the user, initialize form fields with default search criteria values, and setup the
dialog instance by deactivating unused fields or actions the user is not allowed to execute.

```
CONSTRUCT BY NAME where_part ON ...
   BEFORE CONSTRUCT
       MESSAGE "Enter customer search filter"
       CALL DIALOG.setActionActive("clean", FALSE )
   ...
```

The fields are cleared before the `BEFORE CONSTRUCT` block is executed.

You can use the `NEXT FIELD` control instruction in the `BEFORE
CONSTRUCT` block, to jump to a specific field when the dialog starts.

## BEFORE CONSTRUCT block in CONSTRUCT of procedural DIALOG

In a `CONSTRUCT` sub-dialog of a procedural `DIALOG` instruction,
the `BEFORE CONSTRUCT` block is executed when the focus goes to a group of fields
driven by a `CONSTRUCT` sub-dialog. This trigger is only invoked if a field of the
sub-dialog gets the focus, and none of the other fields had the focus.

`BEFORE CONSTRUCT` is executed after the `BEFORE DIALOG`
block and before the `BEFORE FIELD` blocks.

In this example, the `BEFORE CONSTRUCT` block is used to display a message:

```
CONSTRUCT BY NAME sql ON customer.*
  BEFORE CONSTRUCT
    MESSAGE "Enter customer search filter"
```

## Related links

**Related concepts**  

[BEFORE INPUT block](1940-before-input-block.md "BEFORE INPUT block")

[BEFORE DISPLAY block](1973-before-display-block.md "BEFORE DISPLAY block")

[AFTER CONSTRUCT block](2058-after-construct-block.md "AFTER CONSTRUCT block")
