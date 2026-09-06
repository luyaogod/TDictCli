---
title: "INPUT ARRAY programming steps"
source: "fgl-topics/t_fgl_InputArray_007.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > INPUT ARRAY programming steps"
type: "task"
---

# INPUT ARRAY programming steps

> Follow this procedure to use the INPUT ARRAY dialog instruction.

To implement an `INPUT ARRAY` statement:

1. Create a form specification file containing a screen array.
   The screen array identifies the presentation elements to be used by
   the runtime system to display the rows.
2. Make sure that the program controls interruption handling with `DEFER INTERRUPT`, to manage the
   validation/cancellation of the interactive dialog.
3. Define an array of records with the `DEFINE` instruction.
   The members of the program array must correspond to the elements of
   the screen array, by number and data types. If you want to input data
   from a reduced set of columns, you must define a second screen array,
   containing the limited list of form fields, in the form file. You
   can then use the second screen array in an `INPUT ARRAY a FROM
   sa.*` instruction.
4. Open and display the form, using an `OPEN
   WINDOW WITH FORM` or the `OPEN FORM` / `DISPLAY
   FORM` instructions.
5. If needed, fill the program array with data, for example
   with a result set cursor, counting the number of program records being
   filled with retrieved data.
6. Set the `int_flag`
   variable to `FALSE`.
7. Write the [`INPUT ARRAY` dialog
   block](2007-syntax-of-input-array-instruction.md "The INPUT ARRAY supports data entry by users into a screen array and stores the entered data in an array of records.") to handle data input. Set the [`WITHOUT DEFAULTS`](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.") option appropriately.
8. Inside the `INPUT ARRAY` dialog, control the behavior of the instruction with
   control blocks such as `BEFORE INPUT`, `BEFORE INSERT`, `BEFORE
   DELETE`, `BEFORE ROW`, `BEFORE FIELD`,`AFTER
   INSERT`,`AFTER DELETE`,`AFTER FIELD`,`AFTER
   ROW`, `AFTER INPUT` and `ON ACTION` blocks.
9. Get the new number of rows with the `ARR_COUNT()` built-in
   function or with `DIALOG.getArrayLength()`.
10. After the interaction statement block, test the `int_flag` predefined
    variable to check if the dialog was canceled (`int_flag=TRUE`)
    or validated (`int_flag=FALSE`). If the `int_flag` variable
    is `TRUE`, it is recommended that you reset it to `FALSE` to
    not disturb code that relies on this variable to detect interruption
    events from the GUI front-end or TUI console.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
