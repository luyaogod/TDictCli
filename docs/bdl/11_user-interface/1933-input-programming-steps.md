---
title: "INPUT programming steps"
source: "fgl-topics/t_fgl_record_input_005.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > INPUT programming steps"
type: "task"
---

# INPUT programming steps

> Follow this procedure to use the INPUT dialog instruction.

To implement the `INPUT` statement:

1. Create a form specification file, with an optional `screen record`.

   The screen record identifies the presentation elements to be used by the runtime system to
   display the records. If you omit the declaration of the screen record in the form file, the runtime
   system will use the default screen records created by the form compiler for each table listed in the
   `TABLES` section and for the `FORMONLY` pseudo-table.
2. Make sure that the program controls interruption handling with `DEFER INTERRUPT`, to manage the validation/cancellation
   of the interactive dialog.
3. Define a program [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") with the
   `DEFINE` instruction.

   The members of the program record must correspond to the elements of the screen record, by
   number and data types.
4. Open and display the form, using `OPEN
   WINDOW WITH FORM` or the `OPEN FORM` / `DISPLAY
   FORM` instructions.
5. If needed, fill the program record with data, for example, with a [result set cursor](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query.").
6. Set the `int_flag`
   variable to `FALSE`.
7. Implement the `INPUT dialog
   block` to handle data input. Set the [`WITHOUT DEFAULTS`](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.") option appropriately.
8. Inside the `INPUT` statement, control the behavior of the instruction
   with `BEFORE INPUT`, `BEFORE FIELD`, `AFTER FIELD`,
   `AFTER INPUT` and `ON ACTION` blocks.
9. After the interaction statement block, test the `int_flag` predefined
   variable to check if the dialog was canceled (`int_flag=TRUE`) or validated
   (`int_flag=FALSE`).

   If the `int_flag` variable is `TRUE`, you should reset it to
   `FALSE` so as not to disturb code that relies on this variable to detect
   interruption events from the GUI front-end or TUI console.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
