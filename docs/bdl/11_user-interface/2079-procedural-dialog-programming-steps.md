---
title: "Procedural dialog programming steps"
source: "fgl-topics/t_fgl_multiple_dialogs_procedural_dialogs.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Procedural dialog programming steps"
type: "task"
---

# Procedural dialog programming steps

> Follow this procedure to use the DIAOG instruction.

To implement a procedural `DIALOG` block:

1. Create a form specification file containing [screen record(s) and/or screen array(s)](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."). The
   screen records and screen arrays identify the presentation elements to be used by the runtime system
   to display the data models (the content of program variables bound to the `DIALOG`
   blocks).
2. With the `DEFINE` instruction, declare program variables ([records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") and [arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."))
   that will be used as data models. For record lists (`DISPLAY ARRAY` or `INPUT
   ARRAY`), the members of the program array must correspond to the elements of the screen
   array, by number and data types. To handle record lists, use dynamic arrays instead of static
   arrays.
3. Open and display the form, using `OPEN
   WINDOW WITH FORM` or the `OPEN FORM` / `DISPLAY
   FORM` instructions.
4. Fill the program variables (the model) with data. For lists, you typically use a [result set cursor](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query.").
5. Implement the [`DIALOG` instruction
   block](2078-syntax-of-the-procedural-dialog-instruction.md "The DIALOG block is an interactive instruction that executes several sub-dialogs simultaneously.") to handle interaction. Define each sub-dialog with program variables to be used as data
   models. The sub-dialogs will define how variables will be used (display or input).
   1. Inside each sub-dialog instruction, define the behavior with control blocks such as
      `BEFORE DIALOG`, `AFTER ROW`, `BEFORE FIELD`, and
      interaction blocks such as `ON ACTION`.
   2. To end the `DIALOG` instruction, implement an `ON ACTION close`
      or `ON ACTION accept` / `ON ACTION cancel` to handle dialog validation
      and cancellation, with the `ACCEPT DIALOG` and `EXIT DIALOG` control
      instructions. The `int_flag` variable will not be set as in singular
      dialogs.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")

[Using multiple dialogs](2080-using-multiple-dialogs.md "Dialog coding concepts, configuration and code structure.")
