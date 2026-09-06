---
title: "Form field specification in CONSTRUCT"
source: "fgl-topics/c_fgl_Construct_019.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > Form field specification in CONSTRUCT"
type: "concept"
description: "In order to produce an SQL condition, the CONSTRUCT instruction uses a list of database columns that must match form fields for user input. In ATTRIBUTES section of the form file, fields are typically ..."
---

# Form field specification in CONSTRUCT

In order to produce an SQL condition, the `CONSTRUCT` instruction uses a list of
database columns that must match form fields for user input.

In `ATTRIBUTES` section of the form file, fields are typically defined from a
database schema (with the form [tabname.colname](1671-database-column-fields.md "Form fields defined with a table and column name get data type from the database schema file.")), to get the corresponding data type; however, it is also
possible to use [`FORMONLY.field-name`](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.") specifications with the `TYPE
data-type` clause.

Unlike `INPUT`, `DISPLAY ARRAY` and `INPUT ARRAY`,
the `CONSTRUCT` dialog does not use a program variable for each form field.

Only one string variable is required, to hold the SQL condition. Individual field criteria is
available in the input buffers ([`GET_FLDBUF()`](../08_language-basics/0671-get-fldbuf-function.md "The GET_FLDBUF() operator returns as character strings the current values of the specified fields."), [`DIALOG.getFieldBuffer()`](../15_library-reference/3197-ui-dialog-getfieldbuffer.md "Returns the input buffer of the specified field.")).

The list of database columns specified in the `CONSTRUCT` statement will appear
in the SQL condition produced.

## Binding columns and fields by name

The `CONSTRUCT BY NAME variable ON column-list` syntax maps the
field names to database column names by name. Form fields are typically defined in the form by
following a database schema, specifying the column name and data type.

```
SCHEMA stock
DEFINE where_part STRING
...
CONSTRUCT BY NAME where_part ON cust_name, cust_address
  ...
END CONSTRUCT
```

## Binding columns and fields by position

The `CONSTRUCT variable ON column-list FROM field-list`
clause explicitly maps database columns to form fields by position. The form can include other
fields that are not part of the specified column list, but the number of variables or record
members must equal the number of form fields listed in the `FROM` clause. Each
database column must be of the same (or a compatible) data type as the corresponding form
field. When the user enters data, the runtime system checks the entered value against the data
type of the form field.

```
DEFINE where_part STRING
...
CONSTRUCT where_part ON cust_name, cust_address
                     FROM field_02, field_04
  ...
END CONSTRUCT
```

## Related links

**Related concepts**  

[Binding variables to form fields](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")
