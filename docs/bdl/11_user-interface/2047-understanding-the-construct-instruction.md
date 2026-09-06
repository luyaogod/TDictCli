---
title: "Understanding the CONSTRUCT instruction"
source: "fgl-topics/c_fgl_Construct_003.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Understanding the CONSTRUCT instruction"
type: "concept"
---

# Understanding the CONSTRUCT instruction

> The CONSTRUCT instruction provides database query, by entering search filters directly in form fields.

Query by example enables a user to query a database by specifying values (or ranges of values)
for [screen fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.") that correspond to the
database columns.

The runtime system converts the query values entered by the user into a boolean SQL condition
that can be used in the `WHERE` clause of a prepared `SELECT`
statement.

The `CONSTRUCT` statement produces an SQL condition corresponding to all search
criteria that a user specifies in the fields. The instruction fills a character variable with that
SQL condition, and you can use the content of this variable to create the `WHERE`
clause of a `SELECT` statement. The `SELECT` statement must be
executed with the dynamic SQL management instructions `PREPARE` or `DECLARE
ident CURSOR FROM sqltext`:

The `CONSTRUCT` instruction uses the data types of the form field to verify user
input and to produce the SQL condition.

The SQL condition is generated based on the current database session, which defines the type of
database server. Therefore, the program must be connected to a database server before entering the
`CONSTRUCT` block. The generated SQL condition is specific to the database server and
may not be used with other types of database servers.

If no criteria were entered, the string `'1=1'` is assigned to the string variable.
This is a boolean SQL expression that always evaluates to true so that all rows are returned.

The `CONSTRUCT` dialog activates the current form. This is the form most recently
displayed or, if you are using more than one window, the form currently displayed in the current
window. When the `CONSTRUCT` statement completes execution, the form is cleared and
deactivated.

During a `CONSTRUCT` instruction, edit field input is left-aligned, independently
to the form field data type: During an `INPUT`, numeric fields are right-aligned,
during a `CONSTRUCT`, they are left-aligned.

By default the screen field tabbing order is defined by the order of the field names in the
`FROM` clause; by default this is the list of column names in the `ON`
clause when no `FROM` clause is specified. If needed, change the field tabbing order
with the `FIELD ORDER FORM` option and `TABINDEX` field attributes.
For more details, see [Defining the tabbing order](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.").

When the user moves from field to field or changes values, [dialog control blocks](2055-construct-control-blocks.md) such as `BEFORE FIELD` are executed.

When the user clicks on an action view (button), or when an asynchronous event occurs, [dialog interaction blocks](2062-construct-interaction-blocks.md) like `ON ACTION` are
executed.

The code inside a `CONSTRUCT` dialog can use [control instructions](2067-construct-control-instructions.md), [dialog control
functions](2223-dialog-control-functions.md "The language provides several built-in functions and operators to be used in a dialog instruction."), and the [`ui.Dialog`](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.") class, to implement the dialog behavior.

## Related links

**Related concepts**  

[Result set processing](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query.")

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
