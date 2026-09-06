---
title: "Declarative dialog programming steps"
source: "fgl-topics/t_fgl_declarative_dialogs.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Declarative dialog programming steps"
type: "task"
---

# Declarative dialog programming steps

> This procedure describes how to implement a declarative DIALOG block.

To implement a declarative `DIALOG` block:

1. Create a form specification file containing [screen record(s) and/or screen array(s)](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."). The
   screen records and screen arrays identify the presentation elements to be used by the runtime system
   to display the data models (the content of program variables bound to the `DIALOG`
   blocks).
2. Create a dedicated .4gl module to implement the declarative
   `DIALOG` block.
3. With the `TYPE` instruction, declare a user-defined type ([records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") or [arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")).
   for variables that will be used as data models by the sub-dialog. These variables can be defined
   locally in the module implementing the sub-dialog, or they can be defined in the module implementing
   the parent `DIALOG` block, and passed as parameter to the sub-dialog with the
   `SUBDIALOG` instruction.
4. Define the [declarative `DIALOG`
   block](2152-syntax-of-the-declarative-dialog-block.md "The declarative DIALOG block defines an interactive instruction that can be used by a parent DIALOG using the SUBDIALOG clause.") in the module, to handle interaction. Define a sub-dialog with local or parameter
   variables to be used as data models. The sub-dialog will define how variables will be used (display
   or input). 
   1. Inside the sub-dialog instruction, define the behavior with control blocks such as
      `BEFORE ROW`, `AFTER ROW`, `BEFORE FIELD`, and
      interaction blocks such as `ON ACTION`.
5. Reference the declarative dialog with the `SUBDIALOG` clause in a [procedural `DIALOG`
   block](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.").

## Related links

**Related concepts**  

[Using declarative dialogs](2154-using-declarative-dialogs.md "Dialog coding concepts, configuration and code structure.")
