---
title: "The CONSTRUCT sub-dialog"
source: "fgl-topics/c_fgl_DIALOG_subdialog_CONSTRUCT.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Structure of a procedural DIALOG block > The CONSTRUCT sub-dialog"
type: "concept"
---

# The CONSTRUCT sub-dialog

> The CONSTRUCT sub-dialog provides database query by example feature, converting search criteria entered by the user into an SQL WHERE condition that can be used to execute a SELECT statement.

## Defining query by example fields

The `CONSTRUCT` sub-dialog requires a character string variable to hold the
`WHERE` clause, and a list of [screen fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.") where the user can enter search criteria.

```
PRIVATE DEFINE sql_condition STRING
  ...
DIALOG cust_query()
  CONSTRUCT BY NAME sql_condition
      ON customer.cust_name, customer.cust_address 
    BEFORE FIELD cust_name
    ...
  END CONSTRUCT
  ...
END DIALOG
```

Make sure the character string variable is large enough to store all possible SQL conditions.
It is better to use a [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") data
type to avoid any size problems.

`CONSTRUCT` uses the field data types defined
in the current form file to produce the SQL conditions. This is different from other interactive
instructions, where the data types of the program variables define the way to handle input/display.
It is *strongly* recommended (but not mandatory) that the form field data types correspond to
the data types of the program variables used for input. This is implicit if both form fields and
program variables are based on the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").

The `CONSTRUCT` clause can be used in two forms:

1. `CONSTRUCT BY NAME string-variable ON
   column-list`
2. `CONSTRUCT string-variable ON column-list FROM
   field-list`

The `BY NAME` clause implicitly binds the form fields to the columns, where the
form field identifiers match the column names specified in the column-list after the
`ON` keyword. You can specify the individual column names (separated by commas)
or use the `tablename.*` shortcut to include all columns defined for a table in
the database schema file.

The `FROM` clause explicitly binds the form fields listed after the
`FROM` keyword with the column definitions listed after the `ON`
keyword.

In both cases, the name of the columns in *column-list* will be used to produce the SQL
condition in *string-variable*.

## Identifying a CONSTRUCT sub-dialog

The name of a `CONSTRUCT` sub-dialog can be used to qualify [sub-dialog actions](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.") with a prefix. In order to
identify the `CONSTRUCT` sub-dialog with a specific name, use the
`ATTRIBUTES` clause to set the `NAME` attribute:

```
CONSTRUCT BY NAME sql_condition ON customer.*
  ATTRIBUTES (NAME = "q_cust")
  ...
```

## Control blocks in CONSTRUCT

A
Query By Example declared with the `CONSTRUCT` clause
can raise the following triggers:

- [BEFORE CONSTRUCT](2057-before-construct-block.md)
- [BEFORE FIELD](1942-before-field-block.md)
- [AFTER FIELD](1944-after-field-block.md)
- [AFTER
  CONSTRUCT](2058-after-construct-block.md)

In the singular `CONSTRUCT` instruction, `BEFORE CONSTRUCT` and
`AFTER CONSTRUCT` blocks are typically used as initialization and
finalization blocks. In `DIALOG` block, `BEFORE CONSTRUCT` and
`AFTER CONSTRUCT` blocks will be executed each time the focus goes to
(`BEFORE`) or leaves (`AFTER`) the group of fields defined
by this sub-dialog.

## Related links

**Related concepts**  

[Query operators in CONSTRUCT](2052-query-operators-in-construct.md "Query operators in CONSTRUCT")

[CONSTRUCT ATTRIBUTES clause](2093-construct-attributes-clause.md "CONSTRUCT specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")
