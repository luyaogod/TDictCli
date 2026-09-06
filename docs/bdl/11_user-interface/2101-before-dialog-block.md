---
title: "BEFORE DIALOG block"
source: "fgl-topics/c_fgl_DIALOG_block_BEFORE_DIALOG.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > BEFORE DIALOG block"
type: "concept"
description: "Syntax BEFORE DIALOG instruction [...] Usage The BEFORE DIALOG block is executed one time as the first trigger when the DIALOG instruction starts, before the runtime system gives control to the user. ..."
---

# BEFORE DIALOG block

## Syntax

```
BEFORE DIALOG
   instruction [...]
```

## Usage

The `BEFORE DIALOG` block is executed one time as the first trigger when the
`DIALOG` instruction starts, before the runtime system gives control to the user. You
can implement variable initialization and dialog configuration in this block.

Like all control blocks of the procedural `DIALOG / END DIALOG` instruction, the
`BEFORE DIALOG` block must appear after the sub-dialog definitions. For a detailed
description, see the [procedural
`DIALOG` instruction syntax](2078-syntax-of-the-procedural-dialog-instruction.md "The DIALOG block is an interactive instruction that executes several sub-dialogs simultaneously.").

The fields are initialized with the default values, before the `BEFORE
DIALOG` code is executed. When an `INPUT` sub-dialog uses the `WITHOUT
DEFAULTS` option or when inserting/appending a new row in an `INPUT ARRAY`
sub-dialog, the default values are taken from the program variables bound to the fields. Otherwise
(with defaults), the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.")
attributes of the form fields are used.

In the next code example, the `BEFORE DIALOG` block performs some dialog setup and
gives the focus to a specific field:

```
DIALOG
   DISPLAY ARRAY ...
      ...
   END DISPLAY
   INPUT BY NAME ...
      ...
   END INPUT
   ...
   BEFORE DIALOG
     CALL DIALOG.setActionActive("save",FALSE)
     CALL DIALOG.setFieldActive("cust_status", is_admin())
     IF cust_is_new() THEN
        NEXT FIELD cust_name
     END IF
   ...
END DIALOG
```

A `DIALOG` instruction can include no more than one `BEFORE DIALOG`
control block.

## Related links

**Related concepts**  

[The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")

[AFTER DIALOG block](2102-after-dialog-block.md "AFTER DIALOG block")
