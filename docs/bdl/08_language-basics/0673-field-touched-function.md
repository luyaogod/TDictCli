---
title: "FIELD_TOUCHED() [function]"
source: "fgl-topics/c_fgl_operators_FIELD_TOUCHED.html"
breadcrumb: "Language basics > Operators > List of expression elements > Dialog handling operators > FIELD_TOUCHED() [function]"
type: "concept"
---

# FIELD_TOUCHED() [function]

> The FIELD_TOUCHED() operator checks if fields were modified during the dialog execution.

## Syntax

```
FIELD_TOUCHED (
   { [group.]field.*
   | group.*
   | *
   } [,...] )
```

1. group can be a table name, a screen record, a screen array
   or `FORMONLY` as defined in the form.
2. field is the name of the field as defined in the form.

## Usage

`FIELD_TOUCHED()` returns `TRUE` if the value of a screen field (or
multiple fields) has changed since the beginning of the interactive instruction.

When used in an `INPUT ARRAY` instruction, the runtime system assumes that you are
referring to the current row.

Do not confuse the `FIELD_TOUCHED()` operator with the [`fgl_buffertouched()`](../15_library-reference/2736-fgl-buffertouched.md "Returns TRUE if the input buffer was modified in the current field.")
built-in function; which checks a different field modification flag, that is reset when entering the
field.

The operator accepts a list of explicit field names, and supports the
`[group.]*` notation in order to check multiple fields in a single
evaluation. When passing a simple asterisk (`*`) to the operator, the runtime system
will check all fields used by the current dialog.

Without group prefix, the operator will try to find a matching field name, by
ignoring the prefix used in the form file or in the `FROM` clause of the dialog
instruction. Inside a `DIALOG` block, if several sub-dialogs use the same field name,
the operator will use the field in the current sub-dialog.

When used, the group prefix is a table name, screen record, screen array or
the `FORMONLY` keyword as defined in the form file. The group
prefix to be specified depends on the type of variable-to-field binding used by the dialog instruction:

- When using the `BY NAME` clause as in [`INPUT var-list BY NAME`](../11_user-interface/1932-syntax-of-the-input-instruction.md "The INPUT statement supports data entry in fields of the current form.") or [`CONTRUCT BY NAME sql-cond ON
  column-list`](../11_user-interface/2048-syntax-of-construct-instruction.md "The CONSTRUCT instruction provides database query by example, producing a WHERE condition for SELECT."), the field prefix must match the table name specified
  in the form definition file. If no field is found, a second search is done by ignoring the prefix
  and takes the first field of the form with this name.
- When using the `FROM` clause as in [`INPUT var-list FROM field-list`](../11_user-interface/1932-syntax-of-the-input-instruction.md "The INPUT statement supports data entry in fields of the current form."), [`CONTRUCT sql-cond ON
  column-list FROM field-list`](../11_user-interface/2048-syntax-of-construct-instruction.md "The CONSTRUCT instruction provides database query by example, producing a WHERE condition for SELECT.") or [`INPUT ARRAY arr-name FROM
  scr-array.*`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."), the field prefix must match the table name, screen
  record or screen array used in the `FROM` clause.

The `FIELD_TOUCHED()` operator can only be used inside an `INPUT`,
`INPUT ARRAY` and `CONSTRUCT` interaction block.

For more details about the `FIELD_TOUCHED()` operator usage and the understand the
"touched flag" concept, see [Input field modification flag](../11_user-interface/2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.").

## Example

```
INPUT ...
  ...
  AFTER FIELD custname 
    IF FIELD_TOUCHED( customer.custname ) THEN
       MESSAGE "Customer name was changed."
    END IF
  ...
  AFTER INPUT
    IF FIELD_TOUCHED( customer.* ) THEN
       MESSAGE "Customer record was changed."
    END IF
  ...
```

## Related links

**Related concepts**  

[ui.Dialog.getFieldTouched](../15_library-reference/3198-ui-dialog-getfieldtouched.md "Returns the modification flag for a field.")

[ui.Dialog.setFieldTouched](../15_library-reference/3227-ui-dialog-setfieldtouched.md "Sets the modification flag of the specified field.")
