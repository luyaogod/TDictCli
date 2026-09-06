---
title: "Understanding multiple dialogs"
source: "fgl-topics/c_fgl_multiple_dialogs_intro.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Understanding multiple dialogs"
type: "concept"
---

# Understanding multiple dialogs

> Multiple dialogs are defined with DIALOG blocks inside a FUNCTION.

The concept of multiple dialogs refers to the usage of a procedural
`DIALOG` block, to control several elements of a form. During the execution
of a procedural dialog, no other window/form can be accessed: multiple dialogs are in the
category of [modal dialogs](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.").

The `DIALOG` procedural instruction handles different parts of a form
simultaneously, including simple display fields, simple input fields, read-only list of
records, editable list of records, query by example fields, and action views. The
`DIALOG` instruction acts as a collection of singular dialogs working in
parallel.

"Singular interactive instructions" refer to `INPUT`, `CONSTRUCT`,
`DISPLAY ARRAY` and `INPUT ARRAY` independent blocks not surrounded
by the `DIALOG` / `END DIALOG` keywords.
While the `DIALOG` instruction reuses some of the semantics and behaviors of
singular interactive instructions, there are some differences.

Like the singular interactive instructions, `DIALOG` is an interactive instruction.
You can execute a `DIALOG` instruction from one of the singular dialogs, or execute
a singular dialog from a `DIALOG` block. The parent dialog will be disabled until the
child dialog returns.

A `DIALOG` procedural instruction consist of several sub-dialog blocks declared
inside the `DIALOG` instruction, or external dialog blocks declared in scope
outside of the current function. The external dialogs are attached to the current dialog with
the `SUBDIALOG`
clause.

The `DIALOG` instruction binds program variables (such as simple records or
arrays of records) with a screen-record or screen-array defined in a form, allowing the user to
view and update application data.

When a `DIALOG` block executes, it activates the current form (the form most
recently displayed or the form in the current window). When the statement completes execution,
the form is deactivated.

The syntax of the `DIALOG` instruction is very close to singular dialogs, using
common triggers such as `BEFORE FIELD`, `ON ACTION`, and so on.
Despite the similarities, the behavior and semantics of `DIALOG` are a bit different
from singular dialogs.

Understand that the `DIALOG` instruction is not provided to replace singular
dialogs. Singular dialogs are still supported. It is recommended that you use singular dialogs
if no multiple dialog is required.

Unlike singular dialogs, the `DIALOG` instruction does not use the
`int_flag` variable. You must implement `ON ACTION accept` or
`ON ACTION cancel` to handle dialog validation or cancellation. These actions
do not exist by default in `DIALOG`.

Unlike singular dialogs creating automatic accept and cancel actions, by default there is no way
to quit the `DIALOG` instruction. You must implement your own action handler and
execute `EXIT DIALOG` or `ACCEPT DIALOG`.

A good practice is to write a setup dialog function to centralize all field and action
activations for a specific context. Call that setup function at any place in the
`DIALOG` code where the field and action activation rules may change.

While static arrays are supported by the `DIALOG` instruction, it is strongly
recommended that you use dynamic arrays instead. With a dynamic array, the actual number of rows
is automatically defined by the array variable, while static arrays need an additional step to
define the total number of rows.

When needed, use the `UNBUFFERED` mode with multiple dialogs to force model/view
synchronization, and use the `FIELD ORDER FORM` option to follow the
`TABINDEX` definitions in the form file.

This example is of a `DIALOG` procedural instruction that includes both an
`INPUT` and a `DISPLAY ARRAY` sub-dialog, plus a sub-dialog
defined externally and included with the `SUBDIALOG`
keyword:

```
SCHEMA stores 
DEFINE p_customer RECORD LIKE customer.*
DEFINE p_orders DYNAMIC ARRAY OF RECORD LIKE orders.*
FUNCTION customer_dialog()
  DIALOG ATTRIBUTES(UNBUFFERED, FIELD ORDER FORM)
    INPUT BY NAME p_customer.*
      AFTER FIELD cust_name 
        CALL setup_dialog(DIALOG)
    END INPUT
    DISPLAY ARRAY p_orders TO s_orders.*
      BEFORE ROW
        CALL setup_dialog(DIALOG)
    END DISPLAY
    SUBDIALOG common_footer
    ON ACTION close 
      EXIT DIALOG
  END DIALOG
END FUNCTION
```

All elements of the dialog are active at the same time, so you must handle tabbing order
properly. By default - as in singular dialogs - the tabbing order is driven by the binding list
(order of program variables). It is strongly recommended that you use the `FIELD
ORDER FORM` option and the `TABINDEX` field attributes instead.

Like the singular `INPUT ARRAY` instruction, `DIALOG` creates
automatic insert, append, and delete actions. These actions are only active when the focus is in the
list.

When the user moves from field to field, changes values, or browses a list, [dialog control blocks](2099-dialog-control-blocks.md "Dialog control blocks are predefined dialog triggers where you can implement specific code to control the interactive instruction.") such as `BEFORE
FIELD`, `BEFORE ROW`, `BEFORE INPUT`, `BEFORE
DISPLAY` are executed.

When the user clicks on an action view (button), or when an asynchronous event occurs, [dialog interaction blocks](2119-dialog-interaction-blocks.md "Dialog interaction blocks are dialog triggers that can be used to execute specific code when the user executes an action in the dialog. For example, when pressing a button in the form, the corresponding ON ACTION interaction block will be executed.") like `ON
ACTION` are executed.

The code inside a `DIALOG / END DIALOG` dialog can use [control instructions](2136-dialog-control-instructions.md "Dialog control instructions are language instructions dedicated to dialog control, to programmatically force the dialog to behave in a given way."), [dialog control functions](2223-dialog-control-functions.md "The language provides several built-in functions and operators to be used in a dialog instruction."), and the [`ui.Dialog`](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.") class, to implement
the dialog behavior.

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")

[Declarative dialogs (DIALOG - at module level)](2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs.")

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
