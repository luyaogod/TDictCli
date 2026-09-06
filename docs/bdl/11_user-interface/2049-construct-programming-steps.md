---
title: "CONSTRUCT programming steps"
source: "fgl-topics/t_fgl_Construct_007.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > CONSTRUCT programming steps"
type: "task"
---

# CONSTRUCT programming steps

> Follow this procedure to use the CONSTRUCT dialog instruction.

To implement a `CONSTRUCT` statement:

1. Declare a variable with the `DEFINE` statement, it can be `CHAR`,
   `VARCHAR` or `STRING`.
   `STRING` is preferred in order to avoid any size limitation.
2. Open and display the form, using an `OPEN
   WINDOW WITH FORM` or an `OPEN FORM` / `DISPLAY FORM`
   instruction.
3. Set the `int_flag` variable to
   `FALSE`.
4. Define the `CONSTRUCT` block with
   the list of form fields to be used for the query by example. If needed, define dialog control blocks
   to implement rules for the query by example.
5. Inside the `CONSTRUCT` statement, control the behavior of the instruction with
   `BEFORE CONSTRUCT`, `BEFORE FIELD`, `AFTER FIELD`, `AFTER CONSTRUCT` and `ON ACTION` blocks.
6. After the interaction statement block, test the `int_flag` predefined variable
   to check if the dialog was canceled (`int_flag=TRUE`) or validated
   (`int_flag=FALSE`).

   If the `int_flag` variable is `TRUE`, you should reset it to
   `FALSE` to not disturb code that relies on this variable to detect interruption
   events from the GUI front-end or TUI console.
7. To build the complete SQL statement, concatenate "`SELECT
   ... WHERE`" to the string variable that contains the
   boolean SQL expression produced by `CONSTRUCT`.
8. Define a [database cursor](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query.") with the
   `DECLARE FROM` instruction, by using the `SELECT` statement.
9. Execute the cursor and fetch the rows found by the database server. You can for example
   implement a `FOREACH` loop to fill a program array, to be shown by a `DISPLAY ARRAY` statement.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
