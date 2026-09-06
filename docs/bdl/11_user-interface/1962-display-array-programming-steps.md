---
title: "DISPLAY ARRAY programming steps"
source: "fgl-topics/t_fgl_DisplayArray_014.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > DISPLAY ARRAY programming steps"
type: "task"
---

# DISPLAY ARRAY programming steps

> Follow this procedure to use the DISPLAY ARRAY dialog instruction.

To implement a `DISPLAY ARRAY` statement:

1. Create a form specification file containing a [screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."). The screen array identifies the
   presentation elements to be used by the runtime system to display the rows.
2. Make sure that the program controls interruption handling with [`DEFER INTERRUPT`](../09_advanced-features/0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received."), to manage the
   validation/cancellation of the interactive dialog.
3. Define an [array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") of [records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") with the `DEFINE` instruction. The
   members of the program array must correspond to the elements of the screen array, by number and data
   types. Static or a dynamic arrays can be used for the full list mode, but the paged mode requires a
   dynamic array. For new developments, use dynamic arrays in both cases.
4. Open and display the form, using `OPEN
   WINDOW WITH FORM` or the `OPEN
   FORM` / `DISPLAY
   FORM` instructions.
5. If you want to use the full list mode, fill the program
   array with data, typically with a result set cursor, counting the
   number of program records being filled with retrieved data.
6. Set the `int_flag` variable to
   `FALSE`.
7. Implement the [`DISPLAY ARRAY` dialog
   block](1961-syntax-of-display-array-instruction.md "The DISPLAY ARRAY instruction controls the display of a program array on the screen."). When using a static array, specify the number of rows with the `COUNT`
   attribute in the `ATTRIBUTES` clause, or call the `SET_COUNT()`
   function before the dialog block. With dynamic arrays, the number of rows is automatically known by
   the dialog. Consider using the `UNBUFFERED` mode in new developments.
8. If you want to use the [paged mode](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY."),
   define the total number of rows with the `COUNT` attribute (can be -1 for infinite
   number of rows), and add the `ON FILL BUFFER` clause that will contain the code to
   fill the dynamic array with the expected rows from `fgl_dialog_getBufferStart()` to
   `fgl_dialog_getBufferLength()`.
9. If [multi-row selection](2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.") is needed, call
   the `ui.Dialog.setSelectionMode()` method in `BEFORE DISPLAY` to
   enable this mode.
10. Inside the `DISPLAY ARRAY` block, control the behavior of the instruction with
    `BEFORE ROW`, `AFTER ROW` and `ON ACTION` blocks.
11. After the interaction statement block, test the `int_flag` predefined variable
    to check if the dialog was canceled (`int_flag=TRUE`) or validated
    (`int_flag=FALSE`). If the `int_flag` variable is
    `TRUE`, reset it to `FALSE` in order not to disturb code that relies
    on this variable to detect interruption events from the GUI front-end or TUI console.
12. If needed, get the current row with the [`ARR_CURR()`](../15_library-reference/2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.") built-in function after dialog execution. During dialog
    execution, you can also use `DIALOG.getCurrentRow()`.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
