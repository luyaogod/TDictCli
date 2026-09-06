---
title: "INFIELD() [function]"
source: "fgl-topics/c_fgl_operators_INFIELD.html"
breadcrumb: "Language basics > Operators > List of expression elements > Dialog handling operators > INFIELD() [function]"
type: "concept"
---

# INFIELD() [function]

> The INFIELD() operator checks for the current screen field.

## Syntax

```
INFIELD ( [group.]field )
```

1. group can be a table name, a screen record, a screen array
   or `FORMONLY` as defined in the form.
2. field is the name of the field as defined in the form.

## Usage

`INFIELD()` checks for the current field in a `CONSTRUCT`,
`INPUT` or `INPUT ARRAY` dialog.

When used in an `INPUT ARRAY` instruction, the runtime system assumes
that you are referring to the current row.

For a generic equivalent, use the [`DIALOG.getCurrentItem()`](../15_library-reference/3194-ui-dialog-getcurrentitem.md "Returns the current item having focus.") method.

When using
`INFIELD(field)` without group prefix, the
operator will check if the current field name matches, by ignoring the prefix used in the form file
or in the `FROM` clause of the dialog instruction.

When using `INFIELD(group.field)`, the
group prefix must be a table name, screen record, screen array or the
`FORMONLY` keyword as defined in the form, and its usage depends on the type of
variable-to-field binding used by the dialog instruction:

- When using the `BY NAME` clause as in [`INPUT var-list BY NAME`](../11_user-interface/1932-syntax-of-the-input-instruction.md "The INPUT statement supports data entry in fields of the current form.") or [`CONTRUCT BY NAME sql-cond ON
  column-list`](../11_user-interface/2048-syntax-of-construct-instruction.md "The CONSTRUCT instruction provides database query by example, producing a WHERE condition for SELECT."), no field prefix can be used. Even if the field name
  matches, `INFIELD(group.field)` will return
  `FALSE`.
- When using the `FROM` clause as in [`INPUT var-list FROM field-list`](../11_user-interface/1932-syntax-of-the-input-instruction.md "The INPUT statement supports data entry in fields of the current form."), [`CONTRUCT sql-cond ON
  column-list FROM field-list`](../11_user-interface/2048-syntax-of-construct-instruction.md "The CONSTRUCT instruction provides database query by example, producing a WHERE condition for SELECT.") or [`INPUT ARRAY arr-name FROM
  scr-array.*`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."), the field prefix must match the table name, screen
  record or screen array used in the `FROM` clause.

## Example

```
INPUT BY NAME rec.*
  ...
  ON ACTION check
     IF INFIELD( custname ) THEN
        CALL check_customer_name( rec.custname )
  ...
```

## Related links

**Related concepts**  

[Which form item has the focus?](../11_user-interface/2244-which-form-item-has-the-focus.md "Identify what element of the current form has the focus.")

[Query by example (CONSTRUCT)](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")

[Record input (INPUT)](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.")

[Editable record list (INPUT ARRAY)](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")
