---
title: "ACCEPT DIALOG instruction"
source: "fgl-topics/c_fgl_DIALOG_instr_ACCEPT_DIALOG_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control instructions > ACCEPT DIALOG instruction"
type: "concept"
description: "Syntax CANCEL DIALOG Usage The ACCEPT DIALOG statement validates all input fields bound to the DIALOG instruction and leaves the block, if no error is raised. When used in a declarative DIALOG block , ..."
---

# ACCEPT DIALOG instruction

## Syntax

```
CANCEL DIALOG
```

## Usage

The `ACCEPT DIALOG` statement validates all input fields bound to the
`DIALOG` instruction and leaves the block, if no error is raised.

When used in a [declarative
`DIALOG` block](2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs."), the `ACCEPT DIALOG` instruction does only make
sense when the declarative dialog block is included in a procedural dialog block with the
`SUBDIALOG` clause.

When defined in the dialog block, `ON CHANGE`, `AFTER FIELD`,
`AFTER ROW`, `AFTER INPUT`, `AFTER DISPLAY`,
`AFTER CONSTRUCT` control blocks will be executed when `ACCEPT DIALOG`
is performed.

The statements appearing after the `ACCEPT DIALOG` instruction will be
skipped.

You typically code an `ACCEPT DIALOG` in an `ON ACTION accept`
block:

```
ON ACTION accept ACCEPT DIALOG
```

Any usage of `ACCEPT DIALOG` outside an `ON ACTION accept` block is
not intended and its behavior is undefined.

Input field validation is a process that does several successive validation tasks:

1. The current field value is checked, depending on the variable data type (for example, the user
   must input a valid date in a DATE field).
2. [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") field attributes
   are checked for all input fields. This attribute forces the field to have a value set by program or
   entered by the user. If the field contains no value, the constraint is not satisfied. Input values
   are right-trimmed, so if the user inputs only spaces, this corresponds to a NULL value which does
   not fulfill the `NOT NULL` constraint.
3. [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") field attributes
   are checked for all input fields. This attribute forces the field to have a default value, or to be
   modified by the user or by program with a `DISPLAY TO / BY NAME` or
   `DIALOG.setFieldTouched()` call. If the field was not modified during the dialog, the
   `REQUIRED` constraint is not satisfied.
4. [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") field attributes
   are checked for all input fields. This attribute forces the field to contain a value that is listed
   in the include list. If the field contains a value that is not in the list, the constraint is not
   satisfied.

If a field does not satisfy one of these constraints, dialog termination is canceled, an error
message is displayed, and the focus goes to the first field causing a problem.

After input field validation has succeeded, different types of control blocks will be executed,
such as [`AFTER FIELD`](1944-after-field-block.md), [`AFTER ROW`](1976-after-row-block.md), [`AFTER INPUT`](1941-after-input-block.md) and [`AFTER DIALOG`](2102-after-dialog-block.md).

In order to validate some parts of the dialog without leaving the block, use the [`DIALOG.validate()`](../15_library-reference/3233-ui-dialog-validate.md "Checks form level validation rules.") method.

## Related links

**Related concepts**  

[Input field modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.")

[CANCEL DIALOG instruction](2143-cancel-dialog-instruction.md "CANCEL DIALOG instruction")

[EXIT DIALOG instruction](2141-exit-dialog-instruction.md "EXIT DIALOG instruction")

[ui.Dialog.accept](../15_library-reference/3178-ui-dialog-accept.md "Validates and terminates the dialog.")
