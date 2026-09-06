---
title: "The SUBDIALOG clause"
source: "fgl-topics/c_fgl_multiple_dialogs_SUBDIALOG.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Structure of a procedural DIALOG block > The SUBDIALOG clause"
type: "concept"
description: "Purpose of SUBDIALOG The SUBDIALOG clause defines a declarative dialog to be attached to the current procedural DIALOG block . By using form inclusion (with the FORM clause in LAYOUT sections) and ..."
---

# The SUBDIALOG clause

## Purpose of SUBDIALOG

The `SUBDIALOG` clause defines a [declarative dialog](2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs.") to be attached to the
current [procedural `DIALOG`
block](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.").

By using form inclusion (with the `FORM` clause in `LAYOUT`
sections) and declarative dialogs + `SUBDIALOG`, you enforce code reusability in your
application sources.

## Defining the declarative dialog

The declarative dialog is implemented outside the scope of the using `DIALOG`
block, at the same level as a function.

The declarative dialog can be defined in a different module, to be reused in other
`DIALOG` instructions. The sub-dialog module must be imported with the [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") instruction.

Like other module elements such as functions and reports, the name specification is mandatory
when defining a declarative dialog. The name of the declarative dialog will be referenced in a
`SUBDIALOG` clause of a procedural dialog instruction.

In the "comment.4gl" module:

```
PUBLIC TYPE t_comment RECORD
        c_text VARCHAR(200),
        c_checked BOOLEAN
    END RECORD
DIALOG comment_input(rc t_comment INOUT)
   ...
END DIALOG
```

In the using module (note that we use the module prefix here):

```
IMPORT FGL comment
...
FUNCTION mydialog()
   DEFINE r_comment comment.t_comment
   DIALOG ...
      ...
      SUBDIALOG comment.comment_input(r_comment)
      ...
   END DIALOG
END FUNCTION
```

See also [Identifying sub-dialogs in DIALOG](2081-identifying-sub-dialogs-in-dialog.md "Sub-dialogs need to be identified by a name to distinguish the different contexts.").

## Sub-dialogs in form definitions

Implementing a sub-dialog as a declarative dialog in a separate module is typically used in
conjunction with the [FORM](1716-form-clause.md "Reuse the definition of a form in the current form.") clause, in the
`LAYOUT` section of form specification files:

```
LAYOUT
...
FORM "comment"
...
END
```

## Semantics with SUBDIALOG

In terms of semantics, behavior and control block execution, a declarative dialog attached to a
procedural dialog with `SUBDIALOG`, behaves like a sub-dialog that is defined inside
the procedural `DIALOG` block.

For example, [`BEFORE INPUT`](1940-before-input-block.md)
inside an `INPUT` block of a declarative dialog will be executed when the focus goes
to one of the fields of that sub-dialog.

## Scope of dialog instructions

Other sub-dialogs can reference the attached declarative dialog in the current scope.

For example, to execute a [`NEXT
FIELD`](1955-next-field-instruction.md) instruction referencing a field in another
sub-dialog:

```
DIALOG ... -- Parent dialog block
   ...
   NEXT FIELD the_comment  -- Field of the declarative dialog.
   ...
END DIALOG
```

## Scope of DIALOG keyword

When using the `DIALOG` keyword inside a declarative dialog block to use
`ui.Dialog` class methods, it references the current procedural dialog
object:

```
DIALOG comment_input()
   ...
   CALL DIALOG.setFieldActive("the_comment",TRUE)
   ...
END DIALOG
```

## Writing generic code

To be reused by different procedural `DIALOG` instructions, the code of
sub-dialog modules must be generic. However, if the sub-dialog code needs to interact with
the parent `DIALOG`, it must be possible to call a function from the parent
`DIALOG`.

You achieve this by using [function references](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.").
Parent modules can then configure the sub-dialog module at runtime, with callback functions:

1. Create a user-defined `TYPE` with the `FUNCTION` type matching the
   callback function of the using
   module:

   ```
   PUBLIC TYPE t_event_callback FUNCTION (event SMALLINT)
   ```
2. Define a private module variable, with the declared function type. If you want to keep it
   private to the module, define a setter function to assign the variable with the callback function
   reference:

   ```
   PRIVATE DEFINE _event_callback t_event_callback
   ```
3. Define public constants to identify the type of event to be signaled to the parent
   dialog:

   ```
   PUBLIC CONSTANT c_comment_changed = 101
   PUBLIC CONSTANT c_comment_checked = 102
   ```
4. Implement the function to notify the parent
   dialog:

   ```
   PRIVATE FUNCTION _signal_parent(ne SMALLINT) RETURNS ()
       IF _event_callback IS NOT NULL THEN
           CALL _event_callback(ne)
       END IF
   END FUNCTION
   ```
5. Define the sub-dialog with the callback function reference
   parameter:

   ```
   DIALOG comment_input(
       cr t_comment INOUT,
       can_clear STRING,
       event_callback t_event_callback
   )
   ```
6. In the sub0dialog block triggers, call the function to signal events to the parent
   dialog:

   ```
           ON CHANGE is_checked
              CALL _signal_parent(c_comment_checked)
   ```
7. In the parent dialog block, define the callback function to be bound to the
   sub-dialog:

   ```
   FUNCTION event_callback(event SMALLINT)
       CASE event
         WHEN comment.c_comment_checked
            LET msg = "Comment checked..."
         ...
       END CASE
   END FUNCTION
   ```
8. Provide the reference to the callback function as parameter of the sub-dialog, in the
   `SUBDIALOG` instruction of the parent
   `DIALOG`:

   ```
   SUBDIALOG comment_input(cust_comment, TRUE, (FUNCTION event_callback))
   ```

There is no need to implement a lot of complex callback functions: The main purpose is to
indicate to the parent `DIALOG`, that something happened in the sub-dialog. The
parent `DIALOG` can then query the sub-dialog module for more information, as long as
the sub-dialog module provides functions to query its status.

For a complete example, see [Example 3: DIALOG with SUBDIALOG](2149-example-3-dialog-with-subdialog.md).
