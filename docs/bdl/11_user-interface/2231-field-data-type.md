---
title: "Field data type"
source: "fgl-topics/c_fgl_prog_dialogs_field_data_type.html"
breadcrumb: "User interface > User interface programming > Input fields > Field data type"
type: "concept"
---

# Field data type

> Depending on the type of dialog, the field data type is defined by program variables or form specification file.

## Field data types

The field data type defines how a user can input values into a form field. For
example, when defining a field to hold date values, it should only allow date value input.
Therefore,

The data type used by the runtime to control field display and input depends on the type of the
dialog:

- For [`INPUT`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form."), [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") and [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions."), the data type is defined
  by the program variable bound to the field.
- For [`CONSTRUCT`](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form."), the data type is
  defined in the (.per) [form definition
  file](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms."), either by the `TYPE` attribute of a [FORMONLY field](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.") for example, or with the
  corresponding [database column](1671-database-column-fields.md "Form fields defined with a table and column name get data type from the database schema file.") in
  a schema file.

## Field validation rules

Data validation rules can be defined at the form level, such as `NOT NULL`,
`REQUIRED` and
`INCLUDE` attributes.

Data validation constraints are checked when leaving a field, or when the dialog is validated
(for example, with the `ACCEPT DIALOG` instruction inside a `DIALOG`
multiple dialog block).

Trailing blanks entered by the user will be removed when leaving the input field.

## Related links

**Related concepts**  

[Binding variables to form fields](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")

[Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.")

[Field configuration and decoration](2234-field-configuration-and-decoration.md "Form fields can be customized with specific decoration and settings.")
