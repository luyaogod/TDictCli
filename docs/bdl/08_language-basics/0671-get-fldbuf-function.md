---
title: "GET_FLDBUF() [function]"
source: "fgl-topics/c_fgl_operators_GET_FLDBUF.html"
breadcrumb: "Language basics > Operators > List of expression elements > Dialog handling operators > GET_FLDBUF() [function]"
type: "concept"
---

# GET_FLDBUF() [function]

> The GET_FLDBUF() operator returns as character strings the current values of the specified fields.

## Syntax

```
GET_FLDBUF ( [group.]field  [,...] )
```

1. group can be a table name, a screen record, a screen array
   or `FORMONLY` as defined in the form.
2. field is the name of the field as defined in the form.

## Usage

The `GET_FLDBUF()` operator is used to get the value of a screen field before the
input buffer is copied into the associated variable.

Use of the `GET_FLDBUF()` operator is recommended only in dialogs allowing field
input (`INPUT`, `INPUT ARRAY`, `CONSTRUCT`). The
behavior is undefined when used in `DISPLAY ARRAY`.

The `GET_FLDBUF()` operator takes the field names as identifiers, not as string
expressions:

```
LET v = GET_FLDBUF( customer.custname )
```

If multiple fields are specified between parentheses, use the `RETURNING`
clause:

```
CALL GET_FLDBUF( customer.* ) RETURNING rec_customer.*
```

When used in a `INPUT ARRAY` instruction, the runtime system assumes that you are
referring to the current row.

The values returned by this operator are context dependent; it must be used carefully. If
possible, use the variable associated to the input field instead.

When using the `UNBUFFERED` mode, program variables are automatically assigned,
and the `GET_FLDBUF()` operator is not required in most cases.

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

## Example

```
INPUT BY NAME ...
   ...
   ON KEY(CONTROL-Z)
      LET v = GET_FLDBUF( customer.custname )
      IF check_synonyms(v) THEN
         ...
```

## Related links

**Related concepts**  

[The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")

[Editable record list (INPUT ARRAY)](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")

[ui.Dialog.getFieldBuffer](../15_library-reference/3197-ui-dialog-getfieldbuffer.md "Returns the input buffer of the specified field.")
