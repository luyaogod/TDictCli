---
title: "Understanding the INPUT ARRAY instruction"
source: "fgl-topics/c_fgl_InputArray_002.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Understanding the INPUT ARRAY instruction"
type: "concept"
---

# Understanding the INPUT ARRAY instruction

> The INPUT ARRAY dialog controls a list of records than can be directly edited.

`INPUT ARRAY` is designed to browse and modify a list of records, binding a static
or dynamic array model to a [screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
of the current displayed form.

An `INPUT ARRAY` instruction supports additional features,
built-in sort and search, multi-row selection and list modification triggers.
For a detailed description of these features,
see [Table views](2315-table-views.md "Describes how to implement table/list views.").

Use the `INPUT ARRAY` instruction to let the end user update, delete and create
new records in a list, after fetching a result set from the database. The result set is produced
with a database cursor executing a `SELECT` statement. The `SELECT`
SQL statement is usually completed at runtime with a `WHERE` clause produced
from a `CONSTRUCT` dialog.

The `INPUT ARRAY` instruction associates a program array of records with a
screen-array defined in a form so that the user can update the list of records. The
`INPUT ARRAY` statement activates the current form (the form that was most
recently displayed or the form in the current window).

An `INPUT ARRAY` will start with an empty list of rows, unless the [`WITHOUT DEFAULTS`](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.") option is used.
Using `WITHOUT DEFAULTS` clause instructs the `INPUT ARRAY` dialog to
use the current data rows of the program array bound to the dialog. However, when creating a new
row, the field validation rules will apply like in a simple `INPUT` not using the
`WITHOUT DEFAULTS` option.

During the `INPUT ARRAY` execution, the user can edit or delete existing rows,
insert new rows, and move inside the list of records. The program controls the behavior of the
instruction with [dialog control blocks](2013-input-array-control-blocks.md) such as
`BEFORE DELETE`, `BEFORE INSERT`, etc.

When the user clicks on an action view (button), or when an asynchronous event occurs, [dialog interaction blocks](2027-input-array-interaction-blocks.md) like `ON ACTION`
are executed.

The code inside an `INPUT ARRAY` dialog can use [control instructions](2033-input-array-control-instructions.md), [dialog control functions](2223-dialog-control-functions.md "The language provides several built-in functions and operators to be used in a dialog instruction."), and the [`ui.Dialog`](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.") class, to implement
the dialog behavior.

To terminate the `INPUT ARRAY` execution, the user can validate (or cancel)
the dialog to commit (or invalidate) the modifications made in the list of records.

When the statement completes execution, the program must test the `int_flag` variable to check if the dialog was validated (or
canceled) and then use `INSERT`, `DELETE`, or `UPDATE`
SQL statements to modify the appropriate database tables. The database can also be updated during
the execution of the `INPUT ARRAY` statement.

> **Tip:**
>
> A `DISPLAY ARRAY` dialog is in read-only mode for navigating the list of records
> and needs a user action to perform a modification to the record list. An `INPUT
> ARRAY` dialog offers record list navigation with direct modification; no user action is
> required to modify the data.
>
> `DISPLAY ARRAY` is better adapted to modern graphical user interfaces, while
> `INPUT ARRAY` is better adapted to TUI mode ergonomics. It is also easier to program
> a `DISPLAY ARRAY` with modification triggers compared to the implementation of an
> `INPUT ARRAY`, especially when the underlying SQL table must be updated during the
> dialog execution.
>
> For modern GUI applications, we recommend using `DISPLAY ARRAY` with modification
> triggers ([`ON APPEND`](1984-on-append-block.md), [`ON INSERT`](1985-on-insert-block.md), [`ON UPDATE`](1986-on-update-block.md), [`ON DELETE`](1987-on-delete-block.md)).

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
