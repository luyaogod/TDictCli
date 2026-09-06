---
title: "ACCEPT INPUT instruction"
source: "fgl-topics/c_fgl_InputArray_ACCEPT_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY control instructions > ACCEPT INPUT instruction"
type: "concept"
description: "Syntax ACCEPT INPUT Usage The ACCEPT INPUT instruction validates and exits the INPUT ARRAY instruction, if no error is raised. The AFTER FIELD , ON CHANGE , AFTER INSERT , AFTER ROW , etc., control ..."
---

# ACCEPT INPUT instruction

## Syntax

```
ACCEPT INPUT
```

## Usage

The `ACCEPT INPUT` instruction validates and exits the `INPUT
ARRAY` instruction, if no error is raised.

The [`AFTER FIELD`](1944-after-field-block.md), [`ON CHANGE`](1943-on-change-block.md), [`AFTER INSERT`](2021-after-insert-block.md), [`AFTER ROW`](1976-after-row-block.md), etc., control blocks will be
executed.

Statements after the `ACCEPT INPUT` will be skipped.

Input field validation is a process that does several successive validation tasks, as listed
here:

1. The current field value is checked, the check is based on the program variable data type (for
   example, the user must input a valid date in a DATE field).
2. `NOT NULL` field attributes
   are checked for all input fields. This attribute forces the field to have a value set by program or
   entered by the user. If the field contains no value, the constraint is not satisfied. Input values
   are right-trimmed, so if the user inputs only spaces, this corresponds to a NULL value which does
   not fulfill the NOT NULL constraint.
3. `INCLUDE` field attributes
   are checked for all input fields. This attribute forces the field to contain a value that is listed
   in the include list. If the field contains a value that is not in the list, the constraint is not
   satisfied.
4. `REQUIRED` field attributes
   are checked for all input fields. This attribute forces the field to have a default value, or to be
   "touched" by the user or by program. If the field was not edited during the dialog, the constraint
   is not satisfied.

If a field does not satisfy one of these constraints, dialog termination is canceled, an error
message is displayed, and the focus goes to the first field causing a problem.

The `ACCEPT INPUT` instruction can only be used in a singular `INPUT
ARRAY` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog
block.
